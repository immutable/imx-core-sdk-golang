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

// GetSignableTransferResponse struct for GetSignableTransferResponse
type GetSignableTransferResponse struct {
	// Sender of the transfer
	SenderStarkKey string `json:"sender_stark_key"`
	// Message to sign with L1 wallet to confirm transfer request
	SignableMessage string `json:"signable_message"`
	// List of transfer responses without the sender stark key
	SignableResponses []SignableTransferResponseDetails `json:"signable_responses"`
}

// NewGetSignableTransferResponse instantiates a new GetSignableTransferResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableTransferResponse(senderStarkKey string, signableMessage string, signableResponses []SignableTransferResponseDetails) *GetSignableTransferResponse {
	this := GetSignableTransferResponse{}
	this.SenderStarkKey = senderStarkKey
	this.SignableMessage = signableMessage
	this.SignableResponses = signableResponses
	return &this
}

// NewGetSignableTransferResponseWithDefaults instantiates a new GetSignableTransferResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableTransferResponseWithDefaults() *GetSignableTransferResponse {
	this := GetSignableTransferResponse{}
	return &this
}

// GetSenderStarkKey returns the SenderStarkKey field value
func (o *GetSignableTransferResponse) GetSenderStarkKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SenderStarkKey
}

// GetSenderStarkKeyOk returns a tuple with the SenderStarkKey field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponse) GetSenderStarkKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SenderStarkKey, true
}

// SetSenderStarkKey sets field value
func (o *GetSignableTransferResponse) SetSenderStarkKey(v string) {
	o.SenderStarkKey = v
}

// GetSignableMessage returns the SignableMessage field value
func (o *GetSignableTransferResponse) GetSignableMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignableMessage
}

// GetSignableMessageOk returns a tuple with the SignableMessage field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponse) GetSignableMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignableMessage, true
}

// SetSignableMessage sets field value
func (o *GetSignableTransferResponse) SetSignableMessage(v string) {
	o.SignableMessage = v
}

// GetSignableResponses returns the SignableResponses field value
func (o *GetSignableTransferResponse) GetSignableResponses() []SignableTransferResponseDetails {
	if o == nil {
		var ret []SignableTransferResponseDetails
		return ret
	}

	return o.SignableResponses
}

// GetSignableResponsesOk returns a tuple with the SignableResponses field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferResponse) GetSignableResponsesOk() ([]SignableTransferResponseDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.SignableResponses, true
}

// SetSignableResponses sets field value
func (o *GetSignableTransferResponse) SetSignableResponses(v []SignableTransferResponseDetails) {
	o.SignableResponses = v
}

func (o GetSignableTransferResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sender_stark_key"] = o.SenderStarkKey
	}
	if true {
		toSerialize["signable_message"] = o.SignableMessage
	}
	if true {
		toSerialize["signable_responses"] = o.SignableResponses
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableTransferResponse struct {
	value *GetSignableTransferResponse
	isSet bool
}

func (v NullableGetSignableTransferResponse) Get() *GetSignableTransferResponse {
	return v.value
}

func (v *NullableGetSignableTransferResponse) Set(val *GetSignableTransferResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableTransferResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableTransferResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableTransferResponse(val *GetSignableTransferResponse) *NullableGetSignableTransferResponse {
	return &NullableGetSignableTransferResponse{value: val, isSet: true}
}

func (v NullableGetSignableTransferResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableTransferResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


