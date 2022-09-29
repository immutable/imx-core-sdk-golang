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

// CreateMetadataRefreshRequest struct for CreateMetadataRefreshRequest
type CreateMetadataRefreshRequest struct {
	// The collection contract address
	CollectionAddress string `json:"collection_address"`
	// The tokens to refresh
	TokenIds []string `json:"token_ids"`
}

// NewCreateMetadataRefreshRequest instantiates a new CreateMetadataRefreshRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMetadataRefreshRequest(collectionAddress string, tokenIds []string) *CreateMetadataRefreshRequest {
	this := CreateMetadataRefreshRequest{}
	this.CollectionAddress = collectionAddress
	this.TokenIds = tokenIds
	return &this
}

// NewCreateMetadataRefreshRequestWithDefaults instantiates a new CreateMetadataRefreshRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMetadataRefreshRequestWithDefaults() *CreateMetadataRefreshRequest {
	this := CreateMetadataRefreshRequest{}
	return &this
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *CreateMetadataRefreshRequest) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *CreateMetadataRefreshRequest) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *CreateMetadataRefreshRequest) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetTokenIds returns the TokenIds field value
func (o *CreateMetadataRefreshRequest) GetTokenIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TokenIds
}

// GetTokenIdsOk returns a tuple with the TokenIds field value
// and a boolean to check if the value has been set.
func (o *CreateMetadataRefreshRequest) GetTokenIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TokenIds, true
}

// SetTokenIds sets field value
func (o *CreateMetadataRefreshRequest) SetTokenIds(v []string) {
	o.TokenIds = v
}

func (o CreateMetadataRefreshRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["collection_address"] = o.CollectionAddress
	}
	if true {
		toSerialize["token_ids"] = o.TokenIds
	}
	return json.Marshal(toSerialize)
}

type NullableCreateMetadataRefreshRequest struct {
	value *CreateMetadataRefreshRequest
	isSet bool
}

func (v NullableCreateMetadataRefreshRequest) Get() *CreateMetadataRefreshRequest {
	return v.value
}

func (v *NullableCreateMetadataRefreshRequest) Set(val *CreateMetadataRefreshRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMetadataRefreshRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMetadataRefreshRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMetadataRefreshRequest(val *CreateMetadataRefreshRequest) *NullableCreateMetadataRefreshRequest {
	return &NullableCreateMetadataRefreshRequest{value: val, isSet: true}
}

func (v NullableCreateMetadataRefreshRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMetadataRefreshRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


