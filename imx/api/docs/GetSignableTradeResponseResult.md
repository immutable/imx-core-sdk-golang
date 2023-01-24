# GetSignableTradeResponseResult

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

### NewGetSignableTradeResponseResult

`func NewGetSignableTradeResponseResult() *GetSignableTradeResponseResult`

NewGetSignableTradeResponseResult instantiates a new GetSignableTradeResponseResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTradeResponseResultWithDefaults

`func NewGetSignableTradeResponseResultWithDefaults() *GetSignableTradeResponseResult`

NewGetSignableTradeResponseResultWithDefaults instantiates a new GetSignableTradeResponseResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBuy

`func (o *GetSignableTradeResponseResult) GetAmountBuy() string`

GetAmountBuy returns the AmountBuy field if non-nil, zero value otherwise.

### GetAmountBuyOk

`func (o *GetSignableTradeResponseResult) GetAmountBuyOk() (*string, bool)`

GetAmountBuyOk returns a tuple with the AmountBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBuy

`func (o *GetSignableTradeResponseResult) SetAmountBuy(v string)`

SetAmountBuy sets AmountBuy field to given value.

### HasAmountBuy

`func (o *GetSignableTradeResponseResult) HasAmountBuy() bool`

HasAmountBuy returns a boolean if a field has been set.

### GetAmountSell

`func (o *GetSignableTradeResponseResult) GetAmountSell() string`

GetAmountSell returns the AmountSell field if non-nil, zero value otherwise.

### GetAmountSellOk

`func (o *GetSignableTradeResponseResult) GetAmountSellOk() (*string, bool)`

GetAmountSellOk returns a tuple with the AmountSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSell

`func (o *GetSignableTradeResponseResult) SetAmountSell(v string)`

SetAmountSell sets AmountSell field to given value.

### HasAmountSell

`func (o *GetSignableTradeResponseResult) HasAmountSell() bool`

HasAmountSell returns a boolean if a field has been set.

### GetAssetIdBuy

`func (o *GetSignableTradeResponseResult) GetAssetIdBuy() string`

GetAssetIdBuy returns the AssetIdBuy field if non-nil, zero value otherwise.

### GetAssetIdBuyOk

`func (o *GetSignableTradeResponseResult) GetAssetIdBuyOk() (*string, bool)`

GetAssetIdBuyOk returns a tuple with the AssetIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdBuy

`func (o *GetSignableTradeResponseResult) SetAssetIdBuy(v string)`

SetAssetIdBuy sets AssetIdBuy field to given value.

### HasAssetIdBuy

`func (o *GetSignableTradeResponseResult) HasAssetIdBuy() bool`

HasAssetIdBuy returns a boolean if a field has been set.

### GetAssetIdSell

`func (o *GetSignableTradeResponseResult) GetAssetIdSell() string`

GetAssetIdSell returns the AssetIdSell field if non-nil, zero value otherwise.

### GetAssetIdSellOk

`func (o *GetSignableTradeResponseResult) GetAssetIdSellOk() (*string, bool)`

GetAssetIdSellOk returns a tuple with the AssetIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdSell

`func (o *GetSignableTradeResponseResult) SetAssetIdSell(v string)`

SetAssetIdSell sets AssetIdSell field to given value.

### HasAssetIdSell

`func (o *GetSignableTradeResponseResult) HasAssetIdSell() bool`

HasAssetIdSell returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *GetSignableTradeResponseResult) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetSignableTradeResponseResult) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetSignableTradeResponseResult) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *GetSignableTradeResponseResult) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetFeeInfo

`func (o *GetSignableTradeResponseResult) GetFeeInfo() GetSignableTradeResultFeeInfo`

GetFeeInfo returns the FeeInfo field if non-nil, zero value otherwise.

### GetFeeInfoOk

`func (o *GetSignableTradeResponseResult) GetFeeInfoOk() (*GetSignableTradeResultFeeInfo, bool)`

GetFeeInfoOk returns a tuple with the FeeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeInfo

`func (o *GetSignableTradeResponseResult) SetFeeInfo(v GetSignableTradeResultFeeInfo)`

SetFeeInfo sets FeeInfo field to given value.

### HasFeeInfo

`func (o *GetSignableTradeResponseResult) HasFeeInfo() bool`

HasFeeInfo returns a boolean if a field has been set.

### GetNonce

`func (o *GetSignableTradeResponseResult) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetSignableTradeResponseResult) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetSignableTradeResponseResult) SetNonce(v int32)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *GetSignableTradeResponseResult) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetPayloadHash

`func (o *GetSignableTradeResponseResult) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableTradeResponseResult) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableTradeResponseResult) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.

### HasPayloadHash

`func (o *GetSignableTradeResponseResult) HasPayloadHash() bool`

HasPayloadHash returns a boolean if a field has been set.

### GetSignableMessage

`func (o *GetSignableTradeResponseResult) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableTradeResponseResult) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableTradeResponseResult) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.

### HasSignableMessage

`func (o *GetSignableTradeResponseResult) HasSignableMessage() bool`

HasSignableMessage returns a boolean if a field has been set.

### GetStarkKey

`func (o *GetSignableTradeResponseResult) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *GetSignableTradeResponseResult) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *GetSignableTradeResponseResult) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.

### HasStarkKey

`func (o *GetSignableTradeResponseResult) HasStarkKey() bool`

HasStarkKey returns a boolean if a field has been set.

### GetVaultIdBuy

`func (o *GetSignableTradeResponseResult) GetVaultIdBuy() int32`

GetVaultIdBuy returns the VaultIdBuy field if non-nil, zero value otherwise.

### GetVaultIdBuyOk

`func (o *GetSignableTradeResponseResult) GetVaultIdBuyOk() (*int32, bool)`

GetVaultIdBuyOk returns a tuple with the VaultIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdBuy

`func (o *GetSignableTradeResponseResult) SetVaultIdBuy(v int32)`

SetVaultIdBuy sets VaultIdBuy field to given value.

### HasVaultIdBuy

`func (o *GetSignableTradeResponseResult) HasVaultIdBuy() bool`

HasVaultIdBuy returns a boolean if a field has been set.

### GetVaultIdSell

`func (o *GetSignableTradeResponseResult) GetVaultIdSell() int32`

GetVaultIdSell returns the VaultIdSell field if non-nil, zero value otherwise.

### GetVaultIdSellOk

`func (o *GetSignableTradeResponseResult) GetVaultIdSellOk() (*int32, bool)`

GetVaultIdSellOk returns a tuple with the VaultIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdSell

`func (o *GetSignableTradeResponseResult) SetVaultIdSell(v int32)`

SetVaultIdSell sets VaultIdSell field to given value.

### HasVaultIdSell

`func (o *GetSignableTradeResponseResult) HasVaultIdSell() bool`

HasVaultIdSell returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


