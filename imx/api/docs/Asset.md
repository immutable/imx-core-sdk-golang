# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | [**AssetCollection**](AssetCollection.md) |  | 
**CreatedAt** | **NullableString** | Timestamp of when the asset was created | 
**Description** | **NullableString** | Description of this asset | 
**Fees** | Pointer to [**[]Fee**](Fee.md) | Royalties to pay on this asset operations | [optional] 
**Id** | Pointer to **string** | [DEPRECATED] Internal Immutable X Token ID | [optional] 
**ImageUrl** | **NullableString** | URL of the image which should be used for this asset | 
**Metadata** | **map[string]interface{}** | Metadata of this asset | 
**Name** | **NullableString** | Name of this asset | 
**Orders** | Pointer to [**AssetOrders**](AssetOrders.md) |  | [optional] 
**Status** | **string** | Status of this asset (where it is in the system) | 
**TokenAddress** | **string** | Address of the ERC721 contract | 
**TokenId** | **string** | ERC721 Token ID of this asset | 
**UpdatedAt** | **NullableString** | Timestamp of when the asset was updated | 
**Uri** | **NullableString** | URI to access this asset externally to Immutable X | 
**User** | **string** | Ethereum address of the user who owns this asset | 

## Methods

### NewAsset

`func NewAsset(collection AssetCollection, createdAt NullableString, description NullableString, imageUrl NullableString, metadata map[string]interface{}, name NullableString, status string, tokenAddress string, tokenId string, updatedAt NullableString, uri NullableString, user string, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *Asset) GetCollection() AssetCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *Asset) GetCollectionOk() (*AssetCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *Asset) SetCollection(v AssetCollection)`

SetCollection sets Collection field to given value.


### GetCreatedAt

`func (o *Asset) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Asset) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Asset) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Asset) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Asset) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Asset) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Asset) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFees

`func (o *Asset) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Asset) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Asset) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Asset) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Asset) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *Asset) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *Asset) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *Asset) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### SetImageUrlNil

`func (o *Asset) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *Asset) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetMetadata

`func (o *Asset) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Asset) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Asset) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *Asset) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *Asset) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Asset) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Asset) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrders

`func (o *Asset) GetOrders() AssetOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *Asset) GetOrdersOk() (*AssetOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *Asset) SetOrders(v AssetOrders)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *Asset) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetStatus

`func (o *Asset) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Asset) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Asset) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenAddress

`func (o *Asset) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *Asset) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *Asset) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetTokenId

`func (o *Asset) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *Asset) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *Asset) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetUpdatedAt

`func (o *Asset) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Asset) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Asset) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Asset) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Asset) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUri

`func (o *Asset) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Asset) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Asset) SetUri(v string)`

SetUri sets Uri field to given value.


### SetUriNil

`func (o *Asset) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *Asset) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetUser

`func (o *Asset) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Asset) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Asset) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


