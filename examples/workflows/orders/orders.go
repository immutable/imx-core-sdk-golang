package orders

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	converters "immutable.com/imx-core-sdk-golang/utils"
	ordersWorkflow "immutable.com/imx-core-sdk-golang/workflows/orders"
	coreUtils "immutable.com/imx-core-sdk-golang/workflows/utils"
)

const (
	tokenAddress = "" // Provide your token address here, which is also contract address.
	tokenID      = "" // Provide the token id being listed for sale.
)

func Demo_OrdersWorkflow(ctx context.Context, apiClient *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	fnName := utils.GetCurrentFunctionName()
	log.Println("Running ", fnName)

	// Unquantized amount (listing price). See https://immutable.atlassian.net/wiki/spaces/PPS/pages/1895629681/Quantisation+Cheat+Sheet
	amountInt, _ := converters.ParseEtherToWei("0.3")
	amountBuy := amountInt.String()
	amountSell := "1"                                                    // Quantity to list for sale
	ethAddress := l1signer.GetAddress()                                  // Address of the user listing for sale.
	sellToken := coreUtils.NewSignableTokenERC721(tokenID, tokenAddress) // NFT Token
	buyToken := coreUtils.NewSignableTokenEth()                          // The listed asset can be bought with Ethereum
	createOrderRequest := models.GetSignableOrderRequest{
		AmountBuy:           &amountBuy,
		AmountSell:          &amountSell,
		ExpirationTimestamp: 0,
		Fees:                nil,
		TokenBuy:            buyToken,
		TokenSell:           sellToken,
		User:                &ethAddress,
	}

	// Create order will list the given asset for sale.
	createOrderResponse, err := ordersWorkflow.CreateOrder(ctx, apiClient, l1signer, l2signer, createOrderRequest)
	if err != nil {
		log.Fatalf("error calling CreateOrder workflow: %v", err)
	}

	createOrderResponseStr, err := utils.PrettyStruct(createOrderResponse)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CreateOrder response:\n%v\n", createOrderResponseStr)

	cancelOrderRequest := models.GetSignableCancelOrderRequest{
		OrderID: createOrderResponse.OrderID,
	}

	// Cancel Order removes the listed asset from sale. Let's remove the above listed asset from sale.
	cancelOrderResponse, err := ordersWorkflow.CancelOrder(ctx, apiClient, l1signer, l2signer, cancelOrderRequest)
	if err != nil {
		log.Fatalf("error calling CreateOrder workflow: %v", err)
	}

	cancelOrderResponseStr, err := utils.PrettyStruct(cancelOrderResponse)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("CancelOrder response:\n%v\n", cancelOrderResponseStr)

	log.Println("Running Demo_OrdersWorkflow completed")
	log.Println("-------------------------------------------------------")
}
