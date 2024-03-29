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

// CancelOrderResponse struct for CancelOrderResponse
type CancelOrderResponse struct {
	// ID of the cancelled order
	OrderId int32 `json:"order_id"`
	// New status of the order
	Status string `json:"status"`
}

// NewCancelOrderResponse instantiates a new CancelOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelOrderResponse(orderId int32, status string) *CancelOrderResponse {
	this := CancelOrderResponse{}
	this.OrderId = orderId
	this.Status = status
	return &this
}

// NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelOrderResponseWithDefaults() *CancelOrderResponse {
	this := CancelOrderResponse{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *CancelOrderResponse) GetOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *CancelOrderResponse) SetOrderId(v int32) {
	o.OrderId = v
}

// GetStatus returns the Status field value
func (o *CancelOrderResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CancelOrderResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CancelOrderResponse) SetStatus(v string) {
	o.Status = v
}

func (o CancelOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["order_id"] = o.OrderId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableCancelOrderResponse struct {
	value *CancelOrderResponse
	isSet bool
}

func (v NullableCancelOrderResponse) Get() *CancelOrderResponse {
	return v.value
}

func (v *NullableCancelOrderResponse) Set(val *CancelOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelOrderResponse(val *CancelOrderResponse) *NullableCancelOrderResponse {
	return &NullableCancelOrderResponse{value: val, isSet: true}
}

func (v NullableCancelOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


