# TokenDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decimals** | **string** | Number of decimals for token | 
**ImageUrl** | **string** | Url for the icon of the token | 
**Name** | **string** | Full name of the token (e.g. Ether) | 
**Quantum** | **string** | Quantum for token | 
**Symbol** | **string** | Ticker symbol for token (e.g. ETH/USDC/IMX) | 
**TokenAddress** | **string** | Address of the ERC721 contract | 

## Methods

### NewTokenDetails

`func NewTokenDetails(decimals string, imageUrl string, name string, quantum string, symbol string, tokenAddress string, ) *TokenDetails`

NewTokenDetails instantiates a new TokenDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenDetailsWithDefaults

`func NewTokenDetailsWithDefaults() *TokenDetails`

NewTokenDetailsWithDefaults instantiates a new TokenDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecimals

`func (o *TokenDetails) GetDecimals() string`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *TokenDetails) GetDecimalsOk() (*string, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *TokenDetails) SetDecimals(v string)`

SetDecimals sets Decimals field to given value.


### GetImageUrl

`func (o *TokenDetails) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *TokenDetails) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *TokenDetails) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetName

`func (o *TokenDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenDetails) SetName(v string)`

SetName sets Name field to given value.


### GetQuantum

`func (o *TokenDetails) GetQuantum() string`

GetQuantum returns the Quantum field if non-nil, zero value otherwise.

### GetQuantumOk

`func (o *TokenDetails) GetQuantumOk() (*string, bool)`

GetQuantumOk returns a tuple with the Quantum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantum

`func (o *TokenDetails) SetQuantum(v string)`

SetQuantum sets Quantum field to given value.


### GetSymbol

`func (o *TokenDetails) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TokenDetails) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TokenDetails) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.


### GetTokenAddress

`func (o *TokenDetails) GetTokenAddress() string`

GetTokenAddress returns the TokenAddress field if non-nil, zero value otherwise.

### GetTokenAddressOk

`func (o *TokenDetails) GetTokenAddressOk() (*string, bool)`

GetTokenAddressOk returns a tuple with the TokenAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenAddress

`func (o *TokenDetails) SetTokenAddress(v string)`

SetTokenAddress sets TokenAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


