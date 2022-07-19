package minting

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows/minting"
)

var (
	contractAddress string = "" // Provide your collection contract address
)

func Demo_MintingTokens(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer) {
	ethAddress := l1signer.GetAddress()
	tokenID := "5"
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
				Tokens: []*minting.MintableTokenData{
					{
						ID: &tokenID,
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

	request := make([]*minting.UnsignedMintRequest, 1)
	request[0] = &mintableToken

	mintTokensResponse, err := minting.MintTokensWorkflow(ctx, api, l1signer, request)
	if err != nil {
		log.Fatalf("error in minting.MintTokensWorkflow: %v", err)
	}

	mintTokensResponseStr, err := utils.PrettyStruct(mintTokensResponse)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Mint Tokens response:\n%v\n", mintTokensResponseStr)
}
