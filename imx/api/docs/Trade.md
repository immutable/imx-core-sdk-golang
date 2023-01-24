# Trade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**A** | [**TradeA**](TradeA.md) |  | 
**B** | [**TradeB**](TradeB.md) |  | 
**Status** | **string** | Status of this trade | 
**Timestamp** | **NullableString** | Time this trade occurred | 
**TransactionId** | **int32** | Sequential ID of this trade within Immutable X | 

## Methods

### NewTrade

`func NewTrade(a TradeA, b TradeB, status string, timestamp NullableString, transactionId int32, ) *Trade`

NewTrade instantiates a new Trade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeWithDefaults

`func NewTradeWithDefaults() *Trade`

NewTradeWithDefaults instantiates a new Trade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetA

`func (o *Trade) GetA() TradeA`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *Trade) GetAOk() (*TradeA, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *Trade) SetA(v TradeA)`

SetA sets A field to given value.


### GetB

`func (o *Trade) GetB() TradeB`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *Trade) GetBOk() (*TradeB, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *Trade) SetB(v TradeB)`

SetB sets B field to given value.


### GetStatus

`func (o *Trade) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Trade) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Trade) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Trade) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Trade) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Trade) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Trade) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Trade) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetTransactionId

`func (o *Trade) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Trade) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Trade) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


