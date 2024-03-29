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

// RegisterUserRequest struct for RegisterUserRequest
type RegisterUserRequest struct {
	// User email
	Email *string `json:"email,omitempty"`
	// Eth signature
	EthSignature string `json:"eth_signature"`
	// The ether key of the user
	EtherKey string `json:"ether_key"`
	// Public stark key of the user
	StarkKey string `json:"stark_key"`
	// Payload signature
	StarkSignature string `json:"stark_signature"`
}

// NewRegisterUserRequest instantiates a new RegisterUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterUserRequest(ethSignature string, etherKey string, starkKey string, starkSignature string) *RegisterUserRequest {
	this := RegisterUserRequest{}
	this.EthSignature = ethSignature
	this.EtherKey = etherKey
	this.StarkKey = starkKey
	this.StarkSignature = starkSignature
	return &this
}

// NewRegisterUserRequestWithDefaults instantiates a new RegisterUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterUserRequestWithDefaults() *RegisterUserRequest {
	this := RegisterUserRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *RegisterUserRequest) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterUserRequest) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *RegisterUserRequest) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *RegisterUserRequest) SetEmail(v string) {
	o.Email = &v
}

// GetEthSignature returns the EthSignature field value
func (o *RegisterUserRequest) GetEthSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EthSignature
}

// GetEthSignatureOk returns a tuple with the EthSignature field value
// and a boolean to check if the value has been set.
func (o *RegisterUserRequest) GetEthSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EthSignature, true
}

// SetEthSignature sets field value
func (o *RegisterUserRequest) SetEthSignature(v string) {
	o.EthSignature = v
}

// GetEtherKey returns the EtherKey field value
func (o *RegisterUserRequest) GetEtherKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EtherKey
}

// GetEtherKeyOk returns a tuple with the EtherKey field value
// and a boolean to check if the value has been set.
func (o *RegisterUserRequest) GetEtherKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EtherKey, true
}

// SetEtherKey sets field value
func (o *RegisterUserRequest) SetEtherKey(v string) {
	o.EtherKey = v
}

// GetStarkKey returns the StarkKey field value
func (o *RegisterUserRequest) GetStarkKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarkKey
}

// GetStarkKeyOk returns a tuple with the StarkKey field value
// and a boolean to check if the value has been set.
func (o *RegisterUserRequest) GetStarkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarkKey, true
}

// SetStarkKey sets field value
func (o *RegisterUserRequest) SetStarkKey(v string) {
	o.StarkKey = v
}

// GetStarkSignature returns the StarkSignature field value
func (o *RegisterUserRequest) GetStarkSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarkSignature
}

// GetStarkSignatureOk returns a tuple with the StarkSignature field value
// and a boolean to check if the value has been set.
func (o *RegisterUserRequest) GetStarkSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarkSignature, true
}

// SetStarkSignature sets field value
func (o *RegisterUserRequest) SetStarkSignature(v string) {
	o.StarkSignature = v
}

func (o RegisterUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if true {
		toSerialize["eth_signature"] = o.EthSignature
	}
	if true {
		toSerialize["ether_key"] = o.EtherKey
	}
	if true {
		toSerialize["stark_key"] = o.StarkKey
	}
	if true {
		toSerialize["stark_signature"] = o.StarkSignature
	}
	return json.Marshal(toSerialize)
}

type NullableRegisterUserRequest struct {
	value *RegisterUserRequest
	isSet bool
}

func (v NullableRegisterUserRequest) Get() *RegisterUserRequest {
	return v.value
}

func (v *NullableRegisterUserRequest) Set(val *RegisterUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterUserRequest(val *RegisterUserRequest) *NullableRegisterUserRequest {
	return &NullableRegisterUserRequest{value: val, isSet: true}
}

func (v NullableRegisterUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


