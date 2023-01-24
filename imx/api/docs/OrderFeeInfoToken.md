# OrderFeeInfoToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**FeeTokenData**](FeeTokenData.md) |  | [optional] 
**Type** | Pointer to **string** | Fee token type. One of ETH/ERC20 | [optional] 

## Methods

### NewOrderFeeInfoToken

`func NewOrderFeeInfoToken() *OrderFeeInfoToken`

NewOrderFeeInfoToken instantiates a new OrderFeeInfoToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFeeInfoTokenWithDefaults

`func NewOrderFeeInfoTokenWithDefaults() *OrderFeeInfoToken`

NewOrderFeeInfoTokenWithDefaults instantiates a new OrderFeeInfoToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OrderFeeInfoToken) GetData() FeeTokenData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderFeeInfoToken) GetDataOk() (*FeeTokenData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderFeeInfoToken) SetData(v FeeTokenData)`

SetData sets Data field to given value.

### HasData

`func (o *OrderFeeInfoToken) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *OrderFeeInfoToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderFeeInfoToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderFeeInfoToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderFeeInfoToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


