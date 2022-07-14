package minting

import (
	"context"
	"log"
	"strings"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows/minting"
)

var (
	contractAddress string = "" // Provide your collection contract address
)

func Demo_MintingTokens(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer) {

	ethAddress := strings.ToLower(l1signer.GetAddress())
	tokenId := "5"
	blueprint := "123"
	var royaltyPercentage float64 = 1
	var mintableToken = minting.UnsignedMintRequest{
		ContractAddress: &contractAddress,
		Royalties: []*minting.MintFee{
			{
				Percentage: &royaltyPercentage,
				Recipient:  &ethAddress,
			},
		},
		Users: []*minting.User{
			{
				User: &ethAddress,
				Tokens: []*minting.MintableERC721TokenData{
					{
						ID: &tokenId,
						Royalties: []*minting.MintFee{
							{
								Percentage: &royaltyPercentage,
								Recipient:  &ethAddress,
							},
						},
						Blueprint: &blueprint,
					},
				},
			},
		},
	}
	var request []*minting.UnsignedMintRequest
	request = append(request, &mintableToken)

	mintTokensResponse, err := minting.MintTokensWorkflow(ctx, api, l1signer, request)
	if err != nil {
		log.Fatalf("error generating signature from MintRequest data: %v", err)
	}

	// mintTokensResponseStr, err := utils.PrettyStruct(mintTokensResponse)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	log.Printf("Mint Tokens response:\n%v\n", mintTokensResponse)
}
