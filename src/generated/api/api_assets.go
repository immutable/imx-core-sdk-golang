/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type AssetsApi interface {

	/*
	GetAsset Get details of an asset

	Get details of an asset

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tokenAddress Address of the ERC721 contract
	@param tokenId Either ERC721 token ID or internal IMX ID
	@return ApiGetAssetRequest
	*/
	GetAsset(ctx context.Context, tokenAddress string, tokenId string) ApiGetAssetRequest

	// GetAssetExecute executes the request
	//  @return Asset
	GetAssetExecute(r ApiGetAssetRequest) (*Asset, *http.Response, error)

	/*
	ListAssets Get a list of assets

	Get a list of assets

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListAssetsRequest
	*/
	ListAssets(ctx context.Context) ApiListAssetsRequest

	// ListAssetsExecute executes the request
	//  @return ListAssetsResponse
	ListAssetsExecute(r ApiListAssetsRequest) (*ListAssetsResponse, *http.Response, error)
}

// AssetsApiService AssetsApi service
type AssetsApiService service

type ApiGetAssetRequest struct {
	ctx context.Context
	ApiService AssetsApi
	tokenAddress string
	tokenId string
	includeFees *bool
}

// Set flag to include fees associated with the asset
func (r ApiGetAssetRequest) IncludeFees(includeFees bool) ApiGetAssetRequest {
	r.includeFees = &includeFees
	return r
}

func (r ApiGetAssetRequest) Execute() (*Asset, *http.Response, error) {
	return r.ApiService.GetAssetExecute(r)
}

/*
GetAsset Get details of an asset

Get details of an asset

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tokenAddress Address of the ERC721 contract
 @param tokenId Either ERC721 token ID or internal IMX ID
 @return ApiGetAssetRequest
*/
func (a *AssetsApiService) GetAsset(ctx context.Context, tokenAddress string, tokenId string) ApiGetAssetRequest {
	return ApiGetAssetRequest{
		ApiService: a,
		ctx: ctx,
		tokenAddress: tokenAddress,
		tokenId: tokenId,
	}
}

