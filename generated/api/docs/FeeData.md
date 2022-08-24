# FeeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | Pointer to **string** | Address of ERC721/ERC20 contract | [optional] 
**Decimals** | Pointer to **int32** | Number of decimals supported by this asset | [optional] 

## Methods

### NewFeeData

`func NewFeeData() *FeeData`

NewFeeData instantiates a new FeeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeDataWithDefaults

`func NewFeeDataWithDefaults() *FeeData`

NewFeeDataWithDefaults instantiates a new FeeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *FeeData) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *FeeData) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *FeeData) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *FeeData) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetDecimals

`func (o *FeeData) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *FeeData) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *FeeData) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *FeeData) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


