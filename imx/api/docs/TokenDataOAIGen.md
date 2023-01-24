# TokenDataOAIGen

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimals** | Pointer to **int32** | Number of decimals supported by this asset | [optional] 
**Id** | Pointer to **string** | [DEPRECATED] Internal Immutable X Token ID | [optional] 
**Properties** | Pointer to [**TokenDataProperties**](TokenDataProperties.md) |  | [optional] 
**Quantity** | **string** | Quantity of this asset - inclusive of fees for buy order in v1 API and exclusive of fees in v3 API | 
**QuantityWithFees** | **string** | Quantity of this asset with the sum of all fees applied to the asset | 
**Symbol** | Pointer to **string** | Symbol of a token | [optional] 
**TokenAddress** | Pointer to **string** | Address of ERC721/ERC20 contract | [optional] 
**TokenId** | Pointer to **string** | ERC721 Token ID | [optional] 

## Methods

### NewTokenDataOAIGen

`func NewTokenDataOAIGen(quantity string, quantityWithFees string, ) *TokenDataOAIGen`

NewTokenDataOAIGen instantiates a new TokenDataOAIGen object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDataOAIGenWithDefaults

`func NewTokenDataOAIGenWithDefaults() *TokenDataOAIGen`

NewTokenDataOAIGenWithDefaults instantiates a new TokenDataOAIGen object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimals

`func (o *TokenDataOAIGen) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenDataOAIGen) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenDataOAIGen) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *TokenDataOAIGen) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetId

`func (o *TokenDataOAIGen) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TokenDataOAIGen) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TokenDataOAIGen) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TokenDataOAIGen) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProperties

`func (o *TokenDataOAIGen) GetProperties() TokenDataProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TokenDataOAIGen) GetPropertiesOk() (*TokenDataProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TokenDataOAIGen) SetProperties(v TokenDataProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TokenDataOAIGen) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetQuantity

`func (o *TokenDataOAIGen) GetQuantity() string`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *TokenDataOAIGen) GetQuantityOk() (*string, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *TokenDataOAIGen) SetQuantity(v string)`

SetQuantity sets Quantity field to given value.


### GetQuantityWithFees

`func (o *TokenDataOAIGen) GetQuantityWithFees() string`

GetQuantityWithFees returns the QuantityWithFees field if non-nil, zero value otherwise.

### GetQuantityWithFeesOk

`func (o *TokenDataOAIGen) GetQuantityWithFeesOk() (*string, bool)`

GetQuantityWithFeesOk returns a tuple with the QuantityWithFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityWithFees

`func (o *TokenDataOAIGen) SetQuantityWithFees(v string)`

SetQuantityWithFees sets QuantityWithFees field to given value.


### GetSymbol

`func (o *TokenDataOAIGen) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenDataOAIGen) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenDataOAIGen) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TokenDataOAIGen) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetTokenAddress

`func (o *TokenDataOAIGen) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenDataOAIGen) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenDataOAIGen) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.

### HasTokenAddress

`func (o *TokenDataOAIGen) HasTokenAddress() bool`

HasTokenAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *TokenDataOAIGen) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *TokenDataOAIGen) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *TokenDataOAIGen) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *TokenDataOAIGen) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


