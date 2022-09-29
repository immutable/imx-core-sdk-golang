package imx

import (
	"context"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
AddMetadataSchemaToCollection Add metadata schema to collection

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param contractAddress Collection contract address
@param addMetadataSchemaToCollectionRequest The request struct with all the params.
@return SuccessResponse
*/
func (c *Client) AddMetadataSchemaToCollection(
	ctx context.Context,
	l1signer L1Signer,
	contractAddress string,
	addMetadataSchemaToCollectionRequest api.AddMetadataSchemaToCollectionRequest,
) (*api.SuccessResponse, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	successResponse, httpResponse, err := c.metadataAPI.AddMetadataSchemaToCollection(ctx, contractAddress).
		AddMetadataSchemaToCollectionRequest(addMetadataSchemaToCollectionRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return successResponse, nil
}

/*
UpdateMetadataSchemaByName Update metadata schema by name

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1signer Ethereum signer used for ownership authentication.
@param contractAddress Collection contract address
@param metadataSchemaName Metadata schema name
@param metadataSchemaRequest The request struct with all the params.
@return SuccessResponse
*/
func (c *Client) UpdateMetadataSchemaByName(
	ctx context.Context,
	l1signer L1Signer,
	contractAddress, metadataSchemaName string,
	metadataSchemaRequest api.MetadataSchemaRequest,
) (*api.SuccessResponse, error) {
	timestamp, signature, err := getProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		return nil, err
	}

	successResponse, httpResponse, err := c.metadataAPI.UpdateMetadataSchemaByName(ctx, contractAddress, metadataSchemaName).
		MetadataSchemaRequest(metadataSchemaRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return successResponse, nil
}
