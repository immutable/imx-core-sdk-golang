# TokenData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimals** | Pointer to **int32** | Number of decimals supported by this asset | [optional] 
**Id** | Pointer to **string** | [DEPRECATED] Internal Immutable X Token ID | [optional] 
**Properties** | Pointer to [**AssetProperties**](AssetProperties.md) |  | [optional] 
**Quantity** | **string** | Quantity of this asset | 
**QuantityWithFees** | **string** | Quantity of this asset with the sum of all fees applied to the asset | 
**TokenAddress** | Pointer to **string** | Address of ERC721/ERC20 contract | [optional] 
**TokenId** | Pointer to **string** | ERC721 Token ID | [optional] 

## Methods

### NewTokenData

`func NewTokenData(quantity string, quantityWithFees string, ) *TokenData`

NewTokenData instantiates a new TokenData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDataWithDefaults

`func NewTokenDataWithDefaults() *TokenData`

NewTokenDataWithDefaults instantiates a new TokenData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimals

`func (o *TokenData) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenData) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenData) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenData) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetId

`func (o *TokenData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *TokenData) GetProperties() AssetProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenData) GetPropertiesOk() (*AssetProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenData) SetProperties(v AssetProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenData) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuantity

`func (o *TokenData) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TokenData) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TokenData) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetQuantityWithFees

`func (o *TokenData) GetQuantityWithFees() string`

GetQuantityWithFees returns the QuantityWithFees field if non-nil, zero value otherwise.

### GetQuantityWithFeesOk

`func (o *TokenData) GetQuantityWithFeesOk() (*string, bool)`

GetQuantityWithFeesOk returns a tuple with the QuantityWithFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWithFees

`func (o *TokenData) SetQuantityWithFees(v string)`

SetQuantityWithFees sets QuantityWithFees field to given value.


### GetTokenAddress

`func (o *TokenData) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenData) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenData) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TokenData) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *TokenData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenData) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


