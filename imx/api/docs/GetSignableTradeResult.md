# GetSignableTradeResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBuy** | Pointer to **string** | Amount to buy | [optional] 
**AmountSell** | Pointer to **string** | Amount to sell | [optional] 
**AssetIdBuy** | Pointer to **string** | ID of the asset to buy | [optional] 
**AssetIdSell** | Pointer to **string** | ID of the asset to sell | [optional] 
**ExpirationTimestamp** | Pointer to **int32** | Expiration timestamp of this trade | [optional] 
**FeeInfo** | Pointer to [**GetSignableTradeResultFeeInfo**](GetSignableTradeResultFeeInfo.md) |  | [optional] 
**Nonce** | Pointer to **int32** | Nonce of the trade | [optional] 
**PayloadHash** | Pointer to **string** | Payload Hash | [optional] 
**SignableMessage** | Pointer to **string** | Message to sign with L1 wallet to confirm trade request | [optional] 
**StarkKey** | Pointer to **string** | Public stark key of the trading user | [optional] 
**VaultIdBuy** | Pointer to **int32** | ID of the vault into which the traded asset will be placed | [optional] 
**VaultIdSell** | Pointer to **int32** | ID of the vault to sell from | [optional] 

## Methods

### NewGetSignableTradeResult

`func NewGetSignableTradeResult() *GetSignableTradeResult`

NewGetSignableTradeResult instantiates a new GetSignableTradeResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTradeResultWithDefaults

`func NewGetSignableTradeResultWithDefaults() *GetSignableTradeResult`

NewGetSignableTradeResultWithDefaults instantiates a new GetSignableTradeResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBuy

`func (o *GetSignableTradeResult) GetAmountBuy() string`

GetAmountBuy returns the AmountBuy field if non-nil, zero value otherwise.

### GetAmountBuyOk

`func (o *GetSignableTradeResult) GetAmountBuyOk() (*string, bool)`

GetAmountBuyOk returns a tuple with the AmountBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBuy

`func (o *GetSignableTradeResult) SetAmountBuy(v string)`

SetAmountBuy sets AmountBuy field to given value.

### HasAmountBuy

`func (o *GetSignableTradeResult) HasAmountBuy() bool`

HasAmountBuy returns a boolean if a field has been set.

### GetAmountSell

`func (o *GetSignableTradeResult) GetAmountSell() string`

GetAmountSell returns the AmountSell field if non-nil, zero value otherwise.

### GetAmountSellOk

`func (o *GetSignableTradeResult) GetAmountSellOk() (*string, bool)`

GetAmountSellOk returns a tuple with the AmountSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSell

`func (o *GetSignableTradeResult) SetAmountSell(v string)`

SetAmountSell sets AmountSell field to given value.

### HasAmountSell

`func (o *GetSignableTradeResult) HasAmountSell() bool`

HasAmountSell returns a boolean if a field has been set.

### GetAssetIdBuy

`func (o *GetSignableTradeResult) GetAssetIdBuy() string`

GetAssetIdBuy returns the AssetIdBuy field if non-nil, zero value otherwise.

### GetAssetIdBuyOk

`func (o *GetSignableTradeResult) GetAssetIdBuyOk() (*string, bool)`

GetAssetIdBuyOk returns a tuple with the AssetIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdBuy

`func (o *GetSignableTradeResult) SetAssetIdBuy(v string)`

SetAssetIdBuy sets AssetIdBuy field to given value.

### HasAssetIdBuy

`func (o *GetSignableTradeResult) HasAssetIdBuy() bool`

HasAssetIdBuy returns a boolean if a field has been set.

### GetAssetIdSell

`func (o *GetSignableTradeResult) GetAssetIdSell() string`

GetAssetIdSell returns the AssetIdSell field if non-nil, zero value otherwise.

### GetAssetIdSellOk

`func (o *GetSignableTradeResult) GetAssetIdSellOk() (*string, bool)`

GetAssetIdSellOk returns a tuple with the AssetIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdSell

`func (o *GetSignableTradeResult) SetAssetIdSell(v string)`

SetAssetIdSell sets AssetIdSell field to given value.

### HasAssetIdSell

`func (o *GetSignableTradeResult) HasAssetIdSell() bool`

HasAssetIdSell returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *GetSignableTradeResult) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetSignableTradeResult) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetSignableTradeResult) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetSignableTradeResult) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetFeeInfo

`func (o *GetSignableTradeResult) GetFeeInfo() GetSignableTradeResultFeeInfo`

GetFeeInfo returns the FeeInfo field if non-nil, zero value otherwise.

### GetFeeInfoOk

`func (o *GetSignableTradeResult) GetFeeInfoOk() (*GetSignableTradeResultFeeInfo, bool)`

GetFeeInfoOk returns a tuple with the FeeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeInfo

`func (o *GetSignableTradeResult) SetFeeInfo(v GetSignableTradeResultFeeInfo)`

SetFeeInfo sets FeeInfo field to given value.

### HasFeeInfo

`func (o *GetSignableTradeResult) HasFeeInfo() bool`

HasFeeInfo returns a boolean if a field has been set.

### GetNonce

`func (o *GetSignableTradeResult) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetSignableTradeResult) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetSignableTradeResult) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *GetSignableTradeResult) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPayloadHash

`func (o *GetSignableTradeResult) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableTradeResult) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableTradeResult) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.

### HasPayloadHash

`func (o *GetSignableTradeResult) HasPayloadHash() bool`

HasPayloadHash returns a boolean if a field has been set.

### GetSignableMessage

`func (o *GetSignableTradeResult) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableTradeResult) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableTradeResult) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.

### HasSignableMessage

`func (o *GetSignableTradeResult) HasSignableMessage() bool`

HasSignableMessage returns a boolean if a field has been set.

### GetStarkKey

`func (o *GetSignableTradeResult) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *GetSignableTradeResult) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *GetSignableTradeResult) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.

### HasStarkKey

`func (o *GetSignableTradeResult) HasStarkKey() bool`

HasStarkKey returns a boolean if a field has been set.

### GetVaultIdBuy

`func (o *GetSignableTradeResult) GetVaultIdBuy() int32`

GetVaultIdBuy returns the VaultIdBuy field if non-nil, zero value otherwise.

### GetVaultIdBuyOk

`func (o *GetSignableTradeResult) GetVaultIdBuyOk() (*int32, bool)`

GetVaultIdBuyOk returns a tuple with the VaultIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdBuy

`func (o *GetSignableTradeResult) SetVaultIdBuy(v int32)`

SetVaultIdBuy sets VaultIdBuy field to given value.

### HasVaultIdBuy

`func (o *GetSignableTradeResult) HasVaultIdBuy() bool`

HasVaultIdBuy returns a boolean if a field has been set.

### GetVaultIdSell

`func (o *GetSignableTradeResult) GetVaultIdSell() int32`

GetVaultIdSell returns the VaultIdSell field if non-nil, zero value otherwise.

### GetVaultIdSellOk

`func (o *GetSignableTradeResult) GetVaultIdSellOk() (*int32, bool)`

GetVaultIdSellOk returns a tuple with the VaultIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdSell

`func (o *GetSignableTradeResult) SetVaultIdSell(v int32)`

SetVaultIdSell sets VaultIdSell field to given value.

### HasVaultIdSell

`func (o *GetSignableTradeResult) HasVaultIdSell() bool`

HasVaultIdSell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


