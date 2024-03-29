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

// EncodeAssetTokenData struct for EncodeAssetTokenData
type EncodeAssetTokenData struct {
	// Blueprint information of the token to be encoded. Used if token is mintable.
	Blueprint *string `json:"blueprint,omitempty"`
	// ID of the token to be encoded. Used if token is mintable.
	Id *string `json:"id,omitempty"`
	// Contract address of the token to be encoded
	TokenAddress *string `json:"token_address,omitempty"`
	// TokenId of the token to be encoded. Used if token is non-mintable
	TokenId *string `json:"token_id,omitempty"`
}

// NewEncodeAssetTokenData instantiates a new EncodeAssetTokenData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEncodeAssetTokenData() *EncodeAssetTokenData {
	this := EncodeAssetTokenData{}
	return &this
}

// NewEncodeAssetTokenDataWithDefaults instantiates a new EncodeAssetTokenData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEncodeAssetTokenDataWithDefaults() *EncodeAssetTokenData {
	this := EncodeAssetTokenData{}
	return &this
}

// GetBlueprint returns the Blueprint field value if set, zero value otherwise.
func (o *EncodeAssetTokenData) GetBlueprint() string {
	if o == nil || o.Blueprint == nil {
		var ret string
		return ret
	}
	return *o.Blueprint
}

// GetBlueprintOk returns a tuple with the Blueprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodeAssetTokenData) GetBlueprintOk() (*string, bool) {
	if o == nil || o.Blueprint == nil {
		return nil, false
	}
	return o.Blueprint, true
}

// HasBlueprint returns a boolean if a field has been set.
func (o *EncodeAssetTokenData) HasBlueprint() bool {
	if o != nil && o.Blueprint != nil {
		return true
	}

	return false
}

// SetBlueprint gets a reference to the given string and assigns it to the Blueprint field.
func (o *EncodeAssetTokenData) SetBlueprint(v string) {
	o.Blueprint = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EncodeAssetTokenData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodeAssetTokenData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EncodeAssetTokenData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EncodeAssetTokenData) SetId(v string) {
	o.Id = &v
}

// GetTokenAddress returns the TokenAddress field value if set, zero value otherwise.
func (o *EncodeAssetTokenData) GetTokenAddress() string {
	if o == nil || o.TokenAddress == nil {
		var ret string
		return ret
	}
	return *o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodeAssetTokenData) GetTokenAddressOk() (*string, bool) {
	if o == nil || o.TokenAddress == nil {
		return nil, false
	}
	return o.TokenAddress, true
}

// HasTokenAddress returns a boolean if a field has been set.
func (o *EncodeAssetTokenData) HasTokenAddress() bool {
	if o != nil && o.TokenAddress != nil {
		return true
	}

	return false
}

// SetTokenAddress gets a reference to the given string and assigns it to the TokenAddress field.
func (o *EncodeAssetTokenData) SetTokenAddress(v string) {
	o.TokenAddress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *EncodeAssetTokenData) GetTokenId() string {
	if o == nil || o.TokenId == nil {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EncodeAssetTokenData) GetTokenIdOk() (*string, bool) {
	if o == nil || o.TokenId == nil {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *EncodeAssetTokenData) HasTokenId() bool {
	if o != nil && o.TokenId != nil {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *EncodeAssetTokenData) SetTokenId(v string) {
	o.TokenId = &v
}

func (o EncodeAssetTokenData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Blueprint != nil {
		toSerialize["blueprint"] = o.Blueprint
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.TokenAddress != nil {
		toSerialize["token_address"] = o.TokenAddress
	}
	if o.TokenId != nil {
		toSerialize["token_id"] = o.TokenId
	}
	return json.Marshal(toSerialize)
}

type NullableEncodeAssetTokenData struct {
	value *EncodeAssetTokenData
	isSet bool
}

func (v NullableEncodeAssetTokenData) Get() *EncodeAssetTokenData {
	return v.value
}

func (v *NullableEncodeAssetTokenData) Set(val *EncodeAssetTokenData) {
	v.value = val
	v.isSet = true
}

func (v NullableEncodeAssetTokenData) IsSet() bool {
	return v.isSet
}

func (v *NullableEncodeAssetTokenData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncodeAssetTokenData(val *EncodeAssetTokenData) *NullableEncodeAssetTokenData {
	return &NullableEncodeAssetTokenData{value: val, isSet: true}
}

func (v NullableEncodeAssetTokenData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncodeAssetTokenData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


