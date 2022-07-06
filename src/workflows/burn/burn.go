package burn

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/api/client/transfers"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	"immutable.com/imx-core-sdk-golang/workflows/utils"
)

const BurnAddress = "0x0000000000000000000000000000000000000000"

func Burn(
	ctx context.Context,
	api transfers.ClientService,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request types.GetSignableBurnRequest,
) (*models.CreateTransferResponseV1, error) {
	burnAddress := BurnAddress
	getSignableTransferParams := transfers.NewGetSignableTransferV1ParamsWithContext(ctx)
	getSignableTransferParams.SetGetSignableTransferRequest(&models.GetSignableTransferRequestV1{
		Amount:   &request.Amount,
		Sender:   &request.Sender,
		Token:    request.Token,
		Receiver: &burnAddress,
	})
	signableResponse, err := api.GetSignableTransferV1(getSignableTransferParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Api.GetSignableTransferV1`: %v", err)
	}
	data := signableResponse.GetPayload()

	ethSignature, starkSignature, err := utils.CreateSignatures(data.SignableMessage, data.PayloadHash, l1signer, l2signer)
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
	response, err := api.CreateTransferV1(transferRequestParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Api.CreateTransferV1`: %v", err)
	}
	return response.GetPayload(), nil
}

func GetBurn(ctx context.Context, api transfers.ClientService, transferId string) (*models.Transfer, error) {
	params := transfers.NewGetTransferParamsWithContext(ctx)
	params.SetID(transferId)
	response, err := api.GetTransfer(params)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Api.GetTransfer`: %v", err)
	}
	return response.GetPayload(), nil
}
