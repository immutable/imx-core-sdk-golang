# MetadataSchemaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filterable** | Pointer to **bool** | Sets the metadata as filterable | [optional] 
**Name** | **string** | Name of the metadata key | 
**Type** | Pointer to **string** | Type of the metadata. Values: \&quot;enum\&quot;, \&quot;text\&quot;, \&quot;boolean\&quot;, \&quot;continuous\&quot;, \&quot;discrete\&quot; | Default: \&quot;text\&quot;. Src: https://docs.x.github.com/immutable/docs/asset-metadata#property-type-mapping | [optional] 

## Methods

### NewMetadataSchemaRequest

`func NewMetadataSchemaRequest(name string, ) *MetadataSchemaRequest`

NewMetadataSchemaRequest instantiates a new MetadataSchemaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSchemaRequestWithDefaults

`func NewMetadataSchemaRequestWithDefaults() *MetadataSchemaRequest`

NewMetadataSchemaRequestWithDefaults instantiates a new MetadataSchemaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterable

`func (o *MetadataSchemaRequest) GetFilterable() bool`

GetFilterable returns the Filterable field if non-nil, zero value otherwise.

### GetFilterableOk

`func (o *MetadataSchemaRequest) GetFilterableOk() (*bool, bool)`

GetFilterableOk returns a tuple with the Filterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterable

`func (o *MetadataSchemaRequest) SetFilterable(v bool)`

SetFilterable sets Filterable field to given value.

### HasFilterable

`func (o *MetadataSchemaRequest) HasFilterable() bool`

HasFilterable returns a boolean if a field has been set.

### GetName

`func (o *MetadataSchemaRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataSchemaRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataSchemaRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MetadataSchemaRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataSchemaRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataSchemaRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetadataSchemaRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


