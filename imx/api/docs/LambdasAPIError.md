# LambdasAPIError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The error code | [optional] 
**Details** | Pointer to **string** | The error details | [optional] 
**Message** | Pointer to **string** | The error message | [optional] 
**StatusCode** | Pointer to **int32** | The error status code | [optional] 

## Methods

### NewLambdasAPIError

`func NewLambdasAPIError() *LambdasAPIError`

NewLambdasAPIError instantiates a new LambdasAPIError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLambdasAPIErrorWithDefaults

`func NewLambdasAPIErrorWithDefaults() *LambdasAPIError`

NewLambdasAPIErrorWithDefaults instantiates a new LambdasAPIError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *LambdasAPIError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *LambdasAPIError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *LambdasAPIError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *LambdasAPIError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *LambdasAPIError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *LambdasAPIError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *LambdasAPIError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *LambdasAPIError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *LambdasAPIError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LambdasAPIError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LambdasAPIError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LambdasAPIError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatusCode

`func (o *LambdasAPIError) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *LambdasAPIError) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *LambdasAPIError) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *LambdasAPIError) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


