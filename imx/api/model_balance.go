/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.0
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Balance struct for Balance
type Balance struct {
	// Amount which is currently inside the exchange
	Balance string `json:"balance"`
	// Amount which is currently preparing withdrawal from the exchange
	PreparingWithdrawal string `json:"preparing_withdrawal"`
	// Symbol of the token (e.g. ETH, IMX)
	Symbol string `json:"symbol"`
	// Address of the contract that represents this ERC20 token or empty for ETH
	TokenAddress string `json:"token_address"`
	// Amount which is currently withdrawable from the exchange
	Withdrawable string `json:"withdrawable"`
}

// NewBalance instantiates a new Balance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBalance(balance string, preparingWithdrawal string, symbol string, tokenAddress string, withdrawable string) *Balance {
	this := Balance{}
	this.Balance = balance
	this.PreparingWithdrawal = preparingWithdrawal
	this.Symbol = symbol
	this.TokenAddress = tokenAddress
	this.Withdrawable = withdrawable
	return &this
}

// NewBalanceWithDefaults instantiates a new Balance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBalanceWithDefaults() *Balance {
	this := Balance{}
	return &this
}

// GetBalance returns the Balance field value
func (o *Balance) GetBalance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Balance
}

// GetBalanceOk returns a tuple with the Balance field value
// and a boolean to check if the value has been set.
func (o *Balance) GetBalanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Balance, true
}

// SetBalance sets field value
func (o *Balance) SetBalance(v string) {
	o.Balance = v
}

// GetPreparingWithdrawal returns the PreparingWithdrawal field value
func (o *Balance) GetPreparingWithdrawal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PreparingWithdrawal
}

// GetPreparingWithdrawalOk returns a tuple with the PreparingWithdrawal field value
// and a boolean to check if the value has been set.
func (o *Balance) GetPreparingWithdrawalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreparingWithdrawal, true
}

// SetPreparingWithdrawal sets field value
func (o *Balance) SetPreparingWithdrawal(v string) {
	o.PreparingWithdrawal = v
}

// GetSymbol returns the Symbol field value
func (o *Balance) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *Balance) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *Balance) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenAddress returns the TokenAddress field value
func (o *Balance) GetTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value
// and a boolean to check if the value has been set.
func (o *Balance) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAddress, true
}

// SetTokenAddress sets field value
func (o *Balance) SetTokenAddress(v string) {
	o.TokenAddress = v
}

// GetWithdrawable returns the Withdrawable field value
func (o *Balance) GetWithdrawable() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Withdrawable
}

// GetWithdrawableOk returns a tuple with the Withdrawable field value
// and a boolean to check if the value has been set.
func (o *Balance) GetWithdrawableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Withdrawable, true
}

// SetWithdrawable sets field value
func (o *Balance) SetWithdrawable(v string) {
	o.Withdrawable = v
}

func (o Balance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["balance"] = o.Balance
	}
	if true {
		toSerialize["preparing_withdrawal"] = o.PreparingWithdrawal
	}
	if true {
		toSerialize["symbol"] = o.Symbol
	}
	if true {
		toSerialize["token_address"] = o.TokenAddress
	}
	if true {
		toSerialize["withdrawable"] = o.Withdrawable
	}
	return json.Marshal(toSerialize)
}

type NullableBalance struct {
	value *Balance
	isSet bool
}

func (v NullableBalance) Get() *Balance {
	return v.value
}

func (v *NullableBalance) Set(val *Balance) {
	v.value = val
	v.isSet = true
}

func (v NullableBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBalance(val *Balance) *NullableBalance {
	return &NullableBalance{value: val, isSet: true}
}

func (v NullableBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


