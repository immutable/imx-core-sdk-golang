package withdrawals

import (
	"context"
	"fmt"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
)

// PrepareERC20Withdrawal submits a withdrawal request for ERC20 tokens to be included in the generation and submission of the next batch.
// Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner / originator of the asset.
//
// Note: The ERC20 Amount value supplied along with GetSignableWithdrawalRequest should be the same value as you would use on IMX Marketplace UI.
// Any Conversions required are done by SDK.
func PrepareERC20Withdrawal(
	ctx context.Context,
	withdrawalsAPI api.WithdrawalsApi,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableWithdrawalRequest) (*api.CreateWithdrawalResponse, error) {
	// Convert ERC20 tokens amount to unquantized value.
	unquantized, err := utils.ToUnquantized(request.Amount, request.Token.Data["decimals"].(int))
	if err != nil {
		return nil, err
	}
	request.Amount = unquantized.String()

	return prepareWithdrawalHelper(ctx, withdrawalsAPI, l1signer, l2signer, request)
}

// PrepareEthWithdrawal submits a withdrawal request for Eth tokens to be included in the generation and submission of the next batch.
// Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner / originator of the asset.
//
// Note: The Eth Amount value supplied along with GetSignableWithdrawalRequest should be the same value as you would use on IMX Marketplace UI.
// Any Conversions required are done by SDK.
func PrepareEthWithdrawal(
	ctx context.Context,
	withdrawalsAPI api.WithdrawalsApi,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableWithdrawalRequest) (*api.CreateWithdrawalResponse, error) {
	// Convert Eth Amount to unquantized value.
	amountInt, err := utils.ToUnquantized(request.Amount, utils.EtherDecimals)
	if err != nil {
		return nil, err
	}
	request.Amount = amountInt.String()

	return prepareWithdrawalHelper(ctx, withdrawalsAPI, l1signer, l2signer, request)
}

// PrepareERC721Withdrawal submits a withdrawal request for ERC721 tokens to be included in the generation and submission of the next batch.
// Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner / originator of the asset.
func PrepareERC721Withdrawal(
	ctx context.Context,
	withdrawalsAPI api.WithdrawalsApi,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableWithdrawalRequest) (*api.CreateWithdrawalResponse, error) {
	// No amount conversion required for ERC721 tokens as the value will be always 1 for NFTs.
	return prepareWithdrawalHelper(ctx, withdrawalsAPI, l1signer, l2signer, request)
}

// prepareWithdrawalHelper submits a withdrawal request to be included in the generation and submission of the next batch.
// Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner / originator of the asset.
func prepareWithdrawalHelper(
	ctx context.Context,
	withdrawalsAPI api.WithdrawalsApi,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableWithdrawalRequest) (*api.CreateWithdrawalResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableResponse, httpResp, err := withdrawalsAPI.GetSignableWithdrawal(ctx).GetSignableWithdrawalRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Withdrawals.GetSignableWithdrawal`: %v, HTTP response body: %v", err, httpResp.Body)
	}

	ethSignature, starkSignature, err := signers.CreateSignatures(&signableResponse.SignableMessage, &signableResponse.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	withdrawalRequest := api.CreateWithdrawalRequest{
		Amount:         request.Amount,
		AssetId:        signableResponse.AssetId,
		Nonce:          signableResponse.Nonce,
		StarkKey:       signableResponse.StarkKey,
		StarkSignature: starkSignature,
		VaultId:        signableResponse.VaultId,
	}
	apiCreateWithdrawalRequest := withdrawalsAPI.CreateWithdrawal(ctx).XImxEthAddress(ethAddress).XImxEthSignature(ethSignature)
	withdrawalResponse, httpResp, err := apiCreateWithdrawalRequest.CreateWithdrawalRequest(withdrawalRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `apiCreateWithdrawalRequest.CreateWithdrawalRequest`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return withdrawalResponse, nil
}
