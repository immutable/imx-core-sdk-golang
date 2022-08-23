# CreateTradeRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountBuy** | **string** | Amount to buy | 
**AmountSell** | **string** | Amount to sell | 
**AssetIdBuy** | **string** | ID of the asset to buy | 
**AssetIdSell** | **string** | ID of the asset to sell | 
**ExpirationTimestamp** | **int32** | Expiration timestamp for this trade | 
**FeeInfo** | Pointer to [**FeeInfo**](FeeInfo.md) |  | [optional] 
**Fees** | Pointer to [**[]FeeEntry**](FeeEntry.md) | Fee information | [optional] 
**IncludeFees** | Pointer to **bool** | [deprecated] All orders include fees since the introduction of protocol fees | [optional] 
**Nonce** | **int32** | Nonce of the trade | 
**OrderId** | **int32** | ID of the order | 
**StarkKey** | **string** | Public stark key of the user creating trade | 
**StarkSignature** | **string** | Payload signature | 
**VaultIdBuy** | **int32** | ID of the vault into which the traded asset will be placed | 
**VaultIdSell** | **int32** | ID of the vault to sell from | 

## Methods

### NewCreateTradeRequestV1

`func NewCreateTradeRequestV1(amountBuy string, amountSell string, assetIdBuy string, assetIdSell string, expirationTimestamp int32, nonce int32, orderId int32, starkKey string, starkSignature string, vaultIdBuy int32, vaultIdSell int32, ) *CreateTradeRequestV1`

NewCreateTradeRequestV1 instantiates a new CreateTradeRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeRequestV1WithDefaults

`func NewCreateTradeRequestV1WithDefaults() *CreateTradeRequestV1`

NewCreateTradeRequestV1WithDefaults instantiates a new CreateTradeRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountBuy

`func (o *CreateTradeRequestV1) GetAmountBuy() string`

GetAmountBuy returns the AmountBuy field if non-nil, zero value otherwise.

### GetAmountBuyOk

`func (o *CreateTradeRequestV1) GetAmountBuyOk() (*string, bool)`

GetAmountBuyOk returns a tuple with the AmountBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountBuy

`func (o *CreateTradeRequestV1) SetAmountBuy(v string)`

SetAmountBuy sets AmountBuy field to given value.


### GetAmountSell

`func (o *CreateTradeRequestV1) GetAmountSell() string`

GetAmountSell returns the AmountSell field if non-nil, zero value otherwise.

### GetAmountSellOk

`func (o *CreateTradeRequestV1) GetAmountSellOk() (*string, bool)`

GetAmountSellOk returns a tuple with the AmountSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountSell

`func (o *CreateTradeRequestV1) SetAmountSell(v string)`

SetAmountSell sets AmountSell field to given value.


### GetAssetIdBuy

`func (o *CreateTradeRequestV1) GetAssetIdBuy() string`

GetAssetIdBuy returns the AssetIdBuy field if non-nil, zero value otherwise.

### GetAssetIdBuyOk

`func (o *CreateTradeRequestV1) GetAssetIdBuyOk() (*string, bool)`

GetAssetIdBuyOk returns a tuple with the AssetIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdBuy

`func (o *CreateTradeRequestV1) SetAssetIdBuy(v string)`

SetAssetIdBuy sets AssetIdBuy field to given value.


### GetAssetIdSell

`func (o *CreateTradeRequestV1) GetAssetIdSell() string`

GetAssetIdSell returns the AssetIdSell field if non-nil, zero value otherwise.

### GetAssetIdSellOk

`func (o *CreateTradeRequestV1) GetAssetIdSellOk() (*string, bool)`

GetAssetIdSellOk returns a tuple with the AssetIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdSell

`func (o *CreateTradeRequestV1) SetAssetIdSell(v string)`

SetAssetIdSell sets AssetIdSell field to given value.


### GetExpirationTimestamp

`func (o *CreateTradeRequestV1) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *CreateTradeRequestV1) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *CreateTradeRequestV1) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetFeeInfo

`func (o *CreateTradeRequestV1) GetFeeInfo() FeeInfo`

GetFeeInfo returns the FeeInfo field if non-nil, zero value otherwise.

### GetFeeInfoOk

`func (o *CreateTradeRequestV1) GetFeeInfoOk() (*FeeInfo, bool)`

GetFeeInfoOk returns a tuple with the FeeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeInfo

`func (o *CreateTradeRequestV1) SetFeeInfo(v FeeInfo)`

SetFeeInfo sets FeeInfo field to given value.

### HasFeeInfo

`func (o *CreateTradeRequestV1) HasFeeInfo() bool`

HasFeeInfo returns a boolean if a field has been set.

### GetFees

`func (o *CreateTradeRequestV1) GetFees() []FeeEntry`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *CreateTradeRequestV1) GetFeesOk() (*[]FeeEntry, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *CreateTradeRequestV1) SetFees(v []FeeEntry)`

SetFees sets Fees field to given value.

### HasFees

`func (o *CreateTradeRequestV1) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetIncludeFees

`func (o *CreateTradeRequestV1) GetIncludeFees() bool`

GetIncludeFees returns the IncludeFees field if non-nil, zero value otherwise.

### GetIncludeFeesOk

`func (o *CreateTradeRequestV1) GetIncludeFeesOk() (*bool, bool)`

GetIncludeFeesOk returns a tuple with the IncludeFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeFees

`func (o *CreateTradeRequestV1) SetIncludeFees(v bool)`

SetIncludeFees sets IncludeFees field to given value.

### HasIncludeFees

`func (o *CreateTradeRequestV1) HasIncludeFees() bool`

HasIncludeFees returns a boolean if a field has been set.

### GetNonce

`func (o *CreateTradeRequestV1) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateTradeRequestV1) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateTradeRequestV1) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetOrderId

`func (o *CreateTradeRequestV1) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateTradeRequestV1) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateTradeRequestV1) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetStarkKey

`func (o *CreateTradeRequestV1) GetStarkKey() string`

GetStarkKey returns the StarkKey field if non-nil, zero value otherwise.

### GetStarkKeyOk

`func (o *CreateTradeRequestV1) GetStarkKeyOk() (*string, bool)`

GetStarkKeyOk returns a tuple with the StarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkKey

`func (o *CreateTradeRequestV1) SetStarkKey(v string)`

SetStarkKey sets StarkKey field to given value.


### GetStarkSignature

`func (o *CreateTradeRequestV1) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *CreateTradeRequestV1) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *CreateTradeRequestV1) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.


### GetVaultIdBuy

`func (o *CreateTradeRequestV1) GetVaultIdBuy() int32`

GetVaultIdBuy returns the VaultIdBuy field if non-nil, zero value otherwise.

### GetVaultIdBuyOk

`func (o *CreateTradeRequestV1) GetVaultIdBuyOk() (*int32, bool)`

GetVaultIdBuyOk returns a tuple with the VaultIdBuy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdBuy

`func (o *CreateTradeRequestV1) SetVaultIdBuy(v int32)`

SetVaultIdBuy sets VaultIdBuy field to given value.


### GetVaultIdSell

`func (o *CreateTradeRequestV1) GetVaultIdSell() int32`

GetVaultIdSell returns the VaultIdSell field if non-nil, zero value otherwise.

### GetVaultIdSellOk

`func (o *CreateTradeRequestV1) GetVaultIdSellOk() (*int32, bool)`

GetVaultIdSellOk returns a tuple with the VaultIdSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaultIdSell

`func (o *CreateTradeRequestV1) SetVaultIdSell(v int32)`

SetVaultIdSell sets VaultIdSell field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


