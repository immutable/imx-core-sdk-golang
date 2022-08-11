# MintFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Percentage** | **float32** | Fee percentage | 
**Recipient** | **string** | Recipient wallet address | 

## Methods

### NewMintFee

`func NewMintFee(percentage float32, recipient string, ) *MintFee`

NewMintFee instantiates a new MintFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintFeeWithDefaults

`func NewMintFeeWithDefaults() *MintFee`

NewMintFeeWithDefaults instantiates a new MintFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPercentage

`func (o *MintFee) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *MintFee) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *MintFee) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetRecipient

`func (o *MintFee) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *MintFee) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *MintFee) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


