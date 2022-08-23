# Balance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Balance** | **string** | Amount which is currently inside the exchange | 
**PreparingWithdrawal** | **string** | Amount which is currently preparing withdrawal from the exchange | 
**Symbol** | **string** | Symbol of the token (e.g. ETH, IMX) | 
**TokenAddress** | **string** | Address of the contract that represents this ERC20 token or empty for ETH | 
**Withdrawable** | **string** | Amount which is currently withdrawable from the exchange | 

## Methods

### NewBalance

`func NewBalance(balance string, preparingWithdrawal string, symbol string, tokenAddress string, withdrawable string, ) *Balance`

NewBalance instantiates a new Balance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceWithDefaults

`func NewBalanceWithDefaults() *Balance`

NewBalanceWithDefaults instantiates a new Balance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalance

`func (o *Balance) GetBalance() string`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *Balance) GetBalanceOk() (*string, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *Balance) SetBalance(v string)`

SetBalance sets Balance field to given value.


### GetPreparingWithdrawal

`func (o *Balance) GetPreparingWithdrawal() string`

GetPreparingWithdrawal returns the PreparingWithdrawal field if non-nil, zero value otherwise.

### GetPreparingWithdrawalOk

`func (o *Balance) GetPreparingWithdrawalOk() (*string, bool)`

GetPreparingWithdrawalOk returns a tuple with the PreparingWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparingWithdrawal

`func (o *Balance) SetPreparingWithdrawal(v string)`

SetPreparingWithdrawal sets PreparingWithdrawal field to given value.


### GetSymbol

`func (o *Balance) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Balance) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Balance) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenAddress

`func (o *Balance) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *Balance) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *Balance) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetWithdrawable

`func (o *Balance) GetWithdrawable() string`

GetWithdrawable returns the Withdrawable field if non-nil, zero value otherwise.

### GetWithdrawableOk

`func (o *Balance) GetWithdrawableOk() (*string, bool)`

GetWithdrawableOk returns a tuple with the Withdrawable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawable

`func (o *Balance) SetWithdrawable(v string)`

SetWithdrawable sets Withdrawable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


