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

// SignableTransferResponseDetails struct for SignableTransferResponseDetails
type SignableTransferResponseDetails struct {
	// Amount of the asset being transferred
	Amount string `json:"amount"`
	// ID of the asset being transferred
	AssetId string `json:"asset_id"`
	// Timestamp when this transfer will expire
	ExpirationTimestamp int32 `json:"expiration_timestamp"`
	// Nonce of the transfer
	Nonce int32 `json:"nonce"`
	// Hash of the payload to be signed for transfer
	PayloadHash string `json:"payload_hash"`
	// Receiver of the transfer
	ReceiverStarkKey string `json:"receiver_stark_key"`
	// ID of the vault being transferred to
	ReceiverVaultId int32 `json:"receiver_vault_id"`
	// ID of the vault being transferred from
	SenderVaultId int32 `json:"sender_vault_id"`
	Token SignableToken `json:"token"`
}

// NewSignableTransferResponseDetails instantiates a new SignableTransferResponseDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignableTransferResponseDetails(amount string, assetId string, expirationTimestamp int32, nonce int32, payloadHash string, receiverStarkKey string, receiverVaultId int32, senderVaultId int32, token SignableToken) *SignableTransferResponseDetails {
	this := SignableTransferResponseDetails{}
	this.Amount = amount
	this.AssetId = assetId
	this.ExpirationTimestamp = expirationTimestamp
	this.Nonce = nonce
	this.PayloadHash = payloadHash
	this.ReceiverStarkKey = receiverStarkKey
	this.ReceiverVaultId = receiverVaultId
	this.SenderVaultId = senderVaultId
	this.Token = token
	return &this
}

// NewSignableTransferResponseDetailsWithDefaults instantiates a new SignableTransferResponseDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignableTransferResponseDetailsWithDefaults() *SignableTransferResponseDetails {
	this := SignableTransferResponseDetails{}
	return &this
}

// GetAmount returns the Amount field value
func (o *SignableTransferResponseDetails) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *SignableTransferResponseDetails) SetAmount(v string) {
	o.Amount = v
}

// GetAssetId returns the AssetId field value
func (o *SignableTransferResponseDetails) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *SignableTransferResponseDetails) SetAssetId(v string) {
	o.AssetId = v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value
func (o *SignableTransferResponseDetails) GetExpirationTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetExpirationTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestamp, true
}

// SetExpirationTimestamp sets field value
func (o *SignableTransferResponseDetails) SetExpirationTimestamp(v int32) {
	o.ExpirationTimestamp = v
}

// GetNonce returns the Nonce field value
func (o *SignableTransferResponseDetails) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *SignableTransferResponseDetails) SetNonce(v int32) {
	o.Nonce = v
}

// GetPayloadHash returns the PayloadHash field value
func (o *SignableTransferResponseDetails) GetPayloadHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadHash
}

// GetPayloadHashOk returns a tuple with the PayloadHash field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetPayloadHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadHash, true
}

// SetPayloadHash sets field value
func (o *SignableTransferResponseDetails) SetPayloadHash(v string) {
	o.PayloadHash = v
}

// GetReceiverStarkKey returns the ReceiverStarkKey field value
func (o *SignableTransferResponseDetails) GetReceiverStarkKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverStarkKey
}

// GetReceiverStarkKeyOk returns a tuple with the ReceiverStarkKey field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetReceiverStarkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverStarkKey, true
}

// SetReceiverStarkKey sets field value
func (o *SignableTransferResponseDetails) SetReceiverStarkKey(v string) {
	o.ReceiverStarkKey = v
}

// GetReceiverVaultId returns the ReceiverVaultId field value
func (o *SignableTransferResponseDetails) GetReceiverVaultId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReceiverVaultId
}

// GetReceiverVaultIdOk returns a tuple with the ReceiverVaultId field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetReceiverVaultIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverVaultId, true
}

// SetReceiverVaultId sets field value
func (o *SignableTransferResponseDetails) SetReceiverVaultId(v int32) {
	o.ReceiverVaultId = v
}

// GetSenderVaultId returns the SenderVaultId field value
func (o *SignableTransferResponseDetails) GetSenderVaultId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SenderVaultId
}

// GetSenderVaultIdOk returns a tuple with the SenderVaultId field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetSenderVaultIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderVaultId, true
}

// SetSenderVaultId sets field value
func (o *SignableTransferResponseDetails) SetSenderVaultId(v int32) {
	o.SenderVaultId = v
}

// GetToken returns the Token field value
func (o *SignableTransferResponseDetails) GetToken() SignableToken {
	if o == nil {
		var ret SignableToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SignableTransferResponseDetails) GetTokenOk() (*SignableToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SignableTransferResponseDetails) SetToken(v SignableToken) {
	o.Token = v
}

func (o SignableTransferResponseDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["asset_id"] = o.AssetId
	}
	if true {
		toSerialize["expiration_timestamp"] = o.ExpirationTimestamp
	}
	if true {
		toSerialize["nonce"] = o.Nonce
	}
	if true {
		toSerialize["payload_hash"] = o.PayloadHash
	}
	if true {
		toSerialize["receiver_stark_key"] = o.ReceiverStarkKey
	}
	if true {
		toSerialize["receiver_vault_id"] = o.ReceiverVaultId
	}
	if true {
		toSerialize["sender_vault_id"] = o.SenderVaultId
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableSignableTransferResponseDetails struct {
	value *SignableTransferResponseDetails
	isSet bool
}

func (v NullableSignableTransferResponseDetails) Get() *SignableTransferResponseDetails {
	return v.value
}

func (v *NullableSignableTransferResponseDetails) Set(val *SignableTransferResponseDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSignableTransferResponseDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSignableTransferResponseDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignableTransferResponseDetails(val *SignableTransferResponseDetails) *NullableSignableTransferResponseDetails {
	return &NullableSignableTransferResponseDetails{value: val, isSet: true}
}

func (v NullableSignableTransferResponseDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignableTransferResponseDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


