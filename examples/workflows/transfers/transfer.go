package transfers

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	transfersWorkflow "immutable.com/imx-core-sdk-golang/workflows/transfers"
	"immutable.com/imx-core-sdk-golang/workflows/utils"
)

func Demo_TransferWorkflow(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_TransferWorkflow")

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken := utils.NewSignableTokenEth()
	amount := "100000000"
	sender := l1signer.GetAddress()
	receiver := "Set receiver address here"
	transferRequest := models.GetSignableTransferRequestV1{
		Amount:   &amount,
		Sender:   &sender,
		Token:    signableToken,
		Receiver: &receiver,
	}

	response, err := transfersWorkflow.CreateTransfer(ctx, api, l1signer, l2signer, transferRequest)
	if err != nil {
		log.Panicf("error calling transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Println("Running Demo_TransferWorkflow completed")
	log.Println("-------------------------------------------------------")
}
