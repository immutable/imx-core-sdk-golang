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

// NftprimarytransactionWidgetParams struct for NftprimarytransactionWidgetParams
type NftprimarytransactionWidgetParams struct {
	// enums(light, dark)
	Theme *string `json:"theme,omitempty"`
}

// NewNftprimarytransactionWidgetParams instantiates a new NftprimarytransactionWidgetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftprimarytransactionWidgetParams() *NftprimarytransactionWidgetParams {
	this := NftprimarytransactionWidgetParams{}
	return &this
}

// NewNftprimarytransactionWidgetParamsWithDefaults instantiates a new NftprimarytransactionWidgetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftprimarytransactionWidgetParamsWithDefaults() *NftprimarytransactionWidgetParams {
	this := NftprimarytransactionWidgetParams{}
	return &this
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *NftprimarytransactionWidgetParams) GetTheme() string {
	if o == nil || isNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionWidgetParams) GetThemeOk() (*string, bool) {
	if o == nil || isNil(o.Theme) {
    return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *NftprimarytransactionWidgetParams) HasTheme() bool {
	if o != nil && !isNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *NftprimarytransactionWidgetParams) SetTheme(v string) {
	o.Theme = &v
}

func (o NftprimarytransactionWidgetParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	return json.Marshal(toSerialize)
}

type NullableNftprimarytransactionWidgetParams struct {
	value *NftprimarytransactionWidgetParams
	isSet bool
}

func (v NullableNftprimarytransactionWidgetParams) Get() *NftprimarytransactionWidgetParams {
	return v.value
}

func (v *NullableNftprimarytransactionWidgetParams) Set(val *NftprimarytransactionWidgetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableNftprimarytransactionWidgetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableNftprimarytransactionWidgetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftprimarytransactionWidgetParams(val *NftprimarytransactionWidgetParams) *NullableNftprimarytransactionWidgetParams {
	return &NullableNftprimarytransactionWidgetParams{value: val, isSet: true}
}

func (v NullableNftprimarytransactionWidgetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftprimarytransactionWidgetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


