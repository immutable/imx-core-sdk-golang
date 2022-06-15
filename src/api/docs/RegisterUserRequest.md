# RegisterUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EthSignature** | **string** | Eth signature | 
**EtherKey** | **string** | The ether key of the user | 
**StarkKey** | **string** | Public stark key of the user | 
**StarkSignature** | **string** | Payload signature | 

## Methods

### NewRegisterUserRequest

`func NewRegisterUserRequest(ethSignature string, etherKey string, starkKey string, starkSignature string, ) *RegisterUserRequest`

NewRegisterUserRequest instantiates a new RegisterUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisterUserRequestWithDefaults

`func NewRegisterUserRequestWithDefaults() *RegisterUserRequest`

NewRegisterUserRequestWithDefaults instantiates a new RegisterUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEthSignature

`func (o *RegisterUserRequest) GetEthSignature() string`

GetEthSignature returns the EthSignature field if non-nil, zero value otherwise.

### GetEthSignatureOk

`func (o *RegisterUserRequest) GetEthSignatureOk() (*string, bool)`

GetEthSignatureOk returns a tuple with the EthSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEthSignature

`func (o *RegisterUserRequest) SetEthSignature(v string)`

SetEthSignature sets EthSignature field to given value.


### GetEtherKey

`func (o *RegisterUserRequest) GetEtherKey() string`

GetEtherKey returns the EtherKey field if non-nil, zero value otherwise.

### GetEtherKeyOk

`func (o *RegisterUserRequest) GetEtherKeyOk() (*string, bool)`

GetEtherKeyOk returns a tuple with the EtherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherKey

`func (o *RegisterUserRequest) SetEtherKey(v string)`

SetEtherKey sets EtherKey field to given value.


### GetStarkKey

`func (o *RegisterUserRequest) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *RegisterUserRequest) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *RegisterUserRequest) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetStarkSignature

`func (o *RegisterUserRequest) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *RegisterUserRequest) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *RegisterUserRequest) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


