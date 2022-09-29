# CreateMetadataRefreshRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The collection contract address | 
**TokenIds** | **[]string** | The tokens to refresh | 

## Methods

### NewCreateMetadataRefreshRequest

`func NewCreateMetadataRefreshRequest(collectionAddress string, tokenIds []string, ) *CreateMetadataRefreshRequest`

NewCreateMetadataRefreshRequest instantiates a new CreateMetadataRefreshRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMetadataRefreshRequestWithDefaults

`func NewCreateMetadataRefreshRequestWithDefaults() *CreateMetadataRefreshRequest`

NewCreateMetadataRefreshRequestWithDefaults instantiates a new CreateMetadataRefreshRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *CreateMetadataRefreshRequest) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *CreateMetadataRefreshRequest) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *CreateMetadataRefreshRequest) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetTokenIds

`func (o *CreateMetadataRefreshRequest) GetTokenIds() []string`

GetTokenIds returns the TokenIds field if non-nil, zero value otherwise.

### GetTokenIdsOk

`func (o *CreateMetadataRefreshRequest) GetTokenIdsOk() (*[]string, bool)`

GetTokenIdsOk returns a tuple with the TokenIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenIds

`func (o *CreateMetadataRefreshRequest) SetTokenIds(v []string)`

SetTokenIds sets TokenIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


