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

// GetSignableOrderRequest struct for GetSignableOrderRequest
type GetSignableOrderRequest struct {
	// Fee-exclusive amount to buy
	AmountBuy string `json:"amount_buy"`
	// Amount to sell
	AmountSell string `json:"amount_sell"`
	// ExpirationTimestamp in Unix time. Note: will be rounded down to the nearest hour
	ExpirationTimestamp *int32 `json:"expiration_timestamp,omitempty"`
	// Inclusion of either maker or taker fees
	Fees []FeeEntry `json:"fees,omitempty"`
	TokenBuy SignableToken `json:"token_buy"`
	TokenSell SignableToken `json:"token_sell"`
	// Ethereum address of the submitting user
	User string `json:"user"`
}

// NewGetSignableOrderRequest instantiates a new GetSignableOrderRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignableOrderRequest(amountBuy string, amountSell string, tokenBuy SignableToken, tokenSell SignableToken, user string) *GetSignableOrderRequest {
	this := GetSignableOrderRequest{}
	this.AmountBuy = amountBuy
	this.AmountSell = amountSell
	this.TokenBuy = tokenBuy
	this.TokenSell = tokenSell
	this.User = user
	return &this
}

// NewGetSignableOrderRequestWithDefaults instantiates a new GetSignableOrderRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignableOrderRequestWithDefaults() *GetSignableOrderRequest {
	this := GetSignableOrderRequest{}
	return &this
}

// GetAmountBuy returns the AmountBuy field value
func (o *GetSignableOrderRequest) GetAmountBuy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountBuy
}

// GetAmountBuyOk returns a tuple with the AmountBuy field value
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetAmountBuyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountBuy, true
}

// SetAmountBuy sets field value
func (o *GetSignableOrderRequest) SetAmountBuy(v string) {
	o.AmountBuy = v
}

// GetAmountSell returns the AmountSell field value
func (o *GetSignableOrderRequest) GetAmountSell() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AmountSell
}

// GetAmountSellOk returns a tuple with the AmountSell field value
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetAmountSellOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AmountSell, true
}

// SetAmountSell sets field value
func (o *GetSignableOrderRequest) SetAmountSell(v string) {
	o.AmountSell = v
}

// GetExpirationTimestamp returns the ExpirationTimestamp field value if set, zero value otherwise.
func (o *GetSignableOrderRequest) GetExpirationTimestamp() int32 {
	if o == nil || o.ExpirationTimestamp == nil {
		var ret int32
		return ret
	}
	return *o.ExpirationTimestamp
}

// GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetExpirationTimestampOk() (*int32, bool) {
	if o == nil || o.ExpirationTimestamp == nil {
		return nil, false
	}
	return o.ExpirationTimestamp, true
}

// HasExpirationTimestamp returns a boolean if a field has been set.
func (o *GetSignableOrderRequest) HasExpirationTimestamp() bool {
	if o != nil && o.ExpirationTimestamp != nil {
		return true
	}

	return false
}

// SetExpirationTimestamp gets a reference to the given int32 and assigns it to the ExpirationTimestamp field.
func (o *GetSignableOrderRequest) SetExpirationTimestamp(v int32) {
	o.ExpirationTimestamp = &v
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *GetSignableOrderRequest) GetFees() []FeeEntry {
	if o == nil || o.Fees == nil {
		var ret []FeeEntry
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetFeesOk() ([]FeeEntry, bool) {
	if o == nil || o.Fees == nil {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *GetSignableOrderRequest) HasFees() bool {
	if o != nil && o.Fees != nil {
		return true
	}

	return false
}

// SetFees gets a reference to the given []FeeEntry and assigns it to the Fees field.
func (o *GetSignableOrderRequest) SetFees(v []FeeEntry) {
	o.Fees = v
}

// GetTokenBuy returns the TokenBuy field value
func (o *GetSignableOrderRequest) GetTokenBuy() SignableToken {
	if o == nil {
		var ret SignableToken
		return ret
	}

	return o.TokenBuy
}

// GetTokenBuyOk returns a tuple with the TokenBuy field value
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetTokenBuyOk() (*SignableToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenBuy, true
}

// SetTokenBuy sets field value
func (o *GetSignableOrderRequest) SetTokenBuy(v SignableToken) {
	o.TokenBuy = v
}

// GetTokenSell returns the TokenSell field value
func (o *GetSignableOrderRequest) GetTokenSell() SignableToken {
	if o == nil {
		var ret SignableToken
		return ret
	}

	return o.TokenSell
}

// GetTokenSellOk returns a tuple with the TokenSell field value
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetTokenSellOk() (*SignableToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenSell, true
}

// SetTokenSell sets field value
func (o *GetSignableOrderRequest) SetTokenSell(v SignableToken) {
	o.TokenSell = v
}

// GetUser returns the User field value
func (o *GetSignableOrderRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *GetSignableOrderRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *GetSignableOrderRequest) SetUser(v string) {
	o.User = v
}

func (o GetSignableOrderRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount_buy"] = o.AmountBuy
	}
	if true {
		toSerialize["amount_sell"] = o.AmountSell
	}
	if o.ExpirationTimestamp != nil {
		toSerialize["expiration_timestamp"] = o.ExpirationTimestamp
	}
	if o.Fees != nil {
		toSerialize["fees"] = o.Fees
	}
	if true {
		toSerialize["token_buy"] = o.TokenBuy
	}
	if true {
		toSerialize["token_sell"] = o.TokenSell
	}
	if true {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableGetSignableOrderRequest struct {
	value *GetSignableOrderRequest
	isSet bool
}

func (v NullableGetSignableOrderRequest) Get() *GetSignableOrderRequest {
	return v.value
}

func (v *NullableGetSignableOrderRequest) Set(val *GetSignableOrderRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignableOrderRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignableOrderRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignableOrderRequest(val *GetSignableOrderRequest) *NullableGetSignableOrderRequest {
	return &NullableGetSignableOrderRequest{value: val, isSet: true}
}

func (v NullableGetSignableOrderRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignableOrderRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


