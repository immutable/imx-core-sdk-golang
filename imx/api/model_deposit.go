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

// Deposit struct for Deposit
type Deposit struct {
	// Status of this deposit in Immutable X
	Status string `json:"status"`
	// Timestamp of the deposit
	Timestamp string `json:"timestamp"`
	Token Token `json:"token"`
	// Sequential ID of this transaction within Immutable X
	TransactionId int32 `json:"transaction_id"`
	// Ethereum address of the user making this deposit
	User string `json:"user"`
}

// NewDeposit instantiates a new Deposit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeposit(status string, timestamp string, token Token, transactionId int32, user string) *Deposit {
	this := Deposit{}
	this.Status = status
	this.Timestamp = timestamp
	this.Token = token
	this.TransactionId = transactionId
	this.User = user
	return &this
}

// NewDepositWithDefaults instantiates a new Deposit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDepositWithDefaults() *Deposit {
	this := Deposit{}
	return &this
}

// GetStatus returns the Status field value
func (o *Deposit) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Deposit) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
func (o *Deposit) GetTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *Deposit) SetTimestamp(v string) {
	o.Timestamp = v
}

// GetToken returns the Token field value
func (o *Deposit) GetToken() Token {
	if o == nil {
		var ret Token
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetTokenOk() (*Token, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *Deposit) SetToken(v Token) {
	o.Token = v
}

// GetTransactionId returns the TransactionId field value
func (o *Deposit) GetTransactionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetTransactionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Deposit) SetTransactionId(v int32) {
	o.TransactionId = v
}

// GetUser returns the User field value
func (o *Deposit) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Deposit) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Deposit) SetUser(v string) {
	o.User = v
}

func (o Deposit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableDeposit struct {
	value *Deposit
	isSet bool
}

func (v NullableDeposit) Get() *Deposit {
	return v.value
}

func (v *NullableDeposit) Set(val *Deposit) {
	v.value = val
	v.isSet = true
}

func (v NullableDeposit) IsSet() bool {
	return v.isSet
}

func (v *NullableDeposit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeposit(val *Deposit) *NullableDeposit {
	return &NullableDeposit{value: val, isSet: true}
}

func (v NullableDeposit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeposit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


