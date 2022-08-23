package transfers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
	"github.com/immutable/imx-core-sdk-golang/tokens"
	transfersWorkflow "github.com/immutable/imx-core-sdk-golang/workflows/transfers"
)

func DemoTransferWorkflow(ctx context.Context, apiClient *api.APIClient, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/immutable/docs/token-data-object
	signableToken := tokens.NewSignableTokenEth()
	amount := "100000000"
	sender := l1signer.GetAddress()
	receiver := "Set receiver address here"
	transferRequest := api.GetSignableTransferRequestV1{
		Amount:   amount,
		Sender:   sender,
		Token:    *signableToken,
		Receiver: receiver,
	}

	response, err := transfersWorkflow.CreateTransfer(ctx, apiClient, l1signer, l2signer, transferRequest)
	if err != nil {
		log.Panicf("error calling transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
