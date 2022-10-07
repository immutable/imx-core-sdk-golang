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

// MetadataRefreshSummary struct for MetadataRefreshSummary
type MetadataRefreshSummary struct {
	// The number of tokens with failed metadata refreshes
	Failed *int32 `json:"failed,omitempty"`
	// The number of tokens that has not been refreshed yet
	Pending *int32 `json:"pending,omitempty"`
	// The number of tokens with successful metadata refreshes
	Succeeded *int32 `json:"succeeded,omitempty"`
}

// NewMetadataRefreshSummary instantiates a new MetadataRefreshSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataRefreshSummary() *MetadataRefreshSummary {
	this := MetadataRefreshSummary{}
	return &this
}

// NewMetadataRefreshSummaryWithDefaults instantiates a new MetadataRefreshSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataRefreshSummaryWithDefaults() *MetadataRefreshSummary {
	this := MetadataRefreshSummary{}
	return &this
}

// GetFailed returns the Failed field value if set, zero value otherwise.
func (o *MetadataRefreshSummary) GetFailed() int32 {
	if o == nil || o.Failed == nil {
		var ret int32
		return ret
	}
	return *o.Failed
}

// GetFailedOk returns a tuple with the Failed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRefreshSummary) GetFailedOk() (*int32, bool) {
	if o == nil || o.Failed == nil {
		return nil, false
	}
	return o.Failed, true
}

// HasFailed returns a boolean if a field has been set.
func (o *MetadataRefreshSummary) HasFailed() bool {
	if o != nil && o.Failed != nil {
		return true
	}

	return false
}

// SetFailed gets a reference to the given int32 and assigns it to the Failed field.
func (o *MetadataRefreshSummary) SetFailed(v int32) {
	o.Failed = &v
}

// GetPending returns the Pending field value if set, zero value otherwise.
func (o *MetadataRefreshSummary) GetPending() int32 {
	if o == nil || o.Pending == nil {
		var ret int32
		return ret
	}
	return *o.Pending
}

// GetPendingOk returns a tuple with the Pending field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRefreshSummary) GetPendingOk() (*int32, bool) {
	if o == nil || o.Pending == nil {
		return nil, false
	}
	return o.Pending, true
}

// HasPending returns a boolean if a field has been set.
func (o *MetadataRefreshSummary) HasPending() bool {
	if o != nil && o.Pending != nil {
		return true
	}

	return false
}

// SetPending gets a reference to the given int32 and assigns it to the Pending field.
func (o *MetadataRefreshSummary) SetPending(v int32) {
	o.Pending = &v
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise.
func (o *MetadataRefreshSummary) GetSucceeded() int32 {
	if o == nil || o.Succeeded == nil {
		var ret int32
		return ret
	}
	return *o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataRefreshSummary) GetSucceededOk() (*int32, bool) {
	if o == nil || o.Succeeded == nil {
		return nil, false
	}
	return o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *MetadataRefreshSummary) HasSucceeded() bool {
	if o != nil && o.Succeeded != nil {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given int32 and assigns it to the Succeeded field.
func (o *MetadataRefreshSummary) SetSucceeded(v int32) {
	o.Succeeded = &v
}

func (o MetadataRefreshSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Failed != nil {
		toSerialize["failed"] = o.Failed
	}
	if o.Pending != nil {
		toSerialize["pending"] = o.Pending
	}
	if o.Succeeded != nil {
		toSerialize["succeeded"] = o.Succeeded
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataRefreshSummary struct {
	value *MetadataRefreshSummary
	isSet bool
}

func (v NullableMetadataRefreshSummary) Get() *MetadataRefreshSummary {
	return v.value
}

func (v *NullableMetadataRefreshSummary) Set(val *MetadataRefreshSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataRefreshSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataRefreshSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataRefreshSummary(val *MetadataRefreshSummary) *NullableMetadataRefreshSummary {
	return &NullableMetadataRefreshSummary{value: val, isSet: true}
}

func (v NullableMetadataRefreshSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataRefreshSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


