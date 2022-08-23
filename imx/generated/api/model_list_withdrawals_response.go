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

// ListWithdrawalsResponse struct for ListWithdrawalsResponse
type ListWithdrawalsResponse struct {
	// Generated cursor returned by previous query
	Cursor string `json:"cursor"`
	// Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results
	Remaining int32 `json:"remaining"`
	// Withdrawals matching query parameters
	Result []Withdrawal `json:"result"`
}

// NewListWithdrawalsResponse instantiates a new ListWithdrawalsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWithdrawalsResponse(cursor string, remaining int32, result []Withdrawal) *ListWithdrawalsResponse {
	this := ListWithdrawalsResponse{}
	this.Cursor = cursor
	this.Remaining = remaining
	this.Result = result
	return &this
}

// NewListWithdrawalsResponseWithDefaults instantiates a new ListWithdrawalsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWithdrawalsResponseWithDefaults() *ListWithdrawalsResponse {
	this := ListWithdrawalsResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListWithdrawalsResponse) GetCursor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListWithdrawalsResponse) GetCursorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListWithdrawalsResponse) SetCursor(v string) {
	o.Cursor = v
}

// GetRemaining returns the Remaining field value
func (o *ListWithdrawalsResponse) GetRemaining() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value
// and a boolean to check if the value has been set.
func (o *ListWithdrawalsResponse) GetRemainingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Remaining, true
}

// SetRemaining sets field value
func (o *ListWithdrawalsResponse) SetRemaining(v int32) {
	o.Remaining = v
}

// GetResult returns the Result field value
func (o *ListWithdrawalsResponse) GetResult() []Withdrawal {
	if o == nil {
		var ret []Withdrawal
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ListWithdrawalsResponse) GetResultOk() ([]Withdrawal, bool) {
	if o == nil {
		return nil, false
	}
	return o.Result, true
}

// SetResult sets field value
func (o *ListWithdrawalsResponse) SetResult(v []Withdrawal) {
	o.Result = v
}

func (o ListWithdrawalsResponse) MarshalJSON() ([]byte, error) {
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

type NullableListWithdrawalsResponse struct {
	value *ListWithdrawalsResponse
	isSet bool
}

func (v NullableListWithdrawalsResponse) Get() *ListWithdrawalsResponse {
	return v.value
}

func (v *NullableListWithdrawalsResponse) Set(val *ListWithdrawalsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWithdrawalsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWithdrawalsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWithdrawalsResponse(val *ListWithdrawalsResponse) *NullableListWithdrawalsResponse {
	return &NullableListWithdrawalsResponse{value: val, isSet: true}
}

func (v NullableListWithdrawalsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWithdrawalsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


