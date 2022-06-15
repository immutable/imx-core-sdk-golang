# \MintsApi

All URIs are relative to *https://api.dev.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMint**](MintsApi.md#GetMint) | **Get** /v1/mints/{id} | Get details of a mint with the given ID
[**GetMintableTokenDetailsByClientTokenId**](MintsApi.md#GetMintableTokenDetailsByClientTokenId) | **Get** /v1/mintable-token/{token_address}/{token_id} | Get details of a mintable token with the given token address and token ID
[**ListMints**](MintsApi.md#ListMints) | **Get** /v1/mints | Get a list of mints
[**MintTokens**](MintsApi.md#MintTokens) | **Post** /v2/mints | Mint Tokens V2



## GetMint

> Mint GetMint(ctx, id).Execute()

Get details of a mint with the given ID



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
    id := "id_example" // string | Mint ID. This is the transaction_id returned from listMints

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.GetMint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.GetMint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMint`: Mint
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.GetMint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Mint ID. This is the transaction_id returned from listMints | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Mint**](Mint.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMintableTokenDetailsByClientTokenId

> MintableTokenDetails GetMintableTokenDetailsByClientTokenId(ctx, tokenAddress, tokenId).Execute()

Get details of a mintable token with the given token address and token ID



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
    tokenAddress := "tokenAddress_example" // string | Address of the ERC721 contract
    tokenId := "tokenId_example" // string | ERC721 token ID

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.GetMintableTokenDetailsByClientTokenId(context.Background(), tokenAddress, tokenId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.GetMintableTokenDetailsByClientTokenId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMintableTokenDetailsByClientTokenId`: MintableTokenDetails
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.GetMintableTokenDetailsByClientTokenId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tokenAddress** | **string** | Address of the ERC721 contract | 
**tokenId** | **string** | ERC721 token ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMintableTokenDetailsByClientTokenIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**MintableTokenDetails**](MintableTokenDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMints

> ListMintsResponse ListMints(ctx).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenName(tokenName).TokenAddress(tokenAddress).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()

Get a list of mints



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
    user := "user_example" // string | Ethereum address of the user who submitted this mint (optional)
    status := "status_example" // string | Status of this mint (optional)
    updatedMinTimestamp := "updatedMinTimestamp_example" // string | Minimum timestamp for this mint, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    updatedMaxTimestamp := "updatedMaxTimestamp_example" // string | Maximum timestamp for this mint, in ISO 8601 UTC format. Example: '2022-05-27T00:10:22Z' (optional)
    tokenType := "tokenType_example" // string | Token type of the minted asset (optional)
    tokenId := "tokenId_example" // string | ERC721 Token ID of the minted asset (optional)
    assetId := "assetId_example" // string | Internal IMX ID of the minted asset (optional)
    tokenName := "tokenName_example" // string | Token Name of the minted asset (optional)
    tokenAddress := "tokenAddress_example" // string | Token address of the minted asset (optional)
    minQuantity := "minQuantity_example" // string | Min quantity for the minted asset (optional)
    maxQuantity := "maxQuantity_example" // string | Max quantity for the minted asset (optional)
    metadata := "metadata_example" // string | JSON-encoded metadata filters for the minted asset (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.ListMints(context.Background()).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).User(user).Status(status).UpdatedMinTimestamp(updatedMinTimestamp).UpdatedMaxTimestamp(updatedMaxTimestamp).TokenType(tokenType).TokenId(tokenId).AssetId(assetId).TokenName(tokenName).TokenAddress(tokenAddress).MinQuantity(minQuantity).MaxQuantity(maxQuantity).Metadata(metadata).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.ListMints``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMints`: ListMintsResponse
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.ListMints`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 
 **user** | **string** | Ethereum address of the user who submitted this mint | 
 **status** | **string** | Status of this mint | 
 **updatedMinTimestamp** | **string** | Minimum timestamp for this mint, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **updatedMaxTimestamp** | **string** | Maximum timestamp for this mint, in ISO 8601 UTC format. Example: &#39;2022-05-27T00:10:22Z&#39; | 
 **tokenType** | **string** | Token type of the minted asset | 
 **tokenId** | **string** | ERC721 Token ID of the minted asset | 
 **assetId** | **string** | Internal IMX ID of the minted asset | 
 **tokenName** | **string** | Token Name of the minted asset | 
 **tokenAddress** | **string** | Token address of the minted asset | 
 **minQuantity** | **string** | Min quantity for the minted asset | 
 **maxQuantity** | **string** | Max quantity for the minted asset | 
 **metadata** | **string** | JSON-encoded metadata filters for the minted asset | 

### Return type

[**ListMintsResponse**](ListMintsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MintTokens

> MintTokensResponse MintTokens(ctx).MintTokensRequestV2(mintTokensRequestV2).Execute()

Mint Tokens V2



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
    mintTokensRequestV2 := []openapiclient.MintRequest{*openapiclient.NewMintRequest("AuthSignature_example", "ContractAddress_example", []openapiclient.MintUser{*openapiclient.NewMintUser([]openapiclient.MintTokenDataV2{*openapiclient.NewMintTokenDataV2("Id_example")}, "User_example")})} // []MintRequest | details of tokens to mint

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MintsApi.MintTokens(context.Background()).MintTokensRequestV2(mintTokensRequestV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MintsApi.MintTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MintTokens`: MintTokensResponse
    fmt.Fprintf(os.Stdout, "Response from `MintsApi.MintTokens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMintTokensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mintTokensRequestV2** | [**[]MintRequest**](MintRequest.md) | details of tokens to mint | 

### Return type

[**MintTokensResponse**](MintTokensResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

