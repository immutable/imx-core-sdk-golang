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


type MintsApi interface {

	/*
	GetMint Get details of a mint with the given ID

	Get details of a mint with the given ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id Mint ID. This is the transaction_id returned from listMints
	@return ApiGetMintRequest
	*/
	GetMint(ctx context.Context, id string) ApiGetMintRequest

	// GetMintExecute executes the request
	//  @return Mint
	GetMintExecute(r ApiGetMintRequest) (*Mint, *http.Response, error)

	/*
	GetMintableTokenDetailsByClientTokenId Get details of a mintable token with the given token address and token ID

	Get details of a mintable token with the given token address and token ID

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param tokenAddress Address of the ERC721 contract
	@param tokenId ERC721 token ID
	@return ApiGetMintableTokenDetailsByClientTokenIdRequest
	*/
	GetMintableTokenDetailsByClientTokenId(ctx context.Context, tokenAddress string, tokenId string) ApiGetMintableTokenDetailsByClientTokenIdRequest

	// GetMintableTokenDetailsByClientTokenIdExecute executes the request
	//  @return MintableTokenDetails
	GetMintableTokenDetailsByClientTokenIdExecute(r ApiGetMintableTokenDetailsByClientTokenIdRequest) (*MintableTokenDetails, *http.Response, error)

	/*
	ListMints Get a list of mints

	Get a list of mints

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListMintsRequest
	*/
	ListMints(ctx context.Context) ApiListMintsRequest

	// ListMintsExecute executes the request
	//  @return ListMintsResponse
	ListMintsExecute(r ApiListMintsRequest) (*ListMintsResponse, *http.Response, error)

	/*
	MintTokens Mint Tokens V2

	Mint tokens in a batch with fees

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiMintTokensRequest
	*/
	MintTokens(ctx context.Context) ApiMintTokensRequest

	// MintTokensExecute executes the request
	//  @return MintTokensResponse
	MintTokensExecute(r ApiMintTokensRequest) (*MintTokensResponse, *http.Response, error)
}

// MintsApiService MintsApi service
type MintsApiService service

type ApiGetMintRequest struct {
	ctx context.Context
	ApiService MintsApi
	id string
}

func (r ApiGetMintRequest) Execute() (*Mint, *http.Response, error) {
	return r.ApiService.GetMintExecute(r)
}

