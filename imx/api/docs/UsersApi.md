# \UsersApi

All URIs are relative to *https://api.sandbox.x.immutable.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSignableRegistration**](UsersApi.md#GetSignableRegistration) | **Post** /v1/signable-registration | Get operator signature to allow clients to register the user
[**GetSignableRegistrationOffchain**](UsersApi.md#GetSignableRegistrationOffchain) | **Post** /v1/signable-registration-offchain | Get encoded details to allow clients to register the user offchain
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /v1/users/{user} | Get stark keys for a registered user
[**RegisterUser**](UsersApi.md#RegisterUser) | **Post** /v1/users | Registers a user



## GetSignableRegistration

> GetSignableRegistrationResponse GetSignableRegistration(ctx).GetSignableRegistrationRequest(getSignableRegistrationRequest).Execute()

Get operator signature to allow clients to register the user



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
    getSignableRegistrationRequest := *openapiclient.NewGetSignableRegistrationRequest("EtherKey_example", "StarkKey_example") // GetSignableRegistrationRequest | Register User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetSignableRegistration(context.Background()).GetSignableRegistrationRequest(getSignableRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetSignableRegistration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableRegistration`: GetSignableRegistrationResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetSignableRegistration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableRegistrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableRegistrationRequest** | [**GetSignableRegistrationRequest**](GetSignableRegistrationRequest.md) | Register User | 

### Return type

[**GetSignableRegistrationResponse**](GetSignableRegistrationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSignableRegistrationOffchain

> GetSignableRegistrationOffchainResponse GetSignableRegistrationOffchain(ctx).GetSignableRegistrationRequest(getSignableRegistrationRequest).Execute()

Get encoded details to allow clients to register the user offchain



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
    getSignableRegistrationRequest := *openapiclient.NewGetSignableRegistrationRequest("EtherKey_example", "StarkKey_example") // GetSignableRegistrationRequest | Register User Offchain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetSignableRegistrationOffchain(context.Background()).GetSignableRegistrationRequest(getSignableRegistrationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetSignableRegistrationOffchain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSignableRegistrationOffchain`: GetSignableRegistrationOffchainResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetSignableRegistrationOffchain`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSignableRegistrationOffchainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getSignableRegistrationRequest** | [**GetSignableRegistrationRequest**](GetSignableRegistrationRequest.md) | Register User Offchain | 

### Return type

[**GetSignableRegistrationOffchainResponse**](GetSignableRegistrationOffchainResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> GetUsersApiResponse GetUsers(ctx, user).Execute()

Get stark keys for a registered user



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
    user := "user_example" // string | User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.GetUsers(context.Background(), user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.GetUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsers`: GetUsersApiResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.GetUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | User | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetUsersApiResponse**](GetUsersApiResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterUser

> RegisterUserResponse RegisterUser(ctx).RegisterUserRequest(registerUserRequest).Execute()

Registers a user



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
    registerUserRequest := *openapiclient.NewRegisterUserRequest("EthSignature_example", "EtherKey_example", "StarkKey_example", "StarkSignature_example") // RegisterUserRequest | Register User

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersApi.RegisterUser(context.Background()).RegisterUserRequest(registerUserRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersApi.RegisterUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterUser`: RegisterUserResponse
    fmt.Fprintf(os.Stdout, "Response from `UsersApi.RegisterUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registerUserRequest** | [**RegisterUserRequest**](RegisterUserRequest.md) | Register User | 

### Return type

[**RegisterUserResponse**](RegisterUserResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

