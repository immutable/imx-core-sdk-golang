/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
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


type TradesApi interface {

	/*
	CreateTrade Create a Trade between two parties

	Create a Trade

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTradeRequest
	*/
	CreateTrade(ctx context.Context) ApiCreateTradeRequest

	// CreateTradeExecute executes the request
	//  @return CreateTradeResponse
	CreateTradeExecute(r ApiCreateTradeRequest) (*CreateTradeResponse, *http.Response, error)

	/*
	GetSignableTrade Get details a signable trade V3

	Get details a signable trade V3

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSignableTradeRequest
	*/
	GetSignableTrade(ctx context.Context) ApiGetSignableTradeRequest

	// GetSignableTradeExecute executes the request
	//  @return GetSignableTradeResponse
	GetSignableTradeExecute(r ApiGetSignableTradeRequest) (*GetSignableTradeResponse, *http.Response, error)

	/*
	GetTrade Get details of a trade with the given ID

	Get details of a trade with the given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Trade ID
	@return ApiGetTradeRequest
	*/
	GetTrade(ctx context.Context, id string) ApiGetTradeRequest

	// GetTradeExecute executes the request
	//  @return Trade
	GetTradeExecute(r ApiGetTradeRequest) (*Trade, *http.Response, error)

	/*
	ListTrades Get a list of trades

	Get a list of trades

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTradesRequest
	*/
	ListTrades(ctx context.Context) ApiListTradesRequest

	// ListTradesExecute executes the request
	//  @return ListTradesResponse
	ListTradesExecute(r ApiListTradesRequest) (*ListTradesResponse, *http.Response, error)
}

// TradesApiService TradesApi service
type TradesApiService service

type ApiCreateTradeRequest struct {
	ctx context.Context
	ApiService TradesApi
	xImxEthAddress *string
	xImxEthSignature *string
	createTradeRequest *CreateTradeRequestV1
}

// eth address
func (r ApiCreateTradeRequest) XImxEthAddress(xImxEthAddress string) ApiCreateTradeRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// eth signature
func (r ApiCreateTradeRequest) XImxEthSignature(xImxEthSignature string) ApiCreateTradeRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// create a trade
func (r ApiCreateTradeRequest) CreateTradeRequest(createTradeRequest CreateTradeRequestV1) ApiCreateTradeRequest {
	r.createTradeRequest = &createTradeRequest
	return r
}

func (r ApiCreateTradeRequest) Execute() (*CreateTradeResponse, *http.Response, error) {
	return r.ApiService.CreateTradeExecute(r)
}

