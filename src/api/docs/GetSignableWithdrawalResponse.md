# GetSignableWithdrawalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the token we are withdrawing | 
**AssetId** | **string** | ID of the asset to be withdrawn | 
**Nonce** | **int32** | Nonce of this transaction | 
**PayloadHash** | **string** | Encoded payload hash | 
**SignableMessage** | **string** | Message to sign with L1 wallet to verity withdrawal request | 
**StarkKey** | **string** | Public stark key of this user | 
**VaultId** | **int32** | ID of the vault we are withdrawing from | 

## Methods

### NewGetSignableWithdrawalResponse

`func NewGetSignableWithdrawalResponse(amount string, assetId string, nonce int32, payloadHash string, signableMessage string, starkKey string, vaultId int32, ) *GetSignableWithdrawalResponse`

NewGetSignableWithdrawalResponse instantiates a new GetSignableWithdrawalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableWithdrawalResponseWithDefaults

`func NewGetSignableWithdrawalResponseWithDefaults() *GetSignableWithdrawalResponse`

NewGetSignableWithdrawalResponseWithDefaults instantiates a new GetSignableWithdrawalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSignableWithdrawalResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSignableWithdrawalResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSignableWithdrawalResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *GetSignableWithdrawalResponse) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *GetSignableWithdrawalResponse) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *GetSignableWithdrawalResponse) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNonce

`func (o *GetSignableWithdrawalResponse) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetSignableWithdrawalResponse) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetSignableWithdrawalResponse) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetPayloadHash

`func (o *GetSignableWithdrawalResponse) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableWithdrawalResponse) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableWithdrawalResponse) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.


### GetSignableMessage

`func (o *GetSignableWithdrawalResponse) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableWithdrawalResponse) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableWithdrawalResponse) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.


### GetStarkKey

`func (o *GetSignableWithdrawalResponse) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *GetSignableWithdrawalResponse) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *GetSignableWithdrawalResponse) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetVaultId

`func (o *GetSignableWithdrawalResponse) GetVaultId() int32`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *GetSignableWithdrawalResponse) GetVaultIdOk() (*int32, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *GetSignableWithdrawalResponse) SetVaultId(v int32)`

SetVaultId sets VaultId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


