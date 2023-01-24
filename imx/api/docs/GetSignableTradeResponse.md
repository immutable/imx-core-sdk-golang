# GetSignableTradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | TODO: Remove unused fields Code and Message [deprecated] Response code | [optional] 
**Message** | Pointer to **string** | [deprecated] Response message | [optional] 
**Result** | Pointer to [**GetSignableTradeResponseResult**](GetSignableTradeResponseResult.md) |  | [optional] 

## Methods

### NewGetSignableTradeResponse

`func NewGetSignableTradeResponse() *GetSignableTradeResponse`

NewGetSignableTradeResponse instantiates a new GetSignableTradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTradeResponseWithDefaults

`func NewGetSignableTradeResponseWithDefaults() *GetSignableTradeResponse`

NewGetSignableTradeResponseWithDefaults instantiates a new GetSignableTradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetSignableTradeResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetSignableTradeResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetSignableTradeResponse) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetSignableTradeResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *GetSignableTradeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetSignableTradeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetSignableTradeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetSignableTradeResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResult

`func (o *GetSignableTradeResponse) GetResult() GetSignableTradeResponseResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetSignableTradeResponse) GetResultOk() (*GetSignableTradeResponseResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetSignableTradeResponse) SetResult(v GetSignableTradeResponseResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetSignableTradeResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


