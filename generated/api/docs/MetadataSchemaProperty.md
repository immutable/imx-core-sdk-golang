# MetadataSchemaProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filterable** | **bool** | Sets the metadata as filterable | 
**Name** | **string** | Name of the metadata key | 
**Type** | **string** | Type of the metadata. Values: \&quot;enum\&quot;, \&quot;text\&quot;, \&quot;boolean\&quot;, \&quot;continuous\&quot;, \&quot;discrete\&quot; | Default: \&quot;text\&quot;. Src: https://docs.x.immutable.com/docs/asset-metadata#property-type-mapping | 

## Methods

### NewMetadataSchemaProperty

`func NewMetadataSchemaProperty(filterable bool, name string, type_ string, ) *MetadataSchemaProperty`

NewMetadataSchemaProperty instantiates a new MetadataSchemaProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataSchemaPropertyWithDefaults

`func NewMetadataSchemaPropertyWithDefaults() *MetadataSchemaProperty`

NewMetadataSchemaPropertyWithDefaults instantiates a new MetadataSchemaProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterable

`func (o *MetadataSchemaProperty) GetFilterable() bool`

GetFilterable returns the Filterable field if non-nil, zero value otherwise.

### GetFilterableOk

`func (o *MetadataSchemaProperty) GetFilterableOk() (*bool, bool)`

GetFilterableOk returns a tuple with the Filterable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterable

`func (o *MetadataSchemaProperty) SetFilterable(v bool)`

SetFilterable sets Filterable field to given value.


### GetName

`func (o *MetadataSchemaProperty) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetadataSchemaProperty) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetadataSchemaProperty) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *MetadataSchemaProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataSchemaProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataSchemaProperty) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


