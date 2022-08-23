# MintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSignature** | **string** | Signature from authorised minter | 
**ContractAddress** | **string** | minting contract | 
**Royalties** | Pointer to [**[]MintFee**](MintFee.md) | Global contract-level royalty fees | [optional] 
**Users** | [**[]MintUser**](MintUser.md) | Users to mint to | 

## Methods

### NewMintRequest

`func NewMintRequest(authSignature string, contractAddress string, users []MintUser, ) *MintRequest`

NewMintRequest instantiates a new MintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintRequestWithDefaults

`func NewMintRequestWithDefaults() *MintRequest`

NewMintRequestWithDefaults instantiates a new MintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSignature

`func (o *MintRequest) GetAuthSignature() string`

GetAuthSignature returns the AuthSignature field if non-nil, zero value otherwise.

### GetAuthSignatureOk

`func (o *MintRequest) GetAuthSignatureOk() (*string, bool)`

GetAuthSignatureOk returns a tuple with the AuthSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSignature

`func (o *MintRequest) SetAuthSignature(v string)`

SetAuthSignature sets AuthSignature field to given value.


### GetContractAddress

`func (o *MintRequest) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *MintRequest) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *MintRequest) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.


### GetRoyalties

`func (o *MintRequest) GetRoyalties() []MintFee`

GetRoyalties returns the Royalties field if non-nil, zero value otherwise.

### GetRoyaltiesOk

`func (o *MintRequest) GetRoyaltiesOk() (*[]MintFee, bool)`

GetRoyaltiesOk returns a tuple with the Royalties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyalties

`func (o *MintRequest) SetRoyalties(v []MintFee)`

SetRoyalties sets Royalties field to given value.

### HasRoyalties

`func (o *MintRequest) HasRoyalties() bool`

HasRoyalties returns a boolean if a field has been set.

### GetUsers

`func (o *MintRequest) GetUsers() []MintUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *MintRequest) GetUsersOk() (*[]MintUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *MintRequest) SetUsers(v []MintUser)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


