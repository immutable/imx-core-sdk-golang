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

// CreateTransferResponseV1 struct for CreateTransferResponseV1
type CreateTransferResponseV1 struct {
	// [deprecated] Sent signature
	SentSignature string `json:"sent_signature"`
	// [deprecated] The status of transfer
	Status string `json:"status"`
	// [deprecated] Time of the transfer
	Time int32 `json:"time"`
	// ID of the transfer
	TransferId int32 `json:"transfer_id"`
}

// NewCreateTransferResponseV1 instantiates a new CreateTransferResponseV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTransferResponseV1(sentSignature string, status string, time int32, transferId int32) *CreateTransferResponseV1 {
	this := CreateTransferResponseV1{}
	this.SentSignature = sentSignature
	this.Status = status
	this.Time = time
	this.TransferId = transferId
	return &this
}

// NewCreateTransferResponseV1WithDefaults instantiates a new CreateTransferResponseV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTransferResponseV1WithDefaults() *CreateTransferResponseV1 {
	this := CreateTransferResponseV1{}
	return &this
}

// GetSentSignature returns the SentSignature field value
func (o *CreateTransferResponseV1) GetSentSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SentSignature
}

// GetSentSignatureOk returns a tuple with the SentSignature field value
// and a boolean to check if the value has been set.
func (o *CreateTransferResponseV1) GetSentSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SentSignature, true
}

// SetSentSignature sets field value
func (o *CreateTransferResponseV1) SetSentSignature(v string) {
	o.SentSignature = v
}

// GetStatus returns the Status field value
func (o *CreateTransferResponseV1) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CreateTransferResponseV1) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CreateTransferResponseV1) SetStatus(v string) {
	o.Status = v
}

// GetTime returns the Time field value
func (o *CreateTransferResponseV1) GetTime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Time
}

// GetTimeOk returns a tuple with the Time field value
// and a boolean to check if the value has been set.
func (o *CreateTransferResponseV1) GetTimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Time, true
}

// SetTime sets field value
func (o *CreateTransferResponseV1) SetTime(v int32) {
	o.Time = v
}

// GetTransferId returns the TransferId field value
func (o *CreateTransferResponseV1) GetTransferId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *CreateTransferResponseV1) GetTransferIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *CreateTransferResponseV1) SetTransferId(v int32) {
	o.TransferId = v
}

func (o CreateTransferResponseV1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sent_signature"] = o.SentSignature
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["time"] = o.Time
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTransferResponseV1 struct {
	value *CreateTransferResponseV1
	isSet bool
}

func (v NullableCreateTransferResponseV1) Get() *CreateTransferResponseV1 {
	return v.value
}

func (v *NullableCreateTransferResponseV1) Set(val *CreateTransferResponseV1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTransferResponseV1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTransferResponseV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTransferResponseV1(val *CreateTransferResponseV1) *NullableCreateTransferResponseV1 {
	return &NullableCreateTransferResponseV1{value: val, isSet: true}
}

func (v NullableCreateTransferResponseV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTransferResponseV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


