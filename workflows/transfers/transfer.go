package transfers

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/generated/api"
	internalSigners "github.com/immutable/imx-core-sdk-golang/internal/signers"
	"github.com/immutable/imx-core-sdk-golang/signers"
)

// CreateTransfer transfers the request's models.SignableToken from Sender to Receiver.
func CreateTransfer(
	ctx context.Context,
	apiClient *api.APIClient,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableTransferRequestV1,
) (*api.CreateTransferResponseV1, error) {
	data, httpResponse, err := apiClient.TransfersApi.GetSignableTransferV1(ctx).GetSignableTransferRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.GetSignableTransferV1`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := internalSigners.CreateSignatures(&data.SignableMessage, &data.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	response, httpResponse, err := apiClient.TransfersApi.CreateTransferV1(ctx).
		CreateTransferRequest(api.CreateTransferRequestV1{
			Amount:              data.Amount,
			AssetId:             data.AssetId,
			ExpirationTimestamp: data.ExpirationTimestamp,
			Nonce:               data.Nonce,
			ReceiverStarkKey:    data.ReceiverStarkKey,
			ReceiverVaultId:     data.ReceiverVaultId,
			SenderStarkKey:      *data.SenderStarkKey,
			SenderVaultId:       data.SenderVaultId,
			StarkSignature:      starkSignature,
		}).
		XImxEthAddress(ethAddress).
		XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.CreateTransferV1`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}
