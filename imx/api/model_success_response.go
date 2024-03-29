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

// SuccessResponse struct for SuccessResponse
type SuccessResponse struct {
	Result string `json:"result"`
}

// NewSuccessResponse instantiates a new SuccessResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSuccessResponse(result string) *SuccessResponse {
	this := SuccessResponse{}
	this.Result = result
	return &this
}

// NewSuccessResponseWithDefaults instantiates a new SuccessResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuccessResponseWithDefaults() *SuccessResponse {
	this := SuccessResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *SuccessResponse) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *SuccessResponse) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *SuccessResponse) SetResult(v string) {
	o.Result = v
}

func (o SuccessResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableSuccessResponse struct {
	value *SuccessResponse
	isSet bool
}

func (v NullableSuccessResponse) Get() *SuccessResponse {
	return v.value
}

func (v *NullableSuccessResponse) Set(val *SuccessResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSuccessResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSuccessResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuccessResponse(val *SuccessResponse) *NullableSuccessResponse {
	return &NullableSuccessResponse{value: val, isSet: true}
}

func (v NullableSuccessResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuccessResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


