# \WithdrawalsApi

All URIs are relative to *https://api.ropsten.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWithdrawal**](WithdrawalsApi.md#CreateWithdrawal) | **Post** /v1/withdrawals | Creates a withdrawal of a token
[**GetSignableWithdrawal**](WithdrawalsApi.md#GetSignableWithdrawal) | **Post** /v1/signable-withdrawal-details | Gets details of a signable withdrawal
[**GetWithdrawal**](WithdrawalsApi.md#GetWithdrawal) | **Get** /v1/withdrawals/{id} | Gets details of withdrawal with the given ID
[**ListWithdrawals**](WithdrawalsApi.md#ListWithdrawals) | **Get** /v1/withdrawals | Get a list of withdrawals



## CreateWithdrawal

> CreateWithdrawalResponse CreateWithdrawal(ctx).CreateWithdrawalRequest(createWithdrawalRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()

Creates a withdrawal of a token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    createWithdrawalRequest := *openapiclient.NewCreateWithdrawalRequest("Amount_example", "AssetId_example", int32(123), "StarkKey_example", "StarkSignature_example", int32(123)) // CreateWithdrawalRequest | create a withdrawal
    xImxEthAddress := "xImxEthAddress_example" // string | eth address (optional)
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WithdrawalsApi.CreateWithdrawal(context.Background()).CreateWithdrawalRequest(createWithdrawalRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalsApi.CreateWithdrawal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWithdrawal`: CreateWithdrawalResponse
    fmt.Fprintf(os.Stdout, "Response from `WithdrawalsApi.CreateWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWithdrawalRequest** | [**CreateWithdrawalRequest**](CreateWithdrawalRequest.md) | create a withdrawal | 
 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 

### Return type

[**CreateWithdrawalResponse**](CreateWithdrawalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableWithdrawal

> GetSignableWithdrawalResponse GetSignableWithdrawal(ctx).GetSignableWithdrawalRequest(getSignableWithdrawalRequest).Execute()

Gets details of a signable withdrawal



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    getSignableWithdrawalRequest := *openapiclient.NewGetSignableWithdrawalRequest("Amount_example", *openapiclient.NewSignableToken(), "User_example") // GetSignableWithdrawalRequest | get details of signable withdrawal

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WithdrawalsApi.GetSignableWithdrawal(context.Background()).GetSignableWithdrawalRequest(getSignableWithdrawalRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalsApi.GetSignableWithdrawal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableWithdrawal`: GetSignableWithdrawalResponse
    fmt.Fprintf(os.Stdout, "Response from `WithdrawalsApi.GetSignableWithdrawal`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableWithdrawalRequest** | [**GetSignableWithdrawalRequest**](GetSignableWithdrawalRequest.md) | get details of signable withdrawal | 

### Return type

[**GetSignableWithdrawalResponse**](GetSignableWithdrawalResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWithdrawal

> Withdrawal GetWithdrawal(ctx, id).Execute()

Gets details of withdrawal with the given ID



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Withdrawal ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WithdrawalsApi.GetWithdrawal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalsApi.GetWithdrawal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWithdrawal`: Withdrawal
    fmt.Fprintf(os.Stdout, "Response from `WithdrawalsApi.GetWithdrawal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Withdrawal ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWithdrawalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Withdrawal**](Withdrawal.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWithdrawals

> ListWithdrawalsResponse ListWithdrawals(ctx).WithdrawnToWallet(withdrawnToWallet).RollupStatus(rollupStatus).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenAddress(tokenAddress).TokenName(tokenName).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()

Get a list of withdrawals



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    withdrawnToWallet := true // bool | Withdrawal has been transferred to user's Layer 1 wallet (optional)
    rollupStatus := "rollupStatus_example" // string | Status of the on-chain batch confirmation for this withdrawal (optional)
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)
    user := "user_example" // string | Ethereum address of the user who submitted this withdrawal (optional)
    status := "status_example" // string | Status of this withdrawal (optional)
    minTimestamp := "minTimestamp_example" // string | Minimum timestamp for this deposit, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    maxTimestamp := "maxTimestamp_example" // string | Maximum timestamp for this deposit, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    tokenType := "tokenType_example" // string | Token type of the withdrawn asset (optional)
    tokenId := "tokenId_example" // string | ERC721 Token ID of the minted asset (optional)
    assetId := "assetId_example" // string | Internal IMX ID of the minted asset (optional)
    tokenAddress := "tokenAddress_example" // string | Token address of the withdrawn asset (optional)
    tokenName := "tokenName_example" // string | Token name of the withdrawn asset (optional)
    minQuantity := "minQuantity_example" // string | Min quantity for the withdrawn asset (optional)
    maxQuantity := "maxQuantity_example" // string | Max quantity for the withdrawn asset (optional)
    metadata := "metadata_example" // string | JSON-encoded metadata filters for the withdrawn asset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WithdrawalsApi.ListWithdrawals(context.Background()).WithdrawnToWallet(withdrawnToWallet).RollupStatus(rollupStatus).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenAddress(tokenAddress).TokenName(tokenName).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WithdrawalsApi.ListWithdrawals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWithdrawals`: ListWithdrawalsResponse
    fmt.Fprintf(os.Stdout, "Response from `WithdrawalsApi.ListWithdrawals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWithdrawalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **withdrawnToWallet** | **bool** | Withdrawal has been transferred to user&#39;s Layer 1 wallet | 
 **rollupStatus** | **string** | Status of the on-chain batch confirmation for this withdrawal | 
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **user** | **string** | Ethereum address of the user who submitted this withdrawal | 
 **status** | **string** | Status of this withdrawal | 
 **minTimestamp** | **string** | Minimum timestamp for this deposit, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **maxTimestamp** | **string** | Maximum timestamp for this deposit, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **tokenType** | **string** | Token type of the withdrawn asset | 
 **tokenId** | **string** | ERC721 Token ID of the minted asset | 
 **assetId** | **string** | Internal IMX ID of the minted asset | 
 **tokenAddress** | **string** | Token address of the withdrawn asset | 
 **tokenName** | **string** | Token name of the withdrawn asset | 
 **minQuantity** | **string** | Min quantity for the withdrawn asset | 
 **maxQuantity** | **string** | Max quantity for the withdrawn asset | 
 **metadata** | **string** | JSON-encoded metadata filters for the withdrawn asset | 

### Return type

[**ListWithdrawalsResponse**](ListWithdrawalsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

