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

// OrderDetails struct for OrderDetails
type OrderDetails struct {
	// Buy orders for this asset
	BuyOrders []map[string]interface{} `json:"buy_orders,omitempty"`
	// Sell orders for this asset
	SellOrders []map[string]interface{} `json:"sell_orders,omitempty"`
}

// NewOrderDetails instantiates a new OrderDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderDetails() *OrderDetails {
	this := OrderDetails{}
	return &this
}

// NewOrderDetailsWithDefaults instantiates a new OrderDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderDetailsWithDefaults() *OrderDetails {
	this := OrderDetails{}
	return &this
}

// GetBuyOrders returns the BuyOrders field value if set, zero value otherwise.
func (o *OrderDetails) GetBuyOrders() []map[string]interface{} {
	if o == nil || o.BuyOrders == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.BuyOrders
}

// GetBuyOrdersOk returns a tuple with the BuyOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetBuyOrdersOk() ([]map[string]interface{}, bool) {
	if o == nil || o.BuyOrders == nil {
		return nil, false
	}
	return o.BuyOrders, true
}

// HasBuyOrders returns a boolean if a field has been set.
func (o *OrderDetails) HasBuyOrders() bool {
	if o != nil && o.BuyOrders != nil {
		return true
	}

	return false
}

// SetBuyOrders gets a reference to the given []map[string]interface{} and assigns it to the BuyOrders field.
func (o *OrderDetails) SetBuyOrders(v []map[string]interface{}) {
	o.BuyOrders = v
}

// GetSellOrders returns the SellOrders field value if set, zero value otherwise.
func (o *OrderDetails) GetSellOrders() []map[string]interface{} {
	if o == nil || o.SellOrders == nil {
		var ret []map[string]interface{}
		return ret
	}
	return o.SellOrders
}

// GetSellOrdersOk returns a tuple with the SellOrders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderDetails) GetSellOrdersOk() ([]map[string]interface{}, bool) {
	if o == nil || o.SellOrders == nil {
		return nil, false
	}
	return o.SellOrders, true
}

// HasSellOrders returns a boolean if a field has been set.
func (o *OrderDetails) HasSellOrders() bool {
	if o != nil && o.SellOrders != nil {
		return true
	}

	return false
}

// SetSellOrders gets a reference to the given []map[string]interface{} and assigns it to the SellOrders field.
func (o *OrderDetails) SetSellOrders(v []map[string]interface{}) {
	o.SellOrders = v
}

func (o OrderDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuyOrders != nil {
		toSerialize["buy_orders"] = o.BuyOrders
	}
	if o.SellOrders != nil {
		toSerialize["sell_orders"] = o.SellOrders
	}
	return json.Marshal(toSerialize)
}

type NullableOrderDetails struct {
	value *OrderDetails
	isSet bool
}

func (v NullableOrderDetails) Get() *OrderDetails {
	return v.value
}

func (v *NullableOrderDetails) Set(val *OrderDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderDetails(val *OrderDetails) *NullableOrderDetails {
	return &NullableOrderDetails{value: val, isSet: true}
}

func (v NullableOrderDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


