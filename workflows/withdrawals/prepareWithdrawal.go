package withdrawals

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/generated/api"
	internalSigners "github.com/immutable/imx-core-sdk-golang/internal/signers"
	"github.com/immutable/imx-core-sdk-golang/signers"
)

// PrepareWithdrawal submits a withdrawal request for ETH, ERC20 and ERC721 tokens to be included in the generation and submission of the next batch.
// Upon batch confirmation (on-chain state update), the asset is available to be withdrawn by the initial owner / originator of the asset.
func PrepareWithdrawal(
	ctx context.Context,
	withdrawalsAPI api.WithdrawalsApi,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	request api.GetSignableWithdrawalRequest,
) (*api.CreateWithdrawalResponse, error) {
	ethAddress := l1signer.GetAddress()
	request.User = ethAddress
	signableResponse, httpResp, err := withdrawalsAPI.GetSignableWithdrawal(ctx).GetSignableWithdrawalRequest(request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Withdrawals.GetSignableWithdrawal`: %v, HTTP response body: %v", err, httpResp.Body)
	}

	ethSignature, starkSignature, err := internalSigners.CreateSignatures(&signableResponse.SignableMessage, &signableResponse.PayloadHash, l1signer, l2signer)
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
