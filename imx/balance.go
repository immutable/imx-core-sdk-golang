package imx

import (
	"context"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
GetBalance Fetches the token balances of the user

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param owner Address of the owner/user
@param tokenAddress Token address
@return Balance
*/
func (c *Client) GetBalance(ctx context.Context, owner, tokenAddress string) (*api.Balance, error) {
	response, httpResponse, err := c.balancesAPI.GetBalance(ctx, owner, tokenAddress).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}

/*
NewListBalancesRequest Creates a new ApiListBalancesRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListBalancesRequest
*/
func (c *Client) NewListBalancesRequest(ctx context.Context, owner string) api.ApiListBalancesRequest {
	return c.balancesAPI.ListBalances(ctx, owner)
}

/*
ListBalances Get a list of balances for given user

@param req the api request struct with all params populated to make the request
@return ListBalancesResponse
*/
func (c *Client) ListBalances(req *api.ApiListBalancesRequest) (*api.ListBalancesResponse, error) {
	response, httpResponse, err := c.balancesAPI.ListBalancesExecute(*req)
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}
