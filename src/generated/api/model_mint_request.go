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

// MintRequest struct for MintRequest
type MintRequest struct {
	// Signature from authorised minter
	AuthSignature string `json:"auth_signature"`
	// minting contract
	ContractAddress string `json:"contract_address"`
	// Global contract-level royalty fees
	Royalties []MintFee `json:"royalties,omitempty"`
	// Users to mint to
	Users []MintUser `json:"users"`
}

// NewMintRequest instantiates a new MintRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMintRequest(authSignature string, contractAddress string, users []MintUser) *MintRequest {
	this := MintRequest{}
	this.AuthSignature = authSignature
	this.ContractAddress = contractAddress
	this.Users = users
	return &this
}

// NewMintRequestWithDefaults instantiates a new MintRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMintRequestWithDefaults() *MintRequest {
	this := MintRequest{}
	return &this
}

// GetAuthSignature returns the AuthSignature field value
func (o *MintRequest) GetAuthSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthSignature
}

// GetAuthSignatureOk returns a tuple with the AuthSignature field value
// and a boolean to check if the value has been set.
func (o *MintRequest) GetAuthSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthSignature, true
}

// SetAuthSignature sets field value
func (o *MintRequest) SetAuthSignature(v string) {
	o.AuthSignature = v
}

// GetContractAddress returns the ContractAddress field value
func (o *MintRequest) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *MintRequest) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *MintRequest) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetRoyalties returns the Royalties field value if set, zero value otherwise.
func (o *MintRequest) GetRoyalties() []MintFee {
	if o == nil || o.Royalties == nil {
		var ret []MintFee
		return ret
	}
	return o.Royalties
}

// GetRoyaltiesOk returns a tuple with the Royalties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MintRequest) GetRoyaltiesOk() ([]MintFee, bool) {
	if o == nil || o.Royalties == nil {
		return nil, false
	}
	return o.Royalties, true
}

// HasRoyalties returns a boolean if a field has been set.
func (o *MintRequest) HasRoyalties() bool {
	if o != nil && o.Royalties != nil {
		return true
	}

	return false
}

// SetRoyalties gets a reference to the given []MintFee and assigns it to the Royalties field.
func (o *MintRequest) SetRoyalties(v []MintFee) {
	o.Royalties = v
}

// GetUsers returns the Users field value
func (o *MintRequest) GetUsers() []MintUser {
	if o == nil {
		var ret []MintUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *MintRequest) GetUsersOk() ([]MintUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *MintRequest) SetUsers(v []MintUser) {
	o.Users = v
}

func (o MintRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth_signature"] = o.AuthSignature
	}
	if true {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if o.Royalties != nil {
		toSerialize["royalties"] = o.Royalties
	}
	if true {
		toSerialize["users"] = o.Users
	}
	return json.Marshal(toSerialize)
}

type NullableMintRequest struct {
	value *MintRequest
	isSet bool
}

func (v NullableMintRequest) Get() *MintRequest {
	return v.value
}

func (v *NullableMintRequest) Set(val *MintRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMintRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMintRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMintRequest(val *MintRequest) *NullableMintRequest {
	return &NullableMintRequest{value: val, isSet: true}
}

func (v NullableMintRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMintRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


