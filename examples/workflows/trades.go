package main

import (
	"context"
	"encoding/json"
	"log"

	immutable "github.com/immutable/imx-core-sdk-golang"
	"github.com/immutable/imx-core-sdk-golang/api"
)

func demoTradesWorkflow(ctx context.Context, c *immutable.Client, l1signer immutable.L1Signer, l2signer immutable.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	tradeRequest := api.GetSignableTradeRequest{
		Fees:    nil,
		OrderId: 0, // Order ID to execute the trade on.
	}
	tradeRequest.SetExpirationTimestamp(0)
	tradeResponse, err := c.CreateTrade(ctx, l1signer, l2signer, tradeRequest)

	if err != nil {
		log.Panicf("error calling trades workflow: %v", err)
	}

	val, _ := json.MarshalIndent(tradeResponse, "", "  ")
	log.Printf("trade response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
