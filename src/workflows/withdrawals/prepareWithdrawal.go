package withdrawals

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/withdrawals"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
)

// PrepareWithdrawal submits a withdrawal request to be included in the generation and submission of the next batch.
// Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner / originator of the asset.
func PrepareWithdrawal(
	ctx context.Context,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request models.GetSignableWithdrawalRequest,
) (*models.CreateWithdrawalResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = &ethAddress
	getSignableWithdrawalParams := withdrawals.NewGetSignableWithdrawalParamsWithContext(ctx)
	getSignableWithdrawalParams.SetGetSignableWithdrawalRequest(&request)
	signableResponse, err := api.Withdrawals.GetSignableWithdrawal(getSignableWithdrawalParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Withdrawals.GetSignableWithdrawal`: %v", err)
	}
	data := signableResponse.GetPayload()
	ethSignature, starkSignature, err := signers.CreateSignatures(data.SignableMessage, data.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	withdrawalParams := withdrawals.NewCreateWithdrawalParamsWithContext(ctx)
	withdrawalParams.SetCreateWithdrawalRequest(&models.CreateWithdrawalRequest{
		Amount:         request.Amount,
		AssetID:        data.AssetID,
		Nonce:          data.Nonce,
		StarkKey:       data.StarkKey,
		StarkSignature: &starkSignature,
		VaultID:        data.VaultID,
	})
	withdrawalParams.SetXImxEthAddress(&ethAddress)
	withdrawalParams.SetXImxEthSignature(&ethSignature)
	response, err := api.Withdrawals.CreateWithdrawal(withdrawalParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Withdrawals.CreateWithdrawal`: %v", err)
	}
	return response.GetPayload(), nil
}
