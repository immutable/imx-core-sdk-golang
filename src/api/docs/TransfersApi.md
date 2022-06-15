# \TransfersApi

All URIs are relative to *https://api.dev.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTransfer**](TransfersApi.md#CreateTransfer) | **Post** /v2/transfers | Creates a transfer of multiple tokens between two parties
[**CreateTransferV1**](TransfersApi.md#CreateTransferV1) | **Post** /v1/transfers | Creates a transfer of tokens between two parties
[**GetSignableTransfer**](TransfersApi.md#GetSignableTransfer) | **Post** /v2/signable-transfer-details | Gets bulk details of a signable transfer
[**GetSignableTransferV1**](TransfersApi.md#GetSignableTransferV1) | **Post** /v1/signable-transfer-details | Gets details of a signable transfer
[**GetTransfer**](TransfersApi.md#GetTransfer) | **Get** /v1/transfers/{id} | Get details of a transfer with the given ID
[**ListTransfers**](TransfersApi.md#ListTransfers) | **Get** /v1/transfers | Get a list of transfers



## CreateTransfer

> CreateTransferResponse CreateTransfer(ctx).CreateTransferRequestV2(createTransferRequestV2).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()

Creates a transfer of multiple tokens between two parties



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
    createTransferRequestV2 := *openapiclient.NewCreateTransferRequest([]openapiclient.TransferRequest{*openapiclient.NewTransferRequest("Amount_example", "AssetId_example", int32(123), int32(123), "ReceiverStarkKey_example", int32(123), int32(123), "StarkSignature_example")}, "SenderStarkKey_example") // CreateTransferRequest | Create transfer
    xImxEthAddress := "xImxEthAddress_example" // string | eth address (optional)
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.CreateTransfer(context.Background()).CreateTransferRequestV2(createTransferRequestV2).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.CreateTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransfer`: CreateTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.CreateTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransferRequestV2** | [**CreateTransferRequest**](CreateTransferRequest.md) | Create transfer | 
 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 

### Return type

[**CreateTransferResponse**](CreateTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTransferV1

> CreateTransferResponseV1 CreateTransferV1(ctx).CreateTransferRequest(createTransferRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()

Creates a transfer of tokens between two parties



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
    createTransferRequest := *openapiclient.NewCreateTransferRequestV1("Amount_example", "AssetId_example", int32(123), int32(123), "ReceiverStarkKey_example", int32(123), "SenderStarkKey_example", int32(123), "StarkSignature_example") // CreateTransferRequestV1 | Create transfer
    xImxEthAddress := "xImxEthAddress_example" // string | eth address (optional)
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.CreateTransferV1(context.Background()).CreateTransferRequest(createTransferRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.CreateTransferV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTransferV1`: CreateTransferResponseV1
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.CreateTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTransferRequest** | [**CreateTransferRequestV1**](CreateTransferRequestV1.md) | Create transfer | 
 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 

### Return type

[**CreateTransferResponseV1**](CreateTransferResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableTransfer

> GetSignableTransferResponse GetSignableTransfer(ctx).GetSignableTransferRequestV2(getSignableTransferRequestV2).Execute()

Gets bulk details of a signable transfer



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
    getSignableTransferRequestV2 := *openapiclient.NewGetSignableTransferRequest("SenderEtherKey_example", []openapiclient.SignableTransferDetails{*openapiclient.NewSignableTransferDetails("Amount_example", "Receiver_example", *openapiclient.NewSignableToken())}) // GetSignableTransferRequest | get details of signable transfer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.GetSignableTransfer(context.Background()).GetSignableTransferRequestV2(getSignableTransferRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.GetSignableTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableTransfer`: GetSignableTransferResponse
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.GetSignableTransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableTransferRequestV2** | [**GetSignableTransferRequest**](GetSignableTransferRequest.md) | get details of signable transfer | 

### Return type

[**GetSignableTransferResponse**](GetSignableTransferResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableTransferV1

> GetSignableTransferResponseV1 GetSignableTransferV1(ctx).GetSignableTransferRequest(getSignableTransferRequest).Execute()

Gets details of a signable transfer



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
    getSignableTransferRequest := *openapiclient.NewGetSignableTransferRequestV1("Amount_example", "Receiver_example", "Sender_example", *openapiclient.NewSignableToken()) // GetSignableTransferRequestV1 | get details of signable transfer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.GetSignableTransferV1(context.Background()).GetSignableTransferRequest(getSignableTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.GetSignableTransferV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableTransferV1`: GetSignableTransferResponseV1
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.GetSignableTransferV1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableTransferV1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableTransferRequest** | [**GetSignableTransferRequestV1**](GetSignableTransferRequestV1.md) | get details of signable transfer | 

### Return type

[**GetSignableTransferResponseV1**](GetSignableTransferResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransfer

> Transfer GetTransfer(ctx, id).Execute()

Get details of a transfer with the given ID



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
    id := "id_example" // string | Transfer ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.GetTransfer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.GetTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransfer`: Transfer
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.GetTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transfer ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Transfer**](Transfer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTransfers

> ListTransfersResponse ListTransfers(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Receiver(receiver).Status(status).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenAddress(tokenAddress).TokenName(tokenName).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()

Get a list of transfers



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
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)
    user := "user_example" // string | Ethereum address of the user who submitted this transfer (optional)
    receiver := "receiver_example" // string | Ethereum address of the user who received this transfer (optional)
    status := "status_example" // string | Status of this transfer (optional)
    minTimestamp := "minTimestamp_example" // string | Minimum timestamp for this transfer, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    maxTimestamp := "maxTimestamp_example" // string | Maximum timestamp for this transfer, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    tokenType := "tokenType_example" // string | Token type of the transferred asset (optional)
    tokenId := "tokenId_example" // string | ERC721 Token ID of the minted asset (optional)
    assetId := "assetId_example" // string | Internal IMX ID of the minted asset (optional)
    tokenAddress := "tokenAddress_example" // string | Token address of the transferred asset (optional)
    tokenName := "tokenName_example" // string | Token name of the transferred asset (optional)
    minQuantity := "minQuantity_example" // string | Max quantity for the transferred asset (optional)
    maxQuantity := "maxQuantity_example" // string | Max quantity for the transferred asset (optional)
    metadata := "metadata_example" // string | JSON-encoded metadata filters for the transferred asset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TransfersApi.ListTransfers(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Receiver(receiver).Status(status).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenAddress(tokenAddress).TokenName(tokenName).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransfersApi.ListTransfers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTransfers`: ListTransfersResponse
    fmt.Fprintf(os.Stdout, "Response from `TransfersApi.ListTransfers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTransfersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **user** | **string** | Ethereum address of the user who submitted this transfer | 
 **receiver** | **string** | Ethereum address of the user who received this transfer | 
 **status** | **string** | Status of this transfer | 
 **minTimestamp** | **string** | Minimum timestamp for this transfer, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **maxTimestamp** | **string** | Maximum timestamp for this transfer, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **tokenType** | **string** | Token type of the transferred asset | 
 **tokenId** | **string** | ERC721 Token ID of the minted asset | 
 **assetId** | **string** | Internal IMX ID of the minted asset | 
 **tokenAddress** | **string** | Token address of the transferred asset | 
 **tokenName** | **string** | Token name of the transferred asset | 
 **minQuantity** | **string** | Max quantity for the transferred asset | 
 **maxQuantity** | **string** | Max quantity for the transferred asset | 
 **metadata** | **string** | JSON-encoded metadata filters for the transferred asset | 

### Return type

[**ListTransfersResponse**](ListTransfersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

