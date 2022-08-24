# EncodeAssetTokenData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | Pointer to **string** | Blueprint information of the token to be encoded. Used if token is mintable. | [optional] 
**Id** | Pointer to **string** | ID of the token to be encoded. Used if token is mintable. | [optional] 
**TokenAddress** | Pointer to **string** | Contract address of the token to be encoded | [optional] 
**TokenId** | Pointer to **string** | TokenId of the token to be encoded. Used if token is non-mintable | [optional] 

## Methods

### NewEncodeAssetTokenData

`func NewEncodeAssetTokenData() *EncodeAssetTokenData`

NewEncodeAssetTokenData instantiates a new EncodeAssetTokenData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodeAssetTokenDataWithDefaults

`func NewEncodeAssetTokenDataWithDefaults() *EncodeAssetTokenData`

NewEncodeAssetTokenDataWithDefaults instantiates a new EncodeAssetTokenData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *EncodeAssetTokenData) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *EncodeAssetTokenData) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *EncodeAssetTokenData) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *EncodeAssetTokenData) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetId

`func (o *EncodeAssetTokenData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncodeAssetTokenData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncodeAssetTokenData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EncodeAssetTokenData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTokenAddress

`func (o *EncodeAssetTokenData) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *EncodeAssetTokenData) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *EncodeAssetTokenData) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *EncodeAssetTokenData) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *EncodeAssetTokenData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EncodeAssetTokenData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EncodeAssetTokenData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EncodeAssetTokenData) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


