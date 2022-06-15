# \OrdersApi

All URIs are relative to *https://api.dev.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrdersApi.md#CancelOrder) | **Delete** /v1/orders/{id} | cancel an order
[**CreateOrder**](OrdersApi.md#CreateOrder) | **Post** /v1/orders | Create an order
[**GetOrder**](OrdersApi.md#GetOrder) | **Get** /v1/orders/{id} | Get details of an order with the given ID
[**GetSignableCancelOrder**](OrdersApi.md#GetSignableCancelOrder) | **Post** /v1/signable-cancel-order-details | Get details a signable cancel order
[**GetSignableOrder**](OrdersApi.md#GetSignableOrder) | **Post** /v3/signable-order-details | Get details a signable order V3
[**ListOrders**](OrdersApi.md#ListOrders) | **Get** /v1/orders | Get a list of orders



## CancelOrder

> CancelOrderResponse CancelOrder(ctx, id).CancelOrderRequest(cancelOrderRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()

cancel an order



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
    id := "id_example" // string | Order ID to cancel
    cancelOrderRequest := *openapiclient.NewCancelOrderRequest(int32(123), "StarkSignature_example") // CancelOrderRequest | cancel an order
    xImxEthAddress := "xImxEthAddress_example" // string | eth address (optional)
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CancelOrder(context.Background(), id).CancelOrderRequest(cancelOrderRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelOrder`: CancelOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CancelOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Order ID to cancel | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cancelOrderRequest** | [**CancelOrderRequest**](CancelOrderRequest.md) | cancel an order | 
 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 

### Return type

[**CancelOrderResponse**](CancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateOrder

> CreateOrderResponse CreateOrder(ctx).CreateOrderRequest(createOrderRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()

Create an order



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
    createOrderRequest := *openapiclient.NewCreateOrderRequest("AmountBuy_example", "AmountSell_example", "AssetIdBuy_example", "AssetIdSell_example", int32(123), int32(123), "StarkKey_example", "StarkSignature_example", int32(123), int32(123)) // CreateOrderRequest | create an order
    xImxEthAddress := "xImxEthAddress_example" // string | eth address (optional)
    xImxEthSignature := "xImxEthSignature_example" // string | eth signature (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CreateOrder(context.Background()).CreateOrderRequest(createOrderRequest).XImxEthAddress(xImxEthAddress).XImxEthSignature(xImxEthSignature).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrder`: CreateOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createOrderRequest** | [**CreateOrderRequest**](CreateOrderRequest.md) | create an order | 
 **xImxEthAddress** | **string** | eth address | 
 **xImxEthSignature** | **string** | eth signature | 

### Return type

[**CreateOrderResponse**](CreateOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, id).IncludeFees(includeFees).AuxiliaryFeePercentages(auxiliaryFeePercentages).AuxiliaryFeeRecipients(auxiliaryFeeRecipients).Execute()

Get details of an order with the given ID



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
    id := "id_example" // string | Order ID
    includeFees := true // bool | Set flag to true to include fee body for the order (optional)
    auxiliaryFeePercentages := "auxiliaryFeePercentages_example" // string | Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients (optional)
    auxiliaryFeeRecipients := "auxiliaryFeeRecipients_example" // string | Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.GetOrder(context.Background(), id).IncludeFees(includeFees).AuxiliaryFeePercentages(auxiliaryFeePercentages).AuxiliaryFeeRecipients(auxiliaryFeeRecipients).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: Order
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Order ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **includeFees** | **bool** | Set flag to true to include fee body for the order | 
 **auxiliaryFeePercentages** | **string** | Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients | 
 **auxiliaryFeeRecipients** | **string** | Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages | 

### Return type

[**Order**](Order.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableCancelOrder

> GetSignableCancelOrderResponse GetSignableCancelOrder(ctx).GetSignableCancelOrderRequest(getSignableCancelOrderRequest).Execute()

Get details a signable cancel order



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
    getSignableCancelOrderRequest := *openapiclient.NewGetSignableCancelOrderRequest(int32(123)) // GetSignableCancelOrderRequest | get a signable cancel order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.GetSignableCancelOrder(context.Background()).GetSignableCancelOrderRequest(getSignableCancelOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetSignableCancelOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableCancelOrder`: GetSignableCancelOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetSignableCancelOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableCancelOrderRequest** | [**GetSignableCancelOrderRequest**](GetSignableCancelOrderRequest.md) | get a signable cancel order | 

### Return type

[**GetSignableCancelOrderResponse**](GetSignableCancelOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableOrder

> GetSignableOrderResponse GetSignableOrder(ctx).GetSignableOrderRequestV3(getSignableOrderRequestV3).Execute()

Get details a signable order V3



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
    getSignableOrderRequestV3 := *openapiclient.NewGetSignableOrderRequest("AmountBuy_example", "AmountSell_example", *openapiclient.NewSignableToken(), *openapiclient.NewSignableToken(), "User_example") // GetSignableOrderRequest | get a signable order

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.GetSignableOrder(context.Background()).GetSignableOrderRequestV3(getSignableOrderRequestV3).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetSignableOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableOrder`: GetSignableOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetSignableOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableOrderRequestV3** | [**GetSignableOrderRequest**](GetSignableOrderRequest.md) | get a signable order | 

### Return type

[**GetSignableOrderResponse**](GetSignableOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListOrders

> ListOrdersResponse ListOrders(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).BuyTokenType(buyTokenType).BuyTokenId(buyTokenId).BuyAssetId(buyAssetId).BuyTokenAddress(buyTokenAddress).BuyTokenName(buyTokenName).BuyMinQuantity(buyMinQuantity).BuyMaxQuantity(buyMaxQuantity).BuyMetadata(buyMetadata).SellTokenType(sellTokenType).SellTokenId(sellTokenId).SellAssetId(sellAssetId).SellTokenAddress(sellTokenAddress).SellTokenName(sellTokenName).SellMinQuantity(sellMinQuantity).SellMaxQuantity(sellMaxQuantity).SellMetadata(sellMetadata).AuxiliaryFeePercentages(auxiliaryFeePercentages).AuxiliaryFeeRecipients(auxiliaryFeeRecipients).Execute()

Get a list of orders



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
    user := "user_example" // string | Ethereum address of the user who submitted this order (optional)
    status := "status_example" // string | Status of this order (optional)
    minTimestamp := "minTimestamp_example" // string | Minimum created at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    maxTimestamp := "maxTimestamp_example" // string | Maximum created at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    updatedMinTimestamp := "updatedMinTimestamp_example" // string | Minimum updated at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    updatedMaxTimestamp := "updatedMaxTimestamp_example" // string | Maximum updated at timestamp for this order, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    buyTokenType := "buyTokenType_example" // string | Token type of the asset this order buys (optional)
    buyTokenId := "buyTokenId_example" // string | ERC721 Token ID of the asset this order buys (optional)
    buyAssetId := "buyAssetId_example" // string | Internal IMX ID of the asset this order buys (optional)
    buyTokenAddress := "buyTokenAddress_example" // string | Comma separated string of token addresses of the asset this order buys (optional)
    buyTokenName := "buyTokenName_example" // string | Token name of the asset this order buys (optional)
    buyMinQuantity := "buyMinQuantity_example" // string | Min quantity for the asset this order buys (optional)
    buyMaxQuantity := "buyMaxQuantity_example" // string | Max quantity for the asset this order buys (optional)
    buyMetadata := "buyMetadata_example" // string | JSON-encoded metadata filters for the asset this order buys (optional)
    sellTokenType := "sellTokenType_example" // string | Token type of the asset this order sells (optional)
    sellTokenId := "sellTokenId_example" // string | ERC721 Token ID of the asset this order sells (optional)
    sellAssetId := "sellAssetId_example" // string | Internal IMX ID of the asset this order sells (optional)
    sellTokenAddress := "sellTokenAddress_example" // string | Comma separated string of token addresses of the asset this order sells (optional)
    sellTokenName := "sellTokenName_example" // string | Token name of the asset this order sells (optional)
    sellMinQuantity := "sellMinQuantity_example" // string | Min quantity for the asset this order sells (optional)
    sellMaxQuantity := "sellMaxQuantity_example" // string | Max quantity for the asset this order sells (optional)
    sellMetadata := "sellMetadata_example" // string | JSON-encoded metadata filters for the asset this order sells (optional)
    auxiliaryFeePercentages := "auxiliaryFeePercentages_example" // string | Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients (optional)
    auxiliaryFeeRecipients := "auxiliaryFeeRecipients_example" // string | Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.ListOrders(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).MinTimestamp(minTimestamp).MaxTimestamp(maxTimestamp).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).BuyTokenType(buyTokenType).BuyTokenId(buyTokenId).BuyAssetId(buyAssetId).BuyTokenAddress(buyTokenAddress).BuyTokenName(buyTokenName).BuyMinQuantity(buyMinQuantity).BuyMaxQuantity(buyMaxQuantity).BuyMetadata(buyMetadata).SellTokenType(sellTokenType).SellTokenId(sellTokenId).SellAssetId(sellAssetId).SellTokenAddress(sellTokenAddress).SellTokenName(sellTokenName).SellMinQuantity(sellMinQuantity).SellMaxQuantity(sellMaxQuantity).SellMetadata(sellMetadata).AuxiliaryFeePercentages(auxiliaryFeePercentages).AuxiliaryFeeRecipients(auxiliaryFeeRecipients).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrders`: ListOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **user** | **string** | Ethereum address of the user who submitted this order | 
 **status** | **string** | Status of this order | 
 **minTimestamp** | **string** | Minimum created at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **maxTimestamp** | **string** | Maximum created at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **updatedMinTimestamp** | **string** | Minimum updated at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **updatedMaxTimestamp** | **string** | Maximum updated at timestamp for this order, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **buyTokenType** | **string** | Token type of the asset this order buys | 
 **buyTokenId** | **string** | ERC721 Token ID of the asset this order buys | 
 **buyAssetId** | **string** | Internal IMX ID of the asset this order buys | 
 **buyTokenAddress** | **string** | Comma separated string of token addresses of the asset this order buys | 
 **buyTokenName** | **string** | Token name of the asset this order buys | 
 **buyMinQuantity** | **string** | Min quantity for the asset this order buys | 
 **buyMaxQuantity** | **string** | Max quantity for the asset this order buys | 
 **buyMetadata** | **string** | JSON-encoded metadata filters for the asset this order buys | 
 **sellTokenType** | **string** | Token type of the asset this order sells | 
 **sellTokenId** | **string** | ERC721 Token ID of the asset this order sells | 
 **sellAssetId** | **string** | Internal IMX ID of the asset this order sells | 
 **sellTokenAddress** | **string** | Comma separated string of token addresses of the asset this order sells | 
 **sellTokenName** | **string** | Token name of the asset this order sells | 
 **sellMinQuantity** | **string** | Min quantity for the asset this order sells | 
 **sellMaxQuantity** | **string** | Max quantity for the asset this order sells | 
 **sellMetadata** | **string** | JSON-encoded metadata filters for the asset this order sells | 
 **auxiliaryFeePercentages** | **string** | Comma separated string of fee percentages that are to be paired with auxiliary_fee_recipients | 
 **auxiliaryFeeRecipients** | **string** | Comma separated string of fee recipients that are to be paired with auxiliary_fee_percentages | 

### Return type

[**ListOrdersResponse**](ListOrdersResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

