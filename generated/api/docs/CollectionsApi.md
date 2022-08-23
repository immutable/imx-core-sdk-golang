# \CollectionsApi

All URIs are relative to *https://api.ropsten.x.github.com/immutable/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCollection**](CollectionsApi.md#CreateCollection) | **Post** /v1/collections | Create collection
[**GetCollection**](CollectionsApi.md#GetCollection) | **Get** /v1/collections/{address} | Get details of a collection at the given address
[**ListCollectionFilters**](CollectionsApi.md#ListCollectionFilters) | **Get** /v1/collections/{address}/filters | Get a list of collection filters
[**ListCollections**](CollectionsApi.md#ListCollections) | **Get** /v1/collections | Get a list of collections
[**UpdateCollection**](CollectionsApi.md#UpdateCollection) | **Patch** /v1/collections/{address} | Update collection



## CreateCollection

> Collection CreateCollection(ctx).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).CreateCollectionRequest(createCollectionRequest).Execute()

Create collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/generated/api"
    
)

func main() {
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/immutable/docs/generate-imx-signature
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp
    createCollectionRequest := *openapiclient.NewCreateCollectionRequest("ContractAddress_example", "Name_example", "OwnerPublicKey_example", int32(123)) // CreateCollectionRequest | create a collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.CreateCollection(context.Background()).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).CreateCollectionRequest(createCollectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.CreateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMXSignature** | **string** | String created by signing wallet address and timestamp. See https://docs.x.immutable.com/immutable/docs/generate-imx-signature | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 
 **createCollectionRequest** | [**CreateCollectionRequest**](CreateCollectionRequest.md) | create a collection | 

### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollection

> Collection GetCollection(ctx, address).Execute()

Get details of a collection at the given address



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/generated/api"
    
)

func main() {
    address := "address_example" // string | Collection contract address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.GetCollection(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.GetCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Collection contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionFilters

> CollectionFilter ListCollectionFilters(ctx, address).PageSize(pageSize).NextPageToken(nextPageToken).Execute()

Get a list of collection filters



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/generated/api"
    
)

func main() {
    address := "address_example" // string | Collection contract address
    pageSize := int32(56) // int32 | Page size of the result (optional)
    nextPageToken := "nextPageToken_example" // string | Next page token (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.ListCollectionFilters(context.Background(), address).PageSize(pageSize).NextPageToken(nextPageToken).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.ListCollectionFilters``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollectionFilters`: CollectionFilter
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.ListCollectionFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Collection contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | Page size of the result | 
 **nextPageToken** | **string** | Next page token | 

### Return type

[**CollectionFilter**](CollectionFilter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> ListCollectionsResponse ListCollections(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Blacklist(blacklist).Whitelist(whitelist).Keyword(keyword).Execute()

Get a list of collections



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/generated/api"
    
)

func main() {
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)
    blacklist := "blacklist_example" // string | List of collections not to be included, separated by commas (optional)
    whitelist := "whitelist_example" // string | List of collections to be included, separated by commas (optional)
    keyword := "keyword_example" // string | Keyword to search in collection name and description (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.ListCollections(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Blacklist(blacklist).Whitelist(whitelist).Keyword(keyword).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.ListCollections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCollections`: ListCollectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.ListCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **blacklist** | **string** | List of collections not to be included, separated by commas | 
 **whitelist** | **string** | List of collections to be included, separated by commas | 
 **keyword** | **string** | Keyword to search in collection name and description | 

### Return type

[**ListCollectionsResponse**](ListCollectionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCollection

> Collection UpdateCollection(ctx, address).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).UpdateCollectionRequest(updateCollectionRequest).Execute()

Update collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/generated/api"
    
)

func main() {
    address := "address_example" // string | Collection contract address
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp
    updateCollectionRequest := *openapiclient.NewUpdateCollectionRequest() // UpdateCollectionRequest | update a collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CollectionsApi.UpdateCollection(context.Background(), address).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).UpdateCollectionRequest(updateCollectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CollectionsApi.UpdateCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCollection`: Collection
    fmt.Fprintf(os.Stdout, "Response from `CollectionsApi.UpdateCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Collection contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMXSignature** | **string** | String created by signing wallet address and timestamp | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 
 **updateCollectionRequest** | [**UpdateCollectionRequest**](UpdateCollectionRequest.md) | update a collection | 

### Return type

[**Collection**](Collection.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

