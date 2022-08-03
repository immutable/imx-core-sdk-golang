# \ProjectsApi

All URIs are relative to *https://api.ropsten.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProject**](ProjectsApi.md#CreateProject) | **Post** /v1/projects | Create a project
[**GetProject**](ProjectsApi.md#GetProject) | **Get** /v1/projects/{id} | Get a project
[**GetProjects**](ProjectsApi.md#GetProjects) | **Get** /v1/projects | Get projects



## CreateProject

> CreateProjectResponse CreateProject(ctx).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).CreateProjectRequest(createProjectRequest).Execute()

Create a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./api"
)

func main() {
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp
    createProjectRequest := *openapiclient.NewCreateProjectRequest("CompanyName_example", "ContactEmail_example", "Name_example") // CreateProjectRequest | create a project

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.CreateProject(context.Background()).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).CreateProjectRequest(createProjectRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.CreateProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProject`: CreateProjectResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.CreateProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMXSignature** | **string** | String created by signing wallet address and timestamp | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 
 **createProjectRequest** | [**CreateProjectRequest**](CreateProjectRequest.md) | create a project | 

### Return type

[**CreateProjectResponse**](CreateProjectResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProject

> Project GetProject(ctx, id).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).Execute()

Get a project



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./api"
)

func main() {
    id := "id_example" // string | Project ID
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProject(context.Background(), id).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProject`: Project
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **iMXSignature** | **string** | String created by signing wallet address and timestamp | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 

### Return type

[**Project**](Project.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> GetProjectsResponse GetProjects(ctx).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Execute()

Get projects



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./api"
)

func main() {
    iMXSignature := "iMXSignature_example" // string | String created by signing wallet address and timestamp
    iMXTimestamp := "iMXTimestamp_example" // string | Unix Epoc timestamp
    pageSize := int32(56) // int32 | Page size of the result (optional)
    cursor := "cursor_example" // string | Cursor (optional)
    orderBy := "orderBy_example" // string | Property to sort by (optional)
    direction := "direction_example" // string | Direction to sort (asc/desc) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProjectsApi.GetProjects(context.Background()).IMXSignature(iMXSignature).IMXTimestamp(iMXTimestamp).PageSize(pageSize).Cursor(cursor).OrderBy(orderBy).Direction(direction).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProjectsApi.GetProjects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProjects`: GetProjectsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProjectsApi.GetProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **iMXSignature** | **string** | String created by signing wallet address and timestamp | 
 **iMXTimestamp** | **string** | Unix Epoc timestamp | 
 **pageSize** | **int32** | Page size of the result | 
 **cursor** | **string** | Cursor | 
 **orderBy** | **string** | Property to sort by | 
 **direction** | **string** | Direction to sort (asc/desc) | 

### Return type

[**GetProjectsResponse**](GetProjectsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

