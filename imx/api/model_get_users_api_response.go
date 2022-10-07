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

// GetUsersApiResponse struct for GetUsersApiResponse
type GetUsersApiResponse struct {
	// Accounts
	Accounts []string `json:"accounts"`
}

// NewGetUsersApiResponse instantiates a new GetUsersApiResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUsersApiResponse(accounts []string) *GetUsersApiResponse {
	this := GetUsersApiResponse{}
	this.Accounts = accounts
	return &this
}

// NewGetUsersApiResponseWithDefaults instantiates a new GetUsersApiResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUsersApiResponseWithDefaults() *GetUsersApiResponse {
	this := GetUsersApiResponse{}
	return &this
}

// GetAccounts returns the Accounts field value
func (o *GetUsersApiResponse) GetAccounts() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value
// and a boolean to check if the value has been set.
func (o *GetUsersApiResponse) GetAccountsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Accounts, true
}

// SetAccounts sets field value
func (o *GetUsersApiResponse) SetAccounts(v []string) {
	o.Accounts = v
}

func (o GetUsersApiResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accounts"] = o.Accounts
	}
	return json.Marshal(toSerialize)
}

type NullableGetUsersApiResponse struct {
	value *GetUsersApiResponse
	isSet bool
}

func (v NullableGetUsersApiResponse) Get() *GetUsersApiResponse {
	return v.value
}

func (v *NullableGetUsersApiResponse) Set(val *GetUsersApiResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUsersApiResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUsersApiResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUsersApiResponse(val *GetUsersApiResponse) *NullableGetUsersApiResponse {
	return &NullableGetUsersApiResponse{value: val, isSet: true}
}

func (v NullableGetUsersApiResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUsersApiResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


