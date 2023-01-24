# TradeA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | The ID of the order involved in the trade | 
**Sold** | **string** | The amount of that order&#39;s asset this trade sold | 
**TokenAddress** | Pointer to **string** | The contract address of the token that this trade sold | [optional] 
**TokenId** | Pointer to **string** | The ID of the token that this trade sold | [optional] 
**TokenType** | **string** | The type of the token that this trade sold | 

## Methods

### NewTradeA

`func NewTradeA(orderId int32, sold string, tokenType string, ) *TradeA`

NewTradeA instantiates a new TradeA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeAWithDefaults

`func NewTradeAWithDefaults() *TradeA`

NewTradeAWithDefaults instantiates a new TradeA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *TradeA) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *TradeA) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *TradeA) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetSold

`func (o *TradeA) GetSold() string`

GetSold returns the Sold field if non-nil, zero value otherwise.

### GetSoldOk

`func (o *TradeA) GetSoldOk() (*string, bool)`

GetSoldOk returns a tuple with the Sold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSold

`func (o *TradeA) SetSold(v string)`

SetSold sets Sold field to given value.


### GetTokenAddress

`func (o *TradeA) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TradeA) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TradeA) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TradeA) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *TradeA) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TradeA) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TradeA) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TradeA) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTokenType

`func (o *TradeA) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *TradeA) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *TradeA) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


