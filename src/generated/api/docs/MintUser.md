# MintUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | [**[]MintTokenDataV2**](MintTokenDataV2.md) | List of Mint tokens | 
**User** | **string** | User wallet address | 

## Methods

### NewMintUser

`func NewMintUser(tokens []MintTokenDataV2, user string, ) *MintUser`

NewMintUser instantiates a new MintUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintUserWithDefaults

`func NewMintUserWithDefaults() *MintUser`

NewMintUserWithDefaults instantiates a new MintUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokens

`func (o *MintUser) GetTokens() []MintTokenDataV2`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *MintUser) GetTokensOk() (*[]MintTokenDataV2, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *MintUser) SetTokens(v []MintTokenDataV2)`

SetTokens sets Tokens field to given value.


### GetUser

`func (o *MintUser) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *MintUser) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *MintUser) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


