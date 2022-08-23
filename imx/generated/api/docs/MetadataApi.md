# \MetadataApi

All URIs are relative to *https://api.ropsten.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMetadataSchemaToCollection**](MetadataApi.md#AddMetadataSchemaToCollection) | **Post** /v1/collections/{address}/metadata-schema | Add metadata schema to collection
[**GetMetadataSchema**](MetadataApi.md#GetMetadataSchema) | **Get** /v1/collections/{address}/metadata-schema | Get collection metadata schema
[**UpdateMetadataSchemaByName**](MetadataApi.md#UpdateMetadataSchemaByName) | **Patch** /v1/collections/{address}/metadata-schema/{name} | Update metadata schema by name



## AddMetadataSchemaToCollection

> SuccessResponse AddMetadataSchemaToCollection(ctx, address).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).AddMetadataSchemaToCollectionRequest(addMetadataSchemaToCollectionRequest).Execute()

Add metadata schema to collection



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/generated/api"
    
)

func main() {
    address := "address_example" // string | Collection contract address
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp
    addMetadataSchemaToCollectionRequest := *openapiclient.NewAddMetadataSchemaToCollectionRequest([]openapiclient.MetadataSchemaRequest{*openapiclient.NewMetadataSchemaRequest("Name_example")}) // AddMetadataSchemaToCollectionRequest | add metadata schema to a collection

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.AddMetadataSchemaToCollection(context.Background(), address).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).AddMetadataSchemaToCollectionRequest(addMetadataSchemaToCollectionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.AddMetadataSchemaToCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMetadataSchemaToCollection`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.AddMetadataSchemaToCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Collection contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMetadataSchemaToCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMXSignature** | **string** | String created by signing wallet address and timestamp | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 
 **addMetadataSchemaToCollectionRequest** | [**AddMetadataSchemaToCollectionRequest**](AddMetadataSchemaToCollectionRequest.md) | add metadata schema to a collection | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetadataSchema

> []MetadataSchemaProperty GetMetadataSchema(ctx, address).Execute()

Get collection metadata schema



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/generated/api"
    
)

func main() {
    address := "address_example" // string | Collection contract address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.GetMetadataSchema(context.Background(), address).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.GetMetadataSchema``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetadataSchema`: []MetadataSchemaProperty
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.GetMetadataSchema`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Collection contract address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetadataSchemaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]MetadataSchemaProperty**](MetadataSchemaProperty.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetadataSchemaByName

> SuccessResponse UpdateMetadataSchemaByName(ctx, address, name).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).MetadataSchemaRequest(metadataSchemaRequest).Execute()

Update metadata schema by name



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/immutable/imx-core-sdk-golang/imx/generated/api"
    
)

func main() {
    address := "address_example" // string | Collection contract address
    name := "name_example" // string | Metadata schema name
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp
    metadataSchemaRequest := *openapiclient.NewMetadataSchemaRequest("Name_example") // MetadataSchemaRequest | update metadata schema

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetadataApi.UpdateMetadataSchemaByName(context.Background(), address, name).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).MetadataSchemaRequest(metadataSchemaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetadataApi.UpdateMetadataSchemaByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetadataSchemaByName`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `MetadataApi.UpdateMetadataSchemaByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | Collection contract address | 
**name** | **string** | Metadata schema name | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetadataSchemaByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iMXSignature** | **string** | String created by signing wallet address and timestamp | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 
 **metadataSchemaRequest** | [**MetadataSchemaRequest**](MetadataSchemaRequest.md) | update metadata schema | 

### Return type

[**SuccessResponse**](SuccessResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

