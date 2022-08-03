/*
Immutable X API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1
Contact: support@immutable.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// GetSignableTransferRequestV1 struct for GetSignableTransferRequestV1
type GetSignableTransferRequestV1 struct {
	// Amount of the token to transfer
	Amount string `json:"amount"`
	// Ethereum address of the receiving user
	Receiver string `json:"receiver"`
	// Ethereum address of the transferring user
	Sender string `json:"sender"`
	Token SignableToken `json:"token"`
}

// NewGetSignableTransferRequestV1 instantiates a new GetSignableTransferRequestV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableTransferRequestV1(amount string, receiver string, sender string, token SignableToken) *GetSignableTransferRequestV1 {
	this := GetSignableTransferRequestV1{}
	this.Amount = amount
	this.Receiver = receiver
	this.Sender = sender
	this.Token = token
	return &this
}

// NewGetSignableTransferRequestV1WithDefaults instantiates a new GetSignableTransferRequestV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableTransferRequestV1WithDefaults() *GetSignableTransferRequestV1 {
	this := GetSignableTransferRequestV1{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetSignableTransferRequestV1) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferRequestV1) GetAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetSignableTransferRequestV1) SetAmount(v string) {
	o.Amount = v
}

// GetReceiver returns the Receiver field value
func (o *GetSignableTransferRequestV1) GetReceiver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferRequestV1) GetReceiverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *GetSignableTransferRequestV1) SetReceiver(v string) {
	o.Receiver = v
}

// GetSender returns the Sender field value
func (o *GetSignableTransferRequestV1) GetSender() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sender
}

// GetSenderOk returns a tuple with the Sender field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferRequestV1) GetSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sender, true
}

// SetSender sets field value
func (o *GetSignableTransferRequestV1) SetSender(v string) {
	o.Sender = v
}

// GetToken returns the Token field value
func (o *GetSignableTransferRequestV1) GetToken() SignableToken {
	if o == nil {
		var ret SignableToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetSignableTransferRequestV1) GetTokenOk() (*SignableToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GetSignableTransferRequestV1) SetToken(v SignableToken) {
	o.Token = v
}

func (o GetSignableTransferRequestV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	if true {
		toSerialize["sender"] = o.Sender
	}
	if true {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableTransferRequestV1 struct {
	value *GetSignableTransferRequestV1
	isSet bool
}

func (v NullableGetSignableTransferRequestV1) Get() *GetSignableTransferRequestV1 {
	return v.value
}

func (v *NullableGetSignableTransferRequestV1) Set(val *GetSignableTransferRequestV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableTransferRequestV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableTransferRequestV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableTransferRequestV1(val *GetSignableTransferRequestV1) *NullableGetSignableTransferRequestV1 {
	return &NullableGetSignableTransferRequestV1{value: val, isSet: true}
}

func (v NullableGetSignableTransferRequestV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableTransferRequestV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


