package deposits

import (
	"context"
	"fmt"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/tokens"
)

func getSignableDeposit(
	ctx context.Context,
	depositsAPI api.DepositsApi,
	request *api.GetSignableDepositRequest) (*api.GetSignableDepositResponse, error) {
	signableDepositResponse, httpResp, err := depositsAPI.GetSignableDeposit(ctx).GetSignableDepositRequest(*request).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Deposits.GetSignableDeposit`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return signableDepositResponse, nil
}

func NewSignableDepositRequestForEth(amount, user string) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  *tokens.NewSignableTokenEth(),
		User:   user,
	}
}

func NewSignableDepositRequestForERC721(amount, tokenID, tokenAddress, user string) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  *tokens.NewSignableTokenERC721(tokenID, tokenAddress),
		User:   user,
	}
}
