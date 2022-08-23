# CreateOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBuy** | **string** | Amount to buy | 
**AmountSell** | **string** | Amount to sell | 
**AssetIdBuy** | **string** | ID of the asset to buy | 
**AssetIdSell** | **string** | ID of the asset to sell | 
**ExpirationTimestamp** | **int32** | Expiration timestamp for this order | 
**Fees** | Pointer to [**[]FeeEntry**](FeeEntry.md) | Fee information | [optional] 
**IncludeFees** | Pointer to **bool** | Whether to include fees in order | [optional] 
**Nonce** | **int32** | Nonce of the order | 
**StarkKey** | **string** | Public stark key of the user creating order | 
**StarkSignature** | **string** | Payload signature | 
**VaultIdBuy** | **int32** | ID of the vault into which the bought asset will be placed | 
**VaultIdSell** | **int32** | ID of the vault to sell from | 

## Methods

### NewCreateOrderRequest

`func NewCreateOrderRequest(amountBuy string, amountSell string, assetIdBuy string, assetIdSell string, expirationTimestamp int32, nonce int32, starkKey string, starkSignature string, vaultIdBuy int32, vaultIdSell int32, ) *CreateOrderRequest`

NewCreateOrderRequest instantiates a new CreateOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRequestWithDefaults

`func NewCreateOrderRequestWithDefaults() *CreateOrderRequest`

NewCreateOrderRequestWithDefaults instantiates a new CreateOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBuy

`func (o *CreateOrderRequest) GetAmountBuy() string`

GetAmountBuy returns the AmountBuy field if non-nil, zero value otherwise.

### GetAmountBuyOk

`func (o *CreateOrderRequest) GetAmountBuyOk() (*string, bool)`

GetAmountBuyOk returns a tuple with the AmountBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBuy

`func (o *CreateOrderRequest) SetAmountBuy(v string)`

SetAmountBuy sets AmountBuy field to given value.


### GetAmountSell

`func (o *CreateOrderRequest) GetAmountSell() string`

GetAmountSell returns the AmountSell field if non-nil, zero value otherwise.

### GetAmountSellOk

`func (o *CreateOrderRequest) GetAmountSellOk() (*string, bool)`

GetAmountSellOk returns a tuple with the AmountSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSell

`func (o *CreateOrderRequest) SetAmountSell(v string)`

SetAmountSell sets AmountSell field to given value.


### GetAssetIdBuy

`func (o *CreateOrderRequest) GetAssetIdBuy() string`

GetAssetIdBuy returns the AssetIdBuy field if non-nil, zero value otherwise.

### GetAssetIdBuyOk

`func (o *CreateOrderRequest) GetAssetIdBuyOk() (*string, bool)`

GetAssetIdBuyOk returns a tuple with the AssetIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdBuy

`func (o *CreateOrderRequest) SetAssetIdBuy(v string)`

SetAssetIdBuy sets AssetIdBuy field to given value.


### GetAssetIdSell

`func (o *CreateOrderRequest) GetAssetIdSell() string`

GetAssetIdSell returns the AssetIdSell field if non-nil, zero value otherwise.

### GetAssetIdSellOk

`func (o *CreateOrderRequest) GetAssetIdSellOk() (*string, bool)`

GetAssetIdSellOk returns a tuple with the AssetIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdSell

`func (o *CreateOrderRequest) SetAssetIdSell(v string)`

SetAssetIdSell sets AssetIdSell field to given value.


### GetExpirationTimestamp

`func (o *CreateOrderRequest) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *CreateOrderRequest) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *CreateOrderRequest) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetFees

`func (o *CreateOrderRequest) GetFees() []FeeEntry`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CreateOrderRequest) GetFeesOk() (*[]FeeEntry, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CreateOrderRequest) SetFees(v []FeeEntry)`

SetFees sets Fees field to given value.

### HasFees

`func (o *CreateOrderRequest) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetIncludeFees

`func (o *CreateOrderRequest) GetIncludeFees() bool`

GetIncludeFees returns the IncludeFees field if non-nil, zero value otherwise.

### GetIncludeFeesOk

`func (o *CreateOrderRequest) GetIncludeFeesOk() (*bool, bool)`

GetIncludeFeesOk returns a tuple with the IncludeFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFees

`func (o *CreateOrderRequest) SetIncludeFees(v bool)`

SetIncludeFees sets IncludeFees field to given value.

### HasIncludeFees

`func (o *CreateOrderRequest) HasIncludeFees() bool`

HasIncludeFees returns a boolean if a field has been set.

### GetNonce

`func (o *CreateOrderRequest) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateOrderRequest) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateOrderRequest) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetStarkKey

`func (o *CreateOrderRequest) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *CreateOrderRequest) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *CreateOrderRequest) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetStarkSignature

`func (o *CreateOrderRequest) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *CreateOrderRequest) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *CreateOrderRequest) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.


### GetVaultIdBuy

`func (o *CreateOrderRequest) GetVaultIdBuy() int32`

GetVaultIdBuy returns the VaultIdBuy field if non-nil, zero value otherwise.

### GetVaultIdBuyOk

`func (o *CreateOrderRequest) GetVaultIdBuyOk() (*int32, bool)`

GetVaultIdBuyOk returns a tuple with the VaultIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdBuy

`func (o *CreateOrderRequest) SetVaultIdBuy(v int32)`

SetVaultIdBuy sets VaultIdBuy field to given value.


### GetVaultIdSell

`func (o *CreateOrderRequest) GetVaultIdSell() int32`

GetVaultIdSell returns the VaultIdSell field if non-nil, zero value otherwise.

### GetVaultIdSellOk

`func (o *CreateOrderRequest) GetVaultIdSellOk() (*int32, bool)`

GetVaultIdSellOk returns a tuple with the VaultIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdSell

`func (o *CreateOrderRequest) SetVaultIdSell(v int32)`

SetVaultIdSell sets VaultIdSell field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


