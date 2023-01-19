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

// GetSignableDepositRequest struct for GetSignableDepositRequest
type GetSignableDepositRequest struct {
	// Amount of the token the user is depositing
	Amount string `json:"amount"`
	Token SignableToken `json:"token"`
	// User who is depositing
	User string `json:"user"`
}

// NewGetSignableDepositRequest instantiates a new GetSignableDepositRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableDepositRequest(amount string, token SignableToken, user string) *GetSignableDepositRequest {
	this := GetSignableDepositRequest{}
	this.Amount = amount
	this.Token = token
	this.User = user
	return &this
}

// NewGetSignableDepositRequestWithDefaults instantiates a new GetSignableDepositRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableDepositRequestWithDefaults() *GetSignableDepositRequest {
	this := GetSignableDepositRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *GetSignableDepositRequest) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *GetSignableDepositRequest) GetAmountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *GetSignableDepositRequest) SetAmount(v string) {
	o.Amount = v
}

// GetToken returns the Token field value
func (o *GetSignableDepositRequest) GetToken() SignableToken {
	if o == nil {
		var ret SignableToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *GetSignableDepositRequest) GetTokenOk() (*SignableToken, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *GetSignableDepositRequest) SetToken(v SignableToken) {
	o.Token = v
}

// GetUser returns the User field value
func (o *GetSignableDepositRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GetSignableDepositRequest) GetUserOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GetSignableDepositRequest) SetUser(v string) {
	o.User = v
}

func (o GetSignableDepositRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableDepositRequest struct {
	value *GetSignableDepositRequest
	isSet bool
}

func (v NullableGetSignableDepositRequest) Get() *GetSignableDepositRequest {
	return v.value
}

func (v *NullableGetSignableDepositRequest) Set(val *GetSignableDepositRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableDepositRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableDepositRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableDepositRequest(val *GetSignableDepositRequest) *NullableGetSignableDepositRequest {
	return &NullableGetSignableDepositRequest{value: val, isSet: true}
}

func (v NullableGetSignableDepositRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableDepositRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


