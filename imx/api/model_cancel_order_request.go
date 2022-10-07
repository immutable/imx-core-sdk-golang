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

// CancelOrderRequest struct for CancelOrderRequest
type CancelOrderRequest struct {
	// ID of the order
	OrderId int32 `json:"order_id"`
	// Payload signature
	StarkSignature string `json:"stark_signature"`
}

// NewCancelOrderRequest instantiates a new CancelOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderRequest(orderId int32, starkSignature string) *CancelOrderRequest {
	this := CancelOrderRequest{}
	this.OrderId = orderId
	this.StarkSignature = starkSignature
	return &this
}

// NewCancelOrderRequestWithDefaults instantiates a new CancelOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderRequestWithDefaults() *CancelOrderRequest {
	this := CancelOrderRequest{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *CancelOrderRequest) GetOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CancelOrderRequest) SetOrderId(v int32) {
	o.OrderId = v
}

// GetStarkSignature returns the StarkSignature field value
func (o *CancelOrderRequest) GetStarkSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarkSignature
}

// GetStarkSignatureOk returns a tuple with the StarkSignature field value
// and a boolean to check if the value has been set.
func (o *CancelOrderRequest) GetStarkSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarkSignature, true
}

// SetStarkSignature sets field value
func (o *CancelOrderRequest) SetStarkSignature(v string) {
	o.StarkSignature = v
}

func (o CancelOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if true {
		toSerialize["stark_signature"] = o.StarkSignature
	}
	return json.Marshal(toSerialize)
}

type NullableCancelOrderRequest struct {
	value *CancelOrderRequest
	isSet bool
}

func (v NullableCancelOrderRequest) Get() *CancelOrderRequest {
	return v.value
}

func (v *NullableCancelOrderRequest) Set(val *CancelOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderRequest(val *CancelOrderRequest) *NullableCancelOrderRequest {
	return &NullableCancelOrderRequest{value: val, isSet: true}
}

func (v NullableCancelOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


