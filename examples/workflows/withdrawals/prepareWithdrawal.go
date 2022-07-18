package withdrawals

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	converters "immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/workflows/utils"
	withdrawalsWorkflow "immutable.com/imx-core-sdk-golang/workflows/withdrawals"
)

func Demo_PrepareWithdrawalWorkflow(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_PrepareWithdrawalWorkflow")

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken := utils.NewSignableTokenEth()
	amountInt, _ := converters.ParseEtherToWei("0.3")
	amount := amountInt.String()
	withdrawalRequest := models.GetSignableWithdrawalRequest{
		Amount: &amount,
		Token:  signableToken,
	}

	/*
		// Uncomment for ERC721
		signableToken := utils.NewSignableTokenERC721("7", "0x0fb969a08c7c39ba99c1628b59c0b7e5611bd396")
		amount := "1"
		withdrawalRequest := models.GetSignableWithdrawalRequest{
			Amount: &amount,
			Token:  signableToken,
		}
	*/

	response, err := withdrawalsWorkflow.PrepareWithdrawal(ctx, api, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Fatalf("error calling prepare withdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Println("Running Demo_PrepareWithdrawalWorkflow completed")
	log.Println("-------------------------------------------------------")
}
