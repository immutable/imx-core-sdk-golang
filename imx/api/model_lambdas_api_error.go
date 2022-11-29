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

// LambdasAPIError struct for LambdasAPIError
type LambdasAPIError struct {
	// The error code
	Code *string `json:"code,omitempty"`
	// The error details
	Details *string `json:"details,omitempty"`
	// The error message
	Message *string `json:"message,omitempty"`
	// The error status code
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewLambdasAPIError instantiates a new LambdasAPIError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLambdasAPIError() *LambdasAPIError {
	this := LambdasAPIError{}
	return &this
}

// NewLambdasAPIErrorWithDefaults instantiates a new LambdasAPIError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLambdasAPIErrorWithDefaults() *LambdasAPIError {
	this := LambdasAPIError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *LambdasAPIError) GetCode() string {
	if o == nil || o.Code == nil {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LambdasAPIError) GetCodeOk() (*string, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *LambdasAPIError) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *LambdasAPIError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *LambdasAPIError) GetDetails() string {
	if o == nil || o.Details == nil {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LambdasAPIError) GetDetailsOk() (*string, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *LambdasAPIError) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *LambdasAPIError) SetDetails(v string) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LambdasAPIError) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LambdasAPIError) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LambdasAPIError) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LambdasAPIError) SetMessage(v string) {
	o.Message = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *LambdasAPIError) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LambdasAPIError) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *LambdasAPIError) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *LambdasAPIError) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o LambdasAPIError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	return json.Marshal(toSerialize)
}

type NullableLambdasAPIError struct {
	value *LambdasAPIError
	isSet bool
}

func (v NullableLambdasAPIError) Get() *LambdasAPIError {
	return v.value
}

func (v *NullableLambdasAPIError) Set(val *LambdasAPIError) {
	v.value = val
	v.isSet = true
}

func (v NullableLambdasAPIError) IsSet() bool {
	return v.isSet
}

func (v *NullableLambdasAPIError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLambdasAPIError(val *LambdasAPIError) *NullableLambdasAPIError {
	return &NullableLambdasAPIError{value: val, isSet: true}
}

func (v NullableLambdasAPIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLambdasAPIError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


