package imx

import (
	"context"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
CreateOrder will list the given asset for sale on the marketplace.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param l2signer Stark signer to sign the payload hash.
@param request The request struct with all the params.
@return CreateOrderResponse
*/
func (c *Client) CreateOrder(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request *api.GetSignableOrderRequest,
) (*api.CreateOrderResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableOrder, httpResponse, err := c.OrdersAPI.GetSignableOrder(ctx).GetSignableOrderRequestV3(*request).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}

	ethSignature, starkSignature, err := createSignatures(&signableOrder.SignableMessage, &signableOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	includeFees := true
	createOrderResponse, httpResponse, err := c.OrdersAPI.CreateOrder(ctx).
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
		return nil, NewIMXError(httpResponse, err)
	}
	return createOrderResponse, nil
}

/*
CancelOrder will remove the listed asset on marketplace from sale.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param l2signer Stark signer to sign the payload hash.
@param request The request struct with all the params.
@return CancelOrderResponse
*/
func (c *Client) CancelOrder(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableCancelOrderRequest,
) (*api.CancelOrderResponse, error) {
	signableCancelOrder, httpResponse, err := c.OrdersAPI.GetSignableCancelOrder(ctx).GetSignableCancelOrderRequest(request).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}

	ethSignature, starkSignature, err := createSignatures(&signableCancelOrder.SignableMessage, &signableCancelOrder.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	orderID := strconv.FormatInt(int64(request.OrderId), 10)
	cancelOrderResponse, httpResponse, err := c.OrdersAPI.CancelOrder(ctx, orderID).
		CancelOrderRequest(api.CancelOrderRequest{
			OrderId:        request.OrderId,
			StarkSignature: starkSignature,
		}).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return cancelOrderResponse, nil
}

/*
GetOrder Get details of an order with the given ID

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param id Order ID
@return Order
*/
func (c *Client) GetOrder(ctx context.Context, id string) (*api.Order, error) {
	response, httpResponse, err := c.OrdersAPI.GetOrder(ctx, id).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}

/*
NewListOrdersRequest Creates a new ApiListOrdersRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListOrdersRequest
*/
func (c *Client) NewListOrdersRequest(ctx context.Context) api.ApiListOrdersRequest {
	return c.OrdersAPI.ListOrders(ctx)
}

/*
ListOrders Gets a list of orders

@param req the api request struct with all params populated to make the request
@return ListOrdersResponse
*/
func (c *Client) ListOrders(req *api.ApiListOrdersRequest) (*api.ListOrdersResponse, error) {
	response, httpResponse, err := c.OrdersAPI.ListOrdersExecute(*req)
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}
