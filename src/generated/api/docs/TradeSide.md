# TradeSide

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | The ID of the order involved in the trade | 
**Sold** | **string** | The amount of that order&#39;s asset this trade sold | 
**TokenAddress** | Pointer to **string** | The contract address of the token that this trade sold | [optional] 
**TokenId** | Pointer to **string** | The ID of the token that this trade sold | [optional] 
**TokenType** | **string** | The type of the token that this trade sold | 

## Methods

### NewTradeSide

`func NewTradeSide(orderId int32, sold string, tokenType string, ) *TradeSide`

NewTradeSide instantiates a new TradeSide object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeSideWithDefaults

`func NewTradeSideWithDefaults() *TradeSide`

NewTradeSideWithDefaults instantiates a new TradeSide object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *TradeSide) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TradeSide) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TradeSide) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetSold

`func (o *TradeSide) GetSold() string`

GetSold returns the Sold field if non-nil, zero value otherwise.

### GetSoldOk

`func (o *TradeSide) GetSoldOk() (*string, bool)`

GetSoldOk returns a tuple with the Sold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSold

`func (o *TradeSide) SetSold(v string)`

SetSold sets Sold field to given value.


### GetTokenAddress

`func (o *TradeSide) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TradeSide) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TradeSide) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TradeSide) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *TradeSide) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TradeSide) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TradeSide) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TradeSide) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenType

`func (o *TradeSide) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TradeSide) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TradeSide) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


