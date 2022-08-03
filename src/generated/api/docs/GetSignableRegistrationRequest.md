# GetSignableRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EtherKey** | **string** | Ether key of user | 
**StarkKey** | **string** | Public stark key of the user | 

## Methods

### NewGetSignableRegistrationRequest

`func NewGetSignableRegistrationRequest(etherKey string, starkKey string, ) *GetSignableRegistrationRequest`

NewGetSignableRegistrationRequest instantiates a new GetSignableRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableRegistrationRequestWithDefaults

`func NewGetSignableRegistrationRequestWithDefaults() *GetSignableRegistrationRequest`

NewGetSignableRegistrationRequestWithDefaults instantiates a new GetSignableRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEtherKey

`func (o *GetSignableRegistrationRequest) GetEtherKey() string`

GetEtherKey returns the EtherKey field if non-nil, zero value otherwise.

### GetEtherKeyOk

`func (o *GetSignableRegistrationRequest) GetEtherKeyOk() (*string, bool)`

GetEtherKeyOk returns a tuple with the EtherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEtherKey

`func (o *GetSignableRegistrationRequest) SetEtherKey(v string)`

SetEtherKey sets EtherKey field to given value.


### GetStarkKey

`func (o *GetSignableRegistrationRequest) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *GetSignableRegistrationRequest) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *GetSignableRegistrationRequest) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


