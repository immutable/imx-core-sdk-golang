package transfers

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/transfers"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

// CreateBatchTransfer performs a bulk transfer given an array of models.SignableToken and their receivers.
// CreateBatchTransfer currently only accepts ERC721 tokens.
func CreateBatchTransfer(
	ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request models.GetSignableTransferRequest,
) (*models.CreateTransferResponse, error) {
	getSignableTransferParams := transfers.NewGetSignableTransferParamsWithContext(ctx)
	getSignableTransferParams.SetGetSignableTransferRequestV2(&request)
	signableResponse, err := api.Transfers.GetSignableTransfer(getSignableTransferParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Transfers.GetSignableTransfer`: %v", err)
	}
	data := signableResponse.GetPayload()

	ethSignatureBytes, err := l1signer.SignMessage(*data.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error generating EthSignature from SignableMessage: %v", err)
	}

	ethSignature := hexutil.Encode(ethSignatureBytes)
	ethAddress := l1signer.GetAddress()

	transferRequests, err := GetSignedTransferRequests(data.SignableResponses, l2signer)
	if err != nil {
		return nil, err
	}

	transferRequestParams := transfers.NewCreateTransferParamsWithContext(ctx)
	transferRequestParams.SetCreateTransferRequestV2(&models.CreateTransferRequest{
		Requests:       transferRequests,
		SenderStarkKey: data.SenderStarkKey,
	})
	transferRequestParams.SetXImxEthAddress(&ethAddress)
	transferRequestParams.SetXImxEthSignature(&ethSignature)
	response, err := api.Transfers.CreateTransfer(transferRequestParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Transfers.CreateTransfer`: %v", err)
	}
	return response.GetPayload(), nil
}

// GetSignedTransferRequests iterates through signableTransfers, signs each payloadHash and returns an array of pointers to models.TransferRequest.
func GetSignedTransferRequests(signableTransfers []*models.SignableTransferResponseDetails, l2signer signers.L2Signer) ([]*models.TransferRequest, error) {
	mapped := make([]*models.TransferRequest, len(signableTransfers))
	for i, transfer := range signableTransfers {
		starkSignature, err := l2signer.SignMessage(*transfer.PayloadHash)
		if err != nil {
			return nil, fmt.Errorf("error generating StarkSignature from PayloadHash %s for AssetID %s: %v", *transfer.PayloadHash, *transfer.AssetID, err)
		}
		mapped[i] = &models.TransferRequest{
			Amount:              transfer.Amount,
			AssetID:             transfer.AssetID,
			ExpirationTimestamp: transfer.ExpirationTimestamp,
			Nonce:               transfer.Nonce,
			ReceiverStarkKey:    transfer.ReceiverStarkKey,
			ReceiverVaultID:     transfer.ReceiverVaultID,
			SenderVaultID:       transfer.SenderVaultID,
			StarkSignature:      &starkSignature,
		}
	}
	return mapped, nil
}
