package trades

import (
	"context"
	"encoding/json"
	"immutable.com/imx-core-sdk-golang/api/client/trades"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	tradesWorkflow "immutable.com/imx-core-sdk-golang/workflows/trades"
	"log"
)

func Demo_TradesWorkflow(ctx context.Context, tradesApi trades.ClientService, l1signer signers.L1Signer, l2signer signers.L2Signer) {

	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_TradesWorkflow")

	tradeRequest := models.GetSignableTradeRequest{
		ExpirationTimestamp: 0,
		Fees:                nil,
		OrderID:             nil, // Requires orders workflow to generate an order id
	}
	tradeResponse, err := tradesWorkflow.CreateTrade(ctx, tradesApi, l1signer, l2signer, tradeRequest)

	if err != nil {
		log.Fatalf("error calling trades workflow: %v", err)
	}

	val, _ := json.MarshalIndent(tradeResponse, "", "  ")
	log.Printf("trade response:\n%s\n", val)

	log.Println("Running Demo_TradesWorkflow completed")
	log.Println("-------------------------------------------------------")
}
