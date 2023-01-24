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

// TradeA Side A of this trade (the buy order)
type TradeA struct {
	// The ID of the order involved in the trade
	OrderId int32 `json:"order_id"`
	// The amount of that order's asset this trade sold
	Sold string `json:"sold"`
	// The contract address of the token that this trade sold
	TokenAddress *string `json:"token_address,omitempty"`
	// The ID of the token that this trade sold
	TokenId *string `json:"token_id,omitempty"`
	// The type of the token that this trade sold
	TokenType string `json:"token_type"`
}

// NewTradeA instantiates a new TradeA object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTradeA(orderId int32, sold string, tokenType string) *TradeA {
	this := TradeA{}
	this.OrderId = orderId
	this.Sold = sold
	this.TokenType = tokenType
	return &this
}

// NewTradeAWithDefaults instantiates a new TradeA object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeAWithDefaults() *TradeA {
	this := TradeA{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *TradeA) GetOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *TradeA) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *TradeA) SetOrderId(v int32) {
	o.OrderId = v
}

// GetSold returns the Sold field value
func (o *TradeA) GetSold() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sold
}

// GetSoldOk returns a tuple with the Sold field value
// and a boolean to check if the value has been set.
func (o *TradeA) GetSoldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sold, true
}

// SetSold sets field value
func (o *TradeA) SetSold(v string) {
	o.Sold = v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *TradeA) GetTokenAddress() string {
	if o == nil || o.TokenAddress == nil {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeA) GetTokenAddressOk() (*string, bool) {
	if o == nil || o.TokenAddress == nil {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *TradeA) HasTokenAddress() bool {
	if o != nil && o.TokenAddress != nil {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *TradeA) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TradeA) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TradeA) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TradeA) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TradeA) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTokenType returns the TokenType field value
func (o *TradeA) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *TradeA) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *TradeA) SetTokenType(v string) {
	o.TokenType = v
}

func (o TradeA) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if true {
		toSerialize["sold"] = o.Sold
	}
	if o.TokenAddress != nil {
		toSerialize["token_address"] = o.TokenAddress
	}
	if o.TokenId != nil {
		toSerialize["token_id"] = o.TokenId
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	return json.Marshal(toSerialize)
}

type NullableTradeA struct {
	value *TradeA
	isSet bool
}

func (v NullableTradeA) Get() *TradeA {
	return v.value
}

func (v *NullableTradeA) Set(val *TradeA) {
	v.value = val
	v.isSet = true
}

func (v NullableTradeA) IsSet() bool {
	return v.isSet
}

func (v *NullableTradeA) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTradeA(val *TradeA) *NullableTradeA {
	return &NullableTradeA{value: val, isSet: true}
}

func (v NullableTradeA) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTradeA) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


