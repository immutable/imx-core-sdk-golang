# RegisterUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TxHash** | **string** | Immutable signature authorising registration | 

## Methods

### NewRegisterUserResponse

`func NewRegisterUserResponse(txHash string, ) *RegisterUserResponse`

NewRegisterUserResponse instantiates a new RegisterUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserResponseWithDefaults

`func NewRegisterUserResponseWithDefaults() *RegisterUserResponse`

NewRegisterUserResponseWithDefaults instantiates a new RegisterUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTxHash

`func (o *RegisterUserResponse) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *RegisterUserResponse) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *RegisterUserResponse) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


