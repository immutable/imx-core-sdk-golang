# \MetadataRefreshesApi

All URIs are relative to *https://api.sandbox.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAListOfMetadataRefreshes**](MetadataRefreshesApi.md#GetAListOfMetadataRefreshes) | **Get** /v1/metadata-refreshes | 
[**GetMetadataRefreshErrors**](MetadataRefreshesApi.md#GetMetadataRefreshErrors) | **Get** /v1/metadata-refreshes/{refresh_id}/errors | 
[**GetMetadataRefreshResults**](MetadataRefreshesApi.md#GetMetadataRefreshResults) | **Get** /v1/metadata-refreshes/{refresh_id} | 
[**RequestAMetadataRefresh**](MetadataRefreshesApi.md#RequestAMetadataRefresh) | **Post** /v1/metadata-refreshes | 



## GetAListOfMetadataRefreshes

> GetMetadataRefreshes GetAListOfMetadataRefreshes(ctx).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).PageSize(pageSize).Cursor(cursor).CollectionAddress(collectionAddress).Execute()





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
    xImxEthSignature := "xImxEthSignature_example" // string | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
    xImxEthTimestamp := "xImxEthTimestamp_example" // string | Unix Epoc timestamp
    xImxEthAddress := "xImxEthAddress_example" // string | Wallet Address that signed the signature
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    collectionAddress := "collectionAddress_example" // string | Collection address (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataRefreshesApi.GetAListOfMetadataRefreshes(context.Background()).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).PageSize(pageSize).Cursor(cursor).CollectionAddress(collectionAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataRefreshesApi.GetAListOfMetadataRefreshes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAListOfMetadataRefreshes`: GetMetadataRefreshes
    fmt.Fprintf(os.Stdout, "Response from `MetadataRefreshesApi.GetAListOfMetadataRefreshes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAListOfMetadataRefreshesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xImxEthSignature** | **string** | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature | 
 **xImxEthTimestamp** | **string** | Unix Epoc timestamp | 
 **xImxEthAddress** | **string** | Wallet Address that signed the signature | 
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **collectionAddress** | **string** | Collection address | 

### Return type

[**GetMetadataRefreshes**](GetMetadataRefreshes.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataRefreshErrors

> GetMetadataRefreshErrorsResponse GetMetadataRefreshErrors(ctx, refreshId).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).PageSize(pageSize).Cursor(cursor).Execute()





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
    refreshId := "refreshId_example" // string | The metadata refresh ID
    xImxEthSignature := "xImxEthSignature_example" // string | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
    xImxEthTimestamp := "xImxEthTimestamp_example" // string | Unix Epoc timestamp
    xImxEthAddress := "xImxEthAddress_example" // string | Wallet Address that signed the signature
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataRefreshesApi.GetMetadataRefreshErrors(context.Background(), refreshId).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).PageSize(pageSize).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataRefreshesApi.GetMetadataRefreshErrors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataRefreshErrors`: GetMetadataRefreshErrorsResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataRefreshesApi.GetMetadataRefreshErrors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refreshId** | **string** | The metadata refresh ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRefreshErrorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xImxEthSignature** | **string** | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature | 
 **xImxEthTimestamp** | **string** | Unix Epoc timestamp | 
 **xImxEthAddress** | **string** | Wallet Address that signed the signature | 
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 

### Return type

[**GetMetadataRefreshErrorsResponse**](GetMetadataRefreshErrorsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataRefreshResults

> GetMetadataRefreshResponse GetMetadataRefreshResults(ctx, refreshId).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).Execute()





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
    refreshId := "refreshId_example" // string | The metadata refresh ID
    xImxEthSignature := "xImxEthSignature_example" // string | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
    xImxEthTimestamp := "xImxEthTimestamp_example" // string | Unix Epoc timestamp
    xImxEthAddress := "xImxEthAddress_example" // string | Wallet Address that signed the signature

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataRefreshesApi.GetMetadataRefreshResults(context.Background(), refreshId).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataRefreshesApi.GetMetadataRefreshResults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataRefreshResults`: GetMetadataRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataRefreshesApi.GetMetadataRefreshResults`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**refreshId** | **string** | The metadata refresh ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataRefreshResultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xImxEthSignature** | **string** | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature | 
 **xImxEthTimestamp** | **string** | Unix Epoc timestamp | 
 **xImxEthAddress** | **string** | Wallet Address that signed the signature | 

### Return type

[**GetMetadataRefreshResponse**](GetMetadataRefreshResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestAMetadataRefresh

> CreateMetadataRefreshResponse RequestAMetadataRefresh(ctx).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).CreateMetadataRefreshRequest(createMetadataRefreshRequest).Execute()





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
    xImxEthSignature := "xImxEthSignature_example" // string | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature
    xImxEthTimestamp := "xImxEthTimestamp_example" // string | Unix Epoc timestamp
    xImxEthAddress := "xImxEthAddress_example" // string | Wallet Address that signed the signature
    createMetadataRefreshRequest := *openapiclient.NewCreateMetadataRefreshRequest("CollectionAddress_example", []string{"TokenIds_example"}) // CreateMetadataRefreshRequest | Create metadata refresh request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataRefreshesApi.RequestAMetadataRefresh(context.Background()).XImxEthSignature(xImxEthSignature).XImxEthTimestamp(xImxEthTimestamp).XImxEthAddress(xImxEthAddress).CreateMetadataRefreshRequest(createMetadataRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataRefreshesApi.RequestAMetadataRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestAMetadataRefresh`: CreateMetadataRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataRefreshesApi.RequestAMetadataRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestAMetadataRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xImxEthSignature** | **string** | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/docs/generate-imx-signature | 
 **xImxEthTimestamp** | **string** | Unix Epoc timestamp | 
 **xImxEthAddress** | **string** | Wallet Address that signed the signature | 
 **createMetadataRefreshRequest** | [**CreateMetadataRefreshRequest**](CreateMetadataRefreshRequest.md) | Create metadata refresh request | 

### Return type

[**CreateMetadataRefreshResponse**](CreateMetadataRefreshResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

