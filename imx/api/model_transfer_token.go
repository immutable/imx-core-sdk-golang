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

// TransferToken Token transferred by the user
type TransferToken struct {
	Data TokenDataOAIGen `json:"data"`
	// Type of this asset (ETH/ERC20/ERC721)
	Type string `json:"type"`
}

// NewTransferToken instantiates a new TransferToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToken(data TokenDataOAIGen, type_ string) *TransferToken {
	this := TransferToken{}
	this.Data = data
	this.Type = type_
	return &this
}

// NewTransferTokenWithDefaults instantiates a new TransferToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferTokenWithDefaults() *TransferToken {
	this := TransferToken{}
	return &this
}

// GetData returns the Data field value
func (o *TransferToken) GetData() TokenDataOAIGen {
	if o == nil {
		var ret TokenDataOAIGen
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *TransferToken) GetDataOk() (*TokenDataOAIGen, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *TransferToken) SetData(v TokenDataOAIGen) {
	o.Data = v
}

// GetType returns the Type field value
func (o *TransferToken) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferToken) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferToken) SetType(v string) {
	o.Type = v
}

func (o TransferToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableTransferToken struct {
	value *TransferToken
	isSet bool
}

func (v NullableTransferToken) Get() *TransferToken {
	return v.value
}

func (v *NullableTransferToken) Set(val *TransferToken) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToken) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToken(val *TransferToken) *NullableTransferToken {
	return &NullableTransferToken{value: val, isSet: true}
}

func (v NullableTransferToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


