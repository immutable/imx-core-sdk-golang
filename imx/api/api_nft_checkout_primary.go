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


type NftCheckoutPrimaryApi interface {

	/*
	CreateNftPrimary Create NFT primary sale transaction

	Creates a transaction representing minting an NFT with a card payment.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateNftPrimaryRequest
	*/
	CreateNftPrimary(ctx context.Context) ApiCreateNftPrimaryRequest

	// CreateNftPrimaryExecute executes the request
	//  @return NftprimarytransactionCreateResponse
	CreateNftPrimaryExecute(r ApiCreateNftPrimaryRequest) (*NftprimarytransactionCreateResponse, *http.Response, error)

	/*
	GetCurrenciesNFTCheckoutPrimary Get currencies with limits

	Returns a list of supported currencies and their limits

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetCurrenciesNFTCheckoutPrimaryRequest
	*/
	GetCurrenciesNFTCheckoutPrimary(ctx context.Context) ApiGetCurrenciesNFTCheckoutPrimaryRequest

	// GetCurrenciesNFTCheckoutPrimaryExecute executes the request
	//  @return CurrencyWithLimits
	GetCurrenciesNFTCheckoutPrimaryExecute(r ApiGetCurrenciesNFTCheckoutPrimaryRequest) (*CurrencyWithLimits, *http.Response, error)

	/*
	GetNftPrimaryTransaction Get NFT primary sale transaction by id

	given a transaction id, returns the corresponding transaction representing a mint executed from a card payment

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param transactionId Transaction id
	@return ApiGetNftPrimaryTransactionRequest
	*/
	GetNftPrimaryTransaction(ctx context.Context, transactionId string) ApiGetNftPrimaryTransactionRequest

	// GetNftPrimaryTransactionExecute executes the request
	//  @return NftprimarytransactionGetResponse
	GetNftPrimaryTransactionExecute(r ApiGetNftPrimaryTransactionRequest) (*NftprimarytransactionGetResponse, *http.Response, error)

	/*
	GetNftPrimaryTransactions Get a list of NFT primary sales transactions

	Returns a list of NFT primary sales transactions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiGetNftPrimaryTransactionsRequest
	*/
	GetNftPrimaryTransactions(ctx context.Context) ApiGetNftPrimaryTransactionsRequest

	// GetNftPrimaryTransactionsExecute executes the request
	//  @return NftprimarytransactionListTransactionsResponse
	GetNftPrimaryTransactionsExecute(r ApiGetNftPrimaryTransactionsRequest) (*NftprimarytransactionListTransactionsResponse, *http.Response, error)

	/*
	RegisterNftPrimarySalesContract Executes NFT primary sales contract registration

	Registers a new contract for use in the minting with fiat card flow

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiRegisterNftPrimarySalesContractRequest
	*/
	RegisterNftPrimarySalesContract(ctx context.Context) ApiRegisterNftPrimarySalesContractRequest

	// RegisterNftPrimarySalesContractExecute executes the request
	//  @return ContractCreateResponse
	RegisterNftPrimarySalesContractExecute(r ApiRegisterNftPrimarySalesContractRequest) (*ContractCreateResponse, *http.Response, error)
}

// NftCheckoutPrimaryApiService NftCheckoutPrimaryApi service
type NftCheckoutPrimaryApiService service

type ApiCreateNftPrimaryRequest struct {
	ctx context.Context
	ApiService NftCheckoutPrimaryApi
	createAPIRequest *NftprimarytransactionCreateAPIRequest
}

// req
func (r ApiCreateNftPrimaryRequest) CreateAPIRequest(createAPIRequest NftprimarytransactionCreateAPIRequest) ApiCreateNftPrimaryRequest {
	r.createAPIRequest = &createAPIRequest
	return r
}

func (r ApiCreateNftPrimaryRequest) Execute() (*NftprimarytransactionCreateResponse, *http.Response, error) {
	return r.ApiService.CreateNftPrimaryExecute(r)
}

