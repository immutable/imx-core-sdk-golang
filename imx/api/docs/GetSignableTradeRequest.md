# GetSignableTradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTimestamp** | Pointer to **int32** | ExpirationTimestamp in Unix time. Note: will be rounded down to the nearest hour | [optional] 
**Fees** | Pointer to [**[]FeeEntry**](FeeEntry.md) | Inclusion of either maker or taker fees | [optional] 
**OrderId** | **int32** | The ID of the maker order involved | 
**User** | **string** | Ethereum address of the submitting user | 

## Methods

### NewGetSignableTradeRequest

`func NewGetSignableTradeRequest(orderId int32, user string, ) *GetSignableTradeRequest`

NewGetSignableTradeRequest instantiates a new GetSignableTradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTradeRequestWithDefaults

`func NewGetSignableTradeRequestWithDefaults() *GetSignableTradeRequest`

NewGetSignableTradeRequestWithDefaults instantiates a new GetSignableTradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationTimestamp

`func (o *GetSignableTradeRequest) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetSignableTradeRequest) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetSignableTradeRequest) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetSignableTradeRequest) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetFees

`func (o *GetSignableTradeRequest) GetFees() []FeeEntry`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *GetSignableTradeRequest) GetFeesOk() (*[]FeeEntry, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *GetSignableTradeRequest) SetFees(v []FeeEntry)`

SetFees sets Fees field to given value.

### HasFees

`func (o *GetSignableTradeRequest) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetOrderId

`func (o *GetSignableTradeRequest) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSignableTradeRequest) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSignableTradeRequest) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetUser

`func (o *GetSignableTradeRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSignableTradeRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSignableTradeRequest) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


