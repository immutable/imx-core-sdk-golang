# GetSignableRegistrationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorSignature** | **string** | Signature from authorised operator | 
**PayloadHash** | **string** | Hash of the payload to be signed for user registration | 

## Methods

### NewGetSignableRegistrationResponse

`func NewGetSignableRegistrationResponse(operatorSignature string, payloadHash string, ) *GetSignableRegistrationResponse`

NewGetSignableRegistrationResponse instantiates a new GetSignableRegistrationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableRegistrationResponseWithDefaults

`func NewGetSignableRegistrationResponseWithDefaults() *GetSignableRegistrationResponse`

NewGetSignableRegistrationResponseWithDefaults instantiates a new GetSignableRegistrationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatorSignature

`func (o *GetSignableRegistrationResponse) GetOperatorSignature() string`

GetOperatorSignature returns the OperatorSignature field if non-nil, zero value otherwise.

### GetOperatorSignatureOk

`func (o *GetSignableRegistrationResponse) GetOperatorSignatureOk() (*string, bool)`

GetOperatorSignatureOk returns a tuple with the OperatorSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorSignature

`func (o *GetSignableRegistrationResponse) SetOperatorSignature(v string)`

SetOperatorSignature sets OperatorSignature field to given value.


### GetPayloadHash

`func (o *GetSignableRegistrationResponse) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableRegistrationResponse) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableRegistrationResponse) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


