package deposits

import (
	"context"
	"fmt"

	"immutable.com/imx-core-sdk-golang/api/client/deposits"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/workflows/utils"
)

func GetSignableDeposit(
	ctx context.Context,
	depositsAPI deposits.ClientService,
	request *models.GetSignableDepositRequest,
) (*models.GetSignableDepositResponse, error) {
	signableDepositParams := deposits.NewGetSignableDepositParamsWithContext(ctx)
	signableDepositParams.SetGetSignableDepositRequest(request)
	signableDepositOK, err := depositsAPI.GetSignableDeposit(signableDepositParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Deposits.GetSignableDeposit`: %v", err)
	}
	return signableDepositOK.GetPayload(), nil
}

func NewSignableDepositRequestForEth(amount, user string) *models.GetSignableDepositRequest {
	return &models.GetSignableDepositRequest{
		Amount: &amount,
		Token:  utils.NewSignableTokenEth(),
		User:   &user,
	}
}

func NewSignableDepositRequestForERC20(amount, tokenAddress, user, decimals string) *models.GetSignableDepositRequest {
	return &models.GetSignableDepositRequest{
		Amount: &amount,
		Token:  utils.NewSignableTokenERC20(decimals, tokenAddress),
		User:   &user,
	}
}

func NewSignableDepositRequestForERC721(amount, tokenID, tokenAddress, user string) *models.GetSignableDepositRequest {
	return &models.GetSignableDepositRequest{
		Amount: &amount,
		Token:  utils.NewSignableTokenERC721(tokenID, tokenAddress),
		User:   &user,
	}
}
