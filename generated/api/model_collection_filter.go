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

// CollectionFilter struct for CollectionFilter
type CollectionFilter struct {
	// Key of this property
	Key *string `json:"key,omitempty"`
	Range *Range `json:"range,omitempty"`
	// Type of this filter
	Type *string `json:"type,omitempty"`
	// List of possible values for this property
	Value []string `json:"value,omitempty"`
}

// NewCollectionFilter instantiates a new CollectionFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionFilter() *CollectionFilter {
	this := CollectionFilter{}
	return &this
}

// NewCollectionFilterWithDefaults instantiates a new CollectionFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionFilterWithDefaults() *CollectionFilter {
	this := CollectionFilter{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *CollectionFilter) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFilter) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *CollectionFilter) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *CollectionFilter) SetKey(v string) {
	o.Key = &v
}

// GetRange returns the Range field value if set, zero value otherwise.
func (o *CollectionFilter) GetRange() Range {
	if o == nil || o.Range == nil {
		var ret Range
		return ret
	}
	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFilter) GetRangeOk() (*Range, bool) {
	if o == nil || o.Range == nil {
		return nil, false
	}
	return o.Range, true
}

// HasRange returns a boolean if a field has been set.
func (o *CollectionFilter) HasRange() bool {
	if o != nil && o.Range != nil {
		return true
	}

	return false
}

// SetRange gets a reference to the given Range and assigns it to the Range field.
func (o *CollectionFilter) SetRange(v Range) {
	o.Range = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CollectionFilter) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFilter) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CollectionFilter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CollectionFilter) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CollectionFilter) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFilter) GetValueOk() ([]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CollectionFilter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *CollectionFilter) SetValue(v []string) {
	o.Value = v
}

func (o CollectionFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Range != nil {
		toSerialize["range"] = o.Range
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionFilter struct {
	value *CollectionFilter
	isSet bool
}

func (v NullableCollectionFilter) Get() *CollectionFilter {
	return v.value
}

func (v *NullableCollectionFilter) Set(val *CollectionFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionFilter(val *CollectionFilter) *NullableCollectionFilter {
	return &NullableCollectionFilter{value: val, isSet: true}
}

func (v NullableCollectionFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


