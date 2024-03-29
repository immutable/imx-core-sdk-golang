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

// EncodeAssetRequestToken struct for EncodeAssetRequestToken
type EncodeAssetRequestToken struct {
	Data *EncodeAssetTokenData `json:"data,omitempty"`
	// The type of the token to be encoded
	Type *string `json:"type,omitempty"`
}

// NewEncodeAssetRequestToken instantiates a new EncodeAssetRequestToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncodeAssetRequestToken() *EncodeAssetRequestToken {
	this := EncodeAssetRequestToken{}
	return &this
}

// NewEncodeAssetRequestTokenWithDefaults instantiates a new EncodeAssetRequestToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncodeAssetRequestTokenWithDefaults() *EncodeAssetRequestToken {
	this := EncodeAssetRequestToken{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *EncodeAssetRequestToken) GetData() EncodeAssetTokenData {
	if o == nil || o.Data == nil {
		var ret EncodeAssetTokenData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodeAssetRequestToken) GetDataOk() (*EncodeAssetTokenData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *EncodeAssetRequestToken) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given EncodeAssetTokenData and assigns it to the Data field.
func (o *EncodeAssetRequestToken) SetData(v EncodeAssetTokenData) {
	o.Data = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EncodeAssetRequestToken) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodeAssetRequestToken) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EncodeAssetRequestToken) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *EncodeAssetRequestToken) SetType(v string) {
	o.Type = &v
}

func (o EncodeAssetRequestToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableEncodeAssetRequestToken struct {
	value *EncodeAssetRequestToken
	isSet bool
}

func (v NullableEncodeAssetRequestToken) Get() *EncodeAssetRequestToken {
	return v.value
}

func (v *NullableEncodeAssetRequestToken) Set(val *EncodeAssetRequestToken) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodeAssetRequestToken) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodeAssetRequestToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodeAssetRequestToken(val *EncodeAssetRequestToken) *NullableEncodeAssetRequestToken {
	return &NullableEncodeAssetRequestToken{value: val, isSet: true}
}

func (v NullableEncodeAssetRequestToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodeAssetRequestToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


