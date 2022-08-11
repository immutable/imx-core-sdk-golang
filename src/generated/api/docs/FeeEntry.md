# FeeEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**FeePercentage** | Pointer to **float32** |  | [optional] 

## Methods

### NewFeeEntry

`func NewFeeEntry() *FeeEntry`

NewFeeEntry instantiates a new FeeEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeEntryWithDefaults

`func NewFeeEntryWithDefaults() *FeeEntry`

NewFeeEntryWithDefaults instantiates a new FeeEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *FeeEntry) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *FeeEntry) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *FeeEntry) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *FeeEntry) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetFeePercentage

`func (o *FeeEntry) GetFeePercentage() float32`

GetFeePercentage returns the FeePercentage field if non-nil, zero value otherwise.

### GetFeePercentageOk

`func (o *FeeEntry) GetFeePercentageOk() (*float32, bool)`

GetFeePercentageOk returns a tuple with the FeePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeePercentage

`func (o *FeeEntry) SetFeePercentage(v float32)`

SetFeePercentage sets FeePercentage field to given value.

### HasFeePercentage

`func (o *FeeEntry) HasFeePercentage() bool`

HasFeePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


