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

// OrderFeeInfo struct for OrderFeeInfo
type OrderFeeInfo struct {
	// Address of the fee recipient
	Address *string `json:"address,omitempty"`
	// Fee amount
	Amount *string `json:"amount,omitempty"`
	Token *FeeToken `json:"token,omitempty"`
	// Fee type
	Type *string `json:"type,omitempty"`
}

// NewOrderFeeInfo instantiates a new OrderFeeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderFeeInfo() *OrderFeeInfo {
	this := OrderFeeInfo{}
	return &this
}

// NewOrderFeeInfoWithDefaults instantiates a new OrderFeeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderFeeInfoWithDefaults() *OrderFeeInfo {
	this := OrderFeeInfo{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *OrderFeeInfo) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFeeInfo) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *OrderFeeInfo) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *OrderFeeInfo) SetAddress(v string) {
	o.Address = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OrderFeeInfo) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFeeInfo) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OrderFeeInfo) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *OrderFeeInfo) SetAmount(v string) {
	o.Amount = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *OrderFeeInfo) GetToken() FeeToken {
	if o == nil || o.Token == nil {
		var ret FeeToken
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFeeInfo) GetTokenOk() (*FeeToken, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *OrderFeeInfo) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given FeeToken and assigns it to the Token field.
func (o *OrderFeeInfo) SetToken(v FeeToken) {
	o.Token = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrderFeeInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderFeeInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrderFeeInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *OrderFeeInfo) SetType(v string) {
	o.Type = &v
}

func (o OrderFeeInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableOrderFeeInfo struct {
	value *OrderFeeInfo
	isSet bool
}

func (v NullableOrderFeeInfo) Get() *OrderFeeInfo {
	return v.value
}

func (v *NullableOrderFeeInfo) Set(val *OrderFeeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderFeeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderFeeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderFeeInfo(val *OrderFeeInfo) *NullableOrderFeeInfo {
	return &NullableOrderFeeInfo{value: val, isSet: true}
}

func (v NullableOrderFeeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderFeeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


