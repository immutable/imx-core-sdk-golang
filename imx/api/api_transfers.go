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


type TransfersApi interface {

	/*
	CreateTransfer Creates a transfer of multiple tokens between two parties

	Create a new transfer request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTransferRequest
	*/
	CreateTransfer(ctx context.Context) ApiCreateTransferRequest

	// CreateTransferExecute executes the request
	//  @return CreateTransferResponse
	CreateTransferExecute(r ApiCreateTransferRequest) (*CreateTransferResponse, *http.Response, error)

	/*
	CreateTransferV1 Creates a transfer of tokens between two parties

	Create a new transfer request

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateTransferV1Request
	*/
	CreateTransferV1(ctx context.Context) ApiCreateTransferV1Request

	// CreateTransferV1Execute executes the request
	//  @return CreateTransferResponseV1
	CreateTransferV1Execute(r ApiCreateTransferV1Request) (*CreateTransferResponseV1, *http.Response, error)

	/*
	GetSignableTransfer Gets bulk details of a signable transfer

	Gets bulk details of a signable transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSignableTransferRequest
	*/
	GetSignableTransfer(ctx context.Context) ApiGetSignableTransferRequest

	// GetSignableTransferExecute executes the request
	//  @return GetSignableTransferResponse
	GetSignableTransferExecute(r ApiGetSignableTransferRequest) (*GetSignableTransferResponse, *http.Response, error)

	/*
	GetSignableTransferV1 Gets details of a signable transfer

	Gets details of a signable transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSignableTransferV1Request
	*/
	GetSignableTransferV1(ctx context.Context) ApiGetSignableTransferV1Request

	// GetSignableTransferV1Execute executes the request
	//  @return GetSignableTransferResponseV1
	GetSignableTransferV1Execute(r ApiGetSignableTransferV1Request) (*GetSignableTransferResponseV1, *http.Response, error)

	/*
	GetTransfer Get details of a transfer with the given ID

	Get details of a transfer with the given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Transfer ID
	@return ApiGetTransferRequest
	*/
	GetTransfer(ctx context.Context, id string) ApiGetTransferRequest

	// GetTransferExecute executes the request
	//  @return Transfer
	GetTransferExecute(r ApiGetTransferRequest) (*Transfer, *http.Response, error)

	/*
	ListTransfers Get a list of transfers

	Get a list of transfers

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListTransfersRequest
	*/
	ListTransfers(ctx context.Context) ApiListTransfersRequest

	// ListTransfersExecute executes the request
	//  @return ListTransfersResponse
	ListTransfersExecute(r ApiListTransfersRequest) (*ListTransfersResponse, *http.Response, error)
}

// TransfersApiService TransfersApi service
type TransfersApiService service

type ApiCreateTransferRequest struct {
	ctx context.Context
	ApiService TransfersApi
	xImxEthAddress *string
	xImxEthSignature *string
	createTransferRequestV2 *CreateTransferRequest
}

// eth address
func (r ApiCreateTransferRequest) XImxEthAddress(xImxEthAddress string) ApiCreateTransferRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// eth signature
func (r ApiCreateTransferRequest) XImxEthSignature(xImxEthSignature string) ApiCreateTransferRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// Create transfer
func (r ApiCreateTransferRequest) CreateTransferRequestV2(createTransferRequestV2 CreateTransferRequest) ApiCreateTransferRequest {
	r.createTransferRequestV2 = &createTransferRequestV2
	return r
}

func (r ApiCreateTransferRequest) Execute() (*CreateTransferResponse, *http.Response, error) {
	return r.ApiService.CreateTransferExecute(r)
}

