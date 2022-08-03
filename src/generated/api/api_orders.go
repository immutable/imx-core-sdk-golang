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


type OrdersApi interface {

	/*
	CancelOrder Cancel an order

	Cancel an order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Order ID to cancel
	@return ApiCancelOrderRequest
	*/
	CancelOrder(ctx context.Context, id string) ApiCancelOrderRequest

	// CancelOrderExecute executes the request
	//  @return CancelOrderResponse
	CancelOrderExecute(r ApiCancelOrderRequest) (*CancelOrderResponse, *http.Response, error)

	/*
	CreateOrder Create an order

	Create an order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateOrderRequest
	*/
	CreateOrder(ctx context.Context) ApiCreateOrderRequest

	// CreateOrderExecute executes the request
	//  @return CreateOrderResponse
	CreateOrderExecute(r ApiCreateOrderRequest) (*CreateOrderResponse, *http.Response, error)

	/*
	GetOrder Get details of an order with the given ID

	Get details of an order with the given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Order ID
	@return ApiGetOrderRequest
	*/
	GetOrder(ctx context.Context, id string) ApiGetOrderRequest

	// GetOrderExecute executes the request
	//  @return Order
	GetOrderExecute(r ApiGetOrderRequest) (*Order, *http.Response, error)

	/*
	GetSignableCancelOrder Get details a signable cancel order

	Get details a signable cancel order

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSignableCancelOrderRequest
	*/
	GetSignableCancelOrder(ctx context.Context) ApiGetSignableCancelOrderRequest

	// GetSignableCancelOrderExecute executes the request
	//  @return GetSignableCancelOrderResponse
	GetSignableCancelOrderExecute(r ApiGetSignableCancelOrderRequest) (*GetSignableCancelOrderResponse, *http.Response, error)

	/*
	GetSignableOrder Get a signable order request (V3)

	Get a signable order request (V3)

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetSignableOrderRequest
	*/
	GetSignableOrder(ctx context.Context) ApiGetSignableOrderRequest

	// GetSignableOrderExecute executes the request
	//  @return GetSignableOrderResponse
	GetSignableOrderExecute(r ApiGetSignableOrderRequest) (*GetSignableOrderResponse, *http.Response, error)

	/*
	ListOrders Get a list of orders

	Get a list of orders

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListOrdersRequest
	*/
	ListOrders(ctx context.Context) ApiListOrdersRequest

	// ListOrdersExecute executes the request
	//  @return ListOrdersResponse
	ListOrdersExecute(r ApiListOrdersRequest) (*ListOrdersResponse, *http.Response, error)
}

// OrdersApiService OrdersApi service
type OrdersApiService service

type ApiCancelOrderRequest struct {
	ctx context.Context
	ApiService OrdersApi
	xImxEthAddress *string
	xImxEthSignature *string
	id string
	cancelOrderRequest *CancelOrderRequest
}

// eth address
func (r ApiCancelOrderRequest) XImxEthAddress(xImxEthAddress string) ApiCancelOrderRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// eth signature
func (r ApiCancelOrderRequest) XImxEthSignature(xImxEthSignature string) ApiCancelOrderRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// cancel an order
func (r ApiCancelOrderRequest) CancelOrderRequest(cancelOrderRequest CancelOrderRequest) ApiCancelOrderRequest {
	r.cancelOrderRequest = &cancelOrderRequest
	return r
}

func (r ApiCancelOrderRequest) Execute() (*CancelOrderResponse, *http.Response, error) {
	return r.ApiService.CancelOrderExecute(r)
}

