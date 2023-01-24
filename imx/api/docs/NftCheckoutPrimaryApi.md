# \NftCheckoutPrimaryApi

All URIs are relative to *https://api.sandbox.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNftPrimary**](NftCheckoutPrimaryApi.md#CreateNftPrimary) | **Post** /v2/nft/primary | Create NFT primary sale transaction
[**GetCurrenciesNFTCheckoutPrimary**](NftCheckoutPrimaryApi.md#GetCurrenciesNFTCheckoutPrimary) | **Get** /v2/nft/primary/currencies | Get currencies with limits
[**GetNftPrimaryTransaction**](NftCheckoutPrimaryApi.md#GetNftPrimaryTransaction) | **Get** /v2/nft/primary/{transaction_id} | Get NFT primary sale transaction by id
[**GetNftPrimaryTransactions**](NftCheckoutPrimaryApi.md#GetNftPrimaryTransactions) | **Get** /v2/nft/primary | Get a list of NFT primary sales transactions
[**RegisterNftPrimarySalesContract**](NftCheckoutPrimaryApi.md#RegisterNftPrimarySalesContract) | **Post** /v2/nft/primary/register | Executes NFT primary sales contract registration



## CreateNftPrimary

> NftprimarytransactionCreateResponse CreateNftPrimary(ctx).CreateAPIRequest(createAPIRequest).Execute()

Create NFT primary sale transaction



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
    createAPIRequest := *openapiclient.NewNftprimarytransactionCreateAPIRequest() // NftprimarytransactionCreateAPIRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.CreateNftPrimary(context.Background()).CreateAPIRequest(createAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftCheckoutPrimaryApi.CreateNftPrimary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNftPrimary`: NftprimarytransactionCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `NftCheckoutPrimaryApi.CreateNftPrimary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNftPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAPIRequest** | [**NftprimarytransactionCreateAPIRequest**](NftprimarytransactionCreateAPIRequest.md) | req | 

### Return type

[**NftprimarytransactionCreateResponse**](NftprimarytransactionCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCurrenciesNFTCheckoutPrimary

> CurrencyWithLimits GetCurrenciesNFTCheckoutPrimary(ctx).Provider(provider).IncludeLimits(includeLimits).Execute()

Get currencies with limits



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
    provider := "provider_example" // string | Provider name (optional)
    includeLimits := true // bool | Flag if limits should be included in the response or not (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.GetCurrenciesNFTCheckoutPrimary(context.Background()).Provider(provider).IncludeLimits(includeLimits).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftCheckoutPrimaryApi.GetCurrenciesNFTCheckoutPrimary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrenciesNFTCheckoutPrimary`: CurrencyWithLimits
    fmt.Fprintf(os.Stdout, "Response from `NftCheckoutPrimaryApi.GetCurrenciesNFTCheckoutPrimary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrenciesNFTCheckoutPrimaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **provider** | **string** | Provider name | 
 **includeLimits** | **bool** | Flag if limits should be included in the response or not | 

### Return type

[**CurrencyWithLimits**](CurrencyWithLimits.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftPrimaryTransaction

> NftprimarytransactionGetResponse GetNftPrimaryTransaction(ctx, transactionId).Execute()

Get NFT primary sale transaction by id



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
    transactionId := "transactionId_example" // string | Transaction id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.GetNftPrimaryTransaction(context.Background(), transactionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftCheckoutPrimaryApi.GetNftPrimaryTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftPrimaryTransaction`: NftprimarytransactionGetResponse
    fmt.Fprintf(os.Stdout, "Response from `NftCheckoutPrimaryApi.GetNftPrimaryTransaction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transactionId** | **string** | Transaction id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNftPrimaryTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**NftprimarytransactionGetResponse**](NftprimarytransactionGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNftPrimaryTransactions

> NftprimarytransactionListTransactionsResponse GetNftPrimaryTransactions(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).TransactionId(transactionId).ContractAddress(contractAddress).SellerWalletAddress(sellerWalletAddress).UserWalletAddress(userWalletAddress).Status(status).Provider(provider).MintId(mintId).Execute()

Get a list of NFT primary sales transactions



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
    transactionId := "transactionId_example" // string | Transaction id (optional)
    contractAddress := "contractAddress_example" // string | Contract address of the asset (optional)
    sellerWalletAddress := "sellerWalletAddress_example" // string | Ethereum address of the seller (optional)
    userWalletAddress := "userWalletAddress_example" // string | Ethereum address of the user who wants to create transaction (optional)
    status := "status_example" // string | Transaction status (optional)
    provider := "provider_example" // string | Checkout provider name (optional)
    mintId := "mintId_example" // string | Minting transaction ID - see mintTokens response (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.GetNftPrimaryTransactions(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).TransactionId(transactionId).ContractAddress(contractAddress).SellerWalletAddress(sellerWalletAddress).UserWalletAddress(userWalletAddress).Status(status).Provider(provider).MintId(mintId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftCheckoutPrimaryApi.GetNftPrimaryTransactions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftPrimaryTransactions`: NftprimarytransactionListTransactionsResponse
    fmt.Fprintf(os.Stdout, "Response from `NftCheckoutPrimaryApi.GetNftPrimaryTransactions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNftPrimaryTransactionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **transactionId** | **string** | Transaction id | 
 **contractAddress** | **string** | Contract address of the asset | 
 **sellerWalletAddress** | **string** | Ethereum address of the seller | 
 **userWalletAddress** | **string** | Ethereum address of the user who wants to create transaction | 
 **status** | **string** | Transaction status | 
 **provider** | **string** | Checkout provider name | 
 **mintId** | **string** | Minting transaction ID - see mintTokens response | 

### Return type

[**NftprimarytransactionListTransactionsResponse**](NftprimarytransactionListTransactionsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterNftPrimarySalesContract

> ContractCreateResponse RegisterNftPrimarySalesContract(ctx).CreateAPIRequest(createAPIRequest).Execute()

Executes NFT primary sales contract registration



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
    createAPIRequest := *openapiclient.NewContractCreateAPIRequest() // ContractCreateAPIRequest | req

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.RegisterNftPrimarySalesContract(context.Background()).CreateAPIRequest(createAPIRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftCheckoutPrimaryApi.RegisterNftPrimarySalesContract``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterNftPrimarySalesContract`: ContractCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `NftCheckoutPrimaryApi.RegisterNftPrimarySalesContract`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterNftPrimarySalesContractRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAPIRequest** | [**ContractCreateAPIRequest**](ContractCreateAPIRequest.md) | req | 

### Return type

[**ContractCreateResponse**](ContractCreateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

