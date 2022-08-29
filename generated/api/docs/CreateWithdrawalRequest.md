# CreateWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount to withdraw | 
**AssetId** | **string** | The ID of asset the user is withdrawing | 
**Nonce** | **int32** | Nonce of the withdrawal | 
**StarkKey** | **string** | Public stark key of the withdrawing user | 
**StarkSignature** | **string** | Payload signature | 
**VaultId** | **int32** | The ID of the vault the asset belong to | 

## Methods

### NewCreateWithdrawalRequest

`func NewCreateWithdrawalRequest(amount string, assetId string, nonce int32, starkKey string, starkSignature string, vaultId int32, ) *CreateWithdrawalRequest`

NewCreateWithdrawalRequest instantiates a new CreateWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWithdrawalRequestWithDefaults

`func NewCreateWithdrawalRequestWithDefaults() *CreateWithdrawalRequest`

NewCreateWithdrawalRequestWithDefaults instantiates a new CreateWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateWithdrawalRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateWithdrawalRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateWithdrawalRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *CreateWithdrawalRequest) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *CreateWithdrawalRequest) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *CreateWithdrawalRequest) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNonce

`func (o *CreateWithdrawalRequest) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateWithdrawalRequest) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateWithdrawalRequest) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetStarkKey

`func (o *CreateWithdrawalRequest) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *CreateWithdrawalRequest) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *CreateWithdrawalRequest) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetStarkSignature

`func (o *CreateWithdrawalRequest) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *CreateWithdrawalRequest) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *CreateWithdrawalRequest) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.


### GetVaultId

`func (o *CreateWithdrawalRequest) GetVaultId() int32`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *CreateWithdrawalRequest) GetVaultIdOk() (*int32, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *CreateWithdrawalRequest) SetVaultId(v int32)`

SetVaultId sets VaultId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


