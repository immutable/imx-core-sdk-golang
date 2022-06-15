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

// GetSignableTradeResponse struct for GetSignableTradeResponse
type GetSignableTradeResponse struct {
	// Amount to buy
	AmountBuy string `json:"amount_buy"`
	// Amount to sell
	AmountSell string `json:"amount_sell"`
	// ID of the asset to buy
	AssetIdBuy string `json:"asset_id_buy"`
	// ID of the asset to sell
	AssetIdSell string `json:"asset_id_sell"`
	// Expiration timestamp for this order
	ExpirationTimestamp int32 `json:"expiration_timestamp"`
	FeeInfo *FeeInfo `json:"fee_info,omitempty"`
	// Nonce of the order
	Nonce int32 `json:"nonce"`
	// Payload Hash
	PayloadHash string `json:"payload_hash"`
	// Message to sign with L1 wallet to confirm trade request
	SignableMessage string `json:"signable_message"`
	// Public stark key of the created user
	StarkKey string `json:"stark_key"`
	// ID of the vault into which the bought asset will be placed
	VaultIdBuy int32 `json:"vault_id_buy"`
	// ID of the vault to sell from
	VaultIdSell int32 `json:"vault_id_sell"`
}

// NewGetSignableTradeResponse instantiates a new GetSignableTradeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableTradeResponse(amountBuy string, amountSell string, assetIdBuy string, assetIdSell string, expirationTimestamp int32, nonce int32, payloadHash string, signableMessage string, starkKey string, vaultIdBuy int32, vaultIdSell int32) *GetSignableTradeResponse {
	this := GetSignableTradeResponse{}
	this.AmountBuy = amountBuy
	this.AmountSell = amountSell
	this.AssetIdBuy = assetIdBuy
	this.AssetIdSell = assetIdSell
	this.ExpirationTimestamp = expirationTimestamp
	this.Nonce = nonce
	this.PayloadHash = payloadHash
	this.SignableMessage = signableMessage
	this.StarkKey = starkKey
	this.VaultIdBuy = vaultIdBuy
	this.VaultIdSell = vaultIdSell
	return &this
}

// NewGetSignableTradeResponseWithDefaults instantiates a new GetSignableTradeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableTradeResponseWithDefaults() *GetSignableTradeResponse {
	this := GetSignableTradeResponse{}
	return &this
}

// GetAmountBuy returns the AmountBuy field value
func (o *GetSignableTradeResponse) GetAmountBuy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountBuy
}

// GetAmountBuyOk returns a tuple with the AmountBuy field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetAmountBuyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountBuy, true
}

// SetAmountBuy sets field value
func (o *GetSignableTradeResponse) SetAmountBuy(v string) {
	o.AmountBuy = v
}

// GetAmountSell returns the AmountSell field value
func (o *GetSignableTradeResponse) GetAmountSell() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountSell
}

// GetAmountSellOk returns a tuple with the AmountSell field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetAmountSellOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountSell, true
}

// SetAmountSell sets field value
func (o *GetSignableTradeResponse) SetAmountSell(v string) {
	o.AmountSell = v
}

// GetAssetIdBuy returns the AssetIdBuy field value
func (o *GetSignableTradeResponse) GetAssetIdBuy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdBuy
}

// GetAssetIdBuyOk returns a tuple with the AssetIdBuy field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetAssetIdBuyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdBuy, true
}

// SetAssetIdBuy sets field value
func (o *GetSignableTradeResponse) SetAssetIdBuy(v string) {
	o.AssetIdBuy = v
}

// GetAssetIdSell returns the AssetIdSell field value
func (o *GetSignableTradeResponse) GetAssetIdSell() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AssetIdSell
}

// GetAssetIdSellOk returns a tuple with the AssetIdSell field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetAssetIdSellOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssetIdSell, true
}

// SetAssetIdSell sets field value
func (o *GetSignableTradeResponse) SetAssetIdSell(v string) {
	o.AssetIdSell = v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value
func (o *GetSignableTradeResponse) GetExpirationTimestamp() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetExpirationTimestampOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpirationTimestamp, true
}

// SetExpirationTimestamp sets field value
func (o *GetSignableTradeResponse) SetExpirationTimestamp(v int32) {
	o.ExpirationTimestamp = v
}

// GetFeeInfo returns the FeeInfo field value if set, zero value otherwise.
func (o *GetSignableTradeResponse) GetFeeInfo() FeeInfo {
	if o == nil || o.FeeInfo == nil {
		var ret FeeInfo
		return ret
	}
	return *o.FeeInfo
}

