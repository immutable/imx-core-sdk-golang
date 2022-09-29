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

// MintUser struct for MintUser
type MintUser struct {
	// List of Mint tokens
	Tokens []MintTokenDataV2 `json:"tokens"`
	// User wallet address
	User string `json:"user"`
}

// NewMintUser instantiates a new MintUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMintUser(tokens []MintTokenDataV2, user string) *MintUser {
	this := MintUser{}
	this.Tokens = tokens
	this.User = user
	return &this
}

// NewMintUserWithDefaults instantiates a new MintUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMintUserWithDefaults() *MintUser {
	this := MintUser{}
	return &this
}

// GetTokens returns the Tokens field value
func (o *MintUser) GetTokens() []MintTokenDataV2 {
	if o == nil {
		var ret []MintTokenDataV2
		return ret
	}

	return o.Tokens
}

// GetTokensOk returns a tuple with the Tokens field value
// and a boolean to check if the value has been set.
func (o *MintUser) GetTokensOk() ([]MintTokenDataV2, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tokens, true
}

// SetTokens sets field value
func (o *MintUser) SetTokens(v []MintTokenDataV2) {
	o.Tokens = v
}

// GetUser returns the User field value
func (o *MintUser) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *MintUser) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *MintUser) SetUser(v string) {
	o.User = v
}

func (o MintUser) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tokens"] = o.Tokens
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableMintUser struct {
	value *MintUser
	isSet bool
}

func (v NullableMintUser) Get() *MintUser {
	return v.value
}

func (v *NullableMintUser) Set(val *MintUser) {
	v.value = val
	v.isSet = true
}

func (v NullableMintUser) IsSet() bool {
	return v.isSet
}

func (v *NullableMintUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMintUser(val *MintUser) *NullableMintUser {
	return &NullableMintUser{value: val, isSet: true}
}

func (v NullableMintUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMintUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


