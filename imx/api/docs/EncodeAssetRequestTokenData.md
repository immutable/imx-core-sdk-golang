# EncodeAssetRequestTokenData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | Pointer to **string** | Blueprint information of the token to be encoded. Used if token is mintable. | [optional] 
**Id** | Pointer to **string** | ID of the token to be encoded. Used if token is mintable. | [optional] 
**TokenAddress** | Pointer to **string** | Contract address of the token to be encoded | [optional] 
**TokenId** | Pointer to **string** | TokenId of the token to be encoded. Used if token is non-mintable | [optional] 

## Methods

### NewEncodeAssetRequestTokenData

`func NewEncodeAssetRequestTokenData() *EncodeAssetRequestTokenData`

NewEncodeAssetRequestTokenData instantiates a new EncodeAssetRequestTokenData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEncodeAssetRequestTokenDataWithDefaults

`func NewEncodeAssetRequestTokenDataWithDefaults() *EncodeAssetRequestTokenData`

NewEncodeAssetRequestTokenDataWithDefaults instantiates a new EncodeAssetRequestTokenData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *EncodeAssetRequestTokenData) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *EncodeAssetRequestTokenData) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *EncodeAssetRequestTokenData) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.

### HasBlueprint

`func (o *EncodeAssetRequestTokenData) HasBlueprint() bool`

HasBlueprint returns a boolean if a field has been set.

### GetId

`func (o *EncodeAssetRequestTokenData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EncodeAssetRequestTokenData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EncodeAssetRequestTokenData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EncodeAssetRequestTokenData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTokenAddress

`func (o *EncodeAssetRequestTokenData) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *EncodeAssetRequestTokenData) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *EncodeAssetRequestTokenData) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *EncodeAssetRequestTokenData) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *EncodeAssetRequestTokenData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *EncodeAssetRequestTokenData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *EncodeAssetRequestTokenData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *EncodeAssetRequestTokenData) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


