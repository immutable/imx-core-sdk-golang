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

// NftprimarytransactionGetResponse struct for NftprimarytransactionGetResponse
type NftprimarytransactionGetResponse struct {
	// Contract address of the asset
	ContractAddress *string `json:"contract_address,omitempty"`
	// Temporary asset id. Might be a token id if the token id is known or a generic description if it's not
	OfferId *string `json:"offer_id,omitempty"`
	// Provider name
	Provider *string `json:"provider,omitempty"`
	// Ethereum address of the seller
	SellerWalletAddress *string `json:"seller_wallet_address,omitempty"`
	// Transaction status enums(created, waitingPayment, pending, completed, failed)
	Status *string `json:"status,omitempty"`
	// ID of the token that has been successfully minted - might or not be the same as `offer_id`
	TokenId *string `json:"token_id,omitempty"`
	// Transaction id
	TransactionId *string `json:"transaction_id,omitempty"`
	// Wallet address that receives the minted NFT
	UserWalletAddress *string `json:"user_wallet_address,omitempty"`
}

// NewNftprimarytransactionGetResponse instantiates a new NftprimarytransactionGetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftprimarytransactionGetResponse() *NftprimarytransactionGetResponse {
	this := NftprimarytransactionGetResponse{}
	return &this
}

// NewNftprimarytransactionGetResponseWithDefaults instantiates a new NftprimarytransactionGetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftprimarytransactionGetResponseWithDefaults() *NftprimarytransactionGetResponse {
	this := NftprimarytransactionGetResponse{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetContractAddress() string {
	if o == nil || isNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetContractAddressOk() (*string, bool) {
	if o == nil || isNil(o.ContractAddress) {
    return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasContractAddress() bool {
	if o != nil && !isNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *NftprimarytransactionGetResponse) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetOfferId() string {
	if o == nil || isNil(o.OfferId) {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetOfferIdOk() (*string, bool) {
	if o == nil || isNil(o.OfferId) {
    return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasOfferId() bool {
	if o != nil && !isNil(o.OfferId) {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *NftprimarytransactionGetResponse) SetOfferId(v string) {
	o.OfferId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
    return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *NftprimarytransactionGetResponse) SetProvider(v string) {
	o.Provider = &v
}

// GetSellerWalletAddress returns the SellerWalletAddress field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetSellerWalletAddress() string {
	if o == nil || isNil(o.SellerWalletAddress) {
		var ret string
		return ret
	}
	return *o.SellerWalletAddress
}

// GetSellerWalletAddressOk returns a tuple with the SellerWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetSellerWalletAddressOk() (*string, bool) {
	if o == nil || isNil(o.SellerWalletAddress) {
    return nil, false
	}
	return o.SellerWalletAddress, true
}

// HasSellerWalletAddress returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasSellerWalletAddress() bool {
	if o != nil && !isNil(o.SellerWalletAddress) {
		return true
	}

	return false
}

// SetSellerWalletAddress gets a reference to the given string and assigns it to the SellerWalletAddress field.
func (o *NftprimarytransactionGetResponse) SetSellerWalletAddress(v string) {
	o.SellerWalletAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *NftprimarytransactionGetResponse) SetStatus(v string) {
	o.Status = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetTokenId() string {
	if o == nil || isNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || isNil(o.TokenId) {
    return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasTokenId() bool {
	if o != nil && !isNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *NftprimarytransactionGetResponse) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetTransactionId() string {
	if o == nil || isNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.TransactionId) {
    return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasTransactionId() bool {
	if o != nil && !isNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *NftprimarytransactionGetResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetUserWalletAddress returns the UserWalletAddress field value if set, zero value otherwise.
func (o *NftprimarytransactionGetResponse) GetUserWalletAddress() string {
	if o == nil || isNil(o.UserWalletAddress) {
		var ret string
		return ret
	}
	return *o.UserWalletAddress
}

// GetUserWalletAddressOk returns a tuple with the UserWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionGetResponse) GetUserWalletAddressOk() (*string, bool) {
	if o == nil || isNil(o.UserWalletAddress) {
    return nil, false
	}
	return o.UserWalletAddress, true
}

// HasUserWalletAddress returns a boolean if a field has been set.
func (o *NftprimarytransactionGetResponse) HasUserWalletAddress() bool {
	if o != nil && !isNil(o.UserWalletAddress) {
		return true
	}

	return false
}

// SetUserWalletAddress gets a reference to the given string and assigns it to the UserWalletAddress field.
func (o *NftprimarytransactionGetResponse) SetUserWalletAddress(v string) {
	o.UserWalletAddress = &v
}

func (o NftprimarytransactionGetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ContractAddress) {
		toSerialize["contract_address"] = o.ContractAddress
	}
	if !isNil(o.OfferId) {
		toSerialize["offer_id"] = o.OfferId
	}
	if !isNil(o.Provider) {
		toSerialize["provider"] = o.Provider
	}
	if !isNil(o.SellerWalletAddress) {
		toSerialize["seller_wallet_address"] = o.SellerWalletAddress
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !isNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !isNil(o.UserWalletAddress) {
		toSerialize["user_wallet_address"] = o.UserWalletAddress
	}
	return json.Marshal(toSerialize)
}

type NullableNftprimarytransactionGetResponse struct {
	value *NftprimarytransactionGetResponse
	isSet bool
}

func (v NullableNftprimarytransactionGetResponse) Get() *NftprimarytransactionGetResponse {
	return v.value
}

func (v *NullableNftprimarytransactionGetResponse) Set(val *NftprimarytransactionGetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNftprimarytransactionGetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNftprimarytransactionGetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftprimarytransactionGetResponse(val *NftprimarytransactionGetResponse) *NullableNftprimarytransactionGetResponse {
	return &NullableNftprimarytransactionGetResponse{value: val, isSet: true}
}

func (v NullableNftprimarytransactionGetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftprimarytransactionGetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


