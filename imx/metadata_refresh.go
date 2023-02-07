package imx

import (
	"context"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
ListMetadataRefreshes Gets a list of metadata refreshes

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param collectionAddress Collection contract address
@param pageSize The page size of the result
@param cursor The cursor
@return GetMetadataRefreshes
*/
func (c *Client) ListMetadataRefreshes(
	ctx context.Context,
	l1signer L1Signer,
	collectionAddress *string,
	pageSize *int32,
	cursor *string,
) (*api.GetMetadataRefreshes, error) {
	timestamp, signature, err := generateIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.MetadataRefreshesAPI.GetAListOfMetadataRefreshes(ctx)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	if collectionAddress != nil {
		req.CollectionAddress(*collectionAddress)
	}
	if pageSize != nil {
		req.PageSize(*pageSize)
	}
	if cursor != nil {
		req.Cursor(*cursor)
	}

	metadataRefreshes, httpResponse, err := c.MetadataRefreshesAPI.GetAListOfMetadataRefreshesExecute(req)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return metadataRefreshes, nil
}

/*
GetMetadataRefreshErrors Gets metadata refresh errors for the given refresh id

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param collectionAddress Collection contract address
@param pageSize The page size of the result
@param cursor The cursor
@return GetMetadataRefreshes
*/
func (c *Client) GetMetadataRefreshErrors(
	ctx context.Context,
	l1signer L1Signer,
	refreshID string,
	pageSize *int32,
	cursor *string,
) (*api.GetMetadataRefreshErrorsResponse, error) {
	timestamp, signature, err := generateIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.MetadataRefreshesAPI.GetMetadataRefreshErrors(ctx, refreshID)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	if pageSize != nil {
		req.PageSize(*pageSize)
	}
	if cursor != nil {
		req.Cursor(*cursor)
	}

	metadataRefreshErrorsResponse, httpResponse, err := c.MetadataRefreshesAPI.GetMetadataRefreshErrorsExecute(req)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return metadataRefreshErrorsResponse, nil
}

/*
GetMetadataRefreshResults Gets metadata refresh results for the given refresh id

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param collectionAddress Collection contract address
@param pageSize The page size of the result
@param cursor The cursor
@return GetMetadataRefreshes
*/
func (c *Client) GetMetadataRefreshResults(
	ctx context.Context,
	l1signer L1Signer,
	refreshID string,
) (*api.GetMetadataRefreshResponse, error) {
	timestamp, signature, err := generateIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.MetadataRefreshesAPI.GetMetadataRefreshResults(ctx, refreshID)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	metadataRefreshResponse, httpResponse, err := c.MetadataRefreshesAPI.GetMetadataRefreshResultsExecute(req)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return metadataRefreshResponse, nil
}

/*
CreateMetadataRefresh Creates a metadata refresh

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param createMetadataRefreshRequest The request struct with all the params.
@return CreateMetadataRefreshResponse
*/
func (c *Client) CreateMetadataRefresh(
	ctx context.Context,
	l1signer L1Signer,
	createMetadataRefreshRequest *api.CreateMetadataRefreshRequest,
) (*api.CreateMetadataRefreshResponse, error) {
	timestamp, signature, err := generateIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.MetadataRefreshesAPI.RequestAMetadataRefresh(ctx)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	req.CreateMetadataRefreshRequest(*createMetadataRefreshRequest)

	createMetadataRefreshResponse, httpResponse, err := c.MetadataRefreshesAPI.RequestAMetadataRefreshExecute(req)
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return createMetadataRefreshResponse, nil
}