/*
CreateNftPrimary Create NFT primary sale transaction

Creates a transaction representing minting an NFT with a card payment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateNftPrimaryRequest
*/
func (a *NftCheckoutPrimaryApiService) CreateNftPrimary(ctx context.Context) ApiCreateNftPrimaryRequest {
	return ApiCreateNftPrimaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NftprimarytransactionCreateResponse
func (a *NftCheckoutPrimaryApiService) CreateNftPrimaryExecute(r ApiCreateNftPrimaryRequest) (*NftprimarytransactionCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NftprimarytransactionCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NftCheckoutPrimaryApiService.CreateNftPrimary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/nft/primary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAPIRequest == nil {
		return localVarReturnValue, nil, reportError("createAPIRequest is required and must be specified")
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
	localVarPostBody = r.createAPIRequest
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
			var v LambdasAPIError
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

type ApiGetCurrenciesNFTCheckoutPrimaryRequest struct {
	ctx context.Context
	ApiService NftCheckoutPrimaryApi
	provider *string
	includeLimits *bool
}

// Provider name
func (r ApiGetCurrenciesNFTCheckoutPrimaryRequest) Provider(provider string) ApiGetCurrenciesNFTCheckoutPrimaryRequest {
	r.provider = &provider
	return r
}

// Flag if limits should be included in the response or not
func (r ApiGetCurrenciesNFTCheckoutPrimaryRequest) IncludeLimits(includeLimits bool) ApiGetCurrenciesNFTCheckoutPrimaryRequest {
	r.includeLimits = &includeLimits
	return r
}

func (r ApiGetCurrenciesNFTCheckoutPrimaryRequest) Execute() (*CurrencyWithLimits, *http.Response, error) {
	return r.ApiService.GetCurrenciesNFTCheckoutPrimaryExecute(r)
}

/*
GetCurrenciesNFTCheckoutPrimary Get currencies with limits

Returns a list of supported currencies and their limits

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetCurrenciesNFTCheckoutPrimaryRequest
*/
func (a *NftCheckoutPrimaryApiService) GetCurrenciesNFTCheckoutPrimary(ctx context.Context) ApiGetCurrenciesNFTCheckoutPrimaryRequest {
	return ApiGetCurrenciesNFTCheckoutPrimaryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CurrencyWithLimits
func (a *NftCheckoutPrimaryApiService) GetCurrenciesNFTCheckoutPrimaryExecute(r ApiGetCurrenciesNFTCheckoutPrimaryRequest) (*CurrencyWithLimits, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CurrencyWithLimits
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NftCheckoutPrimaryApiService.GetCurrenciesNFTCheckoutPrimary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/nft/primary/currencies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.provider != nil {
		parameterAddToQuery(localVarQueryParams, "provider", r.provider, "")
	}
	if r.includeLimits != nil {
		parameterAddToQuery(localVarQueryParams, "include_limits", r.includeLimits, "")
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
			var v LambdasAPIError
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

type ApiGetNftPrimaryTransactionRequest struct {
	ctx context.Context
	ApiService NftCheckoutPrimaryApi
	transactionId string
}

func (r ApiGetNftPrimaryTransactionRequest) Execute() (*NftprimarytransactionGetResponse, *http.Response, error) {
	return r.ApiService.GetNftPrimaryTransactionExecute(r)
}

/*
GetNftPrimaryTransaction Get NFT primary sale transaction by id

given a transaction id, returns the corresponding transaction representing a mint executed from a card payment

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param transactionId Transaction id
 @return ApiGetNftPrimaryTransactionRequest
*/
func (a *NftCheckoutPrimaryApiService) GetNftPrimaryTransaction(ctx context.Context, transactionId string) ApiGetNftPrimaryTransactionRequest {
	return ApiGetNftPrimaryTransactionRequest{
		ApiService: a,
		ctx: ctx,
		transactionId: transactionId,
	}
}

// Execute executes the request
//  @return NftprimarytransactionGetResponse
func (a *NftCheckoutPrimaryApiService) GetNftPrimaryTransactionExecute(r ApiGetNftPrimaryTransactionRequest) (*NftprimarytransactionGetResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NftprimarytransactionGetResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NftCheckoutPrimaryApiService.GetNftPrimaryTransaction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/nft/primary/{transaction_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"transaction_id"+"}", url.PathEscape(parameterValueToString(r.transactionId, "transactionId")), -1)

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v LambdasAPIError
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

type ApiGetNftPrimaryTransactionsRequest struct {
	ctx context.Context
	ApiService NftCheckoutPrimaryApi
	pageSize *int32
	cursor *string
	orderBy *string
	direction *string
	transactionId *string
	contractAddress *string
	sellerWalletAddress *string
	userWalletAddress *string
	status *string
	provider *string
	mintId *string
}

// Page size of the result
func (r ApiGetNftPrimaryTransactionsRequest) PageSize(pageSize int32) ApiGetNftPrimaryTransactionsRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiGetNftPrimaryTransactionsRequest) Cursor(cursor string) ApiGetNftPrimaryTransactionsRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiGetNftPrimaryTransactionsRequest) OrderBy(orderBy string) ApiGetNftPrimaryTransactionsRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiGetNftPrimaryTransactionsRequest) Direction(direction string) ApiGetNftPrimaryTransactionsRequest {
	r.direction = &direction
	return r
}

// Transaction id
func (r ApiGetNftPrimaryTransactionsRequest) TransactionId(transactionId string) ApiGetNftPrimaryTransactionsRequest {
	r.transactionId = &transactionId
	return r
}

// Contract address of the asset
func (r ApiGetNftPrimaryTransactionsRequest) ContractAddress(contractAddress string) ApiGetNftPrimaryTransactionsRequest {
	r.contractAddress = &contractAddress
	return r
}

// Ethereum address of the seller
func (r ApiGetNftPrimaryTransactionsRequest) SellerWalletAddress(sellerWalletAddress string) ApiGetNftPrimaryTransactionsRequest {
	r.sellerWalletAddress = &sellerWalletAddress
	return r
}

// Ethereum address of the user who wants to create transaction
func (r ApiGetNftPrimaryTransactionsRequest) UserWalletAddress(userWalletAddress string) ApiGetNftPrimaryTransactionsRequest {
	r.userWalletAddress = &userWalletAddress
	return r
}

// Transaction status
func (r ApiGetNftPrimaryTransactionsRequest) Status(status string) ApiGetNftPrimaryTransactionsRequest {
	r.status = &status
	return r
}

// Checkout provider name
func (r ApiGetNftPrimaryTransactionsRequest) Provider(provider string) ApiGetNftPrimaryTransactionsRequest {
	r.provider = &provider
	return r
}

// Minting transaction ID - see mintTokens response
func (r ApiGetNftPrimaryTransactionsRequest) MintId(mintId string) ApiGetNftPrimaryTransactionsRequest {
	r.mintId = &mintId
	return r
}

func (r ApiGetNftPrimaryTransactionsRequest) Execute() (*NftprimarytransactionListTransactionsResponse, *http.Response, error) {
	return r.ApiService.GetNftPrimaryTransactionsExecute(r)
}

/*
GetNftPrimaryTransactions Get a list of NFT primary sales transactions

Returns a list of NFT primary sales transactions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNftPrimaryTransactionsRequest
*/
func (a *NftCheckoutPrimaryApiService) GetNftPrimaryTransactions(ctx context.Context) ApiGetNftPrimaryTransactionsRequest {
	return ApiGetNftPrimaryTransactionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return NftprimarytransactionListTransactionsResponse
func (a *NftCheckoutPrimaryApiService) GetNftPrimaryTransactionsExecute(r ApiGetNftPrimaryTransactionsRequest) (*NftprimarytransactionListTransactionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *NftprimarytransactionListTransactionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NftCheckoutPrimaryApiService.GetNftPrimaryTransactions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/nft/primary"

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
	if r.transactionId != nil {
		parameterAddToQuery(localVarQueryParams, "transaction_id", r.transactionId, "")
	}
	if r.contractAddress != nil {
		parameterAddToQuery(localVarQueryParams, "contract_address", r.contractAddress, "")
	}
	if r.sellerWalletAddress != nil {
		parameterAddToQuery(localVarQueryParams, "seller_wallet_address", r.sellerWalletAddress, "")
	}
	if r.userWalletAddress != nil {
		parameterAddToQuery(localVarQueryParams, "user_wallet_address", r.userWalletAddress, "")
	}
	if r.status != nil {
		parameterAddToQuery(localVarQueryParams, "status", r.status, "")
	}
	if r.provider != nil {
		parameterAddToQuery(localVarQueryParams, "provider", r.provider, "")
	}
	if r.mintId != nil {
		parameterAddToQuery(localVarQueryParams, "mint_id", r.mintId, "")
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
			var v LambdasAPIError
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

type ApiRegisterNftPrimarySalesContractRequest struct {
	ctx context.Context
	ApiService NftCheckoutPrimaryApi
	createAPIRequest *ContractCreateAPIRequest
}

// req
func (r ApiRegisterNftPrimarySalesContractRequest) CreateAPIRequest(createAPIRequest ContractCreateAPIRequest) ApiRegisterNftPrimarySalesContractRequest {
	r.createAPIRequest = &createAPIRequest
	return r
}

func (r ApiRegisterNftPrimarySalesContractRequest) Execute() (*ContractCreateResponse, *http.Response, error) {
	return r.ApiService.RegisterNftPrimarySalesContractExecute(r)
}

/*
RegisterNftPrimarySalesContract Executes NFT primary sales contract registration

Registers a new contract for use in the minting with fiat card flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiRegisterNftPrimarySalesContractRequest
*/
func (a *NftCheckoutPrimaryApiService) RegisterNftPrimarySalesContract(ctx context.Context) ApiRegisterNftPrimarySalesContractRequest {
	return ApiRegisterNftPrimarySalesContractRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ContractCreateResponse
func (a *NftCheckoutPrimaryApiService) RegisterNftPrimarySalesContractExecute(r ApiRegisterNftPrimarySalesContractRequest) (*ContractCreateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ContractCreateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NftCheckoutPrimaryApiService.RegisterNftPrimarySalesContract")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/nft/primary/register"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createAPIRequest == nil {
		return localVarReturnValue, nil, reportError("createAPIRequest is required and must be specified")
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
	localVarPostBody = r.createAPIRequest
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
			var v LambdasAPIError
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
