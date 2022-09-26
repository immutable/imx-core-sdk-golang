package imx

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
Transfer transfers the request's models.SignableToken from Sender to Receiver.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param l2signer Stark signer to sign the payload hash.
@param request The request struct with all the params.
@return CreateTransferResponseV1
*/
func (c *Client) Transfer(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableTransferRequestV1,
) (*api.CreateTransferResponseV1, error) {
	data, httpResponse, err := c.transfersAPI.GetSignableTransferV1(ctx).GetSignableTransferRequest(request).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}

	ethSignature, starkSignature, err := createSignatures(&data.SignableMessage, &data.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	ethAddress := l1signer.GetAddress()
	response, httpResponse, err := c.transfersAPI.CreateTransferV1(ctx).
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
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
BatchNftTransfer performs a bulk transfer of NFTs given an array of models.SignableToken and their receivers.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param l2signer Stark signer to sign the payload hash.
@param request The request struct with all the params.
@return CreateTransferResponse
*/
func (c *Client) BatchNftTransfer(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request api.GetSignableTransferRequest,
) (*api.CreateTransferResponse, error) {
	data, httpResponse, err := c.transfersAPI.GetSignableTransfer(ctx).GetSignableTransferRequestV2(request).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
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

	response, httpResponse, err := c.transfersAPI.CreateTransfer(ctx).
		CreateTransferRequestV2(api.CreateTransferRequest{
			Requests:       transferRequests,
			SenderStarkKey: data.SenderStarkKey,
		}).
		XImxEthAddress(ethAddress).
		XImxEthSignature(ethSignature).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
GetTransfer Get details of a transfer with the given ID

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param id Transfer ID
@return Transfer
*/
func (c *Client) GetTransfer(ctx context.Context, id string) (*api.Transfer, error) {
	response, httpResponse, err := c.transfersAPI.GetTransfer(ctx, id).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
NewListTransfersRequest Creates a new ApiListTransfersRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListTransfersRequest
*/
func (c *Client) NewListTransfersRequest(ctx context.Context) api.ApiListTransfersRequest {
	return c.transfersAPI.ListTransfers(ctx)
}

/*
ListTransfers Gets a list of transfers

@param req the api request struct with all params populated to make the request
@return ListTransfersResponse
*/
func (c *Client) ListTransfers(req *api.ApiListTransfersRequest) (*api.ListTransfersResponse, error) {
	response, httpResponse, err := c.transfersAPI.ListTransfersExecute(*req)
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

// getSignedTransferRequests iterates through signableTransfers, signs each payloadHash and returns an array of pointers to models.TransferRequest.
func getSignedTransferRequests(
	signableTransfers []api.SignableTransferResponseDetails,
	l2signer L2Signer,
) ([]api.TransferRequest, error) {
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
