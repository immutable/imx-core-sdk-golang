package main

import (
	"context"
	"log"

	immutable "github.com/immutable/imx-core-sdk-golang"
)

func demoMintingTokens(ctx context.Context, c *immutable.Client, l1signer immutable.L1Signer, tokenID, tokenAddress string) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	ethAddress := l1signer.GetAddress()
	blueprint := "123"
	var royaltyPercentage float32 = 1
	var mintableToken = immutable.UnsignedMintRequest{
		ContractAddress: tokenAddress,
		Royalties: []immutable.MintFee{
			{
				Percentage: royaltyPercentage,
				Recipient:  ethAddress,
			},
		},
		Users: []immutable.User{
			{
				User: ethAddress,
				Tokens: []immutable.MintableTokenData{
					{
						ID: tokenID,
						Royalties: []immutable.MintFee{
							{
								Percentage: royaltyPercentage,
								Recipient:  ethAddress,
							},
						},
						Blueprint: &blueprint,
					},
				},
			},
		},
	}

	request := make([]immutable.UnsignedMintRequest, 1)
	request[0] = mintableToken

	mintTokensResponse, err := c.Mint(ctx, l1signer, request)
	if err != nil {
		log.Panicf("error in minting.MintTokensWorkflow: %v", err)
	}

	mintTokensResponseStr, err := prettyStruct(mintTokensResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Mint Tokens response:\n%v\n", mintTokensResponseStr)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
