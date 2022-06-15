# MintTokenDataV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | Pointer to **string** | Token metadata blueprint | [optional] 
**Id** | **string** | Token ID | 
**Royalties** | Pointer to [**[]MintFee**](MintFee.md) | List of mint fees | [optional] 

## Methods

### NewMintTokenDataV2

`func NewMintTokenDataV2(id string, ) *MintTokenDataV2`

NewMintTokenDataV2 instantiates a new MintTokenDataV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintTokenDataV2WithDefaults

`func NewMintTokenDataV2WithDefaults() *MintTokenDataV2`

NewMintTokenDataV2WithDefaults instantiates a new MintTokenDataV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *MintTokenDataV2) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *MintTokenDataV2) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *MintTokenDataV2) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *MintTokenDataV2) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetId

`func (o *MintTokenDataV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MintTokenDataV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MintTokenDataV2) SetId(v string)`

SetId sets Id field to given value.


### GetRoyalties

`func (o *MintTokenDataV2) GetRoyalties() []MintFee`

GetRoyalties returns the Royalties field if non-nil, zero value otherwise.

### GetRoyaltiesOk

`func (o *MintTokenDataV2) GetRoyaltiesOk() (*[]MintFee, bool)`

GetRoyaltiesOk returns a tuple with the Royalties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyalties

`func (o *MintTokenDataV2) SetRoyalties(v []MintFee)`

SetRoyalties sets Royalties field to given value.

### HasRoyalties

`func (o *MintTokenDataV2) HasRoyalties() bool`

HasRoyalties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


