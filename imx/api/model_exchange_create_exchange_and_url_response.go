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

// ExchangeCreateExchangeAndURLResponse struct for ExchangeCreateExchangeAndURLResponse
type ExchangeCreateExchangeAndURLResponse struct {
	// Created transaction ID
	Id *int32 `json:"id,omitempty"`
	// Provider name (e.g. moonpay)
	Provider *string `json:"provider,omitempty"`
	// Transaction type
	Type *string `json:"type,omitempty"`
	// Widget URL
	Url *string `json:"url,omitempty"`
	// Ethereum address of the user who created transaction
	WalletAddress *string `json:"wallet_address,omitempty"`
}

// NewExchangeCreateExchangeAndURLResponse instantiates a new ExchangeCreateExchangeAndURLResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExchangeCreateExchangeAndURLResponse() *ExchangeCreateExchangeAndURLResponse {
	this := ExchangeCreateExchangeAndURLResponse{}
	return &this
}

// NewExchangeCreateExchangeAndURLResponseWithDefaults instantiates a new ExchangeCreateExchangeAndURLResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExchangeCreateExchangeAndURLResponseWithDefaults() *ExchangeCreateExchangeAndURLResponse {
	this := ExchangeCreateExchangeAndURLResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ExchangeCreateExchangeAndURLResponse) GetId() int32 {
	if o == nil || isNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeCreateExchangeAndURLResponse) GetIdOk() (*int32, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ExchangeCreateExchangeAndURLResponse) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ExchangeCreateExchangeAndURLResponse) SetId(v int32) {
	o.Id = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *ExchangeCreateExchangeAndURLResponse) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeCreateExchangeAndURLResponse) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *ExchangeCreateExchangeAndURLResponse) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *ExchangeCreateExchangeAndURLResponse) SetProvider(v string) {
	o.Provider = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ExchangeCreateExchangeAndURLResponse) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeCreateExchangeAndURLResponse) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ExchangeCreateExchangeAndURLResponse) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ExchangeCreateExchangeAndURLResponse) SetType(v string) {
	o.Type = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ExchangeCreateExchangeAndURLResponse) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeCreateExchangeAndURLResponse) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
    return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ExchangeCreateExchangeAndURLResponse) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ExchangeCreateExchangeAndURLResponse) SetUrl(v string) {
	o.Url = &v
}

// GetWalletAddress returns the WalletAddress field value if set, zero value otherwise.
func (o *ExchangeCreateExchangeAndURLResponse) GetWalletAddress() string {
	if o == nil || isNil(o.WalletAddress) {
		var ret string
		return ret
	}
	return *o.WalletAddress
}

// GetWalletAddressOk returns a tuple with the WalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExchangeCreateExchangeAndURLResponse) GetWalletAddressOk() (*string, bool) {
	if o == nil || isNil(o.WalletAddress) {
    return nil, false
	}
	return o.WalletAddress, true
}

// HasWalletAddress returns a boolean if a field has been set.
func (o *ExchangeCreateExchangeAndURLResponse) HasWalletAddress() bool {
	if o != nil && !isNil(o.WalletAddress) {
		return true
	}

	return false
}

// SetWalletAddress gets a reference to the given string and assigns it to the WalletAddress field.
func (o *ExchangeCreateExchangeAndURLResponse) SetWalletAddress(v string) {
	o.WalletAddress = &v
}

func (o ExchangeCreateExchangeAndURLResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.WalletAddress) {
		toSerialize["wallet_address"] = o.WalletAddress
	}
	return json.Marshal(toSerialize)
}

type NullableExchangeCreateExchangeAndURLResponse struct {
	value *ExchangeCreateExchangeAndURLResponse
	isSet bool
}

func (v NullableExchangeCreateExchangeAndURLResponse) Get() *ExchangeCreateExchangeAndURLResponse {
	return v.value
}

func (v *NullableExchangeCreateExchangeAndURLResponse) Set(val *ExchangeCreateExchangeAndURLResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableExchangeCreateExchangeAndURLResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableExchangeCreateExchangeAndURLResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExchangeCreateExchangeAndURLResponse(val *ExchangeCreateExchangeAndURLResponse) *NullableExchangeCreateExchangeAndURLResponse {
	return &NullableExchangeCreateExchangeAndURLResponse{value: val, isSet: true}
}

func (v NullableExchangeCreateExchangeAndURLResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExchangeCreateExchangeAndURLResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


