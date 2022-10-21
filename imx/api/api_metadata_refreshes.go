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


type MetadataRefreshesApi interface {

	/*
	GetAListOfMetadataRefreshes Method for GetAListOfMetadataRefreshes

	Get a list of metadata refreshes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetAListOfMetadataRefreshesRequest
	*/
	GetAListOfMetadataRefreshes(ctx context.Context) ApiGetAListOfMetadataRefreshesRequest

	// GetAListOfMetadataRefreshesExecute executes the request
	//  @return GetMetadataRefreshes
	GetAListOfMetadataRefreshesExecute(r ApiGetAListOfMetadataRefreshesRequest) (*GetMetadataRefreshes, *http.Response, error)

	/*
	GetMetadataRefreshErrors Method for GetMetadataRefreshErrors

	Get metadata refresh errors

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param refreshId The metadata refresh ID
	@return ApiGetMetadataRefreshErrorsRequest
	*/
	GetMetadataRefreshErrors(ctx context.Context, refreshId string) ApiGetMetadataRefreshErrorsRequest

	// GetMetadataRefreshErrorsExecute executes the request
	//  @return GetMetadataRefreshErrorsResponse
	GetMetadataRefreshErrorsExecute(r ApiGetMetadataRefreshErrorsRequest) (*GetMetadataRefreshErrorsResponse, *http.Response, error)

	/*
	GetMetadataRefreshResults Method for GetMetadataRefreshResults

	Get metadata refresh results

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param refreshId The metadata refresh ID
	@return ApiGetMetadataRefreshResultsRequest
	*/
	GetMetadataRefreshResults(ctx context.Context, refreshId string) ApiGetMetadataRefreshResultsRequest

	// GetMetadataRefreshResultsExecute executes the request
	//  @return GetMetadataRefreshResponse
	GetMetadataRefreshResultsExecute(r ApiGetMetadataRefreshResultsRequest) (*GetMetadataRefreshResponse, *http.Response, error)

	/*
	RequestAMetadataRefresh Method for RequestAMetadataRefresh

	Request metadata refresh for provided tokens

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRequestAMetadataRefreshRequest
	*/
	RequestAMetadataRefresh(ctx context.Context) ApiRequestAMetadataRefreshRequest

	// RequestAMetadataRefreshExecute executes the request
	//  @return CreateMetadataRefreshResponse
	RequestAMetadataRefreshExecute(r ApiRequestAMetadataRefreshRequest) (*CreateMetadataRefreshResponse, *http.Response, error)
}

// MetadataRefreshesApiService MetadataRefreshesApi service
type MetadataRefreshesApiService service

type ApiGetAListOfMetadataRefreshesRequest struct {
	ctx context.Context
	ApiService MetadataRefreshesApi
	xImxEthSignature *string
	xImxEthTimestamp *string
	xImxEthAddress *string
	pageSize *int32
	cursor *string
	collectionAddress *string
}

// String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
func (r ApiGetAListOfMetadataRefreshesRequest) XImxEthSignature(xImxEthSignature string) ApiGetAListOfMetadataRefreshesRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// Unix Epoc timestamp
func (r ApiGetAListOfMetadataRefreshesRequest) XImxEthTimestamp(xImxEthTimestamp string) ApiGetAListOfMetadataRefreshesRequest {
	r.xImxEthTimestamp = &xImxEthTimestamp
	return r
}

// Wallet Address that signed the signature
func (r ApiGetAListOfMetadataRefreshesRequest) XImxEthAddress(xImxEthAddress string) ApiGetAListOfMetadataRefreshesRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// Page size of the result
func (r ApiGetAListOfMetadataRefreshesRequest) PageSize(pageSize int32) ApiGetAListOfMetadataRefreshesRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiGetAListOfMetadataRefreshesRequest) Cursor(cursor string) ApiGetAListOfMetadataRefreshesRequest {
	r.cursor = &cursor
	return r
}

// Collection address
func (r ApiGetAListOfMetadataRefreshesRequest) CollectionAddress(collectionAddress string) ApiGetAListOfMetadataRefreshesRequest {
	r.collectionAddress = &collectionAddress
	return r
}

func (r ApiGetAListOfMetadataRefreshesRequest) Execute() (*GetMetadataRefreshes, *http.Response, error) {
	return r.ApiService.GetAListOfMetadataRefreshesExecute(r)
}

