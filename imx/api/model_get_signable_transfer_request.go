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

// GetSignableTransferRequest struct for GetSignableTransferRequest
type GetSignableTransferRequest struct {
	// Ethereum address of the transferring user
	SenderEtherKey string `json:"sender_ether_key"`
	// List of signable transfer details
	SignableRequests []SignableTransferDetails `json:"signable_requests"`
}

// NewGetSignableTransferRequest instantiates a new GetSignableTransferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableTransferRequest(senderEtherKey string, signableRequests []SignableTransferDetails) *GetSignableTransferRequest {
	this := GetSignableTransferRequest{}
	this.SenderEtherKey = senderEtherKey
	this.SignableRequests = signableRequests
	return &this
}

// NewGetSignableTransferRequestWithDefaults instantiates a new GetSignableTransferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableTransferRequestWithDefaults() *GetSignableTransferRequest {
	this := GetSignableTransferRequest{}
	return &this
}

// GetSenderEtherKey returns the SenderEtherKey field value
func (o *GetSignableTransferRequest) GetSenderEtherKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderEtherKey
}

// GetSenderEtherKeyOk returns a tuple with the SenderEtherKey field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferRequest) GetSenderEtherKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderEtherKey, true
}

// SetSenderEtherKey sets field value
func (o *GetSignableTransferRequest) SetSenderEtherKey(v string) {
	o.SenderEtherKey = v
}

// GetSignableRequests returns the SignableRequests field value
func (o *GetSignableTransferRequest) GetSignableRequests() []SignableTransferDetails {
	if o == nil {
		var ret []SignableTransferDetails
		return ret
	}

	return o.SignableRequests
}

// GetSignableRequestsOk returns a tuple with the SignableRequests field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferRequest) GetSignableRequestsOk() ([]SignableTransferDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignableRequests, true
}

// SetSignableRequests sets field value
func (o *GetSignableTransferRequest) SetSignableRequests(v []SignableTransferDetails) {
	o.SignableRequests = v
}

func (o GetSignableTransferRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sender_ether_key"] = o.SenderEtherKey
	}
	if true {
		toSerialize["signable_requests"] = o.SignableRequests
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableTransferRequest struct {
	value *GetSignableTransferRequest
	isSet bool
}

func (v NullableGetSignableTransferRequest) Get() *GetSignableTransferRequest {
	return v.value
}

func (v *NullableGetSignableTransferRequest) Set(val *GetSignableTransferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableTransferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableTransferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableTransferRequest(val *GetSignableTransferRequest) *NullableGetSignableTransferRequest {
	return &NullableGetSignableTransferRequest{value: val, isSet: true}
}

func (v NullableGetSignableTransferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableTransferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


