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

// checks if the AssetWithOrders type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetWithOrders{}

// AssetWithOrders struct for AssetWithOrders
type AssetWithOrders struct {
	Collection CollectionDetails `json:"collection"`
	// Timestamp of when the asset was created
	CreatedAt NullableString `json:"created_at"`
	// Description of this asset
	Description NullableString `json:"description"`
	// Royalties to pay on this asset operations
	Fees []Fee `json:"fees,omitempty"`
	// [DEPRECATED] Internal Immutable X Token ID
	Id *string `json:"id,omitempty"`
	// URL of the image which should be used for this asset
	ImageUrl NullableString `json:"image_url"`
	// Metadata of this asset
	Metadata map[string]interface{} `json:"metadata"`
	// Name of this asset
	Name NullableString `json:"name"`
	Orders *OrderDetails `json:"orders,omitempty"`
	// Status of this asset (where it is in the system)
	Status string `json:"status"`
	// Address of the ERC721 contract
	TokenAddress string `json:"token_address"`
	// ERC721 Token ID of this asset
	TokenId string `json:"token_id"`
	// Timestamp of when the asset was updated
	UpdatedAt NullableString `json:"updated_at"`
	// URI to access this asset externally to Immutable X
	Uri NullableString `json:"uri"`
	// Ethereum address of the user who owns this asset
	User string `json:"user"`
}

// NewAssetWithOrders instantiates a new AssetWithOrders object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetWithOrders(collection CollectionDetails, createdAt NullableString, description NullableString, imageUrl NullableString, metadata map[string]interface{}, name NullableString, status string, tokenAddress string, tokenId string, updatedAt NullableString, uri NullableString, user string) *AssetWithOrders {
	this := AssetWithOrders{}
	this.Collection = collection
	this.CreatedAt = createdAt
	this.Description = description
	this.ImageUrl = imageUrl
	this.Metadata = metadata
	this.Name = name
	this.Status = status
	this.TokenAddress = tokenAddress
	this.TokenId = tokenId
	this.UpdatedAt = updatedAt
	this.Uri = uri
	this.User = user
	return &this
}

// NewAssetWithOrdersWithDefaults instantiates a new AssetWithOrders object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetWithOrdersWithDefaults() *AssetWithOrders {
	this := AssetWithOrders{}
	return &this
}

// GetCollection returns the Collection field value
func (o *AssetWithOrders) GetCollection() CollectionDetails {
	if o == nil {
		var ret CollectionDetails
		return ret
	}

	return o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetCollectionOk() (*CollectionDetails, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Collection, true
}

// SetCollection sets field value
func (o *AssetWithOrders) SetCollection(v CollectionDetails) {
	o.Collection = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetWithOrders) GetCreatedAt() string {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *AssetWithOrders) SetCreatedAt(v string) {
	o.CreatedAt.Set(&v)
}

// GetDescription returns the Description field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetWithOrders) GetDescription() string {
	if o == nil || o.Description.Get() == nil {
		var ret string
		return ret
	}

	return *o.Description.Get()
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Description.Get(), o.Description.IsSet()
}

// SetDescription sets field value
func (o *AssetWithOrders) SetDescription(v string) {
	o.Description.Set(&v)
}

// GetFees returns the Fees field value if set, zero value otherwise.
func (o *AssetWithOrders) GetFees() []Fee {
	if o == nil || isNil(o.Fees) {
		var ret []Fee
		return ret
	}
	return o.Fees
}

// GetFeesOk returns a tuple with the Fees field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetFeesOk() ([]Fee, bool) {
	if o == nil || isNil(o.Fees) {
		return nil, false
	}
	return o.Fees, true
}

// HasFees returns a boolean if a field has been set.
func (o *AssetWithOrders) HasFees() bool {
	if o != nil && !isNil(o.Fees) {
		return true
	}

	return false
}

// SetFees gets a reference to the given []Fee and assigns it to the Fees field.
func (o *AssetWithOrders) SetFees(v []Fee) {
	o.Fees = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetWithOrders) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetWithOrders) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetWithOrders) SetId(v string) {
	o.Id = &v
}

