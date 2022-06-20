# \TradesApi

All URIs are relative to *https://api.ropsten.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTrade**](TradesApi.md#CreateTrade) | **Post** /v1/trades | Create a Trade between two parties
[**GetSignableTrade**](TradesApi.md#GetSignableTrade) | **Post** /v3/signable-trade-details | Get details a signable trade V3
[**GetTrade**](TradesApi.md#GetTrade) | **Get** /v1/trades/{id} | Get details of a trade with the given ID
[**ListTrades**](TradesApi.md#ListTrades) | **Get** /v1/trades | Get a list of trades



## CreateTrade

> CreateTradeResponse CreateTrade(ctx).CreateTradeRequest(createTradeRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()

Create a Trade between two parties



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
    createTradeRequest := *openapiclient.NewCreateTradeRequestV1("AmountBuy_example", "AmountSell_example", "AssetIdBuy_example", "AssetIdSell_example", int32(123), int32(123), int32(123), "StarkKey_example", "StarkSignature_example", int32(123), int32(123)) // CreateTradeRequestV1 | create a trade
    xImxEthAddress := "xImxEthAddress_example" // string | eth address (optional)
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TradesApi.CreateTrade(context.Background()).CreateTradeRequest(createTradeRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradesApi.CreateTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTrade`: CreateTradeResponse
    fmt.Fprintf(os.Stdout, "Response from `TradesApi.CreateTrade`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTradeRequest** | [**CreateTradeRequestV1**](CreateTradeRequestV1.md) | create a trade | 
 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 

### Return type

[**CreateTradeResponse**](CreateTradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableTrade

> GetSignableTradeResponse GetSignableTrade(ctx).GetSignableTradeRequest(getSignableTradeRequest).Execute()

Get details a signable trade V3



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
    getSignableTradeRequest := *openapiclient.NewGetSignableTradeRequest(int32(123), "User_example") // GetSignableTradeRequest | get a signable trade

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TradesApi.GetSignableTrade(context.Background()).GetSignableTradeRequest(getSignableTradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradesApi.GetSignableTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableTrade`: GetSignableTradeResponse
    fmt.Fprintf(os.Stdout, "Response from `TradesApi.GetSignableTrade`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableTradeRequest** | [**GetSignableTradeRequest**](GetSignableTradeRequest.md) | get a signable trade | 

### Return type

[**GetSignableTradeResponse**](GetSignableTradeResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrade

> Trade GetTrade(ctx, id).Execute()

Get details of a trade with the given ID



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
    id := "id_example" // string | Trade ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TradesApi.GetTrade(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradesApi.GetTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrade`: Trade
    fmt.Fprintf(os.Stdout, "Response from `TradesApi.GetTrade`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Trade ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Trade**](Trade.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrades

> ListTradesResponse ListTrades(ctx).PartyATokenType(partyATokenType).PartyATokenAddress(partyATokenAddress).PartyATokenId(partyATokenId).PartyBTokenType(partyBTokenType).PartyBTokenAddress(partyBTokenAddress).PartyBTokenId(partyBTokenId).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).Execute()

Get a list of trades



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
    partyATokenType := "partyATokenType_example" // string | Party A's sell token type (optional)
    partyATokenAddress := "partyATokenAddress_example" // string | Party A's sell token address (optional)
    partyATokenId := "partyATokenId_example" // string | Party A's sell token id (optional)
    partyBTokenType := "partyBTokenType_example" // string | Party B's sell token type (optional)
    partyBTokenAddress := "partyBTokenAddress_example" // string | Party B's sell token address (optional)
    partyBTokenId := "partyBTokenId_example" // string | Party B's sell token id (optional)
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)
    minTimestamp := "minTimestamp_example" // string | Minimum timestamp for this trade, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    maxTimestamp := "maxTimestamp_example" // string | Maximum timestamp for this trade, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TradesApi.ListTrades(context.Background()).PartyATokenType(partyATokenType).PartyATokenAddress(partyATokenAddress).PartyATokenId(partyATokenId).PartyBTokenType(partyBTokenType).PartyBTokenAddress(partyBTokenAddress).PartyBTokenId(partyBTokenId).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradesApi.ListTrades``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrades`: ListTradesResponse
    fmt.Fprintf(os.Stdout, "Response from `TradesApi.ListTrades`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTradesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **partyATokenType** | **string** | Party A&#39;s sell token type | 
 **partyATokenAddress** | **string** | Party A&#39;s sell token address | 
 **partyATokenId** | **string** | Party A&#39;s sell token id | 
 **partyBTokenType** | **string** | Party B&#39;s sell token type | 
 **partyBTokenAddress** | **string** | Party B&#39;s sell token address | 
 **partyBTokenId** | **string** | Party B&#39;s sell token id | 
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **minTimestamp** | **string** | Minimum timestamp for this trade, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **maxTimestamp** | **string** | Maximum timestamp for this trade, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 

### Return type

[**ListTradesResponse**](ListTradesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

