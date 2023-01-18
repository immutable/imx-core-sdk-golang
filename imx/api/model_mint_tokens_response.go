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

// checks if the MintTokensResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MintTokensResponse{}

// MintTokensResponse struct for MintTokensResponse
type MintTokensResponse struct {
	// List of mint result details
	Results []MintResultDetails `json:"results"`
}

// NewMintTokensResponse instantiates a new MintTokensResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMintTokensResponse(results []MintResultDetails) *MintTokensResponse {
	this := MintTokensResponse{}
	this.Results = results
	return &this
}

// NewMintTokensResponseWithDefaults instantiates a new MintTokensResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMintTokensResponseWithDefaults() *MintTokensResponse {
	this := MintTokensResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *MintTokensResponse) GetResults() []MintResultDetails {
	if o == nil {
		var ret []MintResultDetails
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *MintTokensResponse) GetResultsOk() ([]MintResultDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *MintTokensResponse) SetResults(v []MintResultDetails) {
	o.Results = v
}

func (o MintTokensResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MintTokensResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableMintTokensResponse struct {
	value *MintTokensResponse
	isSet bool
}

func (v NullableMintTokensResponse) Get() *MintTokensResponse {
	return v.value
}

func (v *NullableMintTokensResponse) Set(val *MintTokensResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMintTokensResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMintTokensResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMintTokensResponse(val *MintTokensResponse) *NullableMintTokensResponse {
	return &NullableMintTokensResponse{value: val, isSet: true}
}

func (v NullableMintTokensResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMintTokensResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


