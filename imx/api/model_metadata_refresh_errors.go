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

// MetadataRefreshErrors struct for MetadataRefreshErrors
type MetadataRefreshErrors struct {
	// Metadata API response for the token
	ClientResponseBody NullableString `json:"client_response_body"`
	// Metadata API response code for the token
	ClientResponseStatusCode NullableString `json:"client_response_status_code"`
	// Requested metadata url for the token
	ClientTokenMetadataUrl string `json:"client_token_metadata_url"`
	// The collection contract address
	CollectionAddress string `json:"collection_address"`
	// When the error was created
	CreatedAt string `json:"created_at"`
	// Metadata refresh error code
	ErrorCode string `json:"error_code"`
	// The token ID
	TokenId string `json:"token_id"`
}

// NewMetadataRefreshErrors instantiates a new MetadataRefreshErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataRefreshErrors(clientResponseBody NullableString, clientResponseStatusCode NullableString, clientTokenMetadataUrl string, collectionAddress string, createdAt string, errorCode string, tokenId string) *MetadataRefreshErrors {
	this := MetadataRefreshErrors{}
	this.ClientResponseBody = clientResponseBody
	this.ClientResponseStatusCode = clientResponseStatusCode
	this.ClientTokenMetadataUrl = clientTokenMetadataUrl
	this.CollectionAddress = collectionAddress
	this.CreatedAt = createdAt
	this.ErrorCode = errorCode
	this.TokenId = tokenId
	return &this
}

// NewMetadataRefreshErrorsWithDefaults instantiates a new MetadataRefreshErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataRefreshErrorsWithDefaults() *MetadataRefreshErrors {
	this := MetadataRefreshErrors{}
	return &this
}

// GetClientResponseBody returns the ClientResponseBody field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetadataRefreshErrors) GetClientResponseBody() string {
	if o == nil || o.ClientResponseBody.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientResponseBody.Get()
}

// GetClientResponseBodyOk returns a tuple with the ClientResponseBody field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataRefreshErrors) GetClientResponseBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientResponseBody.Get(), o.ClientResponseBody.IsSet()
}

// SetClientResponseBody sets field value
func (o *MetadataRefreshErrors) SetClientResponseBody(v string) {
	o.ClientResponseBody.Set(&v)
}

// GetClientResponseStatusCode returns the ClientResponseStatusCode field value
// If the value is explicit nil, the zero value for string will be returned
func (o *MetadataRefreshErrors) GetClientResponseStatusCode() string {
	if o == nil || o.ClientResponseStatusCode.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientResponseStatusCode.Get()
}

// GetClientResponseStatusCodeOk returns a tuple with the ClientResponseStatusCode field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MetadataRefreshErrors) GetClientResponseStatusCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientResponseStatusCode.Get(), o.ClientResponseStatusCode.IsSet()
}

// SetClientResponseStatusCode sets field value
func (o *MetadataRefreshErrors) SetClientResponseStatusCode(v string) {
	o.ClientResponseStatusCode.Set(&v)
}

// GetClientTokenMetadataUrl returns the ClientTokenMetadataUrl field value
func (o *MetadataRefreshErrors) GetClientTokenMetadataUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientTokenMetadataUrl
}

// GetClientTokenMetadataUrlOk returns a tuple with the ClientTokenMetadataUrl field value
// and a boolean to check if the value has been set.
func (o *MetadataRefreshErrors) GetClientTokenMetadataUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientTokenMetadataUrl, true
}

// SetClientTokenMetadataUrl sets field value
func (o *MetadataRefreshErrors) SetClientTokenMetadataUrl(v string) {
	o.ClientTokenMetadataUrl = v
}

// GetCollectionAddress returns the CollectionAddress field value
func (o *MetadataRefreshErrors) GetCollectionAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CollectionAddress
}

// GetCollectionAddressOk returns a tuple with the CollectionAddress field value
// and a boolean to check if the value has been set.
func (o *MetadataRefreshErrors) GetCollectionAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollectionAddress, true
}

// SetCollectionAddress sets field value
func (o *MetadataRefreshErrors) SetCollectionAddress(v string) {
	o.CollectionAddress = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MetadataRefreshErrors) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MetadataRefreshErrors) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MetadataRefreshErrors) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetErrorCode returns the ErrorCode field value
func (o *MetadataRefreshErrors) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *MetadataRefreshErrors) GetErrorCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *MetadataRefreshErrors) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetTokenId returns the TokenId field value
func (o *MetadataRefreshErrors) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *MetadataRefreshErrors) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *MetadataRefreshErrors) SetTokenId(v string) {
	o.TokenId = v
}

func (o MetadataRefreshErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["client_response_body"] = o.ClientResponseBody.Get()
	}
	if true {
		toSerialize["client_response_status_code"] = o.ClientResponseStatusCode.Get()
	}
	if true {
		toSerialize["client_token_metadata_url"] = o.ClientTokenMetadataUrl
	}
	if true {
		toSerialize["collection_address"] = o.CollectionAddress
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if true {
		toSerialize["token_id"] = o.TokenId
	}
	return json.Marshal(toSerialize)
}

type NullableMetadataRefreshErrors struct {
	value *MetadataRefreshErrors
	isSet bool
}

func (v NullableMetadataRefreshErrors) Get() *MetadataRefreshErrors {
	return v.value
}

func (v *NullableMetadataRefreshErrors) Set(val *MetadataRefreshErrors) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataRefreshErrors) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataRefreshErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataRefreshErrors(val *MetadataRefreshErrors) *NullableMetadataRefreshErrors {
	return &NullableMetadataRefreshErrors{value: val, isSet: true}
}

func (v NullableMetadataRefreshErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataRefreshErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

