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

// SignableToken struct for SignableToken
type SignableToken struct {
	// Token data. See https://docs.x.immutable.com/docs/token-data-object\"
	Data map[string]interface{} `json:"data,omitempty"`
	// Type of token
	Type *string `json:"type,omitempty"`
}

// NewSignableToken instantiates a new SignableToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignableToken() *SignableToken {
	this := SignableToken{}
	return &this
}

// NewSignableTokenWithDefaults instantiates a new SignableToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignableTokenWithDefaults() *SignableToken {
	this := SignableToken{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SignableToken) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignableToken) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SignableToken) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *SignableToken) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SignableToken) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignableToken) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SignableToken) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SignableToken) SetType(v string) {
	o.Type = &v
}

func (o SignableToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSignableToken struct {
	value *SignableToken
	isSet bool
}

func (v NullableSignableToken) Get() *SignableToken {
	return v.value
}

func (v *NullableSignableToken) Set(val *SignableToken) {
	v.value = val
	v.isSet = true
}

func (v NullableSignableToken) IsSet() bool {
	return v.isSet
}

func (v *NullableSignableToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignableToken(val *SignableToken) *NullableSignableToken {
	return &NullableSignableToken{value: val, isSet: true}
}

func (v NullableSignableToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignableToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


