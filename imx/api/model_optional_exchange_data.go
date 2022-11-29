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

// OptionalExchangeData struct for OptionalExchangeData
type OptionalExchangeData struct {
	// Provider transaction ID
	ExternalId *string `json:"external_id,omitempty"`
	// Fees amount
	FeesAmount *float32 `json:"fees_amount,omitempty"`
	// Amount that was exchanged from
	FromAmount *float32 `json:"from_amount,omitempty"`
	// Currency that was exchanged from
	FromCurrency *string `json:"from_currency,omitempty"`
	// Provider wallet address that was used for transferring crypto
	ProviderWalletAddress *string `json:"provider_wallet_address,omitempty"`
	// Amount that was exchanged to
	ToAmount *float32 `json:"to_amount,omitempty"`
	// Currency that was exchanged to
	ToCurrency *string `json:"to_currency,omitempty"`
	// Transfer ID
	TransferId *string `json:"transfer_id,omitempty"`
}

// NewOptionalExchangeData instantiates a new OptionalExchangeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionalExchangeData() *OptionalExchangeData {
	this := OptionalExchangeData{}
	return &this
}

// NewOptionalExchangeDataWithDefaults instantiates a new OptionalExchangeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionalExchangeDataWithDefaults() *OptionalExchangeData {
	this := OptionalExchangeData{}
	return &this
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *OptionalExchangeData) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetFeesAmount returns the FeesAmount field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetFeesAmount() float32 {
	if o == nil || o.FeesAmount == nil {
		var ret float32
		return ret
	}
	return *o.FeesAmount
}

// GetFeesAmountOk returns a tuple with the FeesAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetFeesAmountOk() (*float32, bool) {
	if o == nil || o.FeesAmount == nil {
		return nil, false
	}
	return o.FeesAmount, true
}

// HasFeesAmount returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasFeesAmount() bool {
	if o != nil && o.FeesAmount != nil {
		return true
	}

	return false
}

// SetFeesAmount gets a reference to the given float32 and assigns it to the FeesAmount field.
func (o *OptionalExchangeData) SetFeesAmount(v float32) {
	o.FeesAmount = &v
}

// GetFromAmount returns the FromAmount field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetFromAmount() float32 {
	if o == nil || o.FromAmount == nil {
		var ret float32
		return ret
	}
	return *o.FromAmount
}

// GetFromAmountOk returns a tuple with the FromAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetFromAmountOk() (*float32, bool) {
	if o == nil || o.FromAmount == nil {
		return nil, false
	}
	return o.FromAmount, true
}

// HasFromAmount returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasFromAmount() bool {
	if o != nil && o.FromAmount != nil {
		return true
	}

	return false
}

// SetFromAmount gets a reference to the given float32 and assigns it to the FromAmount field.
func (o *OptionalExchangeData) SetFromAmount(v float32) {
	o.FromAmount = &v
}

// GetFromCurrency returns the FromCurrency field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetFromCurrency() string {
	if o == nil || o.FromCurrency == nil {
		var ret string
		return ret
	}
	return *o.FromCurrency
}

// GetFromCurrencyOk returns a tuple with the FromCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetFromCurrencyOk() (*string, bool) {
	if o == nil || o.FromCurrency == nil {
		return nil, false
	}
	return o.FromCurrency, true
}

// HasFromCurrency returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasFromCurrency() bool {
	if o != nil && o.FromCurrency != nil {
		return true
	}

	return false
}

// SetFromCurrency gets a reference to the given string and assigns it to the FromCurrency field.
func (o *OptionalExchangeData) SetFromCurrency(v string) {
	o.FromCurrency = &v
}

// GetProviderWalletAddress returns the ProviderWalletAddress field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetProviderWalletAddress() string {
	if o == nil || o.ProviderWalletAddress == nil {
		var ret string
		return ret
	}
	return *o.ProviderWalletAddress
}

// GetProviderWalletAddressOk returns a tuple with the ProviderWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetProviderWalletAddressOk() (*string, bool) {
	if o == nil || o.ProviderWalletAddress == nil {
		return nil, false
	}
	return o.ProviderWalletAddress, true
}

