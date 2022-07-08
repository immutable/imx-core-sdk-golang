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

func Demo_BatchTransferWorkflow(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer) {

	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_BatchTransferWorkflow")

	signableToken1 := utils.NewSignableTokenERC721("Token ID Here", "Token Address Hex Here")
	signableToken2 := utils.NewSignableTokenERC721("Token ID Here", "Token Address Hex Here")
	amount := "1"
	sender := l1signer.GetAddress()
	receiver := "Receiver Address Here"

	transferRequest1 := models.SignableTransferDetails{
		Amount:   &amount,
		Receiver: &receiver,
		Token:    signableToken1,
	}
	transferRequest2 := models.SignableTransferDetails{
		Amount:   &amount,
		Receiver: &receiver,
		Token:    signableToken2,
	}

	batchTransferRequest := models.GetSignableTransferRequest{
		SenderEtherKey: &sender,
		SignableRequests: []*models.SignableTransferDetails{
			&transferRequest1,
			&transferRequest2,
		},
	}

	response, err := transfersWorkflow.CreateBatchTransfer(ctx, api, l1signer, l2signer, batchTransferRequest)
	if err != nil {
		log.Fatalf("error calling batch transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Println("Running Demo_BatchTransferWorkflow completed")
	log.Println("-------------------------------------------------------")
}
