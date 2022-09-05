package workflow

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
)

func DemoMintingTokens(ctx context.Context, c *imx.Client, l1signer imx.L1Signer, tokenID, tokenAddress string) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	ethAddress := l1signer.GetAddress()
	blueprint := "123"
	var royaltyPercentage float32 = 1
	var mintableToken = imx.UnsignedMintRequest{
		ContractAddress: tokenAddress,
		Royalties: []imx.MintFee{
			{
				Percentage: royaltyPercentage,
				Recipient:  ethAddress,
			},
		},
		Users: []imx.User{
			{
				User: ethAddress,
				Tokens: []imx.MintableTokenData{
					{
						ID: tokenID,
						Royalties: []imx.MintFee{
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

	request := make([]imx.UnsignedMintRequest, 1)
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
