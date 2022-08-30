# MintResultDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | **string** | Contract address of this token | 
**TokenId** | **string** | IMX ID of this token | 
**TxId** | **int32** | Mint Transaction ID | 

## Methods

### NewMintResultDetails

`func NewMintResultDetails(contractAddress string, tokenId string, txId int32, ) *MintResultDetails`

NewMintResultDetails instantiates a new MintResultDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintResultDetailsWithDefaults

`func NewMintResultDetailsWithDefaults() *MintResultDetails`

NewMintResultDetailsWithDefaults instantiates a new MintResultDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *MintResultDetails) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *MintResultDetails) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *MintResultDetails) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetTokenId

`func (o *MintResultDetails) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MintResultDetails) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MintResultDetails) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.


### GetTxId

`func (o *MintResultDetails) GetTxId() int32`

GetTxId returns the TxId field if non-nil, zero value otherwise.

### GetTxIdOk

`func (o *MintResultDetails) GetTxIdOk() (*int32, bool)`

GetTxIdOk returns a tuple with the TxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxId

`func (o *MintResultDetails) SetTxId(v int32)`

SetTxId sets TxId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


