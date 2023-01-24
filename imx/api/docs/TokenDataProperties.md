# TokenDataProperties

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to [**AssetPropertiesCollection**](AssetPropertiesCollection.md) |  | [optional] 
**ImageUrl** | Pointer to **string** | Image URL of this asset | [optional] 
**Name** | Pointer to **string** | Name of this asset | [optional] 

## Methods

### NewTokenDataProperties

`func NewTokenDataProperties() *TokenDataProperties`

NewTokenDataProperties instantiates a new TokenDataProperties object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDataPropertiesWithDefaults

`func NewTokenDataPropertiesWithDefaults() *TokenDataProperties`

NewTokenDataPropertiesWithDefaults instantiates a new TokenDataProperties object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *TokenDataProperties) GetCollection() AssetPropertiesCollection`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *TokenDataProperties) GetCollectionOk() (*AssetPropertiesCollection, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *TokenDataProperties) SetCollection(v AssetPropertiesCollection)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *TokenDataProperties) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetImageUrl

`func (o *TokenDataProperties) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *TokenDataProperties) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *TokenDataProperties) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *TokenDataProperties) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetName

`func (o *TokenDataProperties) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenDataProperties) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenDataProperties) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TokenDataProperties) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


