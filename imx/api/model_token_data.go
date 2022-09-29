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

// TokenData struct for TokenData
type TokenData struct {
	// Number of decimals supported by this asset
	Decimals *int32 `json:"decimals,omitempty"`
	// [DEPRECATED] Internal Immutable X Token ID
	Id *string `json:"id,omitempty"`
	Properties *AssetProperties `json:"properties,omitempty"`
	// Quantity of this asset - inclusive of fees for buy order in v1 API and exclusive of fees in v3 API
	Quantity string `json:"quantity"`
	// Quantity of this asset with the sum of all fees applied to the asset
	QuantityWithFees string `json:"quantity_with_fees"`
	// Symbol of a token
	Symbol *string `json:"symbol,omitempty"`
	// Address of ERC721/ERC20 contract
	TokenAddress *string `json:"token_address,omitempty"`
	// ERC721 Token ID
	TokenId *string `json:"token_id,omitempty"`
}

// NewTokenData instantiates a new TokenData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenData(quantity string, quantityWithFees string) *TokenData {
	this := TokenData{}
	this.Quantity = quantity
	this.QuantityWithFees = quantityWithFees
	return &this
}

// NewTokenDataWithDefaults instantiates a new TokenData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDataWithDefaults() *TokenData {
	this := TokenData{}
	return &this
}

// GetDecimals returns the Decimals field value if set, zero value otherwise.
func (o *TokenData) GetDecimals() int32 {
	if o == nil || o.Decimals == nil {
		var ret int32
		return ret
	}
	return *o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetDecimalsOk() (*int32, bool) {
	if o == nil || o.Decimals == nil {
		return nil, false
	}
	return o.Decimals, true
}

// HasDecimals returns a boolean if a field has been set.
func (o *TokenData) HasDecimals() bool {
	if o != nil && o.Decimals != nil {
		return true
	}

	return false
}

// SetDecimals gets a reference to the given int32 and assigns it to the Decimals field.
func (o *TokenData) SetDecimals(v int32) {
	o.Decimals = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TokenData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TokenData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TokenData) SetId(v string) {
	o.Id = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TokenData) GetProperties() AssetProperties {
	if o == nil || o.Properties == nil {
		var ret AssetProperties
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetPropertiesOk() (*AssetProperties, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TokenData) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given AssetProperties and assigns it to the Properties field.
func (o *TokenData) SetProperties(v AssetProperties) {
	o.Properties = &v
}

// GetQuantity returns the Quantity field value
func (o *TokenData) GetQuantity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value
// and a boolean to check if the value has been set.
func (o *TokenData) GetQuantityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantity, true
}

// SetQuantity sets field value
func (o *TokenData) SetQuantity(v string) {
	o.Quantity = v
}

// GetQuantityWithFees returns the QuantityWithFees field value
func (o *TokenData) GetQuantityWithFees() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QuantityWithFees
}

// GetQuantityWithFeesOk returns a tuple with the QuantityWithFees field value
// and a boolean to check if the value has been set.
func (o *TokenData) GetQuantityWithFeesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuantityWithFees, true
}

// SetQuantityWithFees sets field value
func (o *TokenData) SetQuantityWithFees(v string) {
	o.QuantityWithFees = v
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *TokenData) GetSymbol() string {
	if o == nil || o.Symbol == nil {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetSymbolOk() (*string, bool) {
	if o == nil || o.Symbol == nil {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *TokenData) HasSymbol() bool {
	if o != nil && o.Symbol != nil {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *TokenData) SetSymbol(v string) {
	o.Symbol = &v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *TokenData) GetTokenAddress() string {
	if o == nil || o.TokenAddress == nil {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetTokenAddressOk() (*string, bool) {
	if o == nil || o.TokenAddress == nil {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *TokenData) HasTokenAddress() bool {
	if o != nil && o.TokenAddress != nil {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *TokenData) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *TokenData) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TokenData) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *TokenData) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *TokenData) SetTokenId(v string) {
	o.TokenId = &v
}

func (o TokenData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Decimals != nil {
		toSerialize["decimals"] = o.Decimals
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["quantity"] = o.Quantity
	}
	if true {
		toSerialize["quantity_with_fees"] = o.QuantityWithFees
	}
	if o.Symbol != nil {
		toSerialize["symbol"] = o.Symbol
	}
	if o.TokenAddress != nil {
		toSerialize["token_address"] = o.TokenAddress
	}
	if o.TokenId != nil {
		toSerialize["token_id"] = o.TokenId
	}
	return json.Marshal(toSerialize)
}

type NullableTokenData struct {
	value *TokenData
	isSet bool
}

func (v NullableTokenData) Get() *TokenData {
	return v.value
}

func (v *NullableTokenData) Set(val *TokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenData(val *TokenData) *NullableTokenData {
	return &NullableTokenData{value: val, isSet: true}
}

func (v NullableTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


