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

// checks if the MintResultDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MintResultDetails{}

// MintResultDetails struct for MintResultDetails
type MintResultDetails struct {
	// Contract address of this token
	ContractAddress string `json:"contract_address"`
	// IMX ID of this token
	TokenId string `json:"token_id"`
	// Mint Transaction ID
	TxId int32 `json:"tx_id"`
}

// NewMintResultDetails instantiates a new MintResultDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMintResultDetails(contractAddress string, tokenId string, txId int32) *MintResultDetails {
	this := MintResultDetails{}
	this.ContractAddress = contractAddress
	this.TokenId = tokenId
	this.TxId = txId
	return &this
}

// NewMintResultDetailsWithDefaults instantiates a new MintResultDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMintResultDetailsWithDefaults() *MintResultDetails {
	this := MintResultDetails{}
	return &this
}

// GetContractAddress returns the ContractAddress field value
func (o *MintResultDetails) GetContractAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value
// and a boolean to check if the value has been set.
func (o *MintResultDetails) GetContractAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractAddress, true
}

// SetContractAddress sets field value
func (o *MintResultDetails) SetContractAddress(v string) {
	o.ContractAddress = v
}

// GetTokenId returns the TokenId field value
func (o *MintResultDetails) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *MintResultDetails) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *MintResultDetails) SetTokenId(v string) {
	o.TokenId = v
}

// GetTxId returns the TxId field value
func (o *MintResultDetails) GetTxId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TxId
}

// GetTxIdOk returns a tuple with the TxId field value
// and a boolean to check if the value has been set.
func (o *MintResultDetails) GetTxIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TxId, true
}

// SetTxId sets field value
func (o *MintResultDetails) SetTxId(v int32) {
	o.TxId = v
}

func (o MintResultDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MintResultDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["contract_address"] = o.ContractAddress
	toSerialize["token_id"] = o.TokenId
	toSerialize["tx_id"] = o.TxId
	return toSerialize, nil
}

type NullableMintResultDetails struct {
	value *MintResultDetails
	isSet bool
}

func (v NullableMintResultDetails) Get() *MintResultDetails {
	return v.value
}

func (v *NullableMintResultDetails) Set(val *MintResultDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableMintResultDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableMintResultDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMintResultDetails(val *MintResultDetails) *NullableMintResultDetails {
	return &NullableMintResultDetails{value: val, isSet: true}
}

func (v NullableMintResultDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMintResultDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


