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

// checks if the NftprimarytransactionCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NftprimarytransactionCreateResponse{}

// NftprimarytransactionCreateResponse struct for NftprimarytransactionCreateResponse
type NftprimarytransactionCreateResponse struct {
	// Contract address of the asset to be created
	ContractAddress *string `json:"contract_address,omitempty"`
	// Temporary asset id. Might be a token id if the token id is known or a generic description if it's not
	OfferId *string `json:"offer_id,omitempty"`
	// Provider name
	Provider *string `json:"provider,omitempty"`
	// Wallet address that will receive the payment (in crypto) from the checkout provider for the minted NFT
	SellerWalletAddress *string `json:"seller_wallet_address,omitempty"`
	// ID of the token that has been successfully minted - should be the same as `offer_id`
	TokenId *string `json:"token_id,omitempty"`
	// Transaction id
	TransactionId *string `json:"transaction_id,omitempty"`
	// NFT purchase URL given by the checkout provider that the user can use to complete payment
	Url *string `json:"url,omitempty"`
	// Ethereum address of the user who wants to create transaction
	UserWalletAddress *string `json:"user_wallet_address,omitempty"`
}

// NewNftprimarytransactionCreateResponse instantiates a new NftprimarytransactionCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNftprimarytransactionCreateResponse() *NftprimarytransactionCreateResponse {
	this := NftprimarytransactionCreateResponse{}
	return &this
}

// NewNftprimarytransactionCreateResponseWithDefaults instantiates a new NftprimarytransactionCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNftprimarytransactionCreateResponseWithDefaults() *NftprimarytransactionCreateResponse {
	this := NftprimarytransactionCreateResponse{}
	return &this
}

// GetContractAddress returns the ContractAddress field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetContractAddress() string {
	if o == nil || isNil(o.ContractAddress) {
		var ret string
		return ret
	}
	return *o.ContractAddress
}

// GetContractAddressOk returns a tuple with the ContractAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetContractAddressOk() (*string, bool) {
	if o == nil || isNil(o.ContractAddress) {
		return nil, false
	}
	return o.ContractAddress, true
}

// HasContractAddress returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasContractAddress() bool {
	if o != nil && !isNil(o.ContractAddress) {
		return true
	}

	return false
}

// SetContractAddress gets a reference to the given string and assigns it to the ContractAddress field.
func (o *NftprimarytransactionCreateResponse) SetContractAddress(v string) {
	o.ContractAddress = &v
}

// GetOfferId returns the OfferId field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetOfferId() string {
	if o == nil || isNil(o.OfferId) {
		var ret string
		return ret
	}
	return *o.OfferId
}

// GetOfferIdOk returns a tuple with the OfferId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetOfferIdOk() (*string, bool) {
	if o == nil || isNil(o.OfferId) {
		return nil, false
	}
	return o.OfferId, true
}

// HasOfferId returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasOfferId() bool {
	if o != nil && !isNil(o.OfferId) {
		return true
	}

	return false
}

