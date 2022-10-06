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
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.metadataRefreshesAPI.GetAListOfMetadataRefreshes(ctx)
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

	metadataRefreshes, httpResponse, err := c.metadataRefreshesAPI.GetAListOfMetadataRefreshesExecute(req)
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return metadataRefreshes, nil
}

/*
GetMetadataRefreshErrors Gets a metadata refresh errors for the given refresh id

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
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.metadataRefreshesAPI.GetMetadataRefreshErrors(ctx, refreshID)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	if pageSize != nil {
		req.PageSize(*pageSize)
	}
	if cursor != nil {
		req.Cursor(*cursor)
	}

	metadataRefreshErrorsResponse, httpResponse, err := c.metadataRefreshesAPI.GetMetadataRefreshErrorsExecute(req)
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return metadataRefreshErrorsResponse, nil
}

/*
GetMetadataRefreshResults Gets a metadata refresh results for the given refresh id

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
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.metadataRefreshesAPI.GetMetadataRefreshResults(ctx, refreshID)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	metadataRefreshResponse, httpResponse, err := c.metadataRefreshesAPI.GetMetadataRefreshResultsExecute(req)
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
	timestamp, signature, err := createIMXAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	req := c.metadataRefreshesAPI.RequestAMetadataRefresh(ctx)
	req.XImxEthAddress(l1signer.GetAddress()).
		XImxEthSignature(signature).
		XImxEthTimestamp(timestamp)

	req.CreateMetadataRefreshRequest(*createMetadataRefreshRequest)

	createMetadataRefreshResponse, httpResponse, err := c.metadataRefreshesAPI.RequestAMetadataRefreshExecute(req)
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return createMetadataRefreshResponse, nil
}
