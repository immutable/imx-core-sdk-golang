package trades

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/trades"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

// CreateTrade submits a matched order to the CreateTrade endpoint.
// https://docs.x.immutable.com/reference#/operations/createTrade
func CreateTrade(
	ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request models.GetSignableTradeRequest,
) (*models.CreateTradeResponse, error) {
	ethAddress := l1signer.GetAddress()

	request.User = &ethAddress
	getSignableTradeParams := trades.NewGetSignableTradeParamsWithContext(ctx)
	getSignableTradeParams.SetGetSignableTradeRequest(&request)
	getSignableTradeOk, err := api.Trades.GetSignableTrade(getSignableTradeParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Trades.GetSignableTrade`: %v", err)
	}
	signableTrade := getSignableTradeOk.GetPayload()

	ethSignature, err := l1signer.SignMessage(*signableTrade.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error generating EthSignature from SignableMessage: %v", err)
	}
	ethSignatureEncodedInHex := hexutil.Encode(ethSignature)

	starkSignature, err := l2signer.SignMessage(*signableTrade.PayloadHash)
	if err != nil {
		return nil, fmt.Errorf("error generating StarkSignature from PayloadHash: %v", err)
	}

	createTradeParams := trades.NewCreateTradeParamsWithContext(ctx)
	createTradeParams.SetCreateTradeRequest(&models.CreateTradeRequestV1{
		AmountBuy:           signableTrade.AmountBuy,
		AmountSell:          signableTrade.AmountSell,
		AssetIDBuy:          signableTrade.AssetIDBuy,
		AssetIDSell:         signableTrade.AssetIDSell,
		ExpirationTimestamp: signableTrade.ExpirationTimestamp,
		FeeInfo:             signableTrade.FeeInfo,
		Fees:                request.Fees,
		IncludeFees:         true,
		Nonce:               signableTrade.Nonce,
		OrderID:             request.OrderID,
		StarkKey:            signableTrade.StarkKey,
		StarkSignature:      &starkSignature,
		VaultIDBuy:          signableTrade.VaultIDBuy,
		VaultIDSell:         signableTrade.VaultIDSell,
	})
	createTradeParams.SetXImxEthAddress(&ethAddress)
	createTradeParams.SetXImxEthSignature(&ethSignatureEncodedInHex)
	tradeResponseOK, err := api.Trades.CreateTrade(createTradeParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Trades.CreateTrade`: %v", err)
	}
	return tradeResponseOK.GetPayload(), nil
}
