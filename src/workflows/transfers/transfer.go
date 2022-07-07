package transfers

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/transfers"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

// CreateTransfer transfers the given models.SignableToken from Sender to Receiver.
func CreateTransfer(
	ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request models.GetSignableTransferRequestV1,
) (*models.CreateTransferResponseV1, error) {
	getSignableTransferParams := transfers.NewGetSignableTransferV1ParamsWithContext(ctx)
	getSignableTransferParams.SetGetSignableTransferRequest(&request)
	signableResponse, err := api.Transfers.GetSignableTransferV1(getSignableTransferParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Transfers.GetSignableTransferV1`: %v", err)
	}
	data := signableResponse.GetPayload()

	ethSignature, starkSignature, err := signers.CreateSignatures(data.SignableMessage, data.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	transferRequestParams := transfers.NewCreateTransferV1ParamsWithContext(ctx)
	transferRequestParams.SetCreateTransferRequest(&models.CreateTransferRequestV1{
		Amount:              data.Amount,
		AssetID:             data.AssetID,
		ExpirationTimestamp: data.ExpirationTimestamp,
		Nonce:               data.Nonce,
		ReceiverStarkKey:    data.ReceiverStarkKey,
		ReceiverVaultID:     data.ReceiverVaultID,
		SenderStarkKey:      &data.SenderStarkKey,
		SenderVaultID:       data.SenderVaultID,
		StarkSignature:      &starkSignature,
	})
	transferRequestParams.SetXImxEthAddress(&ethAddress)
	transferRequestParams.SetXImxEthSignature(&ethSignature)
	response, err := api.Transfers.CreateTransferV1(transferRequestParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Transfers.CreateTransferV1`: %v", err)
	}
	return response.GetPayload(), nil
}
