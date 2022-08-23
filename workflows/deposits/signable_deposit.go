package deposits

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/generated/api"
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