/*
GetMint Get details of a mint with the given ID

Get details of a mint with the given ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id Mint ID. This is the transaction_id returned from listMints
 @return ApiGetMintRequest
*/
func (a *MintsApiService) GetMint(ctx context.Context, id string) ApiGetMintRequest {
	return ApiGetMintRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Mint
func (a *MintsApiService) GetMintExecute(r ApiGetMintRequest) (*Mint, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Mint
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MintsApiService.GetMint")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/mints/{id}"
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

type ApiGetMintableTokenDetailsByClientTokenIdRequest struct {
	ctx context.Context
	ApiService MintsApi
	tokenAddress string
	tokenId string
}

func (r ApiGetMintableTokenDetailsByClientTokenIdRequest) Execute() (*MintableTokenDetails, *http.Response, error) {
	return r.ApiService.GetMintableTokenDetailsByClientTokenIdExecute(r)
}

/*
GetMintableTokenDetailsByClientTokenId Get details of a mintable token with the given token address and token ID

Get details of a mintable token with the given token address and token ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tokenAddress Address of the ERC721 contract
 @param tokenId ERC721 token ID
 @return ApiGetMintableTokenDetailsByClientTokenIdRequest
*/
func (a *MintsApiService) GetMintableTokenDetailsByClientTokenId(ctx context.Context, tokenAddress string, tokenId string) ApiGetMintableTokenDetailsByClientTokenIdRequest {
	return ApiGetMintableTokenDetailsByClientTokenIdRequest{
		ApiService: a,
		ctx: ctx,
		tokenAddress: tokenAddress,
		tokenId: tokenId,
	}
}

// Execute executes the request
//  @return MintableTokenDetails
func (a *MintsApiService) GetMintableTokenDetailsByClientTokenIdExecute(r ApiGetMintableTokenDetailsByClientTokenIdRequest) (*MintableTokenDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MintableTokenDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MintsApiService.GetMintableTokenDetailsByClientTokenId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/mintable-token/{token_address}/{token_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"token_address"+"}", url.PathEscape(parameterToString(r.tokenAddress, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"token_id"+"}", url.PathEscape(parameterToString(r.tokenId, "")), -1)

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

type ApiListMintsRequest struct {
	ctx context.Context
	ApiService MintsApi
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
	tokenName *string
	tokenAddress *string
	minQuantity *string
	maxQuantity *string
	metadata *string
}

// Page size of the result
func (r ApiListMintsRequest) PageSize(pageSize int32) ApiListMintsRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor
func (r ApiListMintsRequest) Cursor(cursor string) ApiListMintsRequest {
	r.cursor = &cursor
	return r
}

// Property to sort by
func (r ApiListMintsRequest) OrderBy(orderBy string) ApiListMintsRequest {
	r.orderBy = &orderBy
	return r
}

// Direction to sort (asc/desc)
func (r ApiListMintsRequest) Direction(direction string) ApiListMintsRequest {
	r.direction = &direction
	return r
}

// Ethereum address of the user who submitted this mint
func (r ApiListMintsRequest) User(user string) ApiListMintsRequest {
	r.user = &user
	return r
}

// Status of this mint
func (r ApiListMintsRequest) Status(status string) ApiListMintsRequest {
	r.status = &status
	return r
}

// Minimum timestamp for this mint, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListMintsRequest) MinTimestamp(minTimestamp string) ApiListMintsRequest {
	r.minTimestamp = &minTimestamp
	return r
}

// Maximum timestamp for this mint, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39;
func (r ApiListMintsRequest) MaxTimestamp(maxTimestamp string) ApiListMintsRequest {
	r.maxTimestamp = &maxTimestamp
	return r
}

// Token type of the minted asset
func (r ApiListMintsRequest) TokenType(tokenType string) ApiListMintsRequest {
	r.tokenType = &tokenType
	return r
}

// ERC721 Token ID of the minted asset
func (r ApiListMintsRequest) TokenId(tokenId string) ApiListMintsRequest {
	r.tokenId = &tokenId
	return r
}

// Internal IMX ID of the minted asset
func (r ApiListMintsRequest) AssetId(assetId string) ApiListMintsRequest {
	r.assetId = &assetId
	return r
}

// Token Name of the minted asset
func (r ApiListMintsRequest) TokenName(tokenName string) ApiListMintsRequest {
	r.tokenName = &tokenName
	return r
}

// Token address of the minted asset
func (r ApiListMintsRequest) TokenAddress(tokenAddress string) ApiListMintsRequest {
	r.tokenAddress = &tokenAddress
	return r
}

// Min quantity for the minted asset
func (r ApiListMintsRequest) MinQuantity(minQuantity string) ApiListMintsRequest {
	r.minQuantity = &minQuantity
	return r
}

// Max quantity for the minted asset
func (r ApiListMintsRequest) MaxQuantity(maxQuantity string) ApiListMintsRequest {
	r.maxQuantity = &maxQuantity
	return r
}

// JSON-encoded metadata filters for the minted asset
func (r ApiListMintsRequest) Metadata(metadata string) ApiListMintsRequest {
	r.metadata = &metadata
	return r
}

func (r ApiListMintsRequest) Execute() (*ListMintsResponse, *http.Response, error) {
	return r.ApiService.ListMintsExecute(r)
}

/*
ListMints Get a list of mints

Get a list of mints

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiListMintsRequest
*/
func (a *MintsApiService) ListMints(ctx context.Context) ApiListMintsRequest {
	return ApiListMintsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListMintsResponse
func (a *MintsApiService) ListMintsExecute(r ApiListMintsRequest) (*ListMintsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListMintsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MintsApiService.ListMints")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/mints"

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
	if r.tokenType != nil {
		localVarQueryParams.Add("token_type", parameterToString(*r.tokenType, ""))
	}
	if r.tokenId != nil {
		localVarQueryParams.Add("token_id", parameterToString(*r.tokenId, ""))
	}
	if r.assetId != nil {
		localVarQueryParams.Add("asset_id", parameterToString(*r.assetId, ""))
	}
	if r.tokenName != nil {
		localVarQueryParams.Add("token_name", parameterToString(*r.tokenName, ""))
	}
	if r.tokenAddress != nil {
		localVarQueryParams.Add("token_address", parameterToString(*r.tokenAddress, ""))
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

type ApiMintTokensRequest struct {
	ctx context.Context
	ApiService MintsApi
	mintTokensRequestV2 *[]MintRequest
}

// details of tokens to mint
func (r ApiMintTokensRequest) MintTokensRequestV2(mintTokensRequestV2 []MintRequest) ApiMintTokensRequest {
	r.mintTokensRequestV2 = &mintTokensRequestV2
	return r
}

func (r ApiMintTokensRequest) Execute() (*MintTokensResponse, *http.Response, error) {
	return r.ApiService.MintTokensExecute(r)
}

/*
MintTokens Mint Tokens V2

Mint tokens in a batch with fees

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiMintTokensRequest
*/
func (a *MintsApiService) MintTokens(ctx context.Context) ApiMintTokensRequest {
	return ApiMintTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MintTokensResponse
func (a *MintsApiService) MintTokensExecute(r ApiMintTokensRequest) (*MintTokensResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MintTokensResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "MintsApiService.MintTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v2/mints"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mintTokensRequestV2 == nil {
		return localVarReturnValue, nil, reportError("mintTokensRequestV2 is required and must be specified")
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
	localVarPostBody = r.mintTokensRequestV2
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
