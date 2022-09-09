package imx

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
GetAsset Get details of an asset

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param tokenAddress Address of the ERC721 contract
@param tokenID Either ERC721 token ID or internal IMX ID
@param includeFees optional param.
@return Asset
*/
func (c *Client) GetAsset(ctx context.Context, tokenAddress, tokenID string, includeFees *bool) (*api.Asset, error) {
	apiGetAssetRequest := c.assetsAPI.GetAsset(ctx, tokenAddress, tokenID)
	if includeFees != nil {
		apiGetAssetRequest.IncludeFees(*includeFees)
	}

	response, httpResponse, err := apiGetAssetRequest.Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the details of an asset: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

/*
NewListAssetsRequest Creates a new ApiListAssetsRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListAssetsRequest
*/
func (c *Client) NewListAssetsRequest(ctx context.Context) api.ApiListAssetsRequest {
	return c.assetsAPI.ListAssets(ctx)
}

/*
ListAssets Get a list of assets

@param req the api request struct with all params populated to make the request
@return ListAssetsResponse
*/
func (c *Client) ListAssets(req *api.ApiListAssetsRequest) (*api.ListAssetsResponse, error) {
	response, httpResponse, err := c.assetsAPI.ListAssetsExecute(*req)
	if err != nil {
		return nil, fmt.Errorf("error in getting the list of assets: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}
