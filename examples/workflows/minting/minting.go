package minting

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows/minting"
)

var (
	contractAddress = "0x4B089f5006778FeF128506427235937D60B74DFB"
)

func Demo_MintingTokens(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer) {

	ethAddress := l1signer.GetAddress()
	tokenId := "4"
	var royaltyPercentage float64 = 2

	var request []*models.MintRequest
	mintableToken := models.MintRequest{
		ContractAddress: &contractAddress,
		Royalties: []*models.MintFee{
			{
				Percentage: &royaltyPercentage,
				Recipient:  &ethAddress,
			},
		},
		Users: []*models.MintUser{
			{
				User: &ethAddress,
				Tokens: []*models.MintTokenDataV2{
					{
						ID: &tokenId,
						Royalties: []*models.MintFee{
							{
								Percentage: &royaltyPercentage,
								Recipient:  &ethAddress,
							},
						},
						Blueprint: `{
							"name": "Forest NFT",
							"description": "This is your 4th nft",
							"image_url":"https://gateway.pinata.cloud/ipfs/QmUWHMh2moaTWks7c9k3VxVGjp3Pk4chYeWhdJYYcna4BA",
							"attack": 423,
							"collectable": true,
							"class": "EnumValue4"   
						  }`,
					},
				},
			},
		},
	}
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
