# AssetWithOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | [**CollectionDetails**](CollectionDetails.md) |  | 
**CreatedAt** | **NullableString** | Timestamp of when the asset was created | 
**Description** | **NullableString** | Description of this asset | 
**Fees** | Pointer to [**[]Fee**](Fee.md) | Royalties to pay on this asset operations | [optional] 
**Id** | Pointer to **string** | [DEPRECATED] Internal Immutable X Token ID | [optional] 
**ImageUrl** | **NullableString** | URL of the image which should be used for this asset | 
**Metadata** | **map[string]interface{}** | Metadata of this asset | 
**Name** | **NullableString** | Name of this asset | 
**Orders** | Pointer to [**OrderDetails**](OrderDetails.md) |  | [optional] 
**Status** | **string** | Status of this asset (where it is in the system) | 
**TokenAddress** | **string** | Address of the ERC721 contract | 
**TokenId** | **string** | ERC721 Token ID of this asset | 
**UpdatedAt** | **NullableString** | Timestamp of when the asset was updated | 
**Uri** | **NullableString** | URI to access this asset externally to Immutable X | 
**User** | **string** | Ethereum address of the user who owns this asset | 

## Methods

### NewAssetWithOrders

`func NewAssetWithOrders(collection CollectionDetails, createdAt NullableString, description NullableString, imageUrl NullableString, metadata map[string]interface{}, name NullableString, status string, tokenAddress string, tokenId string, updatedAt NullableString, uri NullableString, user string, ) *AssetWithOrders`

NewAssetWithOrders instantiates a new AssetWithOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithOrdersWithDefaults

`func NewAssetWithOrdersWithDefaults() *AssetWithOrders`

NewAssetWithOrdersWithDefaults instantiates a new AssetWithOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *AssetWithOrders) GetCollection() CollectionDetails`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *AssetWithOrders) GetCollectionOk() (*CollectionDetails, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *AssetWithOrders) SetCollection(v CollectionDetails)`

SetCollection sets Collection field to given value.


### GetCreatedAt

`func (o *AssetWithOrders) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssetWithOrders) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssetWithOrders) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *AssetWithOrders) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *AssetWithOrders) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *AssetWithOrders) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetWithOrders) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetWithOrders) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *AssetWithOrders) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AssetWithOrders) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFees

`func (o *AssetWithOrders) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *AssetWithOrders) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *AssetWithOrders) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *AssetWithOrders) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetId

`func (o *AssetWithOrders) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AssetWithOrders) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AssetWithOrders) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AssetWithOrders) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageUrl

`func (o *AssetWithOrders) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *AssetWithOrders) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *AssetWithOrders) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### SetImageUrlNil

`func (o *AssetWithOrders) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *AssetWithOrders) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetMetadata

`func (o *AssetWithOrders) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AssetWithOrders) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AssetWithOrders) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.


### SetMetadataNil

`func (o *AssetWithOrders) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *AssetWithOrders) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetName

`func (o *AssetWithOrders) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetWithOrders) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetWithOrders) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *AssetWithOrders) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *AssetWithOrders) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrders

`func (o *AssetWithOrders) GetOrders() OrderDetails`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *AssetWithOrders) GetOrdersOk() (*OrderDetails, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *AssetWithOrders) SetOrders(v OrderDetails)`

SetOrders sets Orders field to given value.

### HasOrders

`func (o *AssetWithOrders) HasOrders() bool`

HasOrders returns a boolean if a field has been set.

### GetStatus

`func (o *AssetWithOrders) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssetWithOrders) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssetWithOrders) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTokenAddress

`func (o *AssetWithOrders) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *AssetWithOrders) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *AssetWithOrders) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.


### GetTokenId

`func (o *AssetWithOrders) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *AssetWithOrders) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *AssetWithOrders) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetUpdatedAt

`func (o *AssetWithOrders) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssetWithOrders) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssetWithOrders) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *AssetWithOrders) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *AssetWithOrders) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUri

`func (o *AssetWithOrders) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *AssetWithOrders) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *AssetWithOrders) SetUri(v string)`

SetUri sets Uri field to given value.


### SetUriNil

`func (o *AssetWithOrders) SetUriNil(b bool)`

 SetUriNil sets the value for Uri to be an explicit nil

### UnsetUri
`func (o *AssetWithOrders) UnsetUri()`

UnsetUri ensures that no value is present for Uri, not even an explicit nil
### GetUser

`func (o *AssetWithOrders) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AssetWithOrders) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AssetWithOrders) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


