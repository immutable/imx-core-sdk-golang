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

// checks if the RegisterUserResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterUserResponse{}

// RegisterUserResponse struct for RegisterUserResponse
type RegisterUserResponse struct {
	// Immutable signature authorising registration
	TxHash string `json:"tx_hash"`
}

// NewRegisterUserResponse instantiates a new RegisterUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterUserResponse(txHash string) *RegisterUserResponse {
	this := RegisterUserResponse{}
	this.TxHash = txHash
	return &this
}

// NewRegisterUserResponseWithDefaults instantiates a new RegisterUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterUserResponseWithDefaults() *RegisterUserResponse {
	this := RegisterUserResponse{}
	return &this
}

// GetTxHash returns the TxHash field value
func (o *RegisterUserResponse) GetTxHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TxHash
}

// GetTxHashOk returns a tuple with the TxHash field value
// and a boolean to check if the value has been set.
func (o *RegisterUserResponse) GetTxHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxHash, true
}

// SetTxHash sets field value
func (o *RegisterUserResponse) SetTxHash(v string) {
	o.TxHash = v
}

func (o RegisterUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterUserResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tx_hash"] = o.TxHash
	return toSerialize, nil
}

type NullableRegisterUserResponse struct {
	value *RegisterUserResponse
	isSet bool
}

func (v NullableRegisterUserResponse) Get() *RegisterUserResponse {
	return v.value
}

func (v *NullableRegisterUserResponse) Set(val *RegisterUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterUserResponse(val *RegisterUserResponse) *NullableRegisterUserResponse {
	return &NullableRegisterUserResponse{value: val, isSet: true}
}

func (v NullableRegisterUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


