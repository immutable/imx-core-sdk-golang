package minting

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows/minting"
)

func DemoMintingTokens(ctx context.Context, mintsAPI api.MintsApi, l1signer signers.L1Signer, tokenId, tokenAddress string) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	ethAddress := l1signer.GetAddress()
	blueprint := "123"
	var royaltyPercentage float32 = 1
	var mintableToken = minting.UnsignedMintRequest{
		ContractAddress: tokenAddress,
		Royalties: []minting.MintFee{
			{
				Percentage: royaltyPercentage,
				Recipient:  ethAddress,
			},
		},
		Users: []minting.User{
			{
				User: ethAddress,
				Tokens: []minting.MintableTokenData{
					{
						ID: tokenId,
						Royalties: []minting.MintFee{
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

	request := make([]minting.UnsignedMintRequest, 1)
	request[0] = mintableToken

	mintTokensResponse, err := minting.MintTokensWorkflow(ctx, mintsAPI, l1signer, request)
	if err != nil {
		log.Panicf("error in minting.MintTokensWorkflow: %v", err)
	}

	mintTokensResponseStr, err := utils.PrettyStruct(mintTokensResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Mint Tokens response:\n%v\n", mintTokensResponseStr)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
