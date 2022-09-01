package imx

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

// CreateTrade submits a matched order to the CreateTrade endpoint.
// https://docs.x.immutable.com/reference#/operations/createTrade
func (c *Client) CreateTrade(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableTradeRequest,
) (*api.CreateTradeResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableTrade, httpResponse, err := c.GetSignableTrade(ctx).GetSignableTradeRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TradesApi.GetSignableTrade`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableTrade.SignableMessage, &signableTrade.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	includeFees := true
	createTradeResponse, httpResponse, err := c.TradesApi.CreateTrade(ctx).
		CreateTradeRequest(api.CreateTradeRequestV1{
			AmountBuy:           signableTrade.AmountBuy,
			AmountSell:          signableTrade.AmountSell,
			AssetIdBuy:          signableTrade.AssetIdBuy,
			AssetIdSell:         signableTrade.AssetIdSell,
			ExpirationTimestamp: signableTrade.ExpirationTimestamp,
			FeeInfo:             signableTrade.FeeInfo,
			Fees:                request.Fees,
			IncludeFees:         &includeFees,
			Nonce:               signableTrade.Nonce,
			OrderId:             request.OrderId,
			StarkKey:            signableTrade.StarkKey,
			StarkSignature:      starkSignature,
			VaultIdBuy:          signableTrade.VaultIdBuy,
			VaultIdSell:         signableTrade.VaultIdSell,
		}).
		XImxEthAddress(ethAddress).XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TradesApi.CreateTrade`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return createTradeResponse, nil
}