// HasProviderWalletAddress returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasProviderWalletAddress() bool {
	if o != nil && o.ProviderWalletAddress != nil {
		return true
	}

	return false
}

// SetProviderWalletAddress gets a reference to the given string and assigns it to the ProviderWalletAddress field.
func (o *OptionalExchangeData) SetProviderWalletAddress(v string) {
	o.ProviderWalletAddress = &v
}

// GetToAmount returns the ToAmount field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetToAmount() float32 {
	if o == nil || o.ToAmount == nil {
		var ret float32
		return ret
	}
	return *o.ToAmount
}

// GetToAmountOk returns a tuple with the ToAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetToAmountOk() (*float32, bool) {
	if o == nil || o.ToAmount == nil {
		return nil, false
	}
	return o.ToAmount, true
}

// HasToAmount returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasToAmount() bool {
	if o != nil && o.ToAmount != nil {
		return true
	}

	return false
}

// SetToAmount gets a reference to the given float32 and assigns it to the ToAmount field.
func (o *OptionalExchangeData) SetToAmount(v float32) {
	o.ToAmount = &v
}

// GetToCurrency returns the ToCurrency field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetToCurrency() string {
	if o == nil || o.ToCurrency == nil {
		var ret string
		return ret
	}
	return *o.ToCurrency
}

// GetToCurrencyOk returns a tuple with the ToCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetToCurrencyOk() (*string, bool) {
	if o == nil || o.ToCurrency == nil {
		return nil, false
	}
	return o.ToCurrency, true
}

// HasToCurrency returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasToCurrency() bool {
	if o != nil && o.ToCurrency != nil {
		return true
	}

	return false
}

// SetToCurrency gets a reference to the given string and assigns it to the ToCurrency field.
func (o *OptionalExchangeData) SetToCurrency(v string) {
	o.ToCurrency = &v
}

// GetTransferId returns the TransferId field value if set, zero value otherwise.
func (o *OptionalExchangeData) GetTransferId() string {
	if o == nil || o.TransferId == nil {
		var ret string
		return ret
	}
	return *o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionalExchangeData) GetTransferIdOk() (*string, bool) {
	if o == nil || o.TransferId == nil {
		return nil, false
	}
	return o.TransferId, true
}

// HasTransferId returns a boolean if a field has been set.
func (o *OptionalExchangeData) HasTransferId() bool {
	if o != nil && o.TransferId != nil {
		return true
	}

	return false
}

// SetTransferId gets a reference to the given string and assigns it to the TransferId field.
func (o *OptionalExchangeData) SetTransferId(v string) {
	o.TransferId = &v
}

func (o OptionalExchangeData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExternalId != nil {
		toSerialize["external_id"] = o.ExternalId
	}
	if o.FeesAmount != nil {
		toSerialize["fees_amount"] = o.FeesAmount
	}
	if o.FromAmount != nil {
		toSerialize["from_amount"] = o.FromAmount
	}
	if o.FromCurrency != nil {
		toSerialize["from_currency"] = o.FromCurrency
	}
	if o.ProviderWalletAddress != nil {
		toSerialize["provider_wallet_address"] = o.ProviderWalletAddress
	}
	if o.ToAmount != nil {
		toSerialize["to_amount"] = o.ToAmount
	}
	if o.ToCurrency != nil {
		toSerialize["to_currency"] = o.ToCurrency
	}
	if o.TransferId != nil {
		toSerialize["transfer_id"] = o.TransferId
	}
	return json.Marshal(toSerialize)
}

type NullableOptionalExchangeData struct {
	value *OptionalExchangeData
	isSet bool
}

func (v NullableOptionalExchangeData) Get() *OptionalExchangeData {
	return v.value
}

func (v *NullableOptionalExchangeData) Set(val *OptionalExchangeData) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionalExchangeData) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionalExchangeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionalExchangeData(val *OptionalExchangeData) *NullableOptionalExchangeData {
	return &NullableOptionalExchangeData{value: val, isSet: true}
}

func (v NullableOptionalExchangeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionalExchangeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