/*
CreateTransfer Creates a transfer of multiple tokens between two parties

Create a new transfer request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTransferRequest
*/
func (a *TransfersApiService) CreateTransfer(ctx context.Context) ApiCreateTransferRequest {
	return ApiCreateTransferRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateTransferResponse
func (a *TransfersApiService) CreateTransferExecute(r ApiCreateTransferRequest) (*CreateTransferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateTransferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.CreateTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.createTransferRequestV2 == nil {
		return localVarReturnValue, nil, reportError("createTransferRequestV2 is required and must be specified")
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
	parameterAddToQuery(localVarQueryParams, "x-imx-eth-address", r.xImxEthAddress, "")
	parameterAddToQuery(localVarQueryParams, "x-imx-eth-signature", r.xImxEthSignature, "")
	// body params
	localVarPostBody = r.createTransferRequestV2
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiCreateTransferV1Request struct {
	ctx context.Context
	ApiService TransfersApi
	xImxEthAddress *string
	xImxEthSignature *string
	createTransferRequest *CreateTransferRequestV1
}

// eth address
func (r ApiCreateTransferV1Request) XImxEthAddress(xImxEthAddress string) ApiCreateTransferV1Request {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// eth signature
func (r ApiCreateTransferV1Request) XImxEthSignature(xImxEthSignature string) ApiCreateTransferV1Request {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// Create transfer
func (r ApiCreateTransferV1Request) CreateTransferRequest(createTransferRequest CreateTransferRequestV1) ApiCreateTransferV1Request {
	r.createTransferRequest = &createTransferRequest
	return r
}

func (r ApiCreateTransferV1Request) Execute() (*CreateTransferResponseV1, *http.Response, error) {
	return r.ApiService.CreateTransferV1Execute(r)
}

/*
CreateTransferV1 Creates a transfer of tokens between two parties

Create a new transfer request

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateTransferV1Request
*/
func (a *TransfersApiService) CreateTransferV1(ctx context.Context) ApiCreateTransferV1Request {
	return ApiCreateTransferV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateTransferResponseV1
func (a *TransfersApiService) CreateTransferV1Execute(r ApiCreateTransferV1Request) (*CreateTransferResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateTransferResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.CreateTransferV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.createTransferRequest == nil {
		return localVarReturnValue, nil, reportError("createTransferRequest is required and must be specified")
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
	parameterAddToQuery(localVarQueryParams, "x-imx-eth-address", r.xImxEthAddress, "")
	parameterAddToQuery(localVarQueryParams, "x-imx-eth-signature", r.xImxEthSignature, "")
	// body params
	localVarPostBody = r.createTransferRequest
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiGetSignableTransferRequest struct {
	ctx context.Context
	ApiService TransfersApi
	getSignableTransferRequestV2 *GetSignableTransferRequest
}

// get details of signable transfer
func (r ApiGetSignableTransferRequest) GetSignableTransferRequestV2(getSignableTransferRequestV2 GetSignableTransferRequest) ApiGetSignableTransferRequest {
	r.getSignableTransferRequestV2 = &getSignableTransferRequestV2
	return r
}

func (r ApiGetSignableTransferRequest) Execute() (*GetSignableTransferResponse, *http.Response, error) {
	return r.ApiService.GetSignableTransferExecute(r)
}

/*
GetSignableTransfer Gets bulk details of a signable transfer

Gets bulk details of a signable transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignableTransferRequest
*/
func (a *TransfersApiService) GetSignableTransfer(ctx context.Context) ApiGetSignableTransferRequest {
	return ApiGetSignableTransferRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignableTransferResponse
func (a *TransfersApiService) GetSignableTransferExecute(r ApiGetSignableTransferRequest) (*GetSignableTransferResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignableTransferResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.GetSignableTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/signable-transfer-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getSignableTransferRequestV2 == nil {
		return localVarReturnValue, nil, reportError("getSignableTransferRequestV2 is required and must be specified")
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
	localVarPostBody = r.getSignableTransferRequestV2
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiGetSignableTransferV1Request struct {
	ctx context.Context
	ApiService TransfersApi
	getSignableTransferRequest *GetSignableTransferRequestV1
}

// get details of signable transfer
func (r ApiGetSignableTransferV1Request) GetSignableTransferRequest(getSignableTransferRequest GetSignableTransferRequestV1) ApiGetSignableTransferV1Request {
	r.getSignableTransferRequest = &getSignableTransferRequest
	return r
}

func (r ApiGetSignableTransferV1Request) Execute() (*GetSignableTransferResponseV1, *http.Response, error) {
	return r.ApiService.GetSignableTransferV1Execute(r)
}

/*
GetSignableTransferV1 Gets details of a signable transfer

Gets details of a signable transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignableTransferV1Request
*/
func (a *TransfersApiService) GetSignableTransferV1(ctx context.Context) ApiGetSignableTransferV1Request {
	return ApiGetSignableTransferV1Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignableTransferResponseV1
func (a *TransfersApiService) GetSignableTransferV1Execute(r ApiGetSignableTransferV1Request) (*GetSignableTransferResponseV1, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignableTransferResponseV1
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.GetSignableTransferV1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/signable-transfer-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getSignableTransferRequest == nil {
		return localVarReturnValue, nil, reportError("getSignableTransferRequest is required and must be specified")
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
	localVarPostBody = r.getSignableTransferRequest
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

type ApiGetTransferRequest struct {
	ctx context.Context
	ApiService TransfersApi
	id string
}

func (r ApiGetTransferRequest) Execute() (*Transfer, *http.Response, error) {
	return r.ApiService.GetTransferExecute(r)
}

/*
GetTransfer Get details of a transfer with the given ID

Get details of a transfer with the given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Transfer ID
 @return ApiGetTransferRequest
*/
func (a *TransfersApiService) GetTransfer(ctx context.Context, id string) ApiGetTransferRequest {
	return ApiGetTransferRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Transfer
func (a *TransfersApiService) GetTransferExecute(r ApiGetTransferRequest) (*Transfer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Transfer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.GetTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transfers/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
			var v APIError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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

type ApiListTransfersRequest struct {
	ctx context.Context
	ApiService TransfersApi
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
	user *string
	receiver *string
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

// Page size of the result
func (r ApiListTransfersRequest) PageSize(pageSize int32) ApiListTransfersRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListTransfersRequest) Cursor(cursor string) ApiListTransfersRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListTransfersRequest) OrderBy(orderBy string) ApiListTransfersRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListTransfersRequest) Direction(direction string) ApiListTransfersRequest {
	r.direction = &direction
	return r
}

// Ethereum address of the user who submitted this transfer
func (r ApiListTransfersRequest) User(user string) ApiListTransfersRequest {
	r.user = &user
	return r
}

// Ethereum address of the user who received this transfer
func (r ApiListTransfersRequest) Receiver(receiver string) ApiListTransfersRequest {
	r.receiver = &receiver
	return r
}

// Status of this transfer
func (r ApiListTransfersRequest) Status(status string) ApiListTransfersRequest {
	r.status = &status
	return r
}

// Minimum timestamp for this transfer, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListTransfersRequest) MinTimestamp(minTimestamp string) ApiListTransfersRequest {
	r.minTimestamp = &minTimestamp
	return r
}

// Maximum timestamp for this transfer, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListTransfersRequest) MaxTimestamp(maxTimestamp string) ApiListTransfersRequest {
	r.maxTimestamp = &maxTimestamp
	return r
}

// Token type of the transferred asset
func (r ApiListTransfersRequest) TokenType(tokenType string) ApiListTransfersRequest {
	r.tokenType = &tokenType
	return r
}

// ERC721 Token ID of the minted asset
func (r ApiListTransfersRequest) TokenId(tokenId string) ApiListTransfersRequest {
	r.tokenId = &tokenId
	return r
}

// Internal IMX ID of the minted asset
func (r ApiListTransfersRequest) AssetId(assetId string) ApiListTransfersRequest {
	r.assetId = &assetId
	return r
}

// Token address of the transferred asset
func (r ApiListTransfersRequest) TokenAddress(tokenAddress string) ApiListTransfersRequest {
	r.tokenAddress = &tokenAddress
	return r
}

// Token name of the transferred asset
func (r ApiListTransfersRequest) TokenName(tokenName string) ApiListTransfersRequest {
	r.tokenName = &tokenName
	return r
}

// Max quantity for the transferred asset
func (r ApiListTransfersRequest) MinQuantity(minQuantity string) ApiListTransfersRequest {
	r.minQuantity = &minQuantity
	return r
}

// Max quantity for the transferred asset
func (r ApiListTransfersRequest) MaxQuantity(maxQuantity string) ApiListTransfersRequest {
	r.maxQuantity = &maxQuantity
	return r
}

// JSON-encoded metadata filters for the transferred asset
func (r ApiListTransfersRequest) Metadata(metadata string) ApiListTransfersRequest {
	r.metadata = &metadata
	return r
}

func (r ApiListTransfersRequest) Execute() (*ListTransfersResponse, *http.Response, error) {
	return r.ApiService.ListTransfersExecute(r)
}

/*
ListTransfers Get a list of transfers

Get a list of transfers

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListTransfersRequest
*/
func (a *TransfersApiService) ListTransfers(ctx context.Context) ApiListTransfersRequest {
	return ApiListTransfersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListTransfersResponse
func (a *TransfersApiService) ListTransfersExecute(r ApiListTransfersRequest) (*ListTransfersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListTransfersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TransfersApiService.ListTransfers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/transfers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.pageSize != nil {
		parameterAddToQuery(localVarQueryParams, "page_size", r.pageSize, "")
	}
	if r.cursor != nil {
		parameterAddToQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	if r.orderBy != nil {
		parameterAddToQuery(localVarQueryParams, "order_by", r.orderBy, "")
	}
	if r.direction != nil {
		parameterAddToQuery(localVarQueryParams, "direction", r.direction, "")
	}
	if r.user != nil {
		parameterAddToQuery(localVarQueryParams, "user", r.user, "")
	}
	if r.receiver != nil {
		parameterAddToQuery(localVarQueryParams, "receiver", r.receiver, "")
	}
	if r.status != nil {
		parameterAddToQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.minTimestamp != nil {
		parameterAddToQuery(localVarQueryParams, "min_timestamp", r.minTimestamp, "")
	}
	if r.maxTimestamp != nil {
		parameterAddToQuery(localVarQueryParams, "max_timestamp", r.maxTimestamp, "")
	}
	if r.tokenType != nil {
		parameterAddToQuery(localVarQueryParams, "token_type", r.tokenType, "")
	}
	if r.tokenId != nil {
		parameterAddToQuery(localVarQueryParams, "token_id", r.tokenId, "")
	}
	if r.assetId != nil {
		parameterAddToQuery(localVarQueryParams, "asset_id", r.assetId, "")
	}
	if r.tokenAddress != nil {
		parameterAddToQuery(localVarQueryParams, "token_address", r.tokenAddress, "")
	}
	if r.tokenName != nil {
		parameterAddToQuery(localVarQueryParams, "token_name", r.tokenName, "")
	}
	if r.minQuantity != nil {
		parameterAddToQuery(localVarQueryParams, "min_quantity", r.minQuantity, "")
	}
	if r.maxQuantity != nil {
		parameterAddToQuery(localVarQueryParams, "max_quantity", r.maxQuantity, "")
	}
	if r.metadata != nil {
		parameterAddToQuery(localVarQueryParams, "metadata", r.metadata, "")
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
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
