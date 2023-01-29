# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountSold** | **NullableString** | Amount of the asset already sold by this order | 
**Buy** | [**Token**](Token.md) |  | 
**ExpirationTimestamp** | **NullableString** | Expiration timestamp of this order | 
**Fees** | Pointer to [**[]OrderFeeInfo**](OrderFeeInfo.md) | Fee information for the order | [optional] 
**OrderId** | **int32** | ID of the order | 
**Sell** | [**Token**](Token.md) |  | 
**Status** | **string** | Status of the order | 
**Timestamp** | **NullableString** | Timestamp this order was created | 
**UpdatedTimestamp** | **NullableString** | Updated timestamp of this order | 
**User** | **string** | Ethereum address of the user who submitted the order | 

## Methods

### NewOrder

`func NewOrder(amountSold NullableString, buy Token, expirationTimestamp NullableString, orderId int32, sell Token, status string, timestamp NullableString, updatedTimestamp NullableString, user string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountSold

`func (o *Order) GetAmountSold() string`

GetAmountSold returns the AmountSold field if non-nil, zero value otherwise.

### GetAmountSoldOk

`func (o *Order) GetAmountSoldOk() (*string, bool)`

GetAmountSoldOk returns a tuple with the AmountSold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSold

`func (o *Order) SetAmountSold(v string)`

SetAmountSold sets AmountSold field to given value.


### SetAmountSoldNil

`func (o *Order) SetAmountSoldNil(b bool)`

 SetAmountSoldNil sets the value for AmountSold to be an explicit nil

### UnsetAmountSold
`func (o *Order) UnsetAmountSold()`

UnsetAmountSold ensures that no value is present for AmountSold, not even an explicit nil
### GetBuy

`func (o *Order) GetBuy() Token`

GetBuy returns the Buy field if non-nil, zero value otherwise.

### GetBuyOk

`func (o *Order) GetBuyOk() (*Token, bool)`

GetBuyOk returns a tuple with the Buy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuy

`func (o *Order) SetBuy(v Token)`

SetBuy sets Buy field to given value.


### GetExpirationTimestamp

`func (o *Order) GetExpirationTimestamp() string`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *Order) GetExpirationTimestampOk() (*string, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *Order) SetExpirationTimestamp(v string)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### SetExpirationTimestampNil

`func (o *Order) SetExpirationTimestampNil(b bool)`

 SetExpirationTimestampNil sets the value for ExpirationTimestamp to be an explicit nil

### UnsetExpirationTimestamp
`func (o *Order) UnsetExpirationTimestamp()`

UnsetExpirationTimestamp ensures that no value is present for ExpirationTimestamp, not even an explicit nil
### GetFees

`func (o *Order) GetFees() []OrderFeeInfo`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Order) GetFeesOk() (*[]OrderFeeInfo, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Order) SetFees(v []OrderFeeInfo)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Order) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetOrderId

`func (o *Order) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetSell

`func (o *Order) GetSell() Token`

GetSell returns the Sell field if non-nil, zero value otherwise.

### GetSellOk

`func (o *Order) GetSellOk() (*Token, bool)`

GetSellOk returns a tuple with the Sell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSell

`func (o *Order) SetSell(v Token)`

SetSell sets Sell field to given value.


### GetStatus

`func (o *Order) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Order) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Order) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Order) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Order) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Order) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetUpdatedTimestamp

`func (o *Order) GetUpdatedTimestamp() string`

GetUpdatedTimestamp returns the UpdatedTimestamp field if non-nil, zero value otherwise.

### GetUpdatedTimestampOk

`func (o *Order) GetUpdatedTimestampOk() (*string, bool)`

GetUpdatedTimestampOk returns a tuple with the UpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTimestamp

`func (o *Order) SetUpdatedTimestamp(v string)`

SetUpdatedTimestamp sets UpdatedTimestamp field to given value.


### SetUpdatedTimestampNil

`func (o *Order) SetUpdatedTimestampNil(b bool)`

 SetUpdatedTimestampNil sets the value for UpdatedTimestamp to be an explicit nil

### UnsetUpdatedTimestamp
`func (o *Order) UnsetUpdatedTimestamp()`

UnsetUpdatedTimestamp ensures that no value is present for UpdatedTimestamp, not even an explicit nil
### GetUser

`func (o *Order) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Order) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Order) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


