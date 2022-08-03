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


type WithdrawalsApi interface {

	/*
	CreateWithdrawal Creates a withdrawal of a token

	Creates a withdrawal

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateWithdrawalRequest
	*/
	CreateWithdrawal(ctx context.Context) ApiCreateWithdrawalRequest

	// CreateWithdrawalExecute executes the request
	//  @return CreateWithdrawalResponse
	CreateWithdrawalExecute(r ApiCreateWithdrawalRequest) (*CreateWithdrawalResponse, *http.Response, error)

	/*
	GetSignableWithdrawal Gets details of a signable withdrawal

	Gets details of a signable withdrawal

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSignableWithdrawalRequest
	*/
	GetSignableWithdrawal(ctx context.Context) ApiGetSignableWithdrawalRequest

	// GetSignableWithdrawalExecute executes the request
	//  @return GetSignableWithdrawalResponse
	GetSignableWithdrawalExecute(r ApiGetSignableWithdrawalRequest) (*GetSignableWithdrawalResponse, *http.Response, error)

	/*
	GetWithdrawal Gets details of withdrawal with the given ID

	Gets details of withdrawal with the given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Withdrawal ID
	@return ApiGetWithdrawalRequest
	*/
	GetWithdrawal(ctx context.Context, id string) ApiGetWithdrawalRequest

	// GetWithdrawalExecute executes the request
	//  @return Withdrawal
	GetWithdrawalExecute(r ApiGetWithdrawalRequest) (*Withdrawal, *http.Response, error)

	/*
	ListWithdrawals Get a list of withdrawals

	Get a list of withdrawals

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListWithdrawalsRequest
	*/
	ListWithdrawals(ctx context.Context) ApiListWithdrawalsRequest

	// ListWithdrawalsExecute executes the request
	//  @return ListWithdrawalsResponse
	ListWithdrawalsExecute(r ApiListWithdrawalsRequest) (*ListWithdrawalsResponse, *http.Response, error)
}

// WithdrawalsApiService WithdrawalsApi service
type WithdrawalsApiService service

type ApiCreateWithdrawalRequest struct {
	ctx context.Context
	ApiService WithdrawalsApi
	xImxEthAddress *string
	xImxEthSignature *string
	createWithdrawalRequest *CreateWithdrawalRequest
}

// eth address
func (r ApiCreateWithdrawalRequest) XImxEthAddress(xImxEthAddress string) ApiCreateWithdrawalRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// eth signature
func (r ApiCreateWithdrawalRequest) XImxEthSignature(xImxEthSignature string) ApiCreateWithdrawalRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// create a withdrawal
func (r ApiCreateWithdrawalRequest) CreateWithdrawalRequest(createWithdrawalRequest CreateWithdrawalRequest) ApiCreateWithdrawalRequest {
	r.createWithdrawalRequest = &createWithdrawalRequest
	return r
}

func (r ApiCreateWithdrawalRequest) Execute() (*CreateWithdrawalResponse, *http.Response, error) {
	return r.ApiService.CreateWithdrawalExecute(r)
}

