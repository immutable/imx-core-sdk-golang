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

// FeeToken struct for FeeToken
type FeeToken struct {
	Data *FeeData `json:"data,omitempty"`
	// Fee token type. One of ETH/ERC20
	Type *string `json:"type,omitempty"`
}

// NewFeeToken instantiates a new FeeToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeeToken() *FeeToken {
	this := FeeToken{}
	return &this
}

// NewFeeTokenWithDefaults instantiates a new FeeToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeeTokenWithDefaults() *FeeToken {
	this := FeeToken{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FeeToken) GetData() FeeData {
	if o == nil || isNil(o.Data) {
		var ret FeeData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeToken) GetDataOk() (*FeeData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FeeToken) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given FeeData and assigns it to the Data field.
func (o *FeeToken) SetData(v FeeData) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *FeeToken) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeeToken) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *FeeToken) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *FeeToken) SetType(v string) {
	o.Type = &v
}

func (o FeeToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableFeeToken struct {
	value *FeeToken
	isSet bool
}

func (v NullableFeeToken) Get() *FeeToken {
	return v.value
}

func (v *NullableFeeToken) Set(val *FeeToken) {
	v.value = val
	v.isSet = true
}

func (v NullableFeeToken) IsSet() bool {
	return v.isSet
}

func (v *NullableFeeToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeeToken(val *FeeToken) *NullableFeeToken {
	return &NullableFeeToken{value: val, isSet: true}
}

func (v NullableFeeToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeeToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


