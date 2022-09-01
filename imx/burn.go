package imx

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

// BurnAddress 0x00 where tokens are sent to be removed from circulation
const BurnAddress = "0x0000000000000000000000000000000000000000"

type GetSignableBurnRequest struct {
	// Amount of the token to burn
	Amount string
	// Ethereum address of the user
	Sender string
	// Token to burn
	Token api.SignableToken
}

// Burn permanently removes tokens from circulation.
// To burn an asset, you need to transfer the asset to the burn address.
// We have added burn functions in the SDK for clarity, and these are just wrappers for standard transfer functions to BurnAddress.
// Using the transfer functions with the burn address as a recipient will also work just the same.
func (c *Client) Burn(
	ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	request GetSignableBurnRequest) (*api.CreateTransferResponseV1, error) {
	burnAddress := BurnAddress
	transferRequest := api.GetSignableTransferRequestV1{
		Amount:   request.Amount,
		Sender:   request.Sender,
		Token:    request.Token,
		Receiver: burnAddress,
	}
	return c.CreateTransfer(ctx, l1signer, l2signer, transferRequest)
}

// GetBurn returns the status of the transfer given the transferID.
func (c *Client) GetBurn(ctx context.Context, transferID string) (*api.Transfer, error) {
	response, httpResponse, err := c.TransfersApi.GetTransfer(ctx, transferID).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `TransfersApi.GetTransfer`: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}
