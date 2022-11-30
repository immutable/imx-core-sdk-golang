# CreateCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionImageUrl** | Pointer to **string** | URL of the tile image for this collection | [optional] 
**ContractAddress** | **string** | Ethereum address of the ERC721 contract | 
**Description** | Pointer to **string** | Description of the collection | [optional] 
**IconUrl** | Pointer to **string** | URL of the icon for this collection | [optional] 
**MetadataApiUrl** | Pointer to **string** | URL of the metadata for this collection | [optional] 
**Name** | **string** | Name of the collection | 
**OwnerPublicKey** | **string** | Owner Public Key: The uncompressed public key of the owner of the contract | 
**ProjectId** | **int32** | The collection&#39;s project ID | 

## Methods

### NewCreateCollectionRequest

`func NewCreateCollectionRequest(contractAddress string, name string, ownerPublicKey string, projectId int32, ) *CreateCollectionRequest`

NewCreateCollectionRequest instantiates a new CreateCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCollectionRequestWithDefaults

`func NewCreateCollectionRequestWithDefaults() *CreateCollectionRequest`

NewCreateCollectionRequestWithDefaults instantiates a new CreateCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionImageUrl

`func (o *CreateCollectionRequest) GetCollectionImageUrl() string`

GetCollectionImageUrl returns the CollectionImageUrl field if non-nil, zero value otherwise.

### GetCollectionImageUrlOk

`func (o *CreateCollectionRequest) GetCollectionImageUrlOk() (*string, bool)`

GetCollectionImageUrlOk returns a tuple with the CollectionImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionImageUrl

`func (o *CreateCollectionRequest) SetCollectionImageUrl(v string)`

SetCollectionImageUrl sets CollectionImageUrl field to given value.

### HasCollectionImageUrl

`func (o *CreateCollectionRequest) HasCollectionImageUrl() bool`

HasCollectionImageUrl returns a boolean if a field has been set.

### GetContractAddress

`func (o *CreateCollectionRequest) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *CreateCollectionRequest) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *CreateCollectionRequest) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetDescription

`func (o *CreateCollectionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateCollectionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateCollectionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateCollectionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIconUrl

`func (o *CreateCollectionRequest) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *CreateCollectionRequest) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *CreateCollectionRequest) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.

### HasIconUrl

`func (o *CreateCollectionRequest) HasIconUrl() bool`

HasIconUrl returns a boolean if a field has been set.

### GetMetadataApiUrl

`func (o *CreateCollectionRequest) GetMetadataApiUrl() string`

GetMetadataApiUrl returns the MetadataApiUrl field if non-nil, zero value otherwise.

### GetMetadataApiUrlOk

`func (o *CreateCollectionRequest) GetMetadataApiUrlOk() (*string, bool)`

GetMetadataApiUrlOk returns a tuple with the MetadataApiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataApiUrl

`func (o *CreateCollectionRequest) SetMetadataApiUrl(v string)`

SetMetadataApiUrl sets MetadataApiUrl field to given value.

### HasMetadataApiUrl

`func (o *CreateCollectionRequest) HasMetadataApiUrl() bool`

HasMetadataApiUrl returns a boolean if a field has been set.

### GetName

`func (o *CreateCollectionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCollectionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCollectionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOwnerPublicKey

`func (o *CreateCollectionRequest) GetOwnerPublicKey() string`

GetOwnerPublicKey returns the OwnerPublicKey field if non-nil, zero value otherwise.

### GetOwnerPublicKeyOk

`func (o *CreateCollectionRequest) GetOwnerPublicKeyOk() (*string, bool)`

GetOwnerPublicKeyOk returns a tuple with the OwnerPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerPublicKey

`func (o *CreateCollectionRequest) SetOwnerPublicKey(v string)`

SetOwnerPublicKey sets OwnerPublicKey field to given value.


### GetProjectId

`func (o *CreateCollectionRequest) GetProjectId() int32`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateCollectionRequest) GetProjectIdOk() (*int32, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateCollectionRequest) SetProjectId(v int32)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