// Execute executes the request
//  @return Asset
func (a *AssetsApiService) GetAssetExecute(r ApiGetAssetRequest) (*Asset, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Asset
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.GetAsset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/assets/{token_address}/{token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"token_address"+"}", url.PathEscape(parameterToString(r.tokenAddress, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"token_id"+"}", url.PathEscape(parameterToString(r.tokenId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeFees != nil {
		localVarQueryParams.Add("include_fees", parameterToString(*r.includeFees, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListAssetsRequest struct {
	ctx context.Context
	ApiService AssetsApi
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
	user *string
	status *string
	name *string
	metadata *string
	sellOrders *bool
	buyOrders *bool
	includeFees *bool
	collection *string
	updatedMinTimestamp *string
	updatedMaxTimestamp *string
	auxiliaryFeePercentages *string
	auxiliaryFeeRecipients *string
}

// Page size of the result
func (r ApiListAssetsRequest) PageSize(pageSize int32) ApiListAssetsRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListAssetsRequest) Cursor(cursor string) ApiListAssetsRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListAssetsRequest) OrderBy(orderBy string) ApiListAssetsRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListAssetsRequest) Direction(direction string) ApiListAssetsRequest {
	r.direction = &direction
	return r
}

// Ethereum address of the user who owns these assets
func (r ApiListAssetsRequest) User(user string) ApiListAssetsRequest {
	r.user = &user
	return r
}

// Status of these assets
func (r ApiListAssetsRequest) Status(status string) ApiListAssetsRequest {
	r.status = &status
	return r
}

// Name of the asset to search
func (r ApiListAssetsRequest) Name(name string) ApiListAssetsRequest {
	r.name = &name
	return r
}

// JSON-encoded metadata filters for these asset. Example: {&#39;proto&#39;:[&#39;1147&#39;],&#39;quality&#39;:[&#39;Meteorite&#39;]}
func (r ApiListAssetsRequest) Metadata(metadata string) ApiListAssetsRequest {
	r.metadata = &metadata
	return r
}

// Set flag to true to fetch an array of sell order details with accepted status associated with the asset
func (r ApiListAssetsRequest) SellOrders(sellOrders bool) ApiListAssetsRequest {
	r.sellOrders = &sellOrders
	return r
}

// Set flag to true to fetch an array of buy order details  with accepted status associated with the asset
func (r ApiListAssetsRequest) BuyOrders(buyOrders bool) ApiListAssetsRequest {
	r.buyOrders = &buyOrders
	return r
}

// Set flag to include fees associated with the asset
func (r ApiListAssetsRequest) IncludeFees(includeFees bool) ApiListAssetsRequest {
	r.includeFees = &includeFees
	return r
}

// Collection contract address
func (r ApiListAssetsRequest) Collection(collection string) ApiListAssetsRequest {
	r.collection = &collection
	return r
}

// Minimum timestamp for when these assets were last updated, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListAssetsRequest) UpdatedMinTimestamp(updatedMinTimestamp string) ApiListAssetsRequest {
	r.updatedMinTimestamp = &updatedMinTimestamp
	return r
}

// Maximum timestamp for when these assets were last updated, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListAssetsRequest) UpdatedMaxTimestamp(updatedMaxTimestamp string) ApiListAssetsRequest {
	r.updatedMaxTimestamp = &updatedMaxTimestamp
	return r
}

// Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients
func (r ApiListAssetsRequest) AuxiliaryFeePercentages(auxiliaryFeePercentages string) ApiListAssetsRequest {
	r.auxiliaryFeePercentages = &auxiliaryFeePercentages
	return r
}

// Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages
func (r ApiListAssetsRequest) AuxiliaryFeeRecipients(auxiliaryFeeRecipients string) ApiListAssetsRequest {
	r.auxiliaryFeeRecipients = &auxiliaryFeeRecipients
	return r
}

func (r ApiListAssetsRequest) Execute() (*ListAssetsResponse, *http.Response, error) {
	return r.ApiService.ListAssetsExecute(r)
}

/*
ListAssets Get a list of assets

Get a list of assets

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListAssetsRequest
*/
func (a *AssetsApiService) ListAssets(ctx context.Context) ApiListAssetsRequest {
	return ApiListAssetsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListAssetsResponse
func (a *AssetsApiService) ListAssetsExecute(r ApiListAssetsRequest) (*ListAssetsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListAssetsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AssetsApiService.ListAssets")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/assets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.orderBy != nil {
		localVarQueryParams.Add("order_by", parameterToString(*r.orderBy, ""))
	}
	if r.direction != nil {
		localVarQueryParams.Add("direction", parameterToString(*r.direction, ""))
	}
	if r.user != nil {
		localVarQueryParams.Add("user", parameterToString(*r.user, ""))
	}
	if r.status != nil {
		localVarQueryParams.Add("status", parameterToString(*r.status, ""))
	}
	if r.name != nil {
		localVarQueryParams.Add("name", parameterToString(*r.name, ""))
	}
	if r.metadata != nil {
		localVarQueryParams.Add("metadata", parameterToString(*r.metadata, ""))
	}
	if r.sellOrders != nil {
		localVarQueryParams.Add("sell_orders", parameterToString(*r.sellOrders, ""))
	}
	if r.buyOrders != nil {
		localVarQueryParams.Add("buy_orders", parameterToString(*r.buyOrders, ""))
	}
	if r.includeFees != nil {
		localVarQueryParams.Add("include_fees", parameterToString(*r.includeFees, ""))
	}
	if r.collection != nil {
		localVarQueryParams.Add("collection", parameterToString(*r.collection, ""))
	}
	if r.updatedMinTimestamp != nil {
		localVarQueryParams.Add("updated_min_timestamp", parameterToString(*r.updatedMinTimestamp, ""))
	}
	if r.updatedMaxTimestamp != nil {
		localVarQueryParams.Add("updated_max_timestamp", parameterToString(*r.updatedMaxTimestamp, ""))
	}
	if r.auxiliaryFeePercentages != nil {
		localVarQueryParams.Add("auxiliary_fee_percentages", parameterToString(*r.auxiliaryFeePercentages, ""))
	}
	if r.auxiliaryFeeRecipients != nil {
		localVarQueryParams.Add("auxiliary_fee_recipients", parameterToString(*r.auxiliaryFeeRecipients, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
