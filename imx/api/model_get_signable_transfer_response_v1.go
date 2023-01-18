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

// checks if the GetSignableTransferResponseV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignableTransferResponseV1{}

// GetSignableTransferResponseV1 struct for GetSignableTransferResponseV1
type GetSignableTransferResponseV1 struct {
	// Amount of the asset being transferred
	Amount string `json:"amount"`
	// ID of the asset being transferred
	AssetId string `json:"asset_id"`
	// Token in request to match in SDK implementation
	ExpirationTimestamp int32 `json:"expiration_timestamp"`
	// Nonce of the transfer
	Nonce int32 `json:"nonce"`
	// Hash of the payload
	PayloadHash string `json:"payload_hash"`
	// Receiver of the transfer
	ReceiverStarkKey string `json:"receiver_stark_key"`
	// ID of the vault being transferred to
	ReceiverVaultId int32 `json:"receiver_vault_id"`
	// Sender of the transfer
	SenderStarkKey *string `json:"sender_stark_key,omitempty"`
	// ID of the vault being transferred from
	SenderVaultId int32 `json:"sender_vault_id"`
	// Message to sign with L1 wallet to confirm transfer request
	SignableMessage string `json:"signable_message"`
}

// NewGetSignableTransferResponseV1 instantiates a new GetSignableTransferResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableTransferResponseV1(amount string, assetId string, expirationTimestamp int32, nonce int32, payloadHash string, receiverStarkKey string, receiverVaultId int32, senderVaultId int32, signableMessage string) *GetSignableTransferResponseV1 {
	this := GetSignableTransferResponseV1{}
	this.Amount = amount
	this.AssetId = assetId
	this.ExpirationTimestamp = expirationTimestamp
	this.Nonce = nonce
	this.PayloadHash = payloadHash
	this.ReceiverStarkKey = receiverStarkKey
	this.ReceiverVaultId = receiverVaultId
	this.SenderVaultId = senderVaultId
	this.SignableMessage = signableMessage
	return &this
}

// NewGetSignableTransferResponseV1WithDefaults instantiates a new GetSignableTransferResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableTransferResponseV1WithDefaults() *GetSignableTransferResponseV1 {
	this := GetSignableTransferResponseV1{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetSignableTransferResponseV1) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetSignableTransferResponseV1) SetAmount(v string) {
	o.Amount = v
}

// GetAssetId returns the AssetId field value
func (o *GetSignableTransferResponseV1) GetAssetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetAssetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetId, true
}

// SetAssetId sets field value
func (o *GetSignableTransferResponseV1) SetAssetId(v string) {
	o.AssetId = v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value
func (o *GetSignableTransferResponseV1) GetExpirationTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetExpirationTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestamp, true
}

// SetExpirationTimestamp sets field value
func (o *GetSignableTransferResponseV1) SetExpirationTimestamp(v int32) {
	o.ExpirationTimestamp = v
}

// GetNonce returns the Nonce field value
func (o *GetSignableTransferResponseV1) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetSignableTransferResponseV1) SetNonce(v int32) {
	o.Nonce = v
}

// GetPayloadHash returns the PayloadHash field value
func (o *GetSignableTransferResponseV1) GetPayloadHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadHash
}

// GetPayloadHashOk returns a tuple with the PayloadHash field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetPayloadHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadHash, true
}

// SetPayloadHash sets field value
func (o *GetSignableTransferResponseV1) SetPayloadHash(v string) {
	o.PayloadHash = v
}

// GetReceiverStarkKey returns the ReceiverStarkKey field value
func (o *GetSignableTransferResponseV1) GetReceiverStarkKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReceiverStarkKey
}

// GetReceiverStarkKeyOk returns a tuple with the ReceiverStarkKey field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetReceiverStarkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverStarkKey, true
}

// SetReceiverStarkKey sets field value
func (o *GetSignableTransferResponseV1) SetReceiverStarkKey(v string) {
	o.ReceiverStarkKey = v
}

// GetReceiverVaultId returns the ReceiverVaultId field value
func (o *GetSignableTransferResponseV1) GetReceiverVaultId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReceiverVaultId
}

// GetReceiverVaultIdOk returns a tuple with the ReceiverVaultId field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetReceiverVaultIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceiverVaultId, true
}

// SetReceiverVaultId sets field value
func (o *GetSignableTransferResponseV1) SetReceiverVaultId(v int32) {
	o.ReceiverVaultId = v
}

// GetSenderStarkKey returns the SenderStarkKey field value if set, zero value otherwise.
func (o *GetSignableTransferResponseV1) GetSenderStarkKey() string {
	if o == nil || isNil(o.SenderStarkKey) {
		var ret string
		return ret
	}
	return *o.SenderStarkKey
}

// GetSenderStarkKeyOk returns a tuple with the SenderStarkKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetSenderStarkKeyOk() (*string, bool) {
	if o == nil || isNil(o.SenderStarkKey) {
		return nil, false
	}
	return o.SenderStarkKey, true
}

// HasSenderStarkKey returns a boolean if a field has been set.
func (o *GetSignableTransferResponseV1) HasSenderStarkKey() bool {
	if o != nil && !isNil(o.SenderStarkKey) {
		return true
	}

	return false
}

// SetSenderStarkKey gets a reference to the given string and assigns it to the SenderStarkKey field.
func (o *GetSignableTransferResponseV1) SetSenderStarkKey(v string) {
	o.SenderStarkKey = &v
}

// GetSenderVaultId returns the SenderVaultId field value
func (o *GetSignableTransferResponseV1) GetSenderVaultId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SenderVaultId
}

// GetSenderVaultIdOk returns a tuple with the SenderVaultId field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetSenderVaultIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderVaultId, true
}

// SetSenderVaultId sets field value
func (o *GetSignableTransferResponseV1) SetSenderVaultId(v int32) {
	o.SenderVaultId = v
}

// GetSignableMessage returns the SignableMessage field value
func (o *GetSignableTransferResponseV1) GetSignableMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignableMessage
}

// GetSignableMessageOk returns a tuple with the SignableMessage field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponseV1) GetSignableMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignableMessage, true
}

// SetSignableMessage sets field value
func (o *GetSignableTransferResponseV1) SetSignableMessage(v string) {
	o.SignableMessage = v
}

func (o GetSignableTransferResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignableTransferResponseV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["asset_id"] = o.AssetId
	toSerialize["expiration_timestamp"] = o.ExpirationTimestamp
	toSerialize["nonce"] = o.Nonce
	toSerialize["payload_hash"] = o.PayloadHash
	toSerialize["receiver_stark_key"] = o.ReceiverStarkKey
	toSerialize["receiver_vault_id"] = o.ReceiverVaultId
	if !isNil(o.SenderStarkKey) {
		toSerialize["sender_stark_key"] = o.SenderStarkKey
	}
	toSerialize["sender_vault_id"] = o.SenderVaultId
	toSerialize["signable_message"] = o.SignableMessage
	return toSerialize, nil
}

type NullableGetSignableTransferResponseV1 struct {
	value *GetSignableTransferResponseV1
	isSet bool
}

func (v NullableGetSignableTransferResponseV1) Get() *GetSignableTransferResponseV1 {
	return v.value
}

func (v *NullableGetSignableTransferResponseV1) Set(val *GetSignableTransferResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableTransferResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableTransferResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableTransferResponseV1(val *GetSignableTransferResponseV1) *NullableGetSignableTransferResponseV1 {
	return &NullableGetSignableTransferResponseV1{value: val, isSet: true}
}

func (v NullableGetSignableTransferResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableTransferResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