/*
CreateWithdrawal Creates a withdrawal of a token

Creates a withdrawal

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateWithdrawalRequest
*/
func (a *WithdrawalsApiService) CreateWithdrawal(ctx context.Context) ApiCreateWithdrawalRequest {
	return ApiCreateWithdrawalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateWithdrawalResponse
func (a *WithdrawalsApiService) CreateWithdrawalExecute(r ApiCreateWithdrawalRequest) (*CreateWithdrawalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateWithdrawalResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WithdrawalsApiService.CreateWithdrawal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/withdrawals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.createWithdrawalRequest == nil {
		return localVarReturnValue, nil, reportError("createWithdrawalRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarHeaderParams["x-imx-eth-address"] = parameterToString(*r.xImxEthAddress, "")
	localVarHeaderParams["x-imx-eth-signature"] = parameterToString(*r.xImxEthSignature, "")
	// body params
	localVarPostBody = r.createWithdrawalRequest

	localVarHeaderParams["x-sdk-version"] = "imx-core-sdk-golang-0.1.0"
	
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

type ApiGetSignableWithdrawalRequest struct {
	ctx context.Context
	ApiService WithdrawalsApi
	getSignableWithdrawalRequest *GetSignableWithdrawalRequest
}

// get details of signable withdrawal
func (r ApiGetSignableWithdrawalRequest) GetSignableWithdrawalRequest(getSignableWithdrawalRequest GetSignableWithdrawalRequest) ApiGetSignableWithdrawalRequest {
	r.getSignableWithdrawalRequest = &getSignableWithdrawalRequest
	return r
}

func (r ApiGetSignableWithdrawalRequest) Execute() (*GetSignableWithdrawalResponse, *http.Response, error) {
	return r.ApiService.GetSignableWithdrawalExecute(r)
}

/*
GetSignableWithdrawal Gets details of a signable withdrawal

Gets details of a signable withdrawal

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignableWithdrawalRequest
*/
func (a *WithdrawalsApiService) GetSignableWithdrawal(ctx context.Context) ApiGetSignableWithdrawalRequest {
	return ApiGetSignableWithdrawalRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignableWithdrawalResponse
func (a *WithdrawalsApiService) GetSignableWithdrawalExecute(r ApiGetSignableWithdrawalRequest) (*GetSignableWithdrawalResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignableWithdrawalResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WithdrawalsApiService.GetSignableWithdrawal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/signable-withdrawal-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getSignableWithdrawalRequest == nil {
		return localVarReturnValue, nil, reportError("getSignableWithdrawalRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.getSignableWithdrawalRequest

	localVarHeaderParams["x-sdk-version"] = "imx-core-sdk-golang-0.1.0"
	
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

type ApiGetWithdrawalRequest struct {
	ctx context.Context
	ApiService WithdrawalsApi
	id string
}

func (r ApiGetWithdrawalRequest) Execute() (*Withdrawal, *http.Response, error) {
	return r.ApiService.GetWithdrawalExecute(r)
}

/*
GetWithdrawal Gets details of withdrawal with the given ID

Gets details of withdrawal with the given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Withdrawal ID
 @return ApiGetWithdrawalRequest
*/
func (a *WithdrawalsApiService) GetWithdrawal(ctx context.Context, id string) ApiGetWithdrawalRequest {
	return ApiGetWithdrawalRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Withdrawal
func (a *WithdrawalsApiService) GetWithdrawalExecute(r ApiGetWithdrawalRequest) (*Withdrawal, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Withdrawal
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WithdrawalsApiService.GetWithdrawal")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/withdrawals/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

	localVarHeaderParams["x-sdk-version"] = "imx-core-sdk-golang-0.1.0"
	
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

type ApiListWithdrawalsRequest struct {
	ctx context.Context
	ApiService WithdrawalsApi
	withdrawnToWallet *bool
	rollupStatus *string
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
	user *string
	status *string
	minTimestamp *string
	maxTimestamp *string
	tokenType *string
	tokenId *string
	assetId *string
	tokenAddress *string
	tokenName *string
	minQuantity *string
	maxQuantity *string
	metadata *string
}

// Withdrawal has been transferred to user&#39;s Layer 1 wallet
func (r ApiListWithdrawalsRequest) WithdrawnToWallet(withdrawnToWallet bool) ApiListWithdrawalsRequest {
	r.withdrawnToWallet = &withdrawnToWallet
	return r
}

// Status of the on-chain batch confirmation for this withdrawal
func (r ApiListWithdrawalsRequest) RollupStatus(rollupStatus string) ApiListWithdrawalsRequest {
	r.rollupStatus = &rollupStatus
	return r
}

// Page size of the result
func (r ApiListWithdrawalsRequest) PageSize(pageSize int32) ApiListWithdrawalsRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListWithdrawalsRequest) Cursor(cursor string) ApiListWithdrawalsRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListWithdrawalsRequest) OrderBy(orderBy string) ApiListWithdrawalsRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListWithdrawalsRequest) Direction(direction string) ApiListWithdrawalsRequest {
	r.direction = &direction
	return r
}

// Ethereum address of the user who submitted this withdrawal
func (r ApiListWithdrawalsRequest) User(user string) ApiListWithdrawalsRequest {
	r.user = &user
	return r
}

// Status of this withdrawal
func (r ApiListWithdrawalsRequest) Status(status string) ApiListWithdrawalsRequest {
	r.status = &status
	return r
}

// Minimum timestamp for this deposit, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListWithdrawalsRequest) MinTimestamp(minTimestamp string) ApiListWithdrawalsRequest {
	r.minTimestamp = &minTimestamp
	return r
}

// Maximum timestamp for this deposit, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListWithdrawalsRequest) MaxTimestamp(maxTimestamp string) ApiListWithdrawalsRequest {
	r.maxTimestamp = &maxTimestamp
	return r
}

// Token type of the withdrawn asset
func (r ApiListWithdrawalsRequest) TokenType(tokenType string) ApiListWithdrawalsRequest {
	r.tokenType = &tokenType
	return r
}

// ERC721 Token ID of the minted asset
func (r ApiListWithdrawalsRequest) TokenId(tokenId string) ApiListWithdrawalsRequest {
	r.tokenId = &tokenId
	return r
}

// Internal IMX ID of the minted asset
func (r ApiListWithdrawalsRequest) AssetId(assetId string) ApiListWithdrawalsRequest {
	r.assetId = &assetId
	return r
}

// Token address of the withdrawn asset
func (r ApiListWithdrawalsRequest) TokenAddress(tokenAddress string) ApiListWithdrawalsRequest {
	r.tokenAddress = &tokenAddress
	return r
}

// Token name of the withdrawn asset
func (r ApiListWithdrawalsRequest) TokenName(tokenName string) ApiListWithdrawalsRequest {
	r.tokenName = &tokenName
	return r
}

// Min quantity for the withdrawn asset
func (r ApiListWithdrawalsRequest) MinQuantity(minQuantity string) ApiListWithdrawalsRequest {
	r.minQuantity = &minQuantity
	return r
}

// Max quantity for the withdrawn asset
func (r ApiListWithdrawalsRequest) MaxQuantity(maxQuantity string) ApiListWithdrawalsRequest {
	r.maxQuantity = &maxQuantity
	return r
}

// JSON-encoded metadata filters for the withdrawn asset
func (r ApiListWithdrawalsRequest) Metadata(metadata string) ApiListWithdrawalsRequest {
	r.metadata = &metadata
	return r
}

func (r ApiListWithdrawalsRequest) Execute() (*ListWithdrawalsResponse, *http.Response, error) {
	return r.ApiService.ListWithdrawalsExecute(r)
}

/*
ListWithdrawals Get a list of withdrawals

Get a list of withdrawals

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListWithdrawalsRequest
*/
func (a *WithdrawalsApiService) ListWithdrawals(ctx context.Context) ApiListWithdrawalsRequest {
	return ApiListWithdrawalsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListWithdrawalsResponse
func (a *WithdrawalsApiService) ListWithdrawalsExecute(r ApiListWithdrawalsRequest) (*ListWithdrawalsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListWithdrawalsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WithdrawalsApiService.ListWithdrawals")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/withdrawals"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.withdrawnToWallet != nil {
		localVarQueryParams.Add("withdrawn_to_wallet", parameterToString(*r.withdrawnToWallet, ""))
	}
	if r.rollupStatus != nil {
		localVarQueryParams.Add("rollup_status", parameterToString(*r.rollupStatus, ""))
	}
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
	if r.minTimestamp != nil {
		localVarQueryParams.Add("min_timestamp", parameterToString(*r.minTimestamp, ""))
	}
	if r.maxTimestamp != nil {
		localVarQueryParams.Add("max_timestamp", parameterToString(*r.maxTimestamp, ""))
	}
	if r.tokenType != nil {
		localVarQueryParams.Add("token_type", parameterToString(*r.tokenType, ""))
	}
	if r.tokenId != nil {
		localVarQueryParams.Add("token_id", parameterToString(*r.tokenId, ""))
	}
	if r.assetId != nil {
		localVarQueryParams.Add("asset_id", parameterToString(*r.assetId, ""))
	}
	if r.tokenAddress != nil {
		localVarQueryParams.Add("token_address", parameterToString(*r.tokenAddress, ""))
	}
	if r.tokenName != nil {
		localVarQueryParams.Add("token_name", parameterToString(*r.tokenName, ""))
	}
	if r.minQuantity != nil {
		localVarQueryParams.Add("min_quantity", parameterToString(*r.minQuantity, ""))
	}
	if r.maxQuantity != nil {
		localVarQueryParams.Add("max_quantity", parameterToString(*r.maxQuantity, ""))
	}
	if r.metadata != nil {
		localVarQueryParams.Add("metadata", parameterToString(*r.metadata, ""))
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

	localVarHeaderParams["x-sdk-version"] = "imx-core-sdk-golang-0.1.0"
	
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
