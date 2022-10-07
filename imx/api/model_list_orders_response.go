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

// ListOrdersResponse struct for ListOrdersResponse
type ListOrdersResponse struct {
	// Generated cursor returned by previous query
	Cursor string `json:"cursor"`
	// Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results
	Remaining int32 `json:"remaining"`
	// Orders matching query parameters
	Result []Order `json:"result"`
}

// NewListOrdersResponse instantiates a new ListOrdersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOrdersResponse(cursor string, remaining int32, result []Order) *ListOrdersResponse {
	this := ListOrdersResponse{}
	this.Cursor = cursor
	this.Remaining = remaining
	this.Result = result
	return &this
}

// NewListOrdersResponseWithDefaults instantiates a new ListOrdersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOrdersResponseWithDefaults() *ListOrdersResponse {
	this := ListOrdersResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListOrdersResponse) GetCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListOrdersResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListOrdersResponse) SetCursor(v string) {
	o.Cursor = v
}

// GetRemaining returns the Remaining field value
func (o *ListOrdersResponse) GetRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *ListOrdersResponse) GetRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *ListOrdersResponse) SetRemaining(v int32) {
	o.Remaining = v
}

// GetResult returns the Result field value
func (o *ListOrdersResponse) GetResult() []Order {
	if o == nil {
		var ret []Order
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListOrdersResponse) GetResultOk() ([]Order, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListOrdersResponse) SetResult(v []Order) {
	o.Result = v
}

func (o ListOrdersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	if true {
		toSerialize["remaining"] = o.Remaining
	}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableListOrdersResponse struct {
	value *ListOrdersResponse
	isSet bool
}

func (v NullableListOrdersResponse) Get() *ListOrdersResponse {
	return v.value
}

func (v *NullableListOrdersResponse) Set(val *ListOrdersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListOrdersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListOrdersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOrdersResponse(val *ListOrdersResponse) *NullableListOrdersResponse {
	return &NullableListOrdersResponse{value: val, isSet: true}
}

func (v NullableListOrdersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOrdersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


