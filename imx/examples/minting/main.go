package main

import (
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer := common.CommonInitialise(".env")

	tokenID := envs["MINT_TOKEN_ID"]
	tokenAddress := envs["MINT_TOKEN_ADDRESS"]
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
						Blueprint: blueprint,
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

	mintTokensResponseStr, err := common.PrettyStruct(mintTokensResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Mint Tokens response:\n%v\n", mintTokensResponseStr)
}
