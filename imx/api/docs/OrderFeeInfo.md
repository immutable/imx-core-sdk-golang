# OrderFeeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **string** |  | [optional] 
**Token** | Pointer to [**FeeToken**](FeeToken.md) |  | [optional] 
**Type** | Pointer to **string** | Fee type | [optional] 

## Methods

### NewOrderFeeInfo

`func NewOrderFeeInfo() *OrderFeeInfo`

NewOrderFeeInfo instantiates a new OrderFeeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderFeeInfoWithDefaults

`func NewOrderFeeInfoWithDefaults() *OrderFeeInfo`

NewOrderFeeInfoWithDefaults instantiates a new OrderFeeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *OrderFeeInfo) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *OrderFeeInfo) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *OrderFeeInfo) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *OrderFeeInfo) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetAmount

`func (o *OrderFeeInfo) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *OrderFeeInfo) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *OrderFeeInfo) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *OrderFeeInfo) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetToken

`func (o *OrderFeeInfo) GetToken() FeeToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OrderFeeInfo) GetTokenOk() (*FeeToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OrderFeeInfo) SetToken(v FeeToken)`

SetToken sets Token field to given value.

### HasToken

`func (o *OrderFeeInfo) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *OrderFeeInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderFeeInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderFeeInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OrderFeeInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


