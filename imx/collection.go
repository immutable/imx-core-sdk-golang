package imx

import (
	"context"
	"fmt"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
CreateCollection Create collection

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param createCollectionRequest The request struct with all the params.
@return Collection
*/
func (c *Client) CreateCollection(
	ctx context.Context,
	l1signer L1Signer,
	createCollectionRequest *api.CreateCollectionRequest,
) (*api.Collection, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	createCollectionResponse, httpResp, err := c.collectionsApi.CreateCollection(ctx).
		CreateCollectionRequest(*createCollectionRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `CreateCollection`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return createCollectionResponse, nil
}

/*
CreateCollection Create collection

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param contractAddress Collection contract address
@param updateCollectionRequest The request struct with all the params.
@return Collection
*/
func (c *Client) UpdateCollection(
	ctx context.Context,
	l1signer L1Signer,
	contractAddress string,
	updateCollectionRequest *api.UpdateCollectionRequest,
) (*api.Collection, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	createCollectionResponse, httpResp, err := c.collectionsApi.UpdateCollection(ctx, contractAddress).
		UpdateCollectionRequest(*updateCollectionRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `UpdateCollection`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return createCollectionResponse, nil
}

/*
GetCollection Get details of a collection at the given address

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param collectionContractAddress Collection contract address
@return Collection
*/
func (c *Client) GetCollection(ctx context.Context, collectionContractAddress string) (*api.Collection, error) {
	response, httpResponse, err := c.collectionsApi.GetCollection(ctx, collectionContractAddress).Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the details of a collection: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

/*
NewListBalancesRequest Creates a new ApiListBalancesRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListBalancesRequest
*/
func (c *Client) NewListCollectionsRequest(ctx context.Context, owner string) api.ApiListCollectionsRequest {
	return c.collectionsApi.ListCollections(ctx)
}

/*
ListBalances Get a list of balances for given user

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ListCollectionsResponse
*/
func (c *Client) ListCollections(req api.ApiListCollectionsRequest) (*api.ListCollectionsResponse, error) {
	response, httpResponse, err := c.collectionsApi.ListCollectionsExecute(req)
	if err != nil {
		return nil, fmt.Errorf("error in getting the list of collections: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

/*
NewListBalancesRequest Creates a new ApiListBalancesRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param collectionContractAddress Collection contract address
@return ApiListBalancesRequest
*/
func (c *Client) NewListCollectionFiltersRequest(ctx context.Context, collectionContractAddress string) api.ApiListCollectionFiltersRequest {
	return c.collectionsApi.ListCollectionFilters(ctx, collectionContractAddress)
}

/*
ListBalances Get a list of balances for given user

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return CollectionFilter
*/
func (c *Client) ListCollectionFilters(req api.ApiListCollectionFiltersRequest) (*api.CollectionFilter, error) {
	response, httpResponse, err := c.collectionsApi.ListCollectionFiltersExecute(req)
	if err != nil {
		return nil, fmt.Errorf("error in getting the list of collection Filters: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}