// GetFeeInfoOk returns a tuple with the FeeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetFeeInfoOk() (*FeeInfo, bool) {
	if o == nil || o.FeeInfo == nil {
		return nil, false
	}
	return o.FeeInfo, true
}

// HasFeeInfo returns a boolean if a field has been set.
func (o *GetSignableTradeResponse) HasFeeInfo() bool {
	if o != nil && o.FeeInfo != nil {
		return true
	}

	return false
}

// SetFeeInfo gets a reference to the given FeeInfo and assigns it to the FeeInfo field.
func (o *GetSignableTradeResponse) SetFeeInfo(v FeeInfo) {
	o.FeeInfo = &v
}

// GetNonce returns the Nonce field value
func (o *GetSignableTradeResponse) GetNonce() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nonce
}

// GetNonceOk returns a tuple with the Nonce field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetNonceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nonce, true
}

// SetNonce sets field value
func (o *GetSignableTradeResponse) SetNonce(v int32) {
	o.Nonce = v
}

// GetPayloadHash returns the PayloadHash field value
func (o *GetSignableTradeResponse) GetPayloadHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadHash
}

// GetPayloadHashOk returns a tuple with the PayloadHash field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetPayloadHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadHash, true
}

// SetPayloadHash sets field value
func (o *GetSignableTradeResponse) SetPayloadHash(v string) {
	o.PayloadHash = v
}

// GetSignableMessage returns the SignableMessage field value
func (o *GetSignableTradeResponse) GetSignableMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignableMessage
}

// GetSignableMessageOk returns a tuple with the SignableMessage field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetSignableMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignableMessage, true
}

// SetSignableMessage sets field value
func (o *GetSignableTradeResponse) SetSignableMessage(v string) {
	o.SignableMessage = v
}

// GetStarkKey returns the StarkKey field value
func (o *GetSignableTradeResponse) GetStarkKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StarkKey
}

// GetStarkKeyOk returns a tuple with the StarkKey field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetStarkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StarkKey, true
}

// SetStarkKey sets field value
func (o *GetSignableTradeResponse) SetStarkKey(v string) {
	o.StarkKey = v
}

// GetVaultIdBuy returns the VaultIdBuy field value
func (o *GetSignableTradeResponse) GetVaultIdBuy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VaultIdBuy
}

// GetVaultIdBuyOk returns a tuple with the VaultIdBuy field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetVaultIdBuyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultIdBuy, true
}

// SetVaultIdBuy sets field value
func (o *GetSignableTradeResponse) SetVaultIdBuy(v int32) {
	o.VaultIdBuy = v
}

// GetVaultIdSell returns the VaultIdSell field value
func (o *GetSignableTradeResponse) GetVaultIdSell() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VaultIdSell
}

// GetVaultIdSellOk returns a tuple with the VaultIdSell field value
// and a boolean to check if the value has been set.
func (o *GetSignableTradeResponse) GetVaultIdSellOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VaultIdSell, true
}

// SetVaultIdSell sets field value
func (o *GetSignableTradeResponse) SetVaultIdSell(v int32) {
	o.VaultIdSell = v
}

func (o GetSignableTradeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount_buy"] = o.AmountBuy
	}
	if true {
		toSerialize["amount_sell"] = o.AmountSell
	}
	if true {
		toSerialize["asset_id_buy"] = o.AssetIdBuy
	}
	if true {
		toSerialize["asset_id_sell"] = o.AssetIdSell
	}
	if true {
		toSerialize["expiration_timestamp"] = o.ExpirationTimestamp
	}
	if o.FeeInfo != nil {
		toSerialize["fee_info"] = o.FeeInfo
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
		toSerialize["vault_id_buy"] = o.VaultIdBuy
	}
	if true {
		toSerialize["vault_id_sell"] = o.VaultIdSell
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableTradeResponse struct {
	value *GetSignableTradeResponse
	isSet bool
}

func (v NullableGetSignableTradeResponse) Get() *GetSignableTradeResponse {
	return v.value
}

func (v *NullableGetSignableTradeResponse) Set(val *GetSignableTradeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableTradeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableTradeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableTradeResponse(val *GetSignableTradeResponse) *NullableGetSignableTradeResponse {
	return &NullableGetSignableTradeResponse{value: val, isSet: true}
}

func (v NullableGetSignableTradeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableTradeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


