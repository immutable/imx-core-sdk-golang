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

// checks if the TokenDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TokenDetails{}

// TokenDetails struct for TokenDetails
type TokenDetails struct {
	// Number of decimals for token
	Decimals string `json:"decimals"`
	// Url for the icon of the token
	ImageUrl string `json:"image_url"`
	// Full name of the token (e.g. Ether)
	Name string `json:"name"`
	// Quantum for token
	Quantum string `json:"quantum"`
	// Ticker symbol for token (e.g. ETH/USDC/IMX)
	Symbol string `json:"symbol"`
	// Address of the ERC721 contract
	TokenAddress string `json:"token_address"`
}

// NewTokenDetails instantiates a new TokenDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenDetails(decimals string, imageUrl string, name string, quantum string, symbol string, tokenAddress string) *TokenDetails {
	this := TokenDetails{}
	this.Decimals = decimals
	this.ImageUrl = imageUrl
	this.Name = name
	this.Quantum = quantum
	this.Symbol = symbol
	this.TokenAddress = tokenAddress
	return &this
}

// NewTokenDetailsWithDefaults instantiates a new TokenDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenDetailsWithDefaults() *TokenDetails {
	this := TokenDetails{}
	return &this
}

// GetDecimals returns the Decimals field value
func (o *TokenDetails) GetDecimals() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Decimals
}

// GetDecimalsOk returns a tuple with the Decimals field value
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetDecimalsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decimals, true
}

// SetDecimals sets field value
func (o *TokenDetails) SetDecimals(v string) {
	o.Decimals = v
}

// GetImageUrl returns the ImageUrl field value
func (o *TokenDetails) GetImageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageUrl, true
}

// SetImageUrl sets field value
func (o *TokenDetails) SetImageUrl(v string) {
	o.ImageUrl = v
}

// GetName returns the Name field value
func (o *TokenDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TokenDetails) SetName(v string) {
	o.Name = v
}

// GetQuantum returns the Quantum field value
func (o *TokenDetails) GetQuantum() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Quantum
}

// GetQuantumOk returns a tuple with the Quantum field value
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetQuantumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Quantum, true
}

// SetQuantum sets field value
func (o *TokenDetails) SetQuantum(v string) {
	o.Quantum = v
}

// GetSymbol returns the Symbol field value
func (o *TokenDetails) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *TokenDetails) SetSymbol(v string) {
	o.Symbol = v
}

// GetTokenAddress returns the TokenAddress field value
func (o *TokenDetails) GetTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value
// and a boolean to check if the value has been set.
func (o *TokenDetails) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAddress, true
}

// SetTokenAddress sets field value
func (o *TokenDetails) SetTokenAddress(v string) {
	o.TokenAddress = v
}

func (o TokenDetails) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TokenDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["decimals"] = o.Decimals
	toSerialize["image_url"] = o.ImageUrl
	toSerialize["name"] = o.Name
	toSerialize["quantum"] = o.Quantum
	toSerialize["symbol"] = o.Symbol
	toSerialize["token_address"] = o.TokenAddress
	return toSerialize, nil
}

type NullableTokenDetails struct {
	value *TokenDetails
	isSet bool
}

func (v NullableTokenDetails) Get() *TokenDetails {
	return v.value
}

func (v *NullableTokenDetails) Set(val *TokenDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenDetails(val *TokenDetails) *NullableTokenDetails {
	return &NullableTokenDetails{value: val, isSet: true}
}

func (v NullableTokenDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


