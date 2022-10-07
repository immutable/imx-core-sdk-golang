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

// AssetProperties struct for AssetProperties
type AssetProperties struct {
	Collection *CollectionDetails `json:"collection,omitempty"`
	// Image URL of this asset
	ImageUrl *string `json:"image_url,omitempty"`
	// Name of this asset
	Name *string `json:"name,omitempty"`
}

// NewAssetProperties instantiates a new AssetProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetProperties() *AssetProperties {
	this := AssetProperties{}
	return &this
}

// NewAssetPropertiesWithDefaults instantiates a new AssetProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetPropertiesWithDefaults() *AssetProperties {
	this := AssetProperties{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *AssetProperties) GetCollection() CollectionDetails {
	if o == nil || o.Collection == nil {
		var ret CollectionDetails
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProperties) GetCollectionOk() (*CollectionDetails, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *AssetProperties) HasCollection() bool {
	if o != nil && o.Collection != nil {
		return true
	}

	return false
}

// SetCollection gets a reference to the given CollectionDetails and assigns it to the Collection field.
func (o *AssetProperties) SetCollection(v CollectionDetails) {
	o.Collection = &v
}

// GetImageUrl returns the ImageUrl field value if set, zero value otherwise.
func (o *AssetProperties) GetImageUrl() string {
	if o == nil || o.ImageUrl == nil {
		var ret string
		return ret
	}
	return *o.ImageUrl
}

// GetImageUrlOk returns a tuple with the ImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProperties) GetImageUrlOk() (*string, bool) {
	if o == nil || o.ImageUrl == nil {
		return nil, false
	}
	return o.ImageUrl, true
}

// HasImageUrl returns a boolean if a field has been set.
func (o *AssetProperties) HasImageUrl() bool {
	if o != nil && o.ImageUrl != nil {
		return true
	}

	return false
}

// SetImageUrl gets a reference to the given string and assigns it to the ImageUrl field.
func (o *AssetProperties) SetImageUrl(v string) {
	o.ImageUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetProperties) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetProperties) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetProperties) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetProperties) SetName(v string) {
	o.Name = &v
}

func (o AssetProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.ImageUrl != nil {
		toSerialize["image_url"] = o.ImageUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAssetProperties struct {
	value *AssetProperties
	isSet bool
}

func (v NullableAssetProperties) Get() *AssetProperties {
	return v.value
}

func (v *NullableAssetProperties) Set(val *AssetProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetProperties(val *AssetProperties) *NullableAssetProperties {
	return &NullableAssetProperties{value: val, isSet: true}
}

func (v NullableAssetProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


