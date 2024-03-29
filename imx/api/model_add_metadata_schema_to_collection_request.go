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

// AddMetadataSchemaToCollectionRequest struct for AddMetadataSchemaToCollectionRequest
type AddMetadataSchemaToCollectionRequest struct {
	// Not required from API user
	ContractAddress *string `json:"contract_address,omitempty"`
	// The metadata container
	Metadata []MetadataSchemaRequest `json:"metadata"`
}

// NewAddMetadataSchemaToCollectionRequest instantiates a new AddMetadataSchemaToCollectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddMetadataSchemaToCollectionRequest(metadata []MetadataSchemaRequest) *AddMetadataSchemaToCollectionRequest {
	this := AddMetadataSchemaToCollectionRequest{}
	this.Metadata = metadata
	return &this
}

// NewAddMetadataSchemaToCollectionRequestWithDefaults instantiates a new AddMetadataSchemaToCollectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddMetadataSchemaToCollectionRequestWithDefaults() *AddMetadataSchemaToCollectionRequest {
	this := AddMetadataSchemaToCollectionRequest{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *AddMetadataSchemaToCollectionRequest) GetContractAddress() string {
	if o == nil || o.ContractAddress == nil {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddMetadataSchemaToCollectionRequest) GetContractAddressOk() (*string, bool) {
	if o == nil || o.ContractAddress == nil {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *AddMetadataSchemaToCollectionRequest) HasContractAddress() bool {
	if o != nil && o.ContractAddress != nil {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *AddMetadataSchemaToCollectionRequest) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetMetadata returns the Metadata field value
func (o *AddMetadataSchemaToCollectionRequest) GetMetadata() []MetadataSchemaRequest {
	if o == nil {
		var ret []MetadataSchemaRequest
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *AddMetadataSchemaToCollectionRequest) GetMetadataOk() ([]MetadataSchemaRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *AddMetadataSchemaToCollectionRequest) SetMetadata(v []MetadataSchemaRequest) {
	o.Metadata = v
}

func (o AddMetadataSchemaToCollectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ContractAddress != nil {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableAddMetadataSchemaToCollectionRequest struct {
	value *AddMetadataSchemaToCollectionRequest
	isSet bool
}

func (v NullableAddMetadataSchemaToCollectionRequest) Get() *AddMetadataSchemaToCollectionRequest {
	return v.value
}

func (v *NullableAddMetadataSchemaToCollectionRequest) Set(val *AddMetadataSchemaToCollectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAddMetadataSchemaToCollectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAddMetadataSchemaToCollectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddMetadataSchemaToCollectionRequest(val *AddMetadataSchemaToCollectionRequest) *NullableAddMetadataSchemaToCollectionRequest {
	return &NullableAddMetadataSchemaToCollectionRequest{value: val, isSet: true}
}

func (v NullableAddMetadataSchemaToCollectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddMetadataSchemaToCollectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


