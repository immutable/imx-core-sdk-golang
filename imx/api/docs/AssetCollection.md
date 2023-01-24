# AssetCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IconUrl** | **NullableString** | URL of the icon of the collection | 
**Name** | **string** | Name of the collection | 

## Methods

### NewAssetCollection

`func NewAssetCollection(iconUrl NullableString, name string, ) *AssetCollection`

NewAssetCollection instantiates a new AssetCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetCollectionWithDefaults

`func NewAssetCollectionWithDefaults() *AssetCollection`

NewAssetCollectionWithDefaults instantiates a new AssetCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIconUrl

`func (o *AssetCollection) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *AssetCollection) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *AssetCollection) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### SetIconUrlNil

`func (o *AssetCollection) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *AssetCollection) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetName

`func (o *AssetCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetCollection) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


