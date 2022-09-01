# Fee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Wallet address | 
**Percentage** | **float32** | The percentage of fee | 
**Type** | **string** | Type of fee. Examples: &#x60;royalty&#x60;, &#x60;maker&#x60;, &#x60;taker&#x60; or &#x60;protocol&#x60; | 

## Methods

### NewFee

`func NewFee(address string, percentage float32, type_ string, ) *Fee`

NewFee instantiates a new Fee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeWithDefaults

`func NewFeeWithDefaults() *Fee`

NewFeeWithDefaults instantiates a new Fee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Fee) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Fee) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Fee) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetPercentage

`func (o *Fee) GetPercentage() float32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *Fee) GetPercentageOk() (*float32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *Fee) SetPercentage(v float32)`

SetPercentage sets Percentage field to given value.


### GetType

`func (o *Fee) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Fee) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Fee) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