/*
CancelOrder Cancel an order

Cancel an order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Order ID to cancel
 @return ApiCancelOrderRequest
*/
func (a *OrdersApiService) CancelOrder(ctx context.Context, id string) ApiCancelOrderRequest {
	return ApiCancelOrderRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CancelOrderResponse
func (a *OrdersApiService) CancelOrderExecute(r ApiCancelOrderRequest) (*CancelOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CancelOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.CancelOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/orders/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.cancelOrderRequest == nil {
		return localVarReturnValue, nil, reportError("cancelOrderRequest is required and must be specified")
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
	localVarPostBody = r.cancelOrderRequest

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

type ApiCreateOrderRequest struct {
	ctx context.Context
	ApiService OrdersApi
	xImxEthAddress *string
	xImxEthSignature *string
	createOrderRequest *CreateOrderRequest
}

// eth address
func (r ApiCreateOrderRequest) XImxEthAddress(xImxEthAddress string) ApiCreateOrderRequest {
	r.xImxEthAddress = &xImxEthAddress
	return r
}

// eth signature
func (r ApiCreateOrderRequest) XImxEthSignature(xImxEthSignature string) ApiCreateOrderRequest {
	r.xImxEthSignature = &xImxEthSignature
	return r
}

// create an order
func (r ApiCreateOrderRequest) CreateOrderRequest(createOrderRequest CreateOrderRequest) ApiCreateOrderRequest {
	r.createOrderRequest = &createOrderRequest
	return r
}

func (r ApiCreateOrderRequest) Execute() (*CreateOrderResponse, *http.Response, error) {
	return r.ApiService.CreateOrderExecute(r)
}

/*
CreateOrder Create an order

Create an order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateOrderRequest
*/
func (a *OrdersApiService) CreateOrder(ctx context.Context) ApiCreateOrderRequest {
	return ApiCreateOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateOrderResponse
func (a *OrdersApiService) CreateOrderExecute(r ApiCreateOrderRequest) (*CreateOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.CreateOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/orders"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xImxEthAddress == nil {
		return localVarReturnValue, nil, reportError("xImxEthAddress is required and must be specified")
	}
	if r.xImxEthSignature == nil {
		return localVarReturnValue, nil, reportError("xImxEthSignature is required and must be specified")
	}
	if r.createOrderRequest == nil {
		return localVarReturnValue, nil, reportError("createOrderRequest is required and must be specified")
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
	localVarPostBody = r.createOrderRequest

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

type ApiGetOrderRequest struct {
	ctx context.Context
	ApiService OrdersApi
	id string
	includeFees *bool
	auxiliaryFeePercentages *string
	auxiliaryFeeRecipients *string
}

// Set flag to true to include fee body for the order
func (r ApiGetOrderRequest) IncludeFees(includeFees bool) ApiGetOrderRequest {
	r.includeFees = &includeFees
	return r
}

// Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients
func (r ApiGetOrderRequest) AuxiliaryFeePercentages(auxiliaryFeePercentages string) ApiGetOrderRequest {
	r.auxiliaryFeePercentages = &auxiliaryFeePercentages
	return r
}

// Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages
func (r ApiGetOrderRequest) AuxiliaryFeeRecipients(auxiliaryFeeRecipients string) ApiGetOrderRequest {
	r.auxiliaryFeeRecipients = &auxiliaryFeeRecipients
	return r
}

func (r ApiGetOrderRequest) Execute() (*Order, *http.Response, error) {
	return r.ApiService.GetOrderExecute(r)
}

/*
GetOrder Get details of an order with the given ID

Get details of an order with the given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Order ID
 @return ApiGetOrderRequest
*/
func (a *OrdersApiService) GetOrder(ctx context.Context, id string) ApiGetOrderRequest {
	return ApiGetOrderRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Order
func (a *OrdersApiService) GetOrderExecute(r ApiGetOrderRequest) (*Order, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Order
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.GetOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/orders/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeFees != nil {
		localVarQueryParams.Add("include_fees", parameterToString(*r.includeFees, ""))
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

type ApiGetSignableCancelOrderRequest struct {
	ctx context.Context
	ApiService OrdersApi
	getSignableCancelOrderRequest *GetSignableCancelOrderRequest
}

// get a signable cancel order
func (r ApiGetSignableCancelOrderRequest) GetSignableCancelOrderRequest(getSignableCancelOrderRequest GetSignableCancelOrderRequest) ApiGetSignableCancelOrderRequest {
	r.getSignableCancelOrderRequest = &getSignableCancelOrderRequest
	return r
}

func (r ApiGetSignableCancelOrderRequest) Execute() (*GetSignableCancelOrderResponse, *http.Response, error) {
	return r.ApiService.GetSignableCancelOrderExecute(r)
}

/*
GetSignableCancelOrder Get details a signable cancel order

Get details a signable cancel order

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignableCancelOrderRequest
*/
func (a *OrdersApiService) GetSignableCancelOrder(ctx context.Context) ApiGetSignableCancelOrderRequest {
	return ApiGetSignableCancelOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignableCancelOrderResponse
func (a *OrdersApiService) GetSignableCancelOrderExecute(r ApiGetSignableCancelOrderRequest) (*GetSignableCancelOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignableCancelOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.GetSignableCancelOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/signable-cancel-order-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getSignableCancelOrderRequest == nil {
		return localVarReturnValue, nil, reportError("getSignableCancelOrderRequest is required and must be specified")
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
	localVarPostBody = r.getSignableCancelOrderRequest

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

type ApiGetSignableOrderRequest struct {
	ctx context.Context
	ApiService OrdersApi
	getSignableOrderRequestV3 *GetSignableOrderRequest
}

// get a signable order
func (r ApiGetSignableOrderRequest) GetSignableOrderRequestV3(getSignableOrderRequestV3 GetSignableOrderRequest) ApiGetSignableOrderRequest {
	r.getSignableOrderRequestV3 = &getSignableOrderRequestV3
	return r
}

func (r ApiGetSignableOrderRequest) Execute() (*GetSignableOrderResponse, *http.Response, error) {
	return r.ApiService.GetSignableOrderExecute(r)
}

/*
GetSignableOrder Get a signable order request (V3)

Get a signable order request (V3)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetSignableOrderRequest
*/
func (a *OrdersApiService) GetSignableOrder(ctx context.Context) ApiGetSignableOrderRequest {
	return ApiGetSignableOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSignableOrderResponse
func (a *OrdersApiService) GetSignableOrderExecute(r ApiGetSignableOrderRequest) (*GetSignableOrderResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSignableOrderResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.GetSignableOrder")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v3/signable-order-details"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getSignableOrderRequestV3 == nil {
		return localVarReturnValue, nil, reportError("getSignableOrderRequestV3 is required and must be specified")
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
	localVarPostBody = r.getSignableOrderRequestV3

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

type ApiListOrdersRequest struct {
	ctx context.Context
	ApiService OrdersApi
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
	user *string
	status *string
	minTimestamp *string
	maxTimestamp *string
	updatedMinTimestamp *string
	updatedMaxTimestamp *string
	buyTokenType *string
	buyTokenId *string
	buyAssetId *string
	buyTokenAddress *string
	buyTokenName *string
	buyMinQuantity *string
	buyMaxQuantity *string
	buyMetadata *string
	sellTokenType *string
	sellTokenId *string
	sellAssetId *string
	sellTokenAddress *string
	sellTokenName *string
	sellMinQuantity *string
	sellMaxQuantity *string
	sellMetadata *string
	auxiliaryFeePercentages *string
	auxiliaryFeeRecipients *string
	includeFees *bool
}

// Page size of the result
func (r ApiListOrdersRequest) PageSize(pageSize int32) ApiListOrdersRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListOrdersRequest) Cursor(cursor string) ApiListOrdersRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListOrdersRequest) OrderBy(orderBy string) ApiListOrdersRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListOrdersRequest) Direction(direction string) ApiListOrdersRequest {
	r.direction = &direction
	return r
}

// Ethereum address of the user who submitted this order
func (r ApiListOrdersRequest) User(user string) ApiListOrdersRequest {
	r.user = &user
	return r
}

// Status of this order
func (r ApiListOrdersRequest) Status(status string) ApiListOrdersRequest {
	r.status = &status
	return r
}

// Minimum created at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListOrdersRequest) MinTimestamp(minTimestamp string) ApiListOrdersRequest {
	r.minTimestamp = &minTimestamp
	return r
}

// Maximum created at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListOrdersRequest) MaxTimestamp(maxTimestamp string) ApiListOrdersRequest {
	r.maxTimestamp = &maxTimestamp
	return r
}

// Minimum updated at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListOrdersRequest) UpdatedMinTimestamp(updatedMinTimestamp string) ApiListOrdersRequest {
	r.updatedMinTimestamp = &updatedMinTimestamp
	return r
}

// Maximum updated at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListOrdersRequest) UpdatedMaxTimestamp(updatedMaxTimestamp string) ApiListOrdersRequest {
	r.updatedMaxTimestamp = &updatedMaxTimestamp
	return r
}

// Token type of the asset this order buys
func (r ApiListOrdersRequest) BuyTokenType(buyTokenType string) ApiListOrdersRequest {
	r.buyTokenType = &buyTokenType
	return r
}

// ERC721 Token ID of the asset this order buys
func (r ApiListOrdersRequest) BuyTokenId(buyTokenId string) ApiListOrdersRequest {
	r.buyTokenId = &buyTokenId
	return r
}

// Internal IMX ID of the asset this order buys
func (r ApiListOrdersRequest) BuyAssetId(buyAssetId string) ApiListOrdersRequest {
	r.buyAssetId = &buyAssetId
	return r
}

// Token address of the asset this order buys
func (r ApiListOrdersRequest) BuyTokenAddress(buyTokenAddress string) ApiListOrdersRequest {
	r.buyTokenAddress = &buyTokenAddress
	return r
}

// Token name of the asset this order buys
func (r ApiListOrdersRequest) BuyTokenName(buyTokenName string) ApiListOrdersRequest {
	r.buyTokenName = &buyTokenName
	return r
}

// Min quantity for the asset this order buys
func (r ApiListOrdersRequest) BuyMinQuantity(buyMinQuantity string) ApiListOrdersRequest {
	r.buyMinQuantity = &buyMinQuantity
	return r
}

// Max quantity for the asset this order buys
func (r ApiListOrdersRequest) BuyMaxQuantity(buyMaxQuantity string) ApiListOrdersRequest {
	r.buyMaxQuantity = &buyMaxQuantity
	return r
}

// JSON-encoded metadata filters for the asset this order buys
func (r ApiListOrdersRequest) BuyMetadata(buyMetadata string) ApiListOrdersRequest {
	r.buyMetadata = &buyMetadata
	return r
}

// Token type of the asset this order sells
func (r ApiListOrdersRequest) SellTokenType(sellTokenType string) ApiListOrdersRequest {
	r.sellTokenType = &sellTokenType
	return r
}

// ERC721 Token ID of the asset this order sells
func (r ApiListOrdersRequest) SellTokenId(sellTokenId string) ApiListOrdersRequest {
	r.sellTokenId = &sellTokenId
	return r
}

// Internal IMX ID of the asset this order sells
func (r ApiListOrdersRequest) SellAssetId(sellAssetId string) ApiListOrdersRequest {
	r.sellAssetId = &sellAssetId
	return r
}

// Token address of the asset this order sells
func (r ApiListOrdersRequest) SellTokenAddress(sellTokenAddress string) ApiListOrdersRequest {
	r.sellTokenAddress = &sellTokenAddress
	return r
}

// Token name of the asset this order sells
func (r ApiListOrdersRequest) SellTokenName(sellTokenName string) ApiListOrdersRequest {
	r.sellTokenName = &sellTokenName
	return r
}

// Min quantity for the asset this order sells
func (r ApiListOrdersRequest) SellMinQuantity(sellMinQuantity string) ApiListOrdersRequest {
	r.sellMinQuantity = &sellMinQuantity
	return r
}

// Max quantity for the asset this order sells
func (r ApiListOrdersRequest) SellMaxQuantity(sellMaxQuantity string) ApiListOrdersRequest {
	r.sellMaxQuantity = &sellMaxQuantity
	return r
}

// JSON-encoded metadata filters for the asset this order sells
func (r ApiListOrdersRequest) SellMetadata(sellMetadata string) ApiListOrdersRequest {
	r.sellMetadata = &sellMetadata
	return r
}

// Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients
func (r ApiListOrdersRequest) AuxiliaryFeePercentages(auxiliaryFeePercentages string) ApiListOrdersRequest {
	r.auxiliaryFeePercentages = &auxiliaryFeePercentages
	return r
}

// Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages
func (r ApiListOrdersRequest) AuxiliaryFeeRecipients(auxiliaryFeeRecipients string) ApiListOrdersRequest {
	r.auxiliaryFeeRecipients = &auxiliaryFeeRecipients
	return r
}

// Set flag to true to include fee object for orders
func (r ApiListOrdersRequest) IncludeFees(includeFees bool) ApiListOrdersRequest {
	r.includeFees = &includeFees
	return r
}

func (r ApiListOrdersRequest) Execute() (*ListOrdersResponse, *http.Response, error) {
	return r.ApiService.ListOrdersExecute(r)
}

/*
ListOrders Get a list of orders

Get a list of orders

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListOrdersRequest
*/
func (a *OrdersApiService) ListOrders(ctx context.Context) ApiListOrdersRequest {
	return ApiListOrdersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListOrdersResponse
func (a *OrdersApiService) ListOrdersExecute(r ApiListOrdersRequest) (*ListOrdersResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListOrdersResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OrdersApiService.ListOrders")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/orders"

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
	if r.minTimestamp != nil {
		localVarQueryParams.Add("min_timestamp", parameterToString(*r.minTimestamp, ""))
	}
	if r.maxTimestamp != nil {
		localVarQueryParams.Add("max_timestamp", parameterToString(*r.maxTimestamp, ""))
	}
	if r.updatedMinTimestamp != nil {
		localVarQueryParams.Add("updated_min_timestamp", parameterToString(*r.updatedMinTimestamp, ""))
	}
	if r.updatedMaxTimestamp != nil {
		localVarQueryParams.Add("updated_max_timestamp", parameterToString(*r.updatedMaxTimestamp, ""))
	}
	if r.buyTokenType != nil {
		localVarQueryParams.Add("buy_token_type", parameterToString(*r.buyTokenType, ""))
	}
	if r.buyTokenId != nil {
		localVarQueryParams.Add("buy_token_id", parameterToString(*r.buyTokenId, ""))
	}
	if r.buyAssetId != nil {
		localVarQueryParams.Add("buy_asset_id", parameterToString(*r.buyAssetId, ""))
	}
	if r.buyTokenAddress != nil {
		localVarQueryParams.Add("buy_token_address", parameterToString(*r.buyTokenAddress, ""))
	}
	if r.buyTokenName != nil {
		localVarQueryParams.Add("buy_token_name", parameterToString(*r.buyTokenName, ""))
	}
	if r.buyMinQuantity != nil {
		localVarQueryParams.Add("buy_min_quantity", parameterToString(*r.buyMinQuantity, ""))
	}
	if r.buyMaxQuantity != nil {
		localVarQueryParams.Add("buy_max_quantity", parameterToString(*r.buyMaxQuantity, ""))
	}
	if r.buyMetadata != nil {
		localVarQueryParams.Add("buy_metadata", parameterToString(*r.buyMetadata, ""))
	}
	if r.sellTokenType != nil {
		localVarQueryParams.Add("sell_token_type", parameterToString(*r.sellTokenType, ""))
	}
	if r.sellTokenId != nil {
		localVarQueryParams.Add("sell_token_id", parameterToString(*r.sellTokenId, ""))
	}
	if r.sellAssetId != nil {
		localVarQueryParams.Add("sell_asset_id", parameterToString(*r.sellAssetId, ""))
	}
	if r.sellTokenAddress != nil {
		localVarQueryParams.Add("sell_token_address", parameterToString(*r.sellTokenAddress, ""))
	}
	if r.sellTokenName != nil {
		localVarQueryParams.Add("sell_token_name", parameterToString(*r.sellTokenName, ""))
	}
	if r.sellMinQuantity != nil {
		localVarQueryParams.Add("sell_min_quantity", parameterToString(*r.sellMinQuantity, ""))
	}
	if r.sellMaxQuantity != nil {
		localVarQueryParams.Add("sell_max_quantity", parameterToString(*r.sellMaxQuantity, ""))
	}
	if r.sellMetadata != nil {
		localVarQueryParams.Add("sell_metadata", parameterToString(*r.sellMetadata, ""))
	}
	if r.auxiliaryFeePercentages != nil {
		localVarQueryParams.Add("auxiliary_fee_percentages", parameterToString(*r.auxiliaryFeePercentages, ""))
	}
	if r.auxiliaryFeeRecipients != nil {
		localVarQueryParams.Add("auxiliary_fee_recipients", parameterToString(*r.auxiliaryFeeRecipients, ""))
	}
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
	localVarHTTPHeaderAccepts := []string{"*/*"}

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
