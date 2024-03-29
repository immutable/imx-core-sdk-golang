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

// ContractCreateAPIRequest struct for ContractCreateAPIRequest
type ContractCreateAPIRequest struct {
	ContractAddress *string `json:"contract_address,omitempty"`
	DataUrl *string `json:"data_url,omitempty"`
	MintUrl *string `json:"mint_url,omitempty"`
}

// NewContractCreateAPIRequest instantiates a new ContractCreateAPIRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractCreateAPIRequest() *ContractCreateAPIRequest {
	this := ContractCreateAPIRequest{}
	return &this
}

// NewContractCreateAPIRequestWithDefaults instantiates a new ContractCreateAPIRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractCreateAPIRequestWithDefaults() *ContractCreateAPIRequest {
	this := ContractCreateAPIRequest{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *ContractCreateAPIRequest) GetContractAddress() string {
	if o == nil || o.ContractAddress == nil {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCreateAPIRequest) GetContractAddressOk() (*string, bool) {
	if o == nil || o.ContractAddress == nil {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *ContractCreateAPIRequest) HasContractAddress() bool {
	if o != nil && o.ContractAddress != nil {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *ContractCreateAPIRequest) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetDataUrl returns the DataUrl field value if set, zero value otherwise.
func (o *ContractCreateAPIRequest) GetDataUrl() string {
	if o == nil || o.DataUrl == nil {
		var ret string
		return ret
	}
	return *o.DataUrl
}

// GetDataUrlOk returns a tuple with the DataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCreateAPIRequest) GetDataUrlOk() (*string, bool) {
	if o == nil || o.DataUrl == nil {
		return nil, false
	}
	return o.DataUrl, true
}

// HasDataUrl returns a boolean if a field has been set.
func (o *ContractCreateAPIRequest) HasDataUrl() bool {
	if o != nil && o.DataUrl != nil {
		return true
	}

	return false
}

// SetDataUrl gets a reference to the given string and assigns it to the DataUrl field.
func (o *ContractCreateAPIRequest) SetDataUrl(v string) {
	o.DataUrl = &v
}

// GetMintUrl returns the MintUrl field value if set, zero value otherwise.
func (o *ContractCreateAPIRequest) GetMintUrl() string {
	if o == nil || o.MintUrl == nil {
		var ret string
		return ret
	}
	return *o.MintUrl
}

// GetMintUrlOk returns a tuple with the MintUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCreateAPIRequest) GetMintUrlOk() (*string, bool) {
	if o == nil || o.MintUrl == nil {
		return nil, false
	}
	return o.MintUrl, true
}

// HasMintUrl returns a boolean if a field has been set.
func (o *ContractCreateAPIRequest) HasMintUrl() bool {
	if o != nil && o.MintUrl != nil {
		return true
	}

	return false
}

// SetMintUrl gets a reference to the given string and assigns it to the MintUrl field.
func (o *ContractCreateAPIRequest) SetMintUrl(v string) {
	o.MintUrl = &v
}

func (o ContractCreateAPIRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractAddress != nil {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if o.DataUrl != nil {
		toSerialize["data_url"] = o.DataUrl
	}
	if o.MintUrl != nil {
		toSerialize["mint_url"] = o.MintUrl
	}
	return json.Marshal(toSerialize)
}

type NullableContractCreateAPIRequest struct {
	value *ContractCreateAPIRequest
	isSet bool
}

func (v NullableContractCreateAPIRequest) Get() *ContractCreateAPIRequest {
	return v.value
}

func (v *NullableContractCreateAPIRequest) Set(val *ContractCreateAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCreateAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCreateAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCreateAPIRequest(val *ContractCreateAPIRequest) *NullableContractCreateAPIRequest {
	return &NullableContractCreateAPIRequest{value: val, isSet: true}
}

func (v NullableContractCreateAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCreateAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


