package burn

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	burnWorkflow "immutable.com/imx-core-sdk-golang/workflows/burn"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	"immutable.com/imx-core-sdk-golang/workflows/utils"
)

func Demo_BurnWorkflow(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_BurnWorkflow")

	signableToken := utils.NewSignableTokenEth()
	tradeRequest := types.GetSignableBurnRequest{
		Amount: "100000000",
		Sender: l1signer.GetAddress(),
		Token:  signableToken,
	}

	burnResponse, err := burnWorkflow.Burn(ctx, api, l1signer, l2signer, tradeRequest)
	if err != nil {
		log.Panicf("error calling burn workflow: %v", err)
	}
	val, _ := json.MarshalIndent(burnResponse, "", "  ")
	log.Printf("Burn response:\n%s\n", val)

	transferID := strconv.FormatInt(*burnResponse.TransferID, 10) // GetBurn method requires transferID as a string
	transfer, err := burnWorkflow.GetBurn(ctx, api, transferID)
	if err != nil {
		log.Panicf("error calling GetBurn workflow: %v", err)
	}
	val, _ = json.MarshalIndent(transfer, "", "  ")
	log.Printf("GetBurn response:\n%s\n", val)

	log.Println("Running Demo_BurnWorkflow completed")
	log.Println("-------------------------------------------------------")
}
