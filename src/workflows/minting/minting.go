package minting

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/mints"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

type MintFee struct {
	Recipient  *string  `json:"recipient"`
	Percentage *float64 `json:"percentage" validate:"max=100,gt=0"`
}

type MintableERC721TokenData struct {
	ID        *string    `json:"id"`
	Blueprint *string    `json:"blueprint"`
	Royalties []*MintFee `json:"royalties,omitempty"` // token-level overridable fees (optional)
}

type User struct {
	User   *string                    `json:"ether_key"`
	Tokens []*MintableERC721TokenData `json:"tokens"`
}

type UnsignedMintRequest struct {
	ContractAddress *string    `json:"contract_address" validate:"required,eth_addr"`
	Royalties       []*MintFee `json:"royalties,omitempty" validate:"dive"` // contract-level (optional)
	Users           []*User    `json:"users" validate:"required,dive,min=1"`
	AuthSignature   string     `json:"auth_signature" validate:"required"`
}

func MintTokensWorkflow(ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	unsignedMintRequest []*UnsignedMintRequest) (*models.MintTokensResponse, error) {

	var mintRequest []*models.MintRequest
	for _, eachMintRequst := range unsignedMintRequest {
		mintRequestInBytes, err := json.Marshal(eachMintRequst)
		log.Printf(string(mintRequestInBytes))
		requestHash := crypto.Keccak256Hash(mintRequestInBytes)
		authSignatureInBytes, err := l1signer.SignMessage(requestHash.String())
		if err != nil {
			return nil, err
		}

		authSignatureEncodedInHex := hexutil.Encode(authSignatureInBytes)
		log.Printf("Auth Signature: %v", authSignatureEncodedInHex)

		var mintFees []*models.MintFee
		for _, eachMintFee := range eachMintRequst.Royalties {
			mintFees = append(mintFees, &models.MintFee{
				Percentage: eachMintFee.Percentage,
				Recipient:  eachMintFee.Recipient,
			})
		}

		var mintToUsers []*models.MintUser
		for _, eachMintUser := range eachMintRequst.Users {
			var mintTokens []*models.MintTokenDataV2
			for _, eachMintToken := range eachMintUser.Tokens {

				var mintFeesPerToken []*models.MintFee
				for _, eachMintFeePerToken := range eachMintToken.Royalties {
					mintFeesPerToken = append(mintFeesPerToken, &models.MintFee{
						Percentage: eachMintFeePerToken.Percentage,
						Recipient:  eachMintFeePerToken.Recipient,
					})
				}

				mintTokens = append(mintTokens, &models.MintTokenDataV2{
					Blueprint: *eachMintToken.Blueprint,
					ID:        eachMintToken.ID,
					Royalties: mintFeesPerToken,
				})
			}
			mintToUsers = append(mintToUsers, &models.MintUser{
				Tokens: mintTokens,
				User:   eachMintUser.User,
			})
		}

		mintRequest = append(mintRequest, &models.MintRequest{
			AuthSignature:   &authSignatureEncodedInHex,
			ContractAddress: eachMintRequst.ContractAddress,
			Royalties:       mintFees,
			Users:           mintToUsers,
		})
		if err != nil {
			return nil, fmt.Errorf("error marshaling MintRequest: %v", err)
		}
	}

	mintTokenParams := mints.NewMintTokensParamsWithContext(ctx)
	mintTokenParams.SetMintTokensRequestV2(mintRequest)
	mintTokensOk, err := api.Mints.MintTokens(mintTokenParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Mints.MintTokens`: %v", err)
	}
	return mintTokensOk.GetPayload(), nil
}
