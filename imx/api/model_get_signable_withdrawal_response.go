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

// GetSignableWithdrawalResponse struct for GetSignableWithdrawalResponse
type GetSignableWithdrawalResponse struct {
	// Amount of the token we are withdrawing
	Amount string `json:"amount"`
	// ID of the asset to be withdrawn
	AssetId string `json:"asset_id"`
	// Nonce of this transaction
	Nonce int32 `json:"nonce"`
	// Encoded payload hash
	PayloadHash string `json:"payload_hash"`
	// Message to sign with L1 wallet to verity withdrawal request
	SignableMessage string `json:"signable_message"`
	// Public stark key of this user
	StarkKey string `json:"stark_key"`
	// ID of the vault we are withdrawing from
	VaultId int32 `json:"vault_id"`
}

// NewGetSignableWithdrawalResponse instantiates a new GetSignableWithdrawalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableWithdrawalResponse(amount string, assetId string, nonce int32, payloadHash string, signableMessage string, starkKey string, vaultId int32) *GetSignableWithdrawalResponse {
	this := GetSignableWithdrawalResponse{}
	this.Amount = amount
	this.AssetId = assetId
	this.Nonce = nonce
	this.PayloadHash = payloadHash
	this.SignableMessage = signableMessage
	this.StarkKey = starkKey
	this.VaultId = vaultId
	return &this
}

// NewGetSignableWithdrawalResponseWithDefaults instantiates a new GetSignableWithdrawalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableWithdrawalResponseWithDefaults() *GetSignableWithdrawalResponse {
	this := GetSignableWithdrawalResponse{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetSignableWithdrawalResponse) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetAmountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetSignableWithdrawalResponse) SetAmount(v string) {
	o.Amount = v
}

// GetAssetId returns the AssetId field value
func (o *GetSignableWithdrawalResponse) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetAssetIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *GetSignableWithdrawalResponse) SetAssetId(v string) {
	o.AssetId = v
}

// GetNonce returns the Nonce field value
func (o *GetSignableWithdrawalResponse) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetNonceOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetSignableWithdrawalResponse) SetNonce(v int32) {
	o.Nonce = v
}

// GetPayloadHash returns the PayloadHash field value
func (o *GetSignableWithdrawalResponse) GetPayloadHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadHash
}

// GetPayloadHashOk returns a tuple with the PayloadHash field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetPayloadHashOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PayloadHash, true
}

// SetPayloadHash sets field value
func (o *GetSignableWithdrawalResponse) SetPayloadHash(v string) {
	o.PayloadHash = v
}

// GetSignableMessage returns the SignableMessage field value
func (o *GetSignableWithdrawalResponse) GetSignableMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignableMessage
}

// GetSignableMessageOk returns a tuple with the SignableMessage field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetSignableMessageOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SignableMessage, true
}

// SetSignableMessage sets field value
func (o *GetSignableWithdrawalResponse) SetSignableMessage(v string) {
	o.SignableMessage = v
}

// GetStarkKey returns the StarkKey field value
func (o *GetSignableWithdrawalResponse) GetStarkKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarkKey
}

// GetStarkKeyOk returns a tuple with the StarkKey field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetStarkKeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.StarkKey, true
}

// SetStarkKey sets field value
func (o *GetSignableWithdrawalResponse) SetStarkKey(v string) {
	o.StarkKey = v
}

// GetVaultId returns the VaultId field value
func (o *GetSignableWithdrawalResponse) GetVaultId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VaultId
}

// GetVaultIdOk returns a tuple with the VaultId field value
// and a boolean to check if the value has been set.
func (o *GetSignableWithdrawalResponse) GetVaultIdOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultId, true
}

// SetVaultId sets field value
func (o *GetSignableWithdrawalResponse) SetVaultId(v int32) {
	o.VaultId = v
}

func (o GetSignableWithdrawalResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["asset_id"] = o.AssetId
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["payload_hash"] = o.PayloadHash
	}
	if true {
		toSerialize["signable_message"] = o.SignableMessage
	}
	if true {
		toSerialize["stark_key"] = o.StarkKey
	}
	if true {
		toSerialize["vault_id"] = o.VaultId
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableWithdrawalResponse struct {
	value *GetSignableWithdrawalResponse
	isSet bool
}

func (v NullableGetSignableWithdrawalResponse) Get() *GetSignableWithdrawalResponse {
	return v.value
}

func (v *NullableGetSignableWithdrawalResponse) Set(val *GetSignableWithdrawalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableWithdrawalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableWithdrawalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableWithdrawalResponse(val *GetSignableWithdrawalResponse) *NullableGetSignableWithdrawalResponse {
	return &NullableGetSignableWithdrawalResponse{value: val, isSet: true}
}

func (v NullableGetSignableWithdrawalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableWithdrawalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


