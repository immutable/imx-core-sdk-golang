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


type BalancesApi interface {

	/*
	GetBalance Fetches the token balances of the user

	Fetches the token balances of the user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param owner Address of the owner/user
	@param address Token address
	@return ApiGetBalanceRequest
	*/
	GetBalance(ctx context.Context, owner string, address string) ApiGetBalanceRequest

	// GetBalanceExecute executes the request
	//  @return Balance
	GetBalanceExecute(r ApiGetBalanceRequest) (*Balance, *http.Response, error)

	/*
	ListBalances Get a list of balances for given user

	Get a list of balances for given user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param owner Ethereum wallet address for user
	@return ApiListBalancesRequest
	*/
	ListBalances(ctx context.Context, owner string) ApiListBalancesRequest

	// ListBalancesExecute executes the request
	//  @return ListBalancesResponse
	ListBalancesExecute(r ApiListBalancesRequest) (*ListBalancesResponse, *http.Response, error)
}

// BalancesApiService BalancesApi service
type BalancesApiService service

type ApiGetBalanceRequest struct {
	ctx context.Context
	ApiService BalancesApi
	owner string
	address string
}

func (r ApiGetBalanceRequest) Execute() (*Balance, *http.Response, error) {
	return r.ApiService.GetBalanceExecute(r)
}

/*
GetBalance Fetches the token balances of the user

Fetches the token balances of the user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner Address of the owner/user
 @param address Token address
 @return ApiGetBalanceRequest
*/
func (a *BalancesApiService) GetBalance(ctx context.Context, owner string, address string) ApiGetBalanceRequest {
	return ApiGetBalanceRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
		address: address,
	}
}

// Execute executes the request
//  @return Balance
func (a *BalancesApiService) GetBalanceExecute(r ApiGetBalanceRequest) (*Balance, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Balance
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesApiService.GetBalance")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/balances/{owner}/{address}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"address"+"}", url.PathEscape(parameterToString(r.address, "")), -1)

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

type ApiListBalancesRequest struct {
	ctx context.Context
	ApiService BalancesApi
	owner string
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
}

// Page size of the result
func (r ApiListBalancesRequest) PageSize(pageSize int32) ApiListBalancesRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListBalancesRequest) Cursor(cursor string) ApiListBalancesRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListBalancesRequest) OrderBy(orderBy string) ApiListBalancesRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListBalancesRequest) Direction(direction string) ApiListBalancesRequest {
	r.direction = &direction
	return r
}

func (r ApiListBalancesRequest) Execute() (*ListBalancesResponse, *http.Response, error) {
	return r.ApiService.ListBalancesExecute(r)
}

/*
ListBalances Get a list of balances for given user

Get a list of balances for given user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param owner Ethereum wallet address for user
 @return ApiListBalancesRequest
*/
func (a *BalancesApiService) ListBalances(ctx context.Context, owner string) ApiListBalancesRequest {
	return ApiListBalancesRequest{
		ApiService: a,
		ctx: ctx,
		owner: owner,
	}
}

// Execute executes the request
//  @return ListBalancesResponse
func (a *BalancesApiService) ListBalancesExecute(r ApiListBalancesRequest) (*ListBalancesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListBalancesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BalancesApiService.ListBalances")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/balances/{owner}"
	localVarPath = strings.Replace(localVarPath, "{"+"owner"+"}", url.PathEscape(parameterToString(r.owner, "")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
