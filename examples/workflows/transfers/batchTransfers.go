package transfers

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/tokens"
	transfersWorkflow "immutable.com/imx-core-sdk-golang/workflows/transfers"
)

func DemoBatchTransferWorkflow(ctx context.Context, apiClient *api.APIClient, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// To declare tokens, use tokens.NewSignableTokenERC721 method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken1 := tokens.NewSignableTokenERC721("Token ID Here", "Token Address Hex Here")
	signableToken2 := tokens.NewSignableTokenERC721("Token ID Here", "Token Address Hex Here")
	amount := "1"
	sender := l1signer.GetAddress()
	receiver := "Receiver Address Here"

	transferRequest1 := api.SignableTransferDetails{
		Amount:   amount,
		Receiver: receiver,
		Token:    *signableToken1,
	}
	transferRequest2 := api.SignableTransferDetails{
		Amount:   amount,
		Receiver: receiver,
		Token:    *signableToken2,
	}

	batchTransferRequest := api.GetSignableTransferRequest{
		SenderEtherKey: sender,
		SignableRequests: []api.SignableTransferDetails{
			transferRequest1,
			transferRequest2,
		},
	}

	response, err := transfersWorkflow.CreateBatchTransfer(ctx, apiClient, l1signer, l2signer, batchTransferRequest)
	if err != nil {
		log.Panicf("error calling batch transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
