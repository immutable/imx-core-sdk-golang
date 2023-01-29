# GetSignableTradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBuy** | **string** | Amount to buy | 
**AmountSell** | **string** | Amount to sell | 
**AssetIdBuy** | **string** | ID of the asset to buy | 
**AssetIdSell** | **string** | ID of the asset to sell | 
**ExpirationTimestamp** | **int32** | Expiration timestamp for this order | 
**FeeInfo** | Pointer to [**FeeInfo**](FeeInfo.md) |  | [optional] 
**Nonce** | **int32** | Nonce of the order | 
**PayloadHash** | **string** | Payload Hash | 
**SignableMessage** | **string** | Message to sign with L1 wallet to confirm trade request | 
**StarkKey** | **string** | Public stark key of the created user | 
**VaultIdBuy** | **int32** | ID of the vault into which the bought asset will be placed | 
**VaultIdSell** | **int32** | ID of the vault to sell from | 

## Methods

### NewGetSignableTradeResponse

`func NewGetSignableTradeResponse(amountBuy string, amountSell string, assetIdBuy string, assetIdSell string, expirationTimestamp int32, nonce int32, payloadHash string, signableMessage string, starkKey string, vaultIdBuy int32, vaultIdSell int32, ) *GetSignableTradeResponse`

NewGetSignableTradeResponse instantiates a new GetSignableTradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTradeResponseWithDefaults

`func NewGetSignableTradeResponseWithDefaults() *GetSignableTradeResponse`

NewGetSignableTradeResponseWithDefaults instantiates a new GetSignableTradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBuy

`func (o *GetSignableTradeResponse) GetAmountBuy() string`

GetAmountBuy returns the AmountBuy field if non-nil, zero value otherwise.

### GetAmountBuyOk

`func (o *GetSignableTradeResponse) GetAmountBuyOk() (*string, bool)`

GetAmountBuyOk returns a tuple with the AmountBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBuy

`func (o *GetSignableTradeResponse) SetAmountBuy(v string)`

SetAmountBuy sets AmountBuy field to given value.


### GetAmountSell

`func (o *GetSignableTradeResponse) GetAmountSell() string`

GetAmountSell returns the AmountSell field if non-nil, zero value otherwise.

### GetAmountSellOk

`func (o *GetSignableTradeResponse) GetAmountSellOk() (*string, bool)`

GetAmountSellOk returns a tuple with the AmountSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSell

`func (o *GetSignableTradeResponse) SetAmountSell(v string)`

SetAmountSell sets AmountSell field to given value.


### GetAssetIdBuy

`func (o *GetSignableTradeResponse) GetAssetIdBuy() string`

GetAssetIdBuy returns the AssetIdBuy field if non-nil, zero value otherwise.

### GetAssetIdBuyOk

`func (o *GetSignableTradeResponse) GetAssetIdBuyOk() (*string, bool)`

GetAssetIdBuyOk returns a tuple with the AssetIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdBuy

`func (o *GetSignableTradeResponse) SetAssetIdBuy(v string)`

SetAssetIdBuy sets AssetIdBuy field to given value.


### GetAssetIdSell

`func (o *GetSignableTradeResponse) GetAssetIdSell() string`

GetAssetIdSell returns the AssetIdSell field if non-nil, zero value otherwise.

### GetAssetIdSellOk

`func (o *GetSignableTradeResponse) GetAssetIdSellOk() (*string, bool)`

GetAssetIdSellOk returns a tuple with the AssetIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdSell

`func (o *GetSignableTradeResponse) SetAssetIdSell(v string)`

SetAssetIdSell sets AssetIdSell field to given value.


### GetExpirationTimestamp

`func (o *GetSignableTradeResponse) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetSignableTradeResponse) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetSignableTradeResponse) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetFeeInfo

`func (o *GetSignableTradeResponse) GetFeeInfo() FeeInfo`

GetFeeInfo returns the FeeInfo field if non-nil, zero value otherwise.

### GetFeeInfoOk

`func (o *GetSignableTradeResponse) GetFeeInfoOk() (*FeeInfo, bool)`

GetFeeInfoOk returns a tuple with the FeeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeInfo

`func (o *GetSignableTradeResponse) SetFeeInfo(v FeeInfo)`

SetFeeInfo sets FeeInfo field to given value.

### HasFeeInfo

`func (o *GetSignableTradeResponse) HasFeeInfo() bool`

HasFeeInfo returns a boolean if a field has been set.

### GetNonce

`func (o *GetSignableTradeResponse) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetSignableTradeResponse) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetSignableTradeResponse) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetPayloadHash

`func (o *GetSignableTradeResponse) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableTradeResponse) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableTradeResponse) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.


### GetSignableMessage

`func (o *GetSignableTradeResponse) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableTradeResponse) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableTradeResponse) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.


### GetStarkKey

`func (o *GetSignableTradeResponse) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *GetSignableTradeResponse) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *GetSignableTradeResponse) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetVaultIdBuy

`func (o *GetSignableTradeResponse) GetVaultIdBuy() int32`

GetVaultIdBuy returns the VaultIdBuy field if non-nil, zero value otherwise.

### GetVaultIdBuyOk

`func (o *GetSignableTradeResponse) GetVaultIdBuyOk() (*int32, bool)`

GetVaultIdBuyOk returns a tuple with the VaultIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdBuy

`func (o *GetSignableTradeResponse) SetVaultIdBuy(v int32)`

SetVaultIdBuy sets VaultIdBuy field to given value.


### GetVaultIdSell

`func (o *GetSignableTradeResponse) GetVaultIdSell() int32`

GetVaultIdSell returns the VaultIdSell field if non-nil, zero value otherwise.

### GetVaultIdSellOk

`func (o *GetSignableTradeResponse) GetVaultIdSellOk() (*int32, bool)`

GetVaultIdSellOk returns a tuple with the VaultIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdSell

`func (o *GetSignableTradeResponse) SetVaultIdSell(v int32)`

SetVaultIdSell sets VaultIdSell field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


