package trades

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
)

// CreateTrade submits a matched order to the CreateTrade endpoint.
// https://docs.x.immutable.com/immutable/reference#/operations/createTrade
func CreateTrade(
	ctx context.Context,
	apiClient *api.APIClient,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableTradeRequest,
) (*api.CreateTradeResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableTrade, httpResponse, err := apiClient.TradesApi.GetSignableTrade(ctx).GetSignableTradeRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TradesApi.GetSignableTrade`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, err := l1signer.SignMessage(signableTrade.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error generating EthSignature from SignableMessage: %v", err)
	}
	ethSignatureEncodedInHex := hexutil.Encode(ethSignature)

	starkSignature, err := l2signer.SignMessage(signableTrade.PayloadHash)
	if err != nil {
		return nil, fmt.Errorf("error generating StarkSignature from PayloadHash: %v", err)
	}

	includeFees := true
	createTradeResponse, httpResponse, err := apiClient.TradesApi.CreateTrade(ctx).
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
		XImxEthAddress(ethAddress).XImxEthSignature(ethSignatureEncodedInHex).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TradesApi.CreateTrade`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return createTradeResponse, nil
}
