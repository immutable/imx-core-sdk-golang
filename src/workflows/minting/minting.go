package minting

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/mints"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

// Note: 	These structs need to be in the order defined as below and should match the
//   		implementation in imx-engine to pass the validation step while minting tokens.

type MintFee struct {
	Recipient  *string  `json:"recipient"`
	Percentage *float64 `json:"percentage" validate:"max=100,gt=0"`
}

type MintableTokenData struct {
	ID        *string    `json:"id"`
	Blueprint *string    `json:"blueprint"`
	Royalties []*MintFee `json:"royalties,omitempty"` // token-level overridable fees (optional)
}

type User struct {
	User   *string              `json:"ether_key"`
	Tokens []*MintableTokenData `json:"tokens"`
}

type UnsignedMintRequest struct {
	ContractAddress *string    `json:"contract_address" validate:"required,eth_addr"`
	Royalties       []*MintFee `json:"royalties,omitempty" validate:"dive"` // contract-level (optional)
	Users           []*User    `json:"users" validate:"required,dive,min=1"`
	AuthSignature   string     `json:"auth_signature" validate:"required"`
}

// MintTokensWorkflow assists in minting tokens to the given imx user.
func MintTokensWorkflow(ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	unsignedMintRequest []*UnsignedMintRequest) (*models.MintTokensResponse, error) {

	mintRequest := make([]*models.MintRequest, len(unsignedMintRequest))
	for requestIndex, eachMintRequst := range unsignedMintRequest {
		mintRequestInBytes, err := json.Marshal(eachMintRequst)
		if err != nil {
			return nil, fmt.Errorf("error eachMintRequst MintRequest: %v", err)
		}
		requestHash := crypto.Keccak256Hash(mintRequestInBytes)
		authSignatureInBytes, err := l1signer.SignMessage(requestHash.String())
		if err != nil {
			return nil, fmt.Errorf("error in signing eachMintRequst message: %v", err)
		}

		mintFees := make([]*models.MintFee, len(eachMintRequst.Royalties))
		for index, eachMintFee := range eachMintRequst.Royalties {
			mintFees[index] = &models.MintFee{
				Percentage: eachMintFee.Percentage,
				Recipient:  eachMintFee.Recipient,
			}
		}

		mintToUsers := make([]*models.MintUser, len(eachMintRequst.Users))
		for userIndex, eachMintUser := range eachMintRequst.Users {
			mintTokens := make([]*models.MintTokenDataV2, len(eachMintUser.Tokens))
			for tokenIndex, eachMintToken := range eachMintUser.Tokens {

				mintFeesPerToken := make([]*models.MintFee, len(eachMintToken.Royalties))
				for royaltyIndex, eachMintFeePerToken := range eachMintToken.Royalties {
					mintFeesPerToken[royaltyIndex] = &models.MintFee{
						Percentage: eachMintFeePerToken.Percentage,
						Recipient:  eachMintFeePerToken.Recipient,
					}
				}

				mintTokens[tokenIndex] = &models.MintTokenDataV2{
					Blueprint: *eachMintToken.Blueprint,
					ID:        eachMintToken.ID,
					Royalties: mintFeesPerToken,
				}
			}
			mintToUsers[userIndex] = &models.MintUser{
				Tokens: mintTokens,
				User:   eachMintUser.User,
			}
		}

		authSignatureEncodedInHex := hexutil.Encode(authSignatureInBytes)
		mintRequest[requestIndex] = &models.MintRequest{
			AuthSignature:   &authSignatureEncodedInHex,
			ContractAddress: eachMintRequst.ContractAddress,
			Royalties:       mintFees,
			Users:           mintToUsers,
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
