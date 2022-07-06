package burn

import (
	"context"
	"encoding/json"
	"immutable.com/imx-core-sdk-golang/api/client/transfers"
	"immutable.com/imx-core-sdk-golang/signers"
	burnWorkflow "immutable.com/imx-core-sdk-golang/workflows/burn"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	"immutable.com/imx-core-sdk-golang/workflows/utils"
	"log"
	"strconv"
)

func Demo_BurnWorkflow(ctx context.Context, transfersApi transfers.ClientService, l1signer signers.L1Signer, l2signer signers.L2Signer) {

	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_BurnWorkflow")

	signableToken := utils.NewSignableTokenEth()
	tradeRequest := types.GetSignableBurnRequest{
		Amount: "100000000",
		Sender: l1signer.GetAddress(),
		Token:  signableToken,
	}

	burnResponse, err := burnWorkflow.Burn(ctx, transfersApi, l1signer, l2signer, tradeRequest)
	if err != nil {
		log.Fatalf("error calling burn workflow: %v", err)
	}
	val, _ := json.MarshalIndent(burnResponse, "", "  ")
	log.Printf("Burn response:\n%s\n", val)

	transferId := strconv.FormatInt(*burnResponse.TransferID, 10)
	transfer, err := burnWorkflow.GetBurn(ctx, transfersApi, transferId)
	if err != nil {
		log.Fatalf("error calling GetBurn workflow: %v", err)
	}
	val, _ = json.MarshalIndent(transfer, "", "  ")
	log.Printf("GetBurn response:\n%s\n", val)

	log.Println("Running Demo_BurnWorkflow completed")
	log.Println("-------------------------------------------------------")
}
