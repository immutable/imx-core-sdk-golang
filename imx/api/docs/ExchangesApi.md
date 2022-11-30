# \ExchangesApi

All URIs are relative to *https://api.sandbox.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExchange**](ExchangesApi.md#CreateExchange) | **Post** /v3/exchanges | Create an exchange with URL
[**CreateExchangeTransfer**](ExchangesApi.md#CreateExchangeTransfer) | **Post** /v2/exchanges/{id}/transfers | Accepts a transfer request and updates the transfer_status after processing transfer (offramp)
[**GetExchange**](ExchangesApi.md#GetExchange) | **Get** /v3/exchanges/{id} | Get an exchange by ID
[**GetExchangeSignableTransfer**](ExchangesApi.md#GetExchangeSignableTransfer) | **Post** /v2/exchanges/{id}/signable-transfer-details | Send a request for signable-transfer-details (offramp)
[**GetExchanges**](ExchangesApi.md#GetExchanges) | **Get** /v3/exchanges | Returns a list of exchanges based on the request



## CreateExchange

> ExchangeCreateExchangeAndURLResponse CreateExchange(ctx).CreateExchangeAPIRequest(createExchangeAPIRequest).Execute()

Create an exchange with URL



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
    createExchangeAPIRequest := *openapiclient.NewCreateExchangeAndURLAPIRequest() // CreateExchangeAndURLAPIRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.CreateExchange(context.Background()).CreateExchangeAPIRequest(createExchangeAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.CreateExchange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExchange`: ExchangeCreateExchangeAndURLResponse
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.CreateExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createExchangeAPIRequest** | [**CreateExchangeAndURLAPIRequest**](CreateExchangeAndURLAPIRequest.md) | req | 

### Return type

[**ExchangeCreateExchangeAndURLResponse**](ExchangeCreateExchangeAndURLResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateExchangeTransfer

> CreateTransferResponseV1 CreateExchangeTransfer(ctx, id).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).CreateTransferRequest(createTransferRequest).Execute()

Accepts a transfer request and updates the transfer_status after processing transfer (offramp)



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
    id := "id_example" // string | Transaction ID
    xImxEthAddress := "xImxEthAddress_example" // string | eth address
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature
    createTransferRequest := *openapiclient.NewCreateTransferRequestV1("Amount_example", "AssetId_example", int32(123), int32(123), "ReceiverStarkKey_example", int32(123), "SenderStarkKey_example", int32(123), "StarkSignature_example") // CreateTransferRequestV1 | Create a transfer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.CreateExchangeTransfer(context.Background(), id).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).CreateTransferRequest(createTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.CreateExchangeTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateExchangeTransfer`: CreateTransferResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.CreateExchangeTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExchangeTransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 
 **createTransferRequest** | [**CreateTransferRequestV1**](CreateTransferRequestV1.md) | Create a transfer | 

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


## GetExchange

> Exchange GetExchange(ctx, id).Execute()

Get an exchange by ID



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
    id := "id_example" // string | Exchange ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.GetExchange(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchange`: Exchange
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchange`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Exchange ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Exchange**](Exchange.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetExchangeSignableTransfer

> GetSignableTransferResponseV1 GetExchangeSignableTransfer(ctx, id).GetSignableTransferRequest(getSignableTransferRequest).Execute()

Send a request for signable-transfer-details (offramp)



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
    id := "id_example" // string | Transaction ID
    getSignableTransferRequest := *openapiclient.NewGetSignableTransferRequestV1("Amount_example", "Receiver_example", "Sender_example", *openapiclient.NewSignableToken()) // GetSignableTransferRequestV1 | get details of signable transfer

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.GetExchangeSignableTransfer(context.Background(), id).GetSignableTransferRequest(getSignableTransferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchangeSignableTransfer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchangeSignableTransfer`: GetSignableTransferResponseV1
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchangeSignableTransfer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Transaction ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangeSignableTransferRequest struct via the builder pattern


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


## GetExchanges

> GetTransactionsResponse GetExchanges(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Id(id).WalletAddress(walletAddress).Status(status).Provider(provider).TransferId(transferId).Execute()

Returns a list of exchanges based on the request



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
    direction := "direction_example" // string | Direction to sort (optional)
    id := int32(56) // int32 | Transaction ID (optional)
    walletAddress := "walletAddress_example" // string | Ethereum address of the user who created transaction (optional)
    status := "status_example" // string | Transaction status (optional)
    provider := "provider_example" // string | Provider name (optional)
    transferId := "transferId_example" // string | Transfer ID (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ExchangesApi.GetExchanges(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Id(id).WalletAddress(walletAddress).Status(status).Provider(provider).TransferId(transferId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ExchangesApi.GetExchanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetExchanges`: GetTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ExchangesApi.GetExchanges`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetExchangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort | 
 **id** | **int32** | Transaction ID | 
 **walletAddress** | **string** | Ethereum address of the user who created transaction | 
 **status** | **string** | Transaction status | 
 **provider** | **string** | Provider name | 
 **transferId** | **string** | Transfer ID | 

### Return type

[**GetTransactionsResponse**](GetTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

