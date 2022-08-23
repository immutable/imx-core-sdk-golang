# GetSignableDepositResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount this user is depositing | 
**AssetId** | **string** | ID of the asset this user is depositing | 
**Nonce** | **int32** | Nonce of the deposit | 
**StarkKey** | **string** | Public stark key of the depositing user | 
**VaultId** | **int32** | ID of the vault this user is depositing to | 

## Methods

### NewGetSignableDepositResponse

`func NewGetSignableDepositResponse(amount string, assetId string, nonce int32, starkKey string, vaultId int32, ) *GetSignableDepositResponse`

NewGetSignableDepositResponse instantiates a new GetSignableDepositResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableDepositResponseWithDefaults

`func NewGetSignableDepositResponseWithDefaults() *GetSignableDepositResponse`

NewGetSignableDepositResponseWithDefaults instantiates a new GetSignableDepositResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSignableDepositResponse) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSignableDepositResponse) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSignableDepositResponse) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *GetSignableDepositResponse) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *GetSignableDepositResponse) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *GetSignableDepositResponse) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetNonce

`func (o *GetSignableDepositResponse) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetSignableDepositResponse) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetSignableDepositResponse) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetStarkKey

`func (o *GetSignableDepositResponse) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *GetSignableDepositResponse) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *GetSignableDepositResponse) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetVaultId

`func (o *GetSignableDepositResponse) GetVaultId() int32`

GetVaultId returns the VaultId field if non-nil, zero value otherwise.

### GetVaultIdOk

`func (o *GetSignableDepositResponse) GetVaultIdOk() (*int32, bool)`

GetVaultIdOk returns a tuple with the VaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultId

`func (o *GetSignableDepositResponse) SetVaultId(v int32)`

SetVaultId sets VaultId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


