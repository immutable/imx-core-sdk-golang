# CreateWithdrawalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of the withdrawal | 
**Time** | **int32** | Time of the withdrawal | 
**WithdrawalId** | **int32** | ID of the withdrawal | 

## Methods

### NewCreateWithdrawalResponse

`func NewCreateWithdrawalResponse(status string, time int32, withdrawalId int32, ) *CreateWithdrawalResponse`

NewCreateWithdrawalResponse instantiates a new CreateWithdrawalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWithdrawalResponseWithDefaults

`func NewCreateWithdrawalResponseWithDefaults() *CreateWithdrawalResponse`

NewCreateWithdrawalResponseWithDefaults instantiates a new CreateWithdrawalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateWithdrawalResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateWithdrawalResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateWithdrawalResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTime

`func (o *CreateWithdrawalResponse) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateWithdrawalResponse) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateWithdrawalResponse) SetTime(v int32)`

SetTime sets Time field to given value.


### GetWithdrawalId

`func (o *CreateWithdrawalResponse) GetWithdrawalId() int32`

GetWithdrawalId returns the WithdrawalId field if non-nil, zero value otherwise.

### GetWithdrawalIdOk

`func (o *CreateWithdrawalResponse) GetWithdrawalIdOk() (*int32, bool)`

GetWithdrawalIdOk returns a tuple with the WithdrawalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawalId

`func (o *CreateWithdrawalResponse) SetWithdrawalId(v int32)`

SetWithdrawalId sets WithdrawalId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


