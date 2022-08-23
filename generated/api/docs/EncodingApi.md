# \EncodingApi

All URIs are relative to *https://api.ropsten.x.github.com/immutable/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EncodeAsset**](EncodingApi.md#EncodeAsset) | **Post** /v1/encode/{assetType} | Retrieves the Starkex Encoded format for a given asset



## EncodeAsset

> EncodeAssetResponse EncodeAsset(ctx, assetType).EncodeAssetRequest(encodeAssetRequest).Execute()

Retrieves the Starkex Encoded format for a given asset



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/generated/api"
    
)

func main() {
    assetType := "assetType_example" // string | Asset type to be encoded. (asset/mintable-asset)
    encodeAssetRequest := *openapiclient.NewEncodeAssetRequest(*openapiclient.NewEncodeAssetRequestToken()) // EncodeAssetRequest | Encode Asset

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EncodingApi.EncodeAsset(context.Background(), assetType).EncodeAssetRequest(encodeAssetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EncodingApi.EncodeAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EncodeAsset`: EncodeAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `EncodingApi.EncodeAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetType** | **string** | Asset type to be encoded. (asset/mintable-asset) | 

### Other Parameters

Other parameters are passed through a pointer to a apiEncodeAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **encodeAssetRequest** | [**EncodeAssetRequest**](EncodeAssetRequest.md) | Encode Asset | 

### Return type

[**EncodeAssetResponse**](EncodeAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

