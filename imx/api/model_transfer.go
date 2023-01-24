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

// Transfer struct for Transfer
type Transfer struct {
	// Ethereum address of the user who received this transfer
	Receiver string `json:"receiver"`
	// Status of the transaction
	Status string `json:"status"`
	// Timestamp of the transfer
	Timestamp NullableString `json:"timestamp"`
	Token TransferToken `json:"token"`
	// Sequential transaction ID
	TransactionId int32 `json:"transaction_id"`
	// Ethereum address of the user  who submitted this transfer
	User string `json:"user"`
}

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer(receiver string, status string, timestamp NullableString, token TransferToken, transactionId int32, user string) *Transfer {
	this := Transfer{}
	this.Receiver = receiver
	this.Status = status
	this.Timestamp = timestamp
	this.Token = token
	this.TransactionId = transactionId
	this.User = user
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	return &this
}

// GetReceiver returns the Receiver field value
func (o *Transfer) GetReceiver() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Receiver
}

// GetReceiverOk returns a tuple with the Receiver field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetReceiverOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Receiver, true
}

// SetReceiver sets field value
func (o *Transfer) SetReceiver(v string) {
	o.Receiver = v
}

// GetStatus returns the Status field value
func (o *Transfer) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Transfer) SetStatus(v string) {
	o.Status = v
}

// GetTimestamp returns the Timestamp field value
// If the value is explicit nil, the zero value for string will be returned
func (o *Transfer) GetTimestamp() string {
	if o == nil || o.Timestamp.Get() == nil {
		var ret string
		return ret
	}

	return *o.Timestamp.Get()
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Transfer) GetTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Timestamp.Get(), o.Timestamp.IsSet()
}

// SetTimestamp sets field value
func (o *Transfer) SetTimestamp(v string) {
	o.Timestamp.Set(&v)
}

// GetToken returns the Token field value
func (o *Transfer) GetToken() TransferToken {
	if o == nil {
		var ret TransferToken
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetTokenOk() (*TransferToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *Transfer) SetToken(v TransferToken) {
	o.Token = v
}

// GetTransactionId returns the TransactionId field value
func (o *Transfer) GetTransactionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetTransactionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionId, true
}

// SetTransactionId sets field value
func (o *Transfer) SetTransactionId(v int32) {
	o.TransactionId = v
}

// GetUser returns the User field value
func (o *Transfer) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Transfer) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Transfer) SetUser(v string) {
	o.User = v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["receiver"] = o.Receiver
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp.Get()
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

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


