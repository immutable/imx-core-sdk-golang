package minting

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/signers"
)

// Note: 	These structs need to be in the order defined as below and should match the
//   		implementation in imx-engine to pass the validation step while minting tokens.

type MintFee struct {
	Recipient  string  `json:"recipient"`
	Percentage float32 `json:"percentage" validate:"max=100,gt=0"`
}

type MintableTokenData struct {
	ID        string    `json:"id"`
	Blueprint *string   `json:"blueprint"`
	Royalties []MintFee `json:"royalties,omitempty"` // token-level overridable fees (optional)
}

type User struct {
	User   string              `json:"ether_key"`
	Tokens []MintableTokenData `json:"tokens"`
}

type UnsignedMintRequest struct {
	ContractAddress string    `json:"contract_address" validate:"required,eth_addr"`
	Royalties       []MintFee `json:"royalties,omitempty" validate:"dive"` // contract-level (optional)
	Users           []User    `json:"users" validate:"required,dive,min=1"`
	AuthSignature   string    `json:"auth_signature" validate:"required"`
}

// MintTokensWorkflow assists in minting tokens to the given imx user.
func MintTokensWorkflow(ctx context.Context,
	mintsAPI api.MintsApi,
	l1signer signers.L1Signer,
	unsignedMintRequest []UnsignedMintRequest) (*api.MintTokensResponse, error) {
	mintRequest := make([]api.MintRequest, len(unsignedMintRequest))
	for requestIndex, eachMintRequest := range unsignedMintRequest {
		mintRequestInBytes, err := json.Marshal(eachMintRequest)
		if err != nil {
			return nil, fmt.Errorf("error in json marshaling eachMintRequest: %v", err)
		}
		requestHash := crypto.Keccak256Hash(mintRequestInBytes)
		authSignatureInBytes, err := l1signer.SignMessage(requestHash.String())
		if err != nil {
			return nil, fmt.Errorf("error in signing eachMintRequest message: %v", err)
		}

		mintFees := make([]api.MintFee, len(eachMintRequest.Royalties))
		for index, eachMintFee := range eachMintRequest.Royalties {
			mintFees[index] = api.MintFee{
				Percentage: eachMintFee.Percentage,
				Recipient:  eachMintFee.Recipient,
			}
		}

		mintToUsers := make([]api.MintUser, len(eachMintRequest.Users))
		for userIndex, eachMintUser := range eachMintRequest.Users {
			mintTokens := make([]api.MintTokenDataV2, len(eachMintUser.Tokens))
			for tokenIndex, eachMintToken := range eachMintUser.Tokens {
				mintFeesPerToken := make([]api.MintFee, len(eachMintToken.Royalties))
				for royaltyIndex, eachMintFeePerToken := range eachMintToken.Royalties {
					mintFeesPerToken[royaltyIndex] = api.MintFee{
						Percentage: eachMintFeePerToken.Percentage,
						Recipient:  eachMintFeePerToken.Recipient,
					}
				}

				mintTokens[tokenIndex] = api.MintTokenDataV2{
					Blueprint: eachMintToken.Blueprint,
					Id:        eachMintToken.ID,
					Royalties: mintFeesPerToken,
				}
			}
			mintToUsers[userIndex] = api.MintUser{
				Tokens: mintTokens,
				User:   eachMintUser.User,
			}
		}

		authSignatureEncodedInHex := hexutil.Encode(authSignatureInBytes)
		mintRequest[requestIndex] = api.MintRequest{
			AuthSignature:   authSignatureEncodedInHex,
			ContractAddress: eachMintRequest.ContractAddress,
			Royalties:       mintFees,
			Users:           mintToUsers,
		}
	}

	mintTokensResponse, httpResp, err := mintsAPI.MintTokens(ctx).MintTokensRequestV2(mintRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Mints.MintTokens`: %v, full HTTP response: %v", err, httpResp.Body)
	}
	return mintTokensResponse, nil
}
