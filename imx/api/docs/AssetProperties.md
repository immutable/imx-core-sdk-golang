# AssetProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**AssetPropertiesCollection**](AssetPropertiesCollection.md) |  | [optional] 
**ImageUrl** | Pointer to **string** | Image URL of this asset | [optional] 
**Name** | Pointer to **string** | Name of this asset | [optional] 

## Methods

### NewAssetProperties

`func NewAssetProperties() *AssetProperties`

NewAssetProperties instantiates a new AssetProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetPropertiesWithDefaults

`func NewAssetPropertiesWithDefaults() *AssetProperties`

NewAssetPropertiesWithDefaults instantiates a new AssetProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *AssetProperties) GetCollection() AssetPropertiesCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *AssetProperties) GetCollectionOk() (*AssetPropertiesCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *AssetProperties) SetCollection(v AssetPropertiesCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *AssetProperties) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetImageUrl

`func (o *AssetProperties) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *AssetProperties) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *AssetProperties) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *AssetProperties) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetName

`func (o *AssetProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetProperties) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


