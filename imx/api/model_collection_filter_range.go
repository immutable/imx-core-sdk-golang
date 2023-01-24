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

// CollectionFilterRange Range of values for this property
type CollectionFilterRange struct {
	// Maximum value
	Max *int32 `json:"max,omitempty"`
	// Minimum value
	Min *int32 `json:"min,omitempty"`
}

// NewCollectionFilterRange instantiates a new CollectionFilterRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionFilterRange() *CollectionFilterRange {
	this := CollectionFilterRange{}
	return &this
}

// NewCollectionFilterRangeWithDefaults instantiates a new CollectionFilterRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionFilterRangeWithDefaults() *CollectionFilterRange {
	this := CollectionFilterRange{}
	return &this
}

// GetMax returns the Max field value if set, zero value otherwise.
func (o *CollectionFilterRange) GetMax() int32 {
	if o == nil || o.Max == nil {
		var ret int32
		return ret
	}
	return *o.Max
}

// GetMaxOk returns a tuple with the Max field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFilterRange) GetMaxOk() (*int32, bool) {
	if o == nil || o.Max == nil {
		return nil, false
	}
	return o.Max, true
}

// HasMax returns a boolean if a field has been set.
func (o *CollectionFilterRange) HasMax() bool {
	if o != nil && o.Max != nil {
		return true
	}

	return false
}

// SetMax gets a reference to the given int32 and assigns it to the Max field.
func (o *CollectionFilterRange) SetMax(v int32) {
	o.Max = &v
}

// GetMin returns the Min field value if set, zero value otherwise.
func (o *CollectionFilterRange) GetMin() int32 {
	if o == nil || o.Min == nil {
		var ret int32
		return ret
	}
	return *o.Min
}

// GetMinOk returns a tuple with the Min field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionFilterRange) GetMinOk() (*int32, bool) {
	if o == nil || o.Min == nil {
		return nil, false
	}
	return o.Min, true
}

// HasMin returns a boolean if a field has been set.
func (o *CollectionFilterRange) HasMin() bool {
	if o != nil && o.Min != nil {
		return true
	}

	return false
}

// SetMin gets a reference to the given int32 and assigns it to the Min field.
func (o *CollectionFilterRange) SetMin(v int32) {
	o.Min = &v
}

func (o CollectionFilterRange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Max != nil {
		toSerialize["max"] = o.Max
	}
	if o.Min != nil {
		toSerialize["min"] = o.Min
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionFilterRange struct {
	value *CollectionFilterRange
	isSet bool
}

func (v NullableCollectionFilterRange) Get() *CollectionFilterRange {
	return v.value
}

func (v *NullableCollectionFilterRange) Set(val *CollectionFilterRange) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionFilterRange) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionFilterRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionFilterRange(val *CollectionFilterRange) *NullableCollectionFilterRange {
	return &NullableCollectionFilterRange{value: val, isSet: true}
}

func (v NullableCollectionFilterRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionFilterRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


