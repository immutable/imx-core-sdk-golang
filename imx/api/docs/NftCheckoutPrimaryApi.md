# \NftCheckoutPrimaryApi

All URIs are relative to *https://api.sandbox.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNftPrimary**](NftCheckoutPrimaryApi.md#CreateNftPrimary) | **Post** /v2/nft/primary | Create nft primary transaction
[**GetCurrenciesNFTCheckoutPrimary**](NftCheckoutPrimaryApi.md#GetCurrenciesNFTCheckoutPrimary) | **Get** /v2/nft/primary/currencies | Get currencies with limits
[**GetMintStatusById**](NftCheckoutPrimaryApi.md#GetMintStatusById) | **Get** /v2/{provider}/transaction_status | Get mint status by transaction id
[**GetNftPrimaryTransaction**](NftCheckoutPrimaryApi.md#GetNftPrimaryTransaction) | **Get** /v2/nft/primary/{transaction_id} | Get nft primary transaction by id
[**GetNftPrimaryTransactions**](NftCheckoutPrimaryApi.md#GetNftPrimaryTransactions) | **Get** /v2/nft/primary | Get a list of NFT primary transactions



## CreateNftPrimary

> NftprimarytransactionCreateResponse CreateNftPrimary(ctx).CreateAPIRequest(createAPIRequest).Execute()

Create nft primary transaction



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


## GetMintStatusById

> ProviderGetMintStatusResponse GetMintStatusById(ctx, provider).Id(id).Execute()

Get mint status by transaction id



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
    provider := "provider_example" // string | Provider name
    id := "id_example" // string | transaction id

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.GetMintStatusById(context.Background(), provider).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftCheckoutPrimaryApi.GetMintStatusById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMintStatusById`: ProviderGetMintStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `NftCheckoutPrimaryApi.GetMintStatusById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**provider** | **string** | Provider name | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMintStatusByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **string** | transaction id | 

### Return type

[**ProviderGetMintStatusResponse**](ProviderGetMintStatusResponse.md)

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

Get nft primary transaction by id



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

> NftprimarytransactionListTransactionsResponse GetNftPrimaryTransactions(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).TransactionId(transactionId).ContractAddress(contractAddress).SellerWalletAddress(sellerWalletAddress).WalletAddress(walletAddress).Status(status).Provider(provider).MintId(mintId).Execute()

Get a list of NFT primary transactions



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
    walletAddress := "walletAddress_example" // string | Ethereum address of the user who wants to create transaction (optional)
    status := "status_example" // string | Transaction status (optional)
    provider := "provider_example" // string | Provider name (optional)
    mintId := "mintId_example" // string | Mint id (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftCheckoutPrimaryApi.GetNftPrimaryTransactions(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).TransactionId(transactionId).ContractAddress(contractAddress).SellerWalletAddress(sellerWalletAddress).WalletAddress(walletAddress).Status(status).Provider(provider).MintId(mintId).Execute()
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
 **walletAddress** | **string** | Ethereum address of the user who wants to create transaction | 
 **status** | **string** | Transaction status | 
 **provider** | **string** | Provider name | 
 **mintId** | **string** | Mint id | 

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

