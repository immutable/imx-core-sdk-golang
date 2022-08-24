package transfers

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
)

// CreateBatchTransfer performs a bulk transfer given an array of models.SignableToken and their receivers.
// CreateBatchTransfer currently only accepts ERC721 tokens.
func CreateBatchTransfer(
	ctx context.Context,
	apiClient *api.APIClient,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableTransferRequest,
) (*api.CreateTransferResponse, error) {
	data, httpResponse, err := apiClient.TransfersApi.GetSignableTransfer(ctx).
		GetSignableTransferRequestV2(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.GetSignableTransfer`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignatureBytes, err := l1signer.SignMessage(data.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error generating EthSignature from SignableMessage: %v", err)
	}

	ethSignature := hexutil.Encode(ethSignatureBytes)
	ethAddress := l1signer.GetAddress()

	transferRequests, err := GetSignedTransferRequests(data.SignableResponses, l2signer)
	if err != nil {
		return nil, err
	}

	response, httpResponse, err := apiClient.TransfersApi.CreateTransfer(ctx).
		CreateTransferRequestV2(api.CreateTransferRequest{
			Requests:       transferRequests,
			SenderStarkKey: data.SenderStarkKey,
		}).
		XImxEthAddress(ethAddress).
		XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.CreateTransfer`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

// GetSignedTransferRequests iterates through signableTransfers, signs each payloadHash and returns an array of pointers to models.TransferRequest.
func GetSignedTransferRequests(signableTransfers []api.SignableTransferResponseDetails, l2signer signers.L2Signer) ([]api.TransferRequest, error) {
	mapped := make([]api.TransferRequest, len(signableTransfers))
	for i, transfer := range signableTransfers {
		starkSignature, err := l2signer.SignMessage(transfer.PayloadHash)
		if err != nil {
			return nil, fmt.Errorf("error generating StarkSignature from PayloadHash %s for AssetID %s: %v", transfer.PayloadHash, transfer.AssetId, err)
		}
		mapped[i] = api.TransferRequest{
			Amount:              transfer.Amount,
			AssetId:             transfer.AssetId,
			ExpirationTimestamp: transfer.ExpirationTimestamp,
			Nonce:               transfer.Nonce,
			ReceiverStarkKey:    transfer.ReceiverStarkKey,
			ReceiverVaultId:     transfer.ReceiverVaultId,
			SenderVaultId:       transfer.SenderVaultId,
			StarkSignature:      starkSignature,
		}
	}
	return mapped, nil
}