// SetOfferId gets a reference to the given string and assigns it to the OfferId field.
func (o *NftprimarytransactionCreateResponse) SetOfferId(v string) {
	o.OfferId = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetProvider() string {
	if o == nil || isNil(o.Provider) {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetProviderOk() (*string, bool) {
	if o == nil || isNil(o.Provider) {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasProvider() bool {
	if o != nil && !isNil(o.Provider) {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *NftprimarytransactionCreateResponse) SetProvider(v string) {
	o.Provider = &v
}

// GetSellerWalletAddress returns the SellerWalletAddress field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetSellerWalletAddress() string {
	if o == nil || isNil(o.SellerWalletAddress) {
		var ret string
		return ret
	}
	return *o.SellerWalletAddress
}

// GetSellerWalletAddressOk returns a tuple with the SellerWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetSellerWalletAddressOk() (*string, bool) {
	if o == nil || isNil(o.SellerWalletAddress) {
		return nil, false
	}
	return o.SellerWalletAddress, true
}

// HasSellerWalletAddress returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasSellerWalletAddress() bool {
	if o != nil && !isNil(o.SellerWalletAddress) {
		return true
	}

	return false
}

// SetSellerWalletAddress gets a reference to the given string and assigns it to the SellerWalletAddress field.
func (o *NftprimarytransactionCreateResponse) SetSellerWalletAddress(v string) {
	o.SellerWalletAddress = &v
}

// GetTokenId returns the TokenId field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetTokenId() string {
	if o == nil || isNil(o.TokenId) {
		var ret string
		return ret
	}
	return *o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetTokenIdOk() (*string, bool) {
	if o == nil || isNil(o.TokenId) {
		return nil, false
	}
	return o.TokenId, true
}

// HasTokenId returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasTokenId() bool {
	if o != nil && !isNil(o.TokenId) {
		return true
	}

	return false
}

// SetTokenId gets a reference to the given string and assigns it to the TokenId field.
func (o *NftprimarytransactionCreateResponse) SetTokenId(v string) {
	o.TokenId = &v
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetTransactionId() string {
	if o == nil || isNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.TransactionId) {
		return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasTransactionId() bool {
	if o != nil && !isNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *NftprimarytransactionCreateResponse) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetUrl() string {
	if o == nil || isNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetUrlOk() (*string, bool) {
	if o == nil || isNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasUrl() bool {
	if o != nil && !isNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *NftprimarytransactionCreateResponse) SetUrl(v string) {
	o.Url = &v
}

// GetUserWalletAddress returns the UserWalletAddress field value if set, zero value otherwise.
func (o *NftprimarytransactionCreateResponse) GetUserWalletAddress() string {
	if o == nil || isNil(o.UserWalletAddress) {
		var ret string
		return ret
	}
	return *o.UserWalletAddress
}

// GetUserWalletAddressOk returns a tuple with the UserWalletAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NftprimarytransactionCreateResponse) GetUserWalletAddressOk() (*string, bool) {
	if o == nil || isNil(o.UserWalletAddress) {
		return nil, false
	}
	return o.UserWalletAddress, true
}

// HasUserWalletAddress returns a boolean if a field has been set.
func (o *NftprimarytransactionCreateResponse) HasUserWalletAddress() bool {
	if o != nil && !isNil(o.UserWalletAddress) {
		return true
	}

	return false
}

// SetUserWalletAddress gets a reference to the given string and assigns it to the UserWalletAddress field.
func (o *NftprimarytransactionCreateResponse) SetUserWalletAddress(v string) {
	o.UserWalletAddress = &v
}

func (o NftprimarytransactionCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NftprimarytransactionCreateResponse) ToMap() (map[string]interface{}, error) {
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
	if !isNil(o.TokenId) {
		toSerialize["token_id"] = o.TokenId
	}
	if !isNil(o.TransactionId) {
		toSerialize["transaction_id"] = o.TransactionId
	}
	if !isNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !isNil(o.UserWalletAddress) {
		toSerialize["user_wallet_address"] = o.UserWalletAddress
	}
	return toSerialize, nil
}

type NullableNftprimarytransactionCreateResponse struct {
	value *NftprimarytransactionCreateResponse
	isSet bool
}

func (v NullableNftprimarytransactionCreateResponse) Get() *NftprimarytransactionCreateResponse {
	return v.value
}

func (v *NullableNftprimarytransactionCreateResponse) Set(val *NftprimarytransactionCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableNftprimarytransactionCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableNftprimarytransactionCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNftprimarytransactionCreateResponse(val *NftprimarytransactionCreateResponse) *NullableNftprimarytransactionCreateResponse {
	return &NullableNftprimarytransactionCreateResponse{value: val, isSet: true}
}

func (v NullableNftprimarytransactionCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNftprimarytransactionCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


