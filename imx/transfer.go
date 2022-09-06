package imx

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

// CreateTransfer transfers the request's models.SignableToken from Sender to Receiver.
func (c *Client) CreateTransfer(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableTransferRequestV1,
) (*api.CreateTransferResponseV1, error) {
	data, httpResponse, err := c.transfersApi.GetSignableTransferV1(ctx).GetSignableTransferRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.GetSignableTransferV1`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&data.SignableMessage, &data.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	response, httpResponse, err := c.transfersApi.CreateTransferV1(ctx).
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

// CreateBatchNftTransfer performs a bulk transfer of NFTs given an array of models.SignableToken and their receivers.
func (c *Client) CreateBatchNftTransfer(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableTransferRequest,
) (*api.CreateTransferResponse, error) {
	data, httpResponse, err := c.transfersApi.GetSignableTransfer(ctx).GetSignableTransferRequestV2(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.GetSignableTransfer`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignatureBytes, err := l1signer.SignMessage(data.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error generating EthSignature from SignableMessage: %v", err)
	}

	ethSignature := hexutil.Encode(ethSignatureBytes)
	ethAddress := l1signer.GetAddress()

	transferRequests, err := getSignedTransferRequests(data.SignableResponses, l2signer)
	if err != nil {
		return nil, err
	}

	response, httpResponse, err := c.transfersApi.CreateTransfer(ctx).
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

// getSignedTransferRequests iterates through signableTransfers, signs each payloadHash and returns an array of pointers to models.TransferRequest.
func getSignedTransferRequests(signableTransfers []api.SignableTransferResponseDetails, l2signer L2Signer) ([]api.TransferRequest, error) {
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

/*
GetTransfer Get details of a transfer with the given ID

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param id Transfer ID
@return Transfer
*/
func (c *Client) GetTransfer(ctx context.Context, id string) (*api.Transfer, error) {
	response, httpResponse, err := c.transfersApi.GetTransfer(ctx, id).Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the details of a transfer: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

/*
ListTransfers Gets a list of transfers

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ListTransfersResponse
*/
func (c *Client) ListTransfers(ctx context.Context) (*api.ListTransfersResponse, error) {
	response, httpResponse, err := c.transfersApi.ListTransfers(ctx).Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the list of transfers: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}
