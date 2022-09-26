package workflow

import (
	"context"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

const (
	tokenAddress = "" // Provide your token address here, which is also contract address.
	tokenID      = "" // Provide the token id being listed for sale.
)

func DemoOrdersWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// The amount (listing price) should be in Wei for Eth tokens, see https://docs.starkware.co/starkex-v4/starkex-deep-dive/starkex-specific-concepts
	// and https://eth-converter.com/
	ethAddress := l1signer.GetAddress()                         // Address of the user listing for sale.
	sellToken := imx.SignableERC721Token(tokenID, tokenAddress) // NFT Token
	buyToken := imx.SignableETHToken()                          // The listed asset can be bought with Ethereum
	createOrderRequest := &api.GetSignableOrderRequest{
		AmountBuy:  strconv.FormatUint(amount, 10),
		AmountSell: "1",
		Fees:       nil,
		TokenBuy:   buyToken,
		TokenSell:  sellToken,
		User:       ethAddress,
	}
	createOrderRequest.SetExpirationTimestamp(0)

	// Create order will list the given asset for sale.
	createOrderResponse, err := c.CreateOrder(ctx, l1signer, l2signer, createOrderRequest)
	if err != nil {
		log.Panicf("error calling CreateOrder workflow: %v", err)
	}

	createOrderResponseStr, err := prettyStruct(createOrderResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("CreateOrder response:\n%v\n", createOrderResponseStr)

	cancelOrderRequest := api.GetSignableCancelOrderRequest{
		OrderId: createOrderResponse.OrderId,
	}

	// Cancel Order removes the listed asset from sale. Let's remove the above listed asset from sale.
	cancelOrderResponse, err := c.CancelOrder(ctx, l1signer, l2signer, cancelOrderRequest)
	if err != nil {
		log.Panicf("error calling CreateOrder workflow: %v", err)
	}

	cancelOrderResponseStr, err := prettyStruct(cancelOrderResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("CancelOrder response:\n%v\n", cancelOrderResponseStr)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