/*
GetAListOfMetadataRefreshes Method for GetAListOfMetadataRefreshes

Get a list of metadata refreshes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetAListOfMetadataRefreshesRequest
*/
func (a *MetadataRefreshesApiService) GetAListOfMetadataRefreshes(ctx context.Context) ApiGetAListOfMetadataRefreshesRequest {
	return ApiGetAListOfMetadataRefreshesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetMetadataRefreshes
func (a *MetadataRefreshesApiService) GetAListOfMetadataRefreshesExecute(r ApiGetAListOfMetadataRefreshesRequest) (*GetMetadataRefreshes, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMetadataRefreshes
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataRefreshesApiService.GetAListOfMetadataRefreshes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata-refreshes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.xImxEthTimestamp == nil {
		return localVarReturnValue, nil, reportError("xImxEthTimestamp is required and must be specified")
	}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
	}
	if r.collectionAddress != nil {
		localVarQueryParams.Add("collection_address", parameterToString(*r.collectionAddress, ""))
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
	localVarHeaderParams["x-imx-eth-signature"] = parameterToString(*r.xImxEthSignature, "")
	localVarHeaderParams["x-imx-eth-timestamp"] = parameterToString(*r.xImxEthTimestamp, "")
	localVarHeaderParams["x-imx-eth-address"] = parameterToString(*r.xImxEthAddress, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiGetMetadataRefreshErrorsRequest struct {
	ctx context.Context
	ApiService MetadataRefreshesApi
	refreshId string
	xImxEthSignature *string
	xImxEthTimestamp *string
	xImxEthAddress *string
	pageSize *int32
	cursor *string
}

// String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
func (r ApiGetMetadataRefreshErrorsRequest) XImxEthSignature(xImxEthSignature string) ApiGetMetadataRefreshErrorsRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// Unix Epoc timestamp
func (r ApiGetMetadataRefreshErrorsRequest) XImxEthTimestamp(xImxEthTimestamp string) ApiGetMetadataRefreshErrorsRequest {
	r.xImxEthTimestamp = &xImxEthTimestamp
	return r
}

// Wallet Address that signed the signature
func (r ApiGetMetadataRefreshErrorsRequest) XImxEthAddress(xImxEthAddress string) ApiGetMetadataRefreshErrorsRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// Page size of the result
func (r ApiGetMetadataRefreshErrorsRequest) PageSize(pageSize int32) ApiGetMetadataRefreshErrorsRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiGetMetadataRefreshErrorsRequest) Cursor(cursor string) ApiGetMetadataRefreshErrorsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetMetadataRefreshErrorsRequest) Execute() (*GetMetadataRefreshErrorsResponse, *http.Response, error) {
	return r.ApiService.GetMetadataRefreshErrorsExecute(r)
}

/*
GetMetadataRefreshErrors Method for GetMetadataRefreshErrors

Get metadata refresh errors

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param refreshId The metadata refresh ID
 @return ApiGetMetadataRefreshErrorsRequest
*/
func (a *MetadataRefreshesApiService) GetMetadataRefreshErrors(ctx context.Context, refreshId string) ApiGetMetadataRefreshErrorsRequest {
	return ApiGetMetadataRefreshErrorsRequest{
		ApiService: a,
		ctx: ctx,
		refreshId: refreshId,
	}
}

// Execute executes the request
//  @return GetMetadataRefreshErrorsResponse
func (a *MetadataRefreshesApiService) GetMetadataRefreshErrorsExecute(r ApiGetMetadataRefreshErrorsRequest) (*GetMetadataRefreshErrorsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMetadataRefreshErrorsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataRefreshesApiService.GetMetadataRefreshErrors")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata-refreshes/{refresh_id}/errors"
	localVarPath = strings.Replace(localVarPath, "{"+"refresh_id"+"}", url.PathEscape(parameterToString(r.refreshId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.xImxEthTimestamp == nil {
		return localVarReturnValue, nil, reportError("xImxEthTimestamp is required and must be specified")
	}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}

	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.cursor != nil {
		localVarQueryParams.Add("cursor", parameterToString(*r.cursor, ""))
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
	localVarHeaderParams["x-imx-eth-signature"] = parameterToString(*r.xImxEthSignature, "")
	localVarHeaderParams["x-imx-eth-timestamp"] = parameterToString(*r.xImxEthTimestamp, "")
	localVarHeaderParams["x-imx-eth-address"] = parameterToString(*r.xImxEthAddress, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiGetMetadataRefreshResultsRequest struct {
	ctx context.Context
	ApiService MetadataRefreshesApi
	refreshId string
	xImxEthSignature *string
	xImxEthTimestamp *string
	xImxEthAddress *string
}

// String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
func (r ApiGetMetadataRefreshResultsRequest) XImxEthSignature(xImxEthSignature string) ApiGetMetadataRefreshResultsRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// Unix Epoc timestamp
func (r ApiGetMetadataRefreshResultsRequest) XImxEthTimestamp(xImxEthTimestamp string) ApiGetMetadataRefreshResultsRequest {
	r.xImxEthTimestamp = &xImxEthTimestamp
	return r
}

// Wallet Address that signed the signature
func (r ApiGetMetadataRefreshResultsRequest) XImxEthAddress(xImxEthAddress string) ApiGetMetadataRefreshResultsRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

func (r ApiGetMetadataRefreshResultsRequest) Execute() (*GetMetadataRefreshResponse, *http.Response, error) {
	return r.ApiService.GetMetadataRefreshResultsExecute(r)
}

/*
GetMetadataRefreshResults Method for GetMetadataRefreshResults

Get metadata refresh results

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param refreshId The metadata refresh ID
 @return ApiGetMetadataRefreshResultsRequest
*/
func (a *MetadataRefreshesApiService) GetMetadataRefreshResults(ctx context.Context, refreshId string) ApiGetMetadataRefreshResultsRequest {
	return ApiGetMetadataRefreshResultsRequest{
		ApiService: a,
		ctx: ctx,
		refreshId: refreshId,
	}
}

// Execute executes the request
//  @return GetMetadataRefreshResponse
func (a *MetadataRefreshesApiService) GetMetadataRefreshResultsExecute(r ApiGetMetadataRefreshResultsRequest) (*GetMetadataRefreshResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetMetadataRefreshResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataRefreshesApiService.GetMetadataRefreshResults")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata-refreshes/{refresh_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"refresh_id"+"}", url.PathEscape(parameterToString(r.refreshId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.xImxEthTimestamp == nil {
		return localVarReturnValue, nil, reportError("xImxEthTimestamp is required and must be specified")
	}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
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
	localVarHeaderParams["x-imx-eth-signature"] = parameterToString(*r.xImxEthSignature, "")
	localVarHeaderParams["x-imx-eth-timestamp"] = parameterToString(*r.xImxEthTimestamp, "")
	localVarHeaderParams["x-imx-eth-address"] = parameterToString(*r.xImxEthAddress, "")
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiRequestAMetadataRefreshRequest struct {
	ctx context.Context
	ApiService MetadataRefreshesApi
	xImxEthSignature *string
	xImxEthTimestamp *string
	xImxEthAddress *string
	createMetadataRefreshRequest *CreateMetadataRefreshRequest
}

// String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
func (r ApiRequestAMetadataRefreshRequest) XImxEthSignature(xImxEthSignature string) ApiRequestAMetadataRefreshRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// Unix Epoc timestamp
func (r ApiRequestAMetadataRefreshRequest) XImxEthTimestamp(xImxEthTimestamp string) ApiRequestAMetadataRefreshRequest {
	r.xImxEthTimestamp = &xImxEthTimestamp
	return r
}

// Wallet Address that signed the signature
func (r ApiRequestAMetadataRefreshRequest) XImxEthAddress(xImxEthAddress string) ApiRequestAMetadataRefreshRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// Create metadata refresh request
func (r ApiRequestAMetadataRefreshRequest) CreateMetadataRefreshRequest(createMetadataRefreshRequest CreateMetadataRefreshRequest) ApiRequestAMetadataRefreshRequest {
	r.createMetadataRefreshRequest = &createMetadataRefreshRequest
	return r
}

func (r ApiRequestAMetadataRefreshRequest) Execute() (*CreateMetadataRefreshResponse, *http.Response, error) {
	return r.ApiService.RequestAMetadataRefreshExecute(r)
}

/*
RequestAMetadataRefresh Method for RequestAMetadataRefresh

Request metadata refresh for provided tokens

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRequestAMetadataRefreshRequest
*/
func (a *MetadataRefreshesApiService) RequestAMetadataRefresh(ctx context.Context) ApiRequestAMetadataRefreshRequest {
	return ApiRequestAMetadataRefreshRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateMetadataRefreshResponse
func (a *MetadataRefreshesApiService) RequestAMetadataRefreshExecute(r ApiRequestAMetadataRefreshRequest) (*CreateMetadataRefreshResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateMetadataRefreshResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MetadataRefreshesApiService.RequestAMetadataRefresh")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metadata-refreshes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.xImxEthTimestamp == nil {
		return localVarReturnValue, nil, reportError("xImxEthTimestamp is required and must be specified")
	}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.createMetadataRefreshRequest == nil {
		return localVarReturnValue, nil, reportError("createMetadataRefreshRequest is required and must be specified")
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
	localVarHeaderParams["x-imx-eth-signature"] = parameterToString(*r.xImxEthSignature, "")
	localVarHeaderParams["x-imx-eth-timestamp"] = parameterToString(*r.xImxEthTimestamp, "")
	localVarHeaderParams["x-imx-eth-address"] = parameterToString(*r.xImxEthAddress, "")
	// body params
	localVarPostBody = r.createMetadataRefreshRequest
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
		if localVarHTTPResponse.StatusCode == 401 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
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