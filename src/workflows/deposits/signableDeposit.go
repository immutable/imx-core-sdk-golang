package deposits

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/api/client/deposits"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

func GetSignableDeposit(
	ctx context.Context,
	depositsApi deposits.ClientService,
	request *models.GetSignableDepositRequest,
) (*models.GetSignableDepositResponse, error) {
	signableDepositParams := deposits.NewGetSignableDepositParamsWithContext(ctx)
	signableDepositParams.SetGetSignableDepositRequest(request)
	signableDepositOK, err := depositsApi.GetSignableDeposit(signableDepositParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Deposits.GetSignableDeposit`: %v", err)
	}
	return signableDepositOK.GetPayload(), nil
}

func GetSignableDepositRequestForEth(amount, user string) *models.GetSignableDepositRequest {
	return &models.GetSignableDepositRequest{
		Amount: &amount,
		Token: &models.SignableToken{
			Data: map[string]interface{}{
				"decimals": 18,
			},
			Type: string(types.ETHType),
		},
		User: &user,
	}
}

func GetSignableDepositRequestForERC20(amount, tokenAddress, user string, decimals uint8) *models.GetSignableDepositRequest {
	return &models.GetSignableDepositRequest{
		Amount: &amount,
		Token: &models.SignableToken{
			Data: map[string]interface{}{
				"decimals":      decimals,
				"token_address": tokenAddress,
			},
			Type: string(types.ERC20Type),
		},
		User: &user,
	}
}

func GetSignableDepositRequestForERC721(amount, tokenId, tokenAddress, user string) *models.GetSignableDepositRequest {
	return &models.GetSignableDepositRequest{
		Amount: &amount,
		Token: &models.SignableToken{
			Data: map[string]interface{}{
				"token_id":      tokenId,
				"token_address": tokenAddress,
			},
			Type: string(types.ERC721Type),
		},
		User: &user,
	}
}
