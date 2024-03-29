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

// NftprimarytransactionListTransactionsResponse struct for NftprimarytransactionListTransactionsResponse
type NftprimarytransactionListTransactionsResponse struct {
	// Generated cursor returned by previous query
	Cursor *string `json:"cursor,omitempty"`
	// Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results
	Remaining *int32 `json:"remaining,omitempty"`
	// Transactions matching query parameters
	Result []NftprimarytransactionTransactionData `json:"result,omitempty"`
}

// NewNftprimarytransactionListTransactionsResponse instantiates a new NftprimarytransactionListTransactionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftprimarytransactionListTransactionsResponse() *NftprimarytransactionListTransactionsResponse {
	this := NftprimarytransactionListTransactionsResponse{}
	return &this
}

// NewNftprimarytransactionListTransactionsResponseWithDefaults instantiates a new NftprimarytransactionListTransactionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftprimarytransactionListTransactionsResponseWithDefaults() *NftprimarytransactionListTransactionsResponse {
	this := NftprimarytransactionListTransactionsResponse{}
	return &this
}

// GetCursor returns the Cursor field value if set, zero value otherwise.
func (o *NftprimarytransactionListTransactionsResponse) GetCursor() string {
	if o == nil || o.Cursor == nil {
		var ret string
		return ret
	}
	return *o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionListTransactionsResponse) GetCursorOk() (*string, bool) {
	if o == nil || o.Cursor == nil {
		return nil, false
	}
	return o.Cursor, true
}

// HasCursor returns a boolean if a field has been set.
func (o *NftprimarytransactionListTransactionsResponse) HasCursor() bool {
	if o != nil && o.Cursor != nil {
		return true
	}

	return false
}

// SetCursor gets a reference to the given string and assigns it to the Cursor field.
func (o *NftprimarytransactionListTransactionsResponse) SetCursor(v string) {
	o.Cursor = &v
}

// GetRemaining returns the Remaining field value if set, zero value otherwise.
func (o *NftprimarytransactionListTransactionsResponse) GetRemaining() int32 {
	if o == nil || o.Remaining == nil {
		var ret int32
		return ret
	}
	return *o.Remaining
}

// GetRemainingOk returns a tuple with the Remaining field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionListTransactionsResponse) GetRemainingOk() (*int32, bool) {
	if o == nil || o.Remaining == nil {
		return nil, false
	}
	return o.Remaining, true
}

// HasRemaining returns a boolean if a field has been set.
func (o *NftprimarytransactionListTransactionsResponse) HasRemaining() bool {
	if o != nil && o.Remaining != nil {
		return true
	}

	return false
}

// SetRemaining gets a reference to the given int32 and assigns it to the Remaining field.
func (o *NftprimarytransactionListTransactionsResponse) SetRemaining(v int32) {
	o.Remaining = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *NftprimarytransactionListTransactionsResponse) GetResult() []NftprimarytransactionTransactionData {
	if o == nil || o.Result == nil {
		var ret []NftprimarytransactionTransactionData
		return ret
	}
	return o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionListTransactionsResponse) GetResultOk() ([]NftprimarytransactionTransactionData, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *NftprimarytransactionListTransactionsResponse) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given []NftprimarytransactionTransactionData and assigns it to the Result field.
func (o *NftprimarytransactionListTransactionsResponse) SetResult(v []NftprimarytransactionTransactionData) {
	o.Result = v
}

func (o NftprimarytransactionListTransactionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cursor != nil {
		toSerialize["cursor"] = o.Cursor
	}
	if o.Remaining != nil {
		toSerialize["remaining"] = o.Remaining
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	return json.Marshal(toSerialize)
}

type NullableNftprimarytransactionListTransactionsResponse struct {
	value *NftprimarytransactionListTransactionsResponse
	isSet bool
}

func (v NullableNftprimarytransactionListTransactionsResponse) Get() *NftprimarytransactionListTransactionsResponse {
	return v.value
}

func (v *NullableNftprimarytransactionListTransactionsResponse) Set(val *NftprimarytransactionListTransactionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNftprimarytransactionListTransactionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNftprimarytransactionListTransactionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftprimarytransactionListTransactionsResponse(val *NftprimarytransactionListTransactionsResponse) *NullableNftprimarytransactionListTransactionsResponse {
	return &NullableNftprimarytransactionListTransactionsResponse{value: val, isSet: true}
}

func (v NullableNftprimarytransactionListTransactionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftprimarytransactionListTransactionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


