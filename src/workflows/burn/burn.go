package burn

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/transfers"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	transfersWorkflow "immutable.com/imx-core-sdk-golang/workflows/transfers"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

// BurnAddress 0x00 where tokens are sent to be removed from circulation
const BurnAddress = "0x0000000000000000000000000000000000000000"

// Burn permanently removes tokens from circulation.
// To burn an asset, you need to transfer the asset to the burn address.
// We have added burn functions in the SDK for clarity, and these are just wrappers for standard transfer functions to BurnAddress.
// Using the transfer functions with the burn address as a recipient will also work just the same.
func Burn(
	ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request types.GetSignableBurnRequest,
) (*models.CreateTransferResponseV1, error) {
	burnAddress := BurnAddress
	transferRequest := models.GetSignableTransferRequestV1{
		Amount:   &request.Amount,
		Sender:   &request.Sender,
		Token:    request.Token,
		Receiver: &burnAddress,
	}
	return transfersWorkflow.CreateTransfer(ctx, api, l1signer, l2signer, transferRequest)
}

// GetBurn returns the status of the transfer given the transferId.
func GetBurn(ctx context.Context, api *client.ImmutableXAPI, transferId string) (*models.Transfer, error) {
	params := transfers.NewGetTransferParamsWithContext(ctx)
	params.SetID(transferId)
	response, err := api.Transfers.GetTransfer(params)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Transfers.GetTransfer`: %v", err)
	}
	return response.GetPayload(), nil
}
