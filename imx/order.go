package imx

import (
	"context"
	"fmt"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

// CreateOrder will list the given asset for sale on the marketplace.
func (c *Client) CreateOrder(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request *api.GetSignableOrderRequest) (*api.CreateOrderResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableOrder, httpResponse, err := c.GetSignableOrder(ctx).GetSignableOrderRequestV3(*request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.GetSignableOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableOrder.SignableMessage, &signableOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	includeFees := true
	createOrderResponse, httpResponse, err := c.OrdersApi.CreateOrder(ctx).
		CreateOrderRequest(api.CreateOrderRequest{
			AmountBuy: signableOrder.AmountBuy, // The amount (listing price) should be in Wei for Eth tokens,
			// see https://docs.starkware.co/starkex-v4/starkex-deep-dive/starkex-specific-concepts and https://eth-converter.com/
			AmountSell:          signableOrder.AmountSell,  // Quantity to list for sale
			AssetIdBuy:          signableOrder.AssetIdBuy,  // Token with which this asset can be bought with. For example Eth token can be used to buy this asset.
			AssetIdSell:         signableOrder.AssetIdSell, // Token of the asset being listed for sale. For example, ERC720 is a NFT Token and ERC20 is a fungible token.
			ExpirationTimestamp: signableOrder.ExpirationTimestamp,
			Fees:                request.Fees,
			IncludeFees:         &includeFees,
			Nonce:               signableOrder.Nonce,
			StarkKey:            signableOrder.StarkKey,
			StarkSignature:      starkSignature,
			VaultIdBuy:          signableOrder.VaultIdBuy,
			VaultIdSell:         signableOrder.VaultIdSell,
		}).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.CreateOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return createOrderResponse, nil
}

// CancelOrder will remove the listed asset on marketplace from sale.
func (c *Client) CancelOrder(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableCancelOrderRequest,
) (*api.CancelOrderResponse, error) {
	signableCancelOrder, httpResponse, err := c.GetSignableCancelOrder(ctx).GetSignableCancelOrderRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.GetSignableCancelOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableCancelOrder.SignableMessage, &signableCancelOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	orderID := strconv.FormatInt(int64(request.OrderId), 10)
	cancelOrderResponse, httpResponse, err := c.OrdersApi.CancelOrder(ctx, orderID).
		CancelOrderRequest(api.CancelOrderRequest{
			OrderId:        request.OrderId,
			StarkSignature: starkSignature,
		}).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `OrdersApi.CancelOrder`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return cancelOrderResponse, nil
}