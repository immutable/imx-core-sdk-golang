/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// CreateTradeResponse struct for CreateTradeResponse
type CreateTradeResponse struct {
	// Request ID that returns when a trade initiated through risk-manager
	RequestId *string `json:"request_id,omitempty"`
	// Current status of trade
	Status string `json:"status"`
	// ID of trade within Immutable X
	TradeId int32 `json:"trade_id"`
}

// NewCreateTradeResponse instantiates a new CreateTradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTradeResponse(status string, tradeId int32) *CreateTradeResponse {
	this := CreateTradeResponse{}
	this.Status = status
	this.TradeId = tradeId
	return &this
}

// NewCreateTradeResponseWithDefaults instantiates a new CreateTradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTradeResponseWithDefaults() *CreateTradeResponse {
	this := CreateTradeResponse{}
	return &this
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateTradeResponse) GetRequestId() string {
	if o == nil || o.RequestId == nil {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTradeResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || o.RequestId == nil {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateTradeResponse) HasRequestId() bool {
	if o != nil && o.RequestId != nil {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateTradeResponse) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatus returns the Status field value
func (o *CreateTradeResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateTradeResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateTradeResponse) SetStatus(v string) {
	o.Status = v
}

// GetTradeId returns the TradeId field value
func (o *CreateTradeResponse) GetTradeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TradeId
}

// GetTradeIdOk returns a tuple with the TradeId field value
// and a boolean to check if the value has been set.
func (o *CreateTradeResponse) GetTradeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TradeId, true
}

// SetTradeId sets field value
func (o *CreateTradeResponse) SetTradeId(v int32) {
	o.TradeId = v
}

func (o CreateTradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RequestId != nil {
		toSerialize["request_id"] = o.RequestId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["trade_id"] = o.TradeId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTradeResponse struct {
	value *CreateTradeResponse
	isSet bool
}

func (v NullableCreateTradeResponse) Get() *CreateTradeResponse {
	return v.value
}

func (v *NullableCreateTradeResponse) Set(val *CreateTradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTradeResponse(val *CreateTradeResponse) *NullableCreateTradeResponse {
	return &NullableCreateTradeResponse{value: val, isSet: true}
}

func (v NullableCreateTradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

