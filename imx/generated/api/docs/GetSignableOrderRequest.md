# GetSignableOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBuy** | **string** | Fee-exclusive amount to buy the asset | 
**AmountSell** | **string** | Amount to sell (quantity) | 
**ExpirationTimestamp** | Pointer to **int32** | ExpirationTimestamp in Unix time. Note: will be rounded down to the nearest hour | [optional] 
**Fees** | Pointer to [**[]FeeEntry**](FeeEntry.md) | Inclusion of either maker or taker fees | [optional] 
**TokenBuy** | [**SignableToken**](SignableToken.md) |  | 
**TokenSell** | [**SignableToken**](SignableToken.md) |  | 
**User** | **string** | Ethereum address of the submitting user | 

## Methods

### NewGetSignableOrderRequest

`func NewGetSignableOrderRequest(amountBuy string, amountSell string, tokenBuy SignableToken, tokenSell SignableToken, user string, ) *GetSignableOrderRequest`

NewGetSignableOrderRequest instantiates a new GetSignableOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableOrderRequestWithDefaults

`func NewGetSignableOrderRequestWithDefaults() *GetSignableOrderRequest`

NewGetSignableOrderRequestWithDefaults instantiates a new GetSignableOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBuy

`func (o *GetSignableOrderRequest) GetAmountBuy() string`

GetAmountBuy returns the AmountBuy field if non-nil, zero value otherwise.

### GetAmountBuyOk

`func (o *GetSignableOrderRequest) GetAmountBuyOk() (*string, bool)`

GetAmountBuyOk returns a tuple with the AmountBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBuy

`func (o *GetSignableOrderRequest) SetAmountBuy(v string)`

SetAmountBuy sets AmountBuy field to given value.


### GetAmountSell

`func (o *GetSignableOrderRequest) GetAmountSell() string`

GetAmountSell returns the AmountSell field if non-nil, zero value otherwise.

### GetAmountSellOk

`func (o *GetSignableOrderRequest) GetAmountSellOk() (*string, bool)`

GetAmountSellOk returns a tuple with the AmountSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSell

`func (o *GetSignableOrderRequest) SetAmountSell(v string)`

SetAmountSell sets AmountSell field to given value.


### GetExpirationTimestamp

`func (o *GetSignableOrderRequest) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetSignableOrderRequest) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetSignableOrderRequest) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetSignableOrderRequest) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetFees

`func (o *GetSignableOrderRequest) GetFees() []FeeEntry`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *GetSignableOrderRequest) GetFeesOk() (*[]FeeEntry, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *GetSignableOrderRequest) SetFees(v []FeeEntry)`

SetFees sets Fees field to given value.

### HasFees

`func (o *GetSignableOrderRequest) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetTokenBuy

`func (o *GetSignableOrderRequest) GetTokenBuy() SignableToken`

GetTokenBuy returns the TokenBuy field if non-nil, zero value otherwise.

### GetTokenBuyOk

`func (o *GetSignableOrderRequest) GetTokenBuyOk() (*SignableToken, bool)`

GetTokenBuyOk returns a tuple with the TokenBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenBuy

`func (o *GetSignableOrderRequest) SetTokenBuy(v SignableToken)`

SetTokenBuy sets TokenBuy field to given value.


### GetTokenSell

`func (o *GetSignableOrderRequest) GetTokenSell() SignableToken`

GetTokenSell returns the TokenSell field if non-nil, zero value otherwise.

### GetTokenSellOk

`func (o *GetSignableOrderRequest) GetTokenSellOk() (*SignableToken, bool)`

GetTokenSellOk returns a tuple with the TokenSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenSell

`func (o *GetSignableOrderRequest) SetTokenSell(v SignableToken)`

SetTokenSell sets TokenSell field to given value.


### GetUser

`func (o *GetSignableOrderRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSignableOrderRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSignableOrderRequest) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


