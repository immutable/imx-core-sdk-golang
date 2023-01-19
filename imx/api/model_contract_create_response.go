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

// ContractCreateResponse struct for ContractCreateResponse
type ContractCreateResponse struct {
	WebhookHash *string `json:"webhook_hash,omitempty"`
}

// NewContractCreateResponse instantiates a new ContractCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContractCreateResponse() *ContractCreateResponse {
	this := ContractCreateResponse{}
	return &this
}

// NewContractCreateResponseWithDefaults instantiates a new ContractCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContractCreateResponseWithDefaults() *ContractCreateResponse {
	this := ContractCreateResponse{}
	return &this
}

// GetWebhookHash returns the WebhookHash field value if set, zero value otherwise.
func (o *ContractCreateResponse) GetWebhookHash() string {
	if o == nil || isNil(o.WebhookHash) {
		var ret string
		return ret
	}
	return *o.WebhookHash
}

// GetWebhookHashOk returns a tuple with the WebhookHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContractCreateResponse) GetWebhookHashOk() (*string, bool) {
	if o == nil || isNil(o.WebhookHash) {
    return nil, false
	}
	return o.WebhookHash, true
}

// HasWebhookHash returns a boolean if a field has been set.
func (o *ContractCreateResponse) HasWebhookHash() bool {
	if o != nil && !isNil(o.WebhookHash) {
		return true
	}

	return false
}

// SetWebhookHash gets a reference to the given string and assigns it to the WebhookHash field.
func (o *ContractCreateResponse) SetWebhookHash(v string) {
	o.WebhookHash = &v
}

func (o ContractCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.WebhookHash) {
		toSerialize["webhook_hash"] = o.WebhookHash
	}
	return json.Marshal(toSerialize)
}

type NullableContractCreateResponse struct {
	value *ContractCreateResponse
	isSet bool
}

func (v NullableContractCreateResponse) Get() *ContractCreateResponse {
	return v.value
}

func (v *NullableContractCreateResponse) Set(val *ContractCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableContractCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableContractCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContractCreateResponse(val *ContractCreateResponse) *NullableContractCreateResponse {
	return &NullableContractCreateResponse{value: val, isSet: true}
}

func (v NullableContractCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContractCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


