# AddMetadataSchemaToCollectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | Pointer to **string** | Not required from API user | [optional] 
**Metadata** | [**[]MetadataSchemaRequest**](MetadataSchemaRequest.md) | The metadata container | 

## Methods

### NewAddMetadataSchemaToCollectionRequest

`func NewAddMetadataSchemaToCollectionRequest(metadata []MetadataSchemaRequest, ) *AddMetadataSchemaToCollectionRequest`

NewAddMetadataSchemaToCollectionRequest instantiates a new AddMetadataSchemaToCollectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMetadataSchemaToCollectionRequestWithDefaults

`func NewAddMetadataSchemaToCollectionRequestWithDefaults() *AddMetadataSchemaToCollectionRequest`

NewAddMetadataSchemaToCollectionRequestWithDefaults instantiates a new AddMetadataSchemaToCollectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *AddMetadataSchemaToCollectionRequest) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *AddMetadataSchemaToCollectionRequest) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *AddMetadataSchemaToCollectionRequest) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *AddMetadataSchemaToCollectionRequest) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetMetadata

`func (o *AddMetadataSchemaToCollectionRequest) GetMetadata() []MetadataSchemaRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddMetadataSchemaToCollectionRequest) GetMetadataOk() (*[]MetadataSchemaRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddMetadataSchemaToCollectionRequest) SetMetadata(v []MetadataSchemaRequest)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


