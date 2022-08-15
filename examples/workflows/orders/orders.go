package orders

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/tokens"
	converters "immutable.com/imx-core-sdk-golang/utils"
	ordersWorkflow "immutable.com/imx-core-sdk-golang/workflows/orders"
)

const (
	tokenAddress = "" // Provide your token address here, which is also contract address.
	tokenID      = "" // Provide the token id being listed for sale.
)

func DemoOrdersWorkflow(ctx context.Context, apiClient *api.APIClient, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// The amount (listing price) should be in Wei for Eth tokens, see https://docs.starkware.co/starkex-v4/starkex-deep-dive/starkex-specific-concepts
	// and https://eth-converter.com/
	amountBuy, _ := converters.ToUnquantized("0.3", converters.EtherDecimals)
	ethAddress := l1signer.GetAddress()                               // Address of the user listing for sale.
	sellToken := tokens.NewSignableTokenERC721(tokenID, tokenAddress) // NFT Token
	buyToken := tokens.NewSignableTokenEth()                          // The listed asset can be bought with Ethereum
	createOrderRequest := &api.GetSignableOrderRequest{
		AmountBuy:  amountBuy.String(),
		AmountSell: "1",
		Fees:       nil,
		TokenBuy:   *buyToken,
		TokenSell:  *sellToken,
		User:       ethAddress,
	}
	createOrderRequest.SetExpirationTimestamp(0)

	// Create order will list the given asset for sale.
	createOrderResponse, err := ordersWorkflow.CreateOrder(ctx, apiClient, l1signer, l2signer, createOrderRequest)
	if err != nil {
		log.Panicf("error calling CreateOrder workflow: %v", err)
	}

	createOrderResponseStr, err := utils.PrettyStruct(createOrderResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("CreateOrder response:\n%v\n", createOrderResponseStr)

	cancelOrderRequest := api.GetSignableCancelOrderRequest{
		OrderId: createOrderResponse.OrderId,
	}

	// Cancel Order removes the listed asset from sale. Let's remove the above listed asset from sale.
	cancelOrderResponse, err := ordersWorkflow.CancelOrder(ctx, apiClient, l1signer, l2signer, cancelOrderRequest)
	if err != nil {
		log.Panicf("error calling CreateOrder workflow: %v", err)
	}

	cancelOrderResponseStr, err := utils.PrettyStruct(cancelOrderResponse)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("CancelOrder response:\n%v\n", cancelOrderResponseStr)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
