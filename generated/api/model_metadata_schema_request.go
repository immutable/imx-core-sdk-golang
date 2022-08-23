/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1
Contact: support@github.com/immutable/
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// MetadataSchemaRequest struct for MetadataSchemaRequest
type MetadataSchemaRequest struct {
	// Sets the metadata as filterable
	Filterable *bool `json:"filterable,omitempty"`
	// Name of the metadata key
	Name string `json:"name"`
	// Type of the metadata. Values: \"enum\", \"text\", \"boolean\", \"continuous\", \"discrete\" | Default: \"text\". Src: https://docs.x.immutable.com/immutable/docs/asset-metadata#property-type-mapping
	Type *string `json:"type,omitempty"`
}

// NewMetadataSchemaRequest instantiates a new MetadataSchemaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataSchemaRequest(name string) *MetadataSchemaRequest {
	this := MetadataSchemaRequest{}
	this.Name = name
	return &this
}

// NewMetadataSchemaRequestWithDefaults instantiates a new MetadataSchemaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataSchemaRequestWithDefaults() *MetadataSchemaRequest {
	this := MetadataSchemaRequest{}
	return &this
}

// GetFilterable returns the Filterable field value if set, zero value otherwise.
func (o *MetadataSchemaRequest) GetFilterable() bool {
	if o == nil || o.Filterable == nil {
		var ret bool
		return ret
	}
	return *o.Filterable
}

// GetFilterableOk returns a tuple with the Filterable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSchemaRequest) GetFilterableOk() (*bool, bool) {
	if o == nil || o.Filterable == nil {
		return nil, false
	}
	return o.Filterable, true
}

// HasFilterable returns a boolean if a field has been set.
func (o *MetadataSchemaRequest) HasFilterable() bool {
	if o != nil && o.Filterable != nil {
		return true
	}

	return false
}

// SetFilterable gets a reference to the given bool and assigns it to the Filterable field.
func (o *MetadataSchemaRequest) SetFilterable(v bool) {
	o.Filterable = &v
}

// GetName returns the Name field value
func (o *MetadataSchemaRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MetadataSchemaRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MetadataSchemaRequest) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MetadataSchemaRequest) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataSchemaRequest) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MetadataSchemaRequest) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *MetadataSchemaRequest) SetType(v string) {
	o.Type = &v
}

func (o MetadataSchemaRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filterable != nil {
		toSerialize["filterable"] = o.Filterable
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataSchemaRequest struct {
	value *MetadataSchemaRequest
	isSet bool
}

func (v NullableMetadataSchemaRequest) Get() *MetadataSchemaRequest {
	return v.value
}

func (v *NullableMetadataSchemaRequest) Set(val *MetadataSchemaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataSchemaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataSchemaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataSchemaRequest(val *MetadataSchemaRequest) *NullableMetadataSchemaRequest {
	return &NullableMetadataSchemaRequest{value: val, isSet: true}
}

func (v NullableMetadataSchemaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataSchemaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


