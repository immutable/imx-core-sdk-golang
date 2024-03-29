# CollectionFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | Key of this property | [optional] 
**Range** | Pointer to [**Range**](Range.md) |  | [optional] 
**Type** | Pointer to **string** | Type of this filter | [optional] 
**Value** | Pointer to **[]string** | List of possible values for this property | [optional] 

## Methods

### NewCollectionFilter

`func NewCollectionFilter() *CollectionFilter`

NewCollectionFilter instantiates a new CollectionFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionFilterWithDefaults

`func NewCollectionFilterWithDefaults() *CollectionFilter`

NewCollectionFilterWithDefaults instantiates a new CollectionFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CollectionFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CollectionFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CollectionFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CollectionFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetRange

`func (o *CollectionFilter) GetRange() Range`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *CollectionFilter) GetRangeOk() (*Range, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *CollectionFilter) SetRange(v Range)`

SetRange sets Range field to given value.

### HasRange

`func (o *CollectionFilter) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetType

`func (o *CollectionFilter) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CollectionFilter) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CollectionFilter) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CollectionFilter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *CollectionFilter) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CollectionFilter) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CollectionFilter) SetValue(v []string)`

SetValue sets Value field to given value.

### HasValue

`func (o *CollectionFilter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


