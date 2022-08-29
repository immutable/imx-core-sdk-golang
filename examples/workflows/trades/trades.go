package trades

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
	tradesWorkflow "github.com/immutable/imx-core-sdk-golang/workflows/trades"
)

func DemoTradesWorkflow(ctx context.Context, apiClient *api.APIClient, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	tradeRequest := api.GetSignableTradeRequest{
		Fees:    nil,
		OrderId: 0, // Order ID to execute the trade on.
	}
	tradeRequest.SetExpirationTimestamp(0)
	tradeResponse, err := tradesWorkflow.CreateTrade(ctx, apiClient, l1signer, l2signer, tradeRequest)

	if err != nil {
		log.Panicf("error calling trades workflow: %v", err)
	}

	val, _ := json.MarshalIndent(tradeResponse, "", "  ")
	log.Printf("trade response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
