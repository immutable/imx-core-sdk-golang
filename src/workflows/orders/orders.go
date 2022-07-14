package orders

import (
	"context"
	"fmt"
	"strconv"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/orders"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

// CreateOrder will list the given asset for sale on the marketplace.
func CreateOrder(ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request models.GetSignableOrderRequest) (*models.CreateOrderResponse, error) {

	ethAddress := l1signer.GetAddress()

	request.User = &ethAddress
	getSignableOrderParams := orders.NewGetSignableOrderParamsWithContext(ctx)
	getSignableOrderParams.SetGetSignableOrderRequestV3(&request)
	getSignableOrderOK, err := api.Orders.GetSignableOrder(getSignableOrderParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `ordersClient.GetSignableOrder`: %v", err)
	}

	signableOrder := getSignableOrderOK.GetPayload()
	ethSignature, starkSignature, err := signers.CreateSignatures(signableOrder.SignableMessage, signableOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	createOrderParams := orders.NewCreateOrderParamsWithContext(ctx)
	createOrderParams.SetCreateOrderRequest(&models.CreateOrderRequest{
		AmountBuy:           signableOrder.AmountBuy,   // Unquantized amount (listing price). For more info in quantum, see https://docs.starkware.co/starkex-v4/starkex-deep-dive/starkex-specific-concepts
		AmountSell:          signableOrder.AmountSell,  // Quantity to list for sale
		AssetIDBuy:          signableOrder.AssetIDBuy,  // Token with which this asset can be bought with. For example Eth token can be used to buy this asset.
		AssetIDSell:         signableOrder.AssetIDSell, // Token of the asset being listed for sale. For example, ERC720 is a NFT Token and ERC20 is a fungible token.
		ExpirationTimestamp: signableOrder.ExpirationTimestamp,
		Fees:                request.Fees,
		IncludeFees:         true,
		Nonce:               signableOrder.Nonce,
		StarkKey:            signableOrder.StarkKey,
		StarkSignature:      &starkSignature,
		VaultIDBuy:          signableOrder.VaultIDBuy,
		VaultIDSell:         signableOrder.VaultIDSell,
	})
	createOrderParams.SetXImxEthAddress(&ethAddress)
	createOrderParams.SetXImxEthSignature(&ethSignature)
	createOrderOk, err := api.Orders.CreateOrder(createOrderParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `ordersClient.CreateOrder`: %v", err)
	}
	return createOrderOk.GetPayload(), nil
}

// CancelOrder will remove the listed asset on marketplace from sale.
func CancelOrder(ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request models.GetSignableCancelOrderRequest) (*models.CancelOrderResponse, error) {

	getSignableCancelOrderParams := orders.NewGetSignableCancelOrderParamsWithContext(ctx)

	getSignableCancelOrderParams.SetGetSignableCancelOrderRequest(&request)
	getSignableCancelOrderOk, err := api.Orders.GetSignableCancelOrder(getSignableCancelOrderParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `ordersClient.GetSignableCancelOrder`: %v", err)
	}

	signableCancelOrder := getSignableCancelOrderOk.GetPayload()
	ethSignature, starkSignature, err := signers.CreateSignatures(signableCancelOrder.SignableMessage, signableCancelOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	cancelOrderParams := orders.NewCancelOrderParamsWithContext(ctx)
	cancelOrderParams.SetID(strconv.FormatInt(*(request.OrderID), 10))
	cancelOrderParams.SetCancelOrderRequest(&models.CancelOrderRequest{
		OrderID:        request.OrderID,
		StarkSignature: &starkSignature,
	})

	ethAddress := l1signer.GetAddress()
	cancelOrderParams.SetXImxEthAddress(&ethAddress)
	cancelOrderParams.SetXImxEthSignature(&ethSignature)
	cancelOrderOk, err := api.Orders.CancelOrder(cancelOrderParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `ordersClient.CancelOrder`: %v", err)
	}
	return cancelOrderOk.GetPayload(), nil
}