// GetImageUrl returns the ImageUrl field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetWithOrders) GetImageUrl() string {
	if o == nil || o.ImageUrl.Get() == nil {
		var ret string
		return ret
	}

	return *o.ImageUrl.Get()
}

// GetImageUrlOk returns a tuple with the ImageUrl field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetImageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ImageUrl.Get(), o.ImageUrl.IsSet()
}

// SetImageUrl sets field value
func (o *AssetWithOrders) SetImageUrl(v string) {
	o.ImageUrl.Set(&v)
}

// GetMetadata returns the Metadata field value
// If the value is explicit nil, the zero value for map[string]interface{} will be returned
func (o *AssetWithOrders) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *AssetWithOrders) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetWithOrders) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *AssetWithOrders) SetName(v string) {
	o.Name.Set(&v)
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *AssetWithOrders) GetOrders() OrderDetails {
	if o == nil || isNil(o.Orders) {
		var ret OrderDetails
		return ret
	}
	return *o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetOrdersOk() (*OrderDetails, bool) {
	if o == nil || isNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *AssetWithOrders) HasOrders() bool {
	if o != nil && !isNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given OrderDetails and assigns it to the Orders field.
func (o *AssetWithOrders) SetOrders(v OrderDetails) {
	o.Orders = &v
}

// GetStatus returns the Status field value
func (o *AssetWithOrders) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AssetWithOrders) SetStatus(v string) {
	o.Status = v
}

// GetTokenAddress returns the TokenAddress field value
func (o *AssetWithOrders) GetTokenAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAddress
}

// GetTokenAddressOk returns a tuple with the TokenAddress field value
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetTokenAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAddress, true
}

// SetTokenAddress sets field value
func (o *AssetWithOrders) SetTokenAddress(v string) {
	o.TokenAddress = v
}

// GetTokenId returns the TokenId field value
func (o *AssetWithOrders) GetTokenId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenId
}

// GetTokenIdOk returns a tuple with the TokenId field value
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetTokenIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenId, true
}

// SetTokenId sets field value
func (o *AssetWithOrders) SetTokenId(v string) {
	o.TokenId = v
}

// GetUpdatedAt returns the UpdatedAt field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetWithOrders) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt.Get() == nil {
		var ret string
		return ret
	}

	return *o.UpdatedAt.Get()
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpdatedAt.Get(), o.UpdatedAt.IsSet()
}

// SetUpdatedAt sets field value
func (o *AssetWithOrders) SetUpdatedAt(v string) {
	o.UpdatedAt.Set(&v)
}

// GetUri returns the Uri field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AssetWithOrders) GetUri() string {
	if o == nil || o.Uri.Get() == nil {
		var ret string
		return ret
	}

	return *o.Uri.Get()
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetWithOrders) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uri.Get(), o.Uri.IsSet()
}

// SetUri sets field value
func (o *AssetWithOrders) SetUri(v string) {
	o.Uri.Set(&v)
}

// GetUser returns the User field value
func (o *AssetWithOrders) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *AssetWithOrders) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *AssetWithOrders) SetUser(v string) {
	o.User = v
}

func (o AssetWithOrders) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetWithOrders) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["collection"] = o.Collection
	toSerialize["created_at"] = o.CreatedAt.Get()
	toSerialize["description"] = o.Description.Get()
	if !isNil(o.Fees) {
		toSerialize["fees"] = o.Fees
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["image_url"] = o.ImageUrl.Get()
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name.Get()
	if !isNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	toSerialize["status"] = o.Status
	toSerialize["token_address"] = o.TokenAddress
	toSerialize["token_id"] = o.TokenId
	toSerialize["updated_at"] = o.UpdatedAt.Get()
	toSerialize["uri"] = o.Uri.Get()
	toSerialize["user"] = o.User
	return toSerialize, nil
}

type NullableAssetWithOrders struct {
	value *AssetWithOrders
	isSet bool
}

func (v NullableAssetWithOrders) Get() *AssetWithOrders {
	return v.value
}

func (v *NullableAssetWithOrders) Set(val *AssetWithOrders) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetWithOrders) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetWithOrders) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetWithOrders(val *AssetWithOrders) *NullableAssetWithOrders {
	return &NullableAssetWithOrders{value: val, isSet: true}
}

func (v NullableAssetWithOrders) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetWithOrders) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


