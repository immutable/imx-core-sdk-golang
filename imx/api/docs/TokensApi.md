# \TokensApi

All URIs are relative to *https://api.sandbox.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetToken**](TokensApi.md#GetToken) | **Get** /v1/tokens/{address} | Get details of a token
[**ListTokens**](TokensApi.md#ListTokens) | **Get** /v1/tokens | Get a list of tokens



## GetToken

> TokenDetails GetToken(ctx, address).Execute()

Get details of a token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/api"
    
)

func main() {
    address := "address_example" // string | Token Contract Address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.GetToken(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToken`: TokenDetails
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Token Contract Address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenDetails**](TokenDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> ListTokensResponse ListTokens(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Address(address).Symbols(symbols).Execute()

Get a list of tokens



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/api"
    
)

func main() {
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)
    address := "address_example" // string | Contract address of the token (optional)
    symbols := "symbols_example" // string | Token symbols for the token, e.g. ?symbols=IMX,ETH (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TokensApi.ListTokens(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Address(address).Symbols(symbols).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensApi.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: ListTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensApi.ListTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **address** | **string** | Contract address of the token | 
 **symbols** | **string** | Token symbols for the token, e.g. ?symbols&#x3D;IMX,ETH | 

### Return type

[**ListTokensResponse**](ListTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

