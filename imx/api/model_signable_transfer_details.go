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

// SignableTransferDetails struct for SignableTransferDetails
type SignableTransferDetails struct {
	// Amount of the token to transfer
	Amount string `json:"amount"`
	// Ethereum address of the receiving user
	Receiver string `json:"receiver"`
	Token SignableToken `json:"token"`
}

// NewSignableTransferDetails instantiates a new SignableTransferDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignableTransferDetails(amount string, receiver string, token SignableToken) *SignableTransferDetails {
	this := SignableTransferDetails{}
	this.Amount = amount
	this.Receiver = receiver
	this.Token = token
	return &this
}

// NewSignableTransferDetailsWithDefaults instantiates a new SignableTransferDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignableTransferDetailsWithDefaults() *SignableTransferDetails {
	this := SignableTransferDetails{}
	return &this
}

// GetAmount returns the Amount field value
func (o *SignableTransferDetails) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *SignableTransferDetails) GetAmountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *SignableTransferDetails) SetAmount(v string) {
	o.Amount = v
}

// GetReceiver returns the Receiver field value
func (o *SignableTransferDetails) GetReceiver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *SignableTransferDetails) GetReceiverOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *SignableTransferDetails) SetReceiver(v string) {
	o.Receiver = v
}

// GetToken returns the Token field value
func (o *SignableTransferDetails) GetToken() SignableToken {
	if o == nil {
		var ret SignableToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *SignableTransferDetails) GetTokenOk() (*SignableToken, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *SignableTransferDetails) SetToken(v SignableToken) {
	o.Token = v
}

func (o SignableTransferDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableSignableTransferDetails struct {
	value *SignableTransferDetails
	isSet bool
}

func (v NullableSignableTransferDetails) Get() *SignableTransferDetails {
	return v.value
}

func (v *NullableSignableTransferDetails) Set(val *SignableTransferDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSignableTransferDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSignableTransferDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignableTransferDetails(val *SignableTransferDetails) *NullableSignableTransferDetails {
	return &NullableSignableTransferDetails{value: val, isSet: true}
}

func (v NullableSignableTransferDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignableTransferDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


