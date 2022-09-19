package imx

import (
	"context"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
CreateCollection Creates a new collection

A collection refers to a series of NFTs, minted under a project, and corresponds to a specific deployed smart contract.
All minted assets belong to a collection, and in order to mint assets on L2 you must first register your collection (smart contract) with Immutable X.
* Each collection belongs to a project.
* Each collection may contain many similar or different NFTs.

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

	createCollectionResponse, httpResponse, err := c.collectionsAPI.CreateCollection(ctx).
		CreateCollectionRequest(*createCollectionRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return createCollectionResponse, nil
}

/*
UpdateCollection Updates an existing collection

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

	createCollectionResponse, httpResponse, err := c.collectionsAPI.UpdateCollection(ctx, contractAddress).
		UpdateCollectionRequest(*updateCollectionRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
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
	response, httpResponse, err := c.collectionsAPI.GetCollection(ctx, collectionContractAddress).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
NewListCollectionsRequest Creates a new ApiListCollectionsRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListCollectionsRequest
*/
func (c *Client) NewListCollectionsRequest(ctx context.Context) api.ApiListCollectionsRequest {
	return c.collectionsAPI.ListCollections(ctx)
}

/*
ListCollections Get a list of collections

@param req the api request struct with all params populated to make the request
@return ListCollectionsResponse
*/
func (c *Client) ListCollections(req *api.ApiListCollectionsRequest) (*api.ListCollectionsResponse, error) {
	response, httpResponse, err := c.collectionsAPI.ListCollectionsExecute(*req)
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
NewListCollectionFiltersRequest Creates a new ApiListCollectionFiltersRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param collectionContractAddress Collection contract address
@return ApiListCollectionFiltersRequest
*/
func (c *Client) NewListCollectionFiltersRequest(ctx context.Context, collectionContractAddress string) api.ApiListCollectionFiltersRequest {
	return c.collectionsAPI.ListCollectionFilters(ctx, collectionContractAddress)
}

/*
ListCollectionFilters Get a list of collection filters

@param req the api request struct with all params populated to make the request
@return CollectionFilter
*/
func (c *Client) ListCollectionFilters(req *api.ApiListCollectionFiltersRequest) (*api.CollectionFilter, error) {
	response, httpResponse, err := c.collectionsAPI.ListCollectionFiltersExecute(*req)
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}
