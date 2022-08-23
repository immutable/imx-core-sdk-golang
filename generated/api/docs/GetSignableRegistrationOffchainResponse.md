# GetSignableRegistrationOffchainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PayloadHash** | **string** | Hash of the payload to be signed for user registration offchain | 
**SignableMessage** | **string** | Message to sign with L1 wallet to register user offchain | 

## Methods

### NewGetSignableRegistrationOffchainResponse

`func NewGetSignableRegistrationOffchainResponse(payloadHash string, signableMessage string, ) *GetSignableRegistrationOffchainResponse`

NewGetSignableRegistrationOffchainResponse instantiates a new GetSignableRegistrationOffchainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableRegistrationOffchainResponseWithDefaults

`func NewGetSignableRegistrationOffchainResponseWithDefaults() *GetSignableRegistrationOffchainResponse`

NewGetSignableRegistrationOffchainResponseWithDefaults instantiates a new GetSignableRegistrationOffchainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPayloadHash

`func (o *GetSignableRegistrationOffchainResponse) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableRegistrationOffchainResponse) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableRegistrationOffchainResponse) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.


### GetSignableMessage

`func (o *GetSignableRegistrationOffchainResponse) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableRegistrationOffchainResponse) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableRegistrationOffchainResponse) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


