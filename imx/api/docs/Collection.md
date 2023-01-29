# Collection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Ethereum address of the ERC721 contract | 
**CollectionImageUrl** | **NullableString** | URL of the tile image for this collection | 
**CreatedAt** | **NullableString** | Timestamp of when the collection was created | 
**Description** | **NullableString** | Description of the collection | 
**IconUrl** | **NullableString** | URL of the icon for this collection | 
**MetadataApiUrl** | **NullableString** | URL of the metadata for this collection | 
**Name** | **string** | Name of the collection | 
**ProjectId** | **int32** | The collection&#39;s project ID | 
**ProjectOwnerAddress** | **string** | Project owner address | 
**UpdatedAt** | **NullableString** | Timestamp of when the collection was updated | 

## Methods

### NewCollection

`func NewCollection(address string, collectionImageUrl NullableString, createdAt NullableString, description NullableString, iconUrl NullableString, metadataApiUrl NullableString, name string, projectId int32, projectOwnerAddress string, updatedAt NullableString, ) *Collection`

NewCollection instantiates a new Collection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionWithDefaults

`func NewCollectionWithDefaults() *Collection`

NewCollectionWithDefaults instantiates a new Collection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Collection) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Collection) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Collection) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetCollectionImageUrl

`func (o *Collection) GetCollectionImageUrl() string`

GetCollectionImageUrl returns the CollectionImageUrl field if non-nil, zero value otherwise.

### GetCollectionImageUrlOk

`func (o *Collection) GetCollectionImageUrlOk() (*string, bool)`

GetCollectionImageUrlOk returns a tuple with the CollectionImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionImageUrl

`func (o *Collection) SetCollectionImageUrl(v string)`

SetCollectionImageUrl sets CollectionImageUrl field to given value.


### SetCollectionImageUrlNil

`func (o *Collection) SetCollectionImageUrlNil(b bool)`

 SetCollectionImageUrlNil sets the value for CollectionImageUrl to be an explicit nil

### UnsetCollectionImageUrl
`func (o *Collection) UnsetCollectionImageUrl()`

UnsetCollectionImageUrl ensures that no value is present for CollectionImageUrl, not even an explicit nil
### GetCreatedAt

`func (o *Collection) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Collection) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Collection) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *Collection) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *Collection) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetDescription

`func (o *Collection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Collection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Collection) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *Collection) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *Collection) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *Collection) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *Collection) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *Collection) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### SetIconUrlNil

`func (o *Collection) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *Collection) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetMetadataApiUrl

`func (o *Collection) GetMetadataApiUrl() string`

GetMetadataApiUrl returns the MetadataApiUrl field if non-nil, zero value otherwise.

### GetMetadataApiUrlOk

`func (o *Collection) GetMetadataApiUrlOk() (*string, bool)`

GetMetadataApiUrlOk returns a tuple with the MetadataApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataApiUrl

`func (o *Collection) SetMetadataApiUrl(v string)`

SetMetadataApiUrl sets MetadataApiUrl field to given value.


### SetMetadataApiUrlNil

`func (o *Collection) SetMetadataApiUrlNil(b bool)`

 SetMetadataApiUrlNil sets the value for MetadataApiUrl to be an explicit nil

### UnsetMetadataApiUrl
`func (o *Collection) UnsetMetadataApiUrl()`

UnsetMetadataApiUrl ensures that no value is present for MetadataApiUrl, not even an explicit nil
### GetName

`func (o *Collection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Collection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Collection) SetName(v string)`

SetName sets Name field to given value.


### GetProjectId

`func (o *Collection) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Collection) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Collection) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.


### GetProjectOwnerAddress

`func (o *Collection) GetProjectOwnerAddress() string`

GetProjectOwnerAddress returns the ProjectOwnerAddress field if non-nil, zero value otherwise.

### GetProjectOwnerAddressOk

`func (o *Collection) GetProjectOwnerAddressOk() (*string, bool)`

GetProjectOwnerAddressOk returns a tuple with the ProjectOwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOwnerAddress

`func (o *Collection) SetProjectOwnerAddress(v string)`

SetProjectOwnerAddress sets ProjectOwnerAddress field to given value.


### GetUpdatedAt

`func (o *Collection) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Collection) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Collection) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### SetUpdatedAtNil

`func (o *Collection) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *Collection) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


