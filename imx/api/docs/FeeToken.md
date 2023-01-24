# FeeToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**FeeTokenData**](FeeTokenData.md) |  | [optional] 
**Type** | Pointer to **string** | Fee token type. One of ETH/ERC20 | [optional] 

## Methods

### NewFeeToken

`func NewFeeToken() *FeeToken`

NewFeeToken instantiates a new FeeToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeTokenWithDefaults

`func NewFeeTokenWithDefaults() *FeeToken`

NewFeeTokenWithDefaults instantiates a new FeeToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *FeeToken) GetData() FeeTokenData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *FeeToken) GetDataOk() (*FeeTokenData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *FeeToken) SetData(v FeeTokenData)`

SetData sets Data field to given value.

### HasData

`func (o *FeeToken) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *FeeToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeeToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeeToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *FeeToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


