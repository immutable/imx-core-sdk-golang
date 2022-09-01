# \DepositsApi

All URIs are relative to *https://api.ropsten.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeposit**](DepositsApi.md#GetDeposit) | **Get** /v1/deposits/{id} | Get details of a deposit with the given ID
[**GetSignableDeposit**](DepositsApi.md#GetSignableDeposit) | **Post** /v1/signable-deposit-details | Gets details of a signable deposit
[**ListDeposits**](DepositsApi.md#ListDeposits) | **Get** /v1/deposits | Get a list of deposits



## GetDeposit

> Deposit GetDeposit(ctx, id).Execute()

Get details of a deposit with the given ID



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
    id := "id_example" // string | Deposit ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepositsApi.GetDeposit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepositsApi.GetDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeposit`: Deposit
    fmt.Fprintf(os.Stdout, "Response from `DepositsApi.GetDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Deposit ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Deposit**](Deposit.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableDeposit

> GetSignableDepositResponse GetSignableDeposit(ctx).GetSignableDepositRequest(getSignableDepositRequest).Execute()

Gets details of a signable deposit



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
    getSignableDepositRequest := *openapiclient.NewGetSignableDepositRequest("Amount_example", *openapiclient.NewSignableToken(), "User_example") // GetSignableDepositRequest | Get details of signable deposit

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepositsApi.GetSignableDeposit(context.Background()).GetSignableDepositRequest(getSignableDepositRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepositsApi.GetSignableDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableDeposit`: GetSignableDepositResponse
    fmt.Fprintf(os.Stdout, "Response from `DepositsApi.GetSignableDeposit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableDepositRequest** | [**GetSignableDepositRequest**](GetSignableDepositRequest.md) | Get details of signable deposit | 

### Return type

[**GetSignableDepositResponse**](GetSignableDepositResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeposits

> ListDepositsResponse ListDeposits(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenAddress(tokenAddress).TokenName(tokenName).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()

Get a list of deposits



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
    user := "user_example" // string | Ethereum address of the user who submitted this deposit (optional)
    status := "status_example" // string | Status of this deposit (optional)
    updatedMinTimestamp := "updatedMinTimestamp_example" // string | Minimum timestamp for this deposit, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    updatedMaxTimestamp := "updatedMaxTimestamp_example" // string | Maximum timestamp for this deposit, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    tokenType := "tokenType_example" // string | Token type of the deposited asset (optional)
    tokenId := "tokenId_example" // string | ERC721 Token ID of the minted asset (optional)
    assetId := "assetId_example" // string | Internal IMX ID of the minted asset (optional)
    tokenAddress := "tokenAddress_example" // string | Token address of the deposited asset (optional)
    tokenName := "tokenName_example" // string | Token name of the deposited asset (optional)
    minQuantity := "minQuantity_example" // string | Min quantity for the deposited asset (optional)
    maxQuantity := "maxQuantity_example" // string | Max quantity for the deposited asset (optional)
    metadata := "metadata_example" // string | JSON-encoded metadata filters for the deposited asset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DepositsApi.ListDeposits(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenAddress(tokenAddress).TokenName(tokenName).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DepositsApi.ListDeposits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeposits`: ListDepositsResponse
    fmt.Fprintf(os.Stdout, "Response from `DepositsApi.ListDeposits`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDepositsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **user** | **string** | Ethereum address of the user who submitted this deposit | 
 **status** | **string** | Status of this deposit | 
 **updatedMinTimestamp** | **string** | Minimum timestamp for this deposit, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **updatedMaxTimestamp** | **string** | Maximum timestamp for this deposit, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **tokenType** | **string** | Token type of the deposited asset | 
 **tokenId** | **string** | ERC721 Token ID of the minted asset | 
 **assetId** | **string** | Internal IMX ID of the minted asset | 
 **tokenAddress** | **string** | Token address of the deposited asset | 
 **tokenName** | **string** | Token name of the deposited asset | 
 **minQuantity** | **string** | Min quantity for the deposited asset | 
 **maxQuantity** | **string** | Max quantity for the deposited asset | 
 **metadata** | **string** | JSON-encoded metadata filters for the deposited asset | 

### Return type

[**ListDepositsResponse**](ListDepositsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

