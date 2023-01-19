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

// Trade struct for Trade
type Trade struct {
	A TradeSide `json:"a"`
	B TradeSide `json:"b"`
	// Status of this trade
	Status string `json:"status"`
	// Time this trade occurred
	Timestamp NullableString `json:"timestamp"`
	// Sequential ID of this trade within Immutable X
	TransactionId int32 `json:"transaction_id"`
}

// NewTrade instantiates a new Trade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrade(a TradeSide, b TradeSide, status string, timestamp NullableString, transactionId int32) *Trade {
	this := Trade{}
	this.A = a
	this.B = b
	this.Status = status
	this.Timestamp = timestamp
	this.TransactionId = transactionId
	return &this
}

// NewTradeWithDefaults instantiates a new Trade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTradeWithDefaults() *Trade {
	this := Trade{}
	return &this
}

// GetA returns the A field value
func (o *Trade) GetA() TradeSide {
	if o == nil {
		var ret TradeSide
		return ret
	}

	return o.A
}

// GetAOk returns a tuple with the A field value
// and a boolean to check if the value has been set.
func (o *Trade) GetAOk() (*TradeSide, bool) {
	if o == nil {
    return nil, false
	}
	return &o.A, true
}

// SetA sets field value
func (o *Trade) SetA(v TradeSide) {
	o.A = v
}

// GetB returns the B field value
func (o *Trade) GetB() TradeSide {
	if o == nil {
		var ret TradeSide
		return ret
	}

	return o.B
}

// GetBOk returns a tuple with the B field value
// and a boolean to check if the value has been set.
func (o *Trade) GetBOk() (*TradeSide, bool) {
	if o == nil {
    return nil, false
	}
	return &o.B, true
}

// SetB sets field value
func (o *Trade) SetB(v TradeSide) {
	o.B = v
}

// GetStatus returns the Status field value
func (o *Trade) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Trade) GetStatusOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Trade) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Trade) GetTimestamp() string {
	if o == nil || o.Timestamp.Get() == nil {
		var ret string
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Trade) GetTimestampOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *Trade) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}

// GetTransactionId returns the TransactionId field value
func (o *Trade) GetTransactionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Trade) GetTransactionIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Trade) SetTransactionId(v int32) {
	o.TransactionId = v
}

func (o Trade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["a"] = o.A
	}
	if true {
		toSerialize["b"] = o.B
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp.Get()
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	return json.Marshal(toSerialize)
}

type NullableTrade struct {
	value *Trade
	isSet bool
}

func (v NullableTrade) Get() *Trade {
	return v.value
}

func (v *NullableTrade) Set(val *Trade) {
	v.value = val
	v.isSet = true
}

func (v NullableTrade) IsSet() bool {
	return v.isSet
}

func (v *NullableTrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrade(val *Trade) *NullableTrade {
	return &NullableTrade{value: val, isSet: true}
}

func (v NullableTrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


