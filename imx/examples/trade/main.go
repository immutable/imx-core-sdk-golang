package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, _, c, l1signer, l2signer := common.CommonInitialise()

	// Order ID to execute the trade on.
	//
	// Note: 	This orderID to be obtained by performing the CreateOrder using the seller account.
	orderID := int32(166854)

	// Trade demo.
	DemoTradesWorkflow(ctx, c, orderID, l1signer, l2signer)
}

// Trade workflow demo.
//
// Note: 	This trade request should be performed by the buyer account. Please make sure you use a different account to seller account.
// 			The trade will be performed at the listed price as per the OrderId.
// 			Please make sure the buyer imx account contains sufficient funds to perform the purchase (trade).
//
// Tip: 	You can use Deposit workflow to deposit tokens from your L1 Ethereum wallet to L2 Imx Wallet.
//
func DemoTradesWorkflow(ctx context.Context, c *imx.Client, orderID int32, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	tradeRequest := api.GetSignableTradeRequest{
		Fees:    nil,
		OrderId: orderID,
	}
	tradeRequest.SetExpirationTimestamp(0)
	tradeResponse, err := c.CreateTrade(ctx, l1signer, l2signer, tradeRequest)

	if err != nil {
		log.Panicf("error calling trades workflow: %v", err)
	}

	val, _ := json.MarshalIndent(tradeResponse, "", "  ")
	log.Printf("trade response:\n%s\n", val)

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
