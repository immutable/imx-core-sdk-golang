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

// GetSignableOrderRequestTokenSell Token to sell
type GetSignableOrderRequestTokenSell struct {
	// Token data. See https://docs.x.immutable.com/docs/token-data-object
	Data map[string]interface{} `json:"data,omitempty"`
	// Type of token
	Type *string `json:"type,omitempty"`
}

// NewGetSignableOrderRequestTokenSell instantiates a new GetSignableOrderRequestTokenSell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableOrderRequestTokenSell() *GetSignableOrderRequestTokenSell {
	this := GetSignableOrderRequestTokenSell{}
	return &this
}

// NewGetSignableOrderRequestTokenSellWithDefaults instantiates a new GetSignableOrderRequestTokenSell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableOrderRequestTokenSellWithDefaults() *GetSignableOrderRequestTokenSell {
	this := GetSignableOrderRequestTokenSell{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetSignableOrderRequestTokenSell) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequestTokenSell) GetDataOk() (map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetSignableOrderRequestTokenSell) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *GetSignableOrderRequestTokenSell) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetSignableOrderRequestTokenSell) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequestTokenSell) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GetSignableOrderRequestTokenSell) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetSignableOrderRequestTokenSell) SetType(v string) {
	o.Type = &v
}

func (o GetSignableOrderRequestTokenSell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableOrderRequestTokenSell struct {
	value *GetSignableOrderRequestTokenSell
	isSet bool
}

func (v NullableGetSignableOrderRequestTokenSell) Get() *GetSignableOrderRequestTokenSell {
	return v.value
}

func (v *NullableGetSignableOrderRequestTokenSell) Set(val *GetSignableOrderRequestTokenSell) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableOrderRequestTokenSell) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableOrderRequestTokenSell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableOrderRequestTokenSell(val *GetSignableOrderRequestTokenSell) *NullableGetSignableOrderRequestTokenSell {
	return &NullableGetSignableOrderRequestTokenSell{value: val, isSet: true}
}

func (v NullableGetSignableOrderRequestTokenSell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableOrderRequestTokenSell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