/*
CreateTrade Create a Trade between two parties

Create a Trade

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTradeRequest
*/
func (a *TradesApiService) CreateTrade(ctx context.Context) ApiCreateTradeRequest {
	return ApiCreateTradeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateTradeResponse
func (a *TradesApiService) CreateTradeExecute(r ApiCreateTradeRequest) (*CreateTradeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateTradeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TradesApiService.CreateTrade")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/trades"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.createTradeRequest == nil {
		return localVarReturnValue, nil, reportError("createTradeRequest is required and must be specified")
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
	localVarPostBody = r.createTradeRequest
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

type ApiGetSignableTradeRequest struct {
	ctx context.Context
	ApiService TradesApi
	getSignableTradeRequest *GetSignableTradeRequest
}

// get a signable trade
func (r ApiGetSignableTradeRequest) GetSignableTradeRequest(getSignableTradeRequest GetSignableTradeRequest) ApiGetSignableTradeRequest {
	r.getSignableTradeRequest = &getSignableTradeRequest
	return r
}

func (r ApiGetSignableTradeRequest) Execute() (*GetSignableTradeResponse, *http.Response, error) {
	return r.ApiService.GetSignableTradeExecute(r)
}

/*
GetSignableTrade Get details a signable trade V3

Get details a signable trade V3

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignableTradeRequest
*/
func (a *TradesApiService) GetSignableTrade(ctx context.Context) ApiGetSignableTradeRequest {
	return ApiGetSignableTradeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignableTradeResponse
func (a *TradesApiService) GetSignableTradeExecute(r ApiGetSignableTradeRequest) (*GetSignableTradeResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignableTradeResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TradesApiService.GetSignableTrade")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/signable-trade-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getSignableTradeRequest == nil {
		return localVarReturnValue, nil, reportError("getSignableTradeRequest is required and must be specified")
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
	localVarPostBody = r.getSignableTradeRequest
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

type ApiGetTradeRequest struct {
	ctx context.Context
	ApiService TradesApi
	id string
}

func (r ApiGetTradeRequest) Execute() (*Trade, *http.Response, error) {
	return r.ApiService.GetTradeExecute(r)
}

/*
GetTrade Get details of a trade with the given ID

Get details of a trade with the given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Trade ID
 @return ApiGetTradeRequest
*/
func (a *TradesApiService) GetTrade(ctx context.Context, id string) ApiGetTradeRequest {
	return ApiGetTradeRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Trade
func (a *TradesApiService) GetTradeExecute(r ApiGetTradeRequest) (*Trade, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Trade
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TradesApiService.GetTrade")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/trades/{id}"
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

type ApiListTradesRequest struct {
	ctx context.Context
	ApiService TradesApi
	partyATokenType *string
	partyATokenAddress *string
	partyBTokenType *string
	partyBTokenAddress *string
	partyBTokenId *string
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
	minTimestamp *string
	maxTimestamp *string
}

// Party A&#39;s (buy order) token type of currency used to buy
func (r ApiListTradesRequest) PartyATokenType(partyATokenType string) ApiListTradesRequest {
	r.partyATokenType = &partyATokenType
	return r
}

// Party A&#39;s (buy order) token address of currency used to buy
func (r ApiListTradesRequest) PartyATokenAddress(partyATokenAddress string) ApiListTradesRequest {
	r.partyATokenAddress = &partyATokenAddress
	return r
}

// Party B&#39;s (sell order) token type of NFT sold - always ERC721
func (r ApiListTradesRequest) PartyBTokenType(partyBTokenType string) ApiListTradesRequest {
	r.partyBTokenType = &partyBTokenType
	return r
}

// Party B&#39;s (sell order) collection address of NFT sold
func (r ApiListTradesRequest) PartyBTokenAddress(partyBTokenAddress string) ApiListTradesRequest {
	r.partyBTokenAddress = &partyBTokenAddress
	return r
}

// Party B&#39;s (sell order) token id of NFT sold
func (r ApiListTradesRequest) PartyBTokenId(partyBTokenId string) ApiListTradesRequest {
	r.partyBTokenId = &partyBTokenId
	return r
}

// Page size of the result
func (r ApiListTradesRequest) PageSize(pageSize int32) ApiListTradesRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListTradesRequest) Cursor(cursor string) ApiListTradesRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListTradesRequest) OrderBy(orderBy string) ApiListTradesRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListTradesRequest) Direction(direction string) ApiListTradesRequest {
	r.direction = &direction
	return r
}

// Minimum timestamp for this trade, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListTradesRequest) MinTimestamp(minTimestamp string) ApiListTradesRequest {
	r.minTimestamp = &minTimestamp
	return r
}

// Maximum timestamp for this trade, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListTradesRequest) MaxTimestamp(maxTimestamp string) ApiListTradesRequest {
	r.maxTimestamp = &maxTimestamp
	return r
}

func (r ApiListTradesRequest) Execute() (*ListTradesResponse, *http.Response, error) {
	return r.ApiService.ListTradesExecute(r)
}

/*
ListTrades Get a list of trades

Get a list of trades

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListTradesRequest
*/
func (a *TradesApiService) ListTrades(ctx context.Context) ApiListTradesRequest {
	return ApiListTradesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListTradesResponse
func (a *TradesApiService) ListTradesExecute(r ApiListTradesRequest) (*ListTradesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListTradesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TradesApiService.ListTrades")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/trades"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.partyATokenType != nil {
		localVarQueryParams.Add("party_a_token_type", parameterToString(*r.partyATokenType, ""))
	}
	if r.partyATokenAddress != nil {
		localVarQueryParams.Add("party_a_token_address", parameterToString(*r.partyATokenAddress, ""))
	}
	if r.partyBTokenType != nil {
		localVarQueryParams.Add("party_b_token_type", parameterToString(*r.partyBTokenType, ""))
	}
	if r.partyBTokenAddress != nil {
		localVarQueryParams.Add("party_b_token_address", parameterToString(*r.partyBTokenAddress, ""))
	}
	if r.partyBTokenId != nil {
		localVarQueryParams.Add("party_b_token_id", parameterToString(*r.partyBTokenId, ""))
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
	if r.minTimestamp != nil {
		localVarQueryParams.Add("min_timestamp", parameterToString(*r.minTimestamp, ""))
	}
	if r.maxTimestamp != nil {
		localVarQueryParams.Add("max_timestamp", parameterToString(*r.maxTimestamp, ""))
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
