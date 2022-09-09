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
@param tokenId Either ERC721 token ID or internal IMX ID
@param includeFees optional param.
@return Asset
*/
func (c *Client) GetAsset(ctx context.Context, tokenAddress string, tokenId string, includeFees *bool) (*api.Asset, error) {
	apiGetAssetRequest := c.assetsApi.GetAsset(ctx, tokenAddress, tokenId)
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
	return c.assetsApi.ListAssets(ctx)
}

/*
ListAssets Get a list of assets

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ListAssetsResponse
*/
func (c *Client) ListAssets(req api.ApiListAssetsRequest) (*api.ListAssetsResponse, error) {
	response, httpResponse, err := c.assetsApi.ListAssetsExecute(req)
	if err != nil {
		return nil, fmt.Errorf("error in getting the list of assets: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}
