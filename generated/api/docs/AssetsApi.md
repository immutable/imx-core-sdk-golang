# \AssetsApi

All URIs are relative to *https://api.ropsten.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAsset**](AssetsApi.md#GetAsset) | **Get** /v1/assets/{token_address}/{token_id} | Get details of an asset
[**ListAssets**](AssetsApi.md#ListAssets) | **Get** /v1/assets | Get a list of assets



## GetAsset

> Asset GetAsset(ctx, tokenAddress, tokenId).IncludeFees(includeFees).Execute()

Get details of an asset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/generated/api"
    
)

func main() {
    tokenAddress := "tokenAddress_example" // string | Address of the ERC721 contract
    tokenId := "tokenId_example" // string | Either ERC721 token ID or internal IMX ID
    includeFees := true // bool | Set flag to include fees associated with the asset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.GetAsset(context.Background(), tokenAddress, tokenId).IncludeFees(includeFees).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.GetAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsset`: Asset
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.GetAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenAddress** | **string** | Address of the ERC721 contract | 
**tokenId** | **string** | Either ERC721 token ID or internal IMX ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeFees** | **bool** | Set flag to include fees associated with the asset | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssets

> ListAssetsResponse ListAssets(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).Name(name).Metadata(metadata).SellOrders(sellOrders).BuyOrders(buyOrders).IncludeFees(includeFees).Collection(collection).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).AuxiliaryFeePercentages(auxiliaryFeePercentages).AuxiliaryFeeRecipients(auxiliaryFeeRecipients).Execute()

Get a list of assets



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/generated/api"
    
)

func main() {
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)
    user := "user_example" // string | Ethereum address of the user who owns these assets (optional)
    status := "status_example" // string | Status of these assets (optional)
    name := "name_example" // string | Name of the asset to search (optional)
    metadata := "metadata_example" // string | JSON-encoded metadata filters for these asset. Example: {'proto':['1147'],'quality':['Meteorite']} (optional)
    sellOrders := true // bool | Set flag to true to fetch an array of sell order details with accepted status associated with the asset (optional)
    buyOrders := true // bool | Set flag to true to fetch an array of buy order details  with accepted status associated with the asset (optional)
    includeFees := true // bool | Set flag to include fees associated with the asset (optional)
    collection := "collection_example" // string | Collection contract address (optional)
    updatedMinTimestamp := "updatedMinTimestamp_example" // string | Minimum timestamp for when these assets were last updated, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    updatedMaxTimestamp := "updatedMaxTimestamp_example" // string | Maximum timestamp for when these assets were last updated, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    auxiliaryFeePercentages := "auxiliaryFeePercentages_example" // string | Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients (optional)
    auxiliaryFeeRecipients := "auxiliaryFeeRecipients_example" // string | Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AssetsApi.ListAssets(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).Name(name).Metadata(metadata).SellOrders(sellOrders).BuyOrders(buyOrders).IncludeFees(includeFees).Collection(collection).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).AuxiliaryFeePercentages(auxiliaryFeePercentages).AuxiliaryFeeRecipients(auxiliaryFeeRecipients).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.ListAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssets`: ListAssetsResponse
    fmt.Fprintf(os.Stdout, "Response from `AssetsApi.ListAssets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **user** | **string** | Ethereum address of the user who owns these assets | 
 **status** | **string** | Status of these assets | 
 **name** | **string** | Name of the asset to search | 
 **metadata** | **string** | JSON-encoded metadata filters for these asset. Example: {&#39;proto&#39;:[&#39;1147&#39;],&#39;quality&#39;:[&#39;Meteorite&#39;]} | 
 **sellOrders** | **bool** | Set flag to true to fetch an array of sell order details with accepted status associated with the asset | 
 **buyOrders** | **bool** | Set flag to true to fetch an array of buy order details  with accepted status associated with the asset | 
 **includeFees** | **bool** | Set flag to include fees associated with the asset | 
 **collection** | **string** | Collection contract address | 
 **updatedMinTimestamp** | **string** | Minimum timestamp for when these assets were last updated, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **updatedMaxTimestamp** | **string** | Maximum timestamp for when these assets were last updated, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **auxiliaryFeePercentages** | **string** | Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients | 
 **auxiliaryFeeRecipients** | **string** | Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages | 

### Return type

[**ListAssetsResponse**](ListAssetsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

