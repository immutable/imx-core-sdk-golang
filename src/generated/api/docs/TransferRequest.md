# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount to transfer | 
**AssetId** | **string** | ID of the asset to transfer | 
**ExpirationTimestamp** | **int32** | Expiration timestamp for this transfer | 
**Nonce** | **int32** | Nonce of the transfer | 
**ReceiverStarkKey** | **string** | Public stark key of the user receiving the transfer | 
**ReceiverVaultId** | **int32** | ID of the vault into which the asset will be transferred to | 
**SenderVaultId** | **int32** | ID of the vault into which the asset is from | 
**StarkSignature** | **string** | Transfer payload signature | 

## Methods

### NewTransferRequest

`func NewTransferRequest(amount string, assetId string, expirationTimestamp int32, nonce int32, receiverStarkKey string, receiverVaultId int32, senderVaultId int32, starkSignature string, ) *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *TransferRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *TransferRequest) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *TransferRequest) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *TransferRequest) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetExpirationTimestamp

`func (o *TransferRequest) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *TransferRequest) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *TransferRequest) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetNonce

`func (o *TransferRequest) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *TransferRequest) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *TransferRequest) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetReceiverStarkKey

`func (o *TransferRequest) GetReceiverStarkKey() string`

GetReceiverStarkKey returns the ReceiverStarkKey field if non-nil, zero value otherwise.

### GetReceiverStarkKeyOk

`func (o *TransferRequest) GetReceiverStarkKeyOk() (*string, bool)`

GetReceiverStarkKeyOk returns a tuple with the ReceiverStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverStarkKey

`func (o *TransferRequest) SetReceiverStarkKey(v string)`

SetReceiverStarkKey sets ReceiverStarkKey field to given value.


### GetReceiverVaultId

`func (o *TransferRequest) GetReceiverVaultId() int32`

GetReceiverVaultId returns the ReceiverVaultId field if non-nil, zero value otherwise.

### GetReceiverVaultIdOk

`func (o *TransferRequest) GetReceiverVaultIdOk() (*int32, bool)`

GetReceiverVaultIdOk returns a tuple with the ReceiverVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverVaultId

`func (o *TransferRequest) SetReceiverVaultId(v int32)`

SetReceiverVaultId sets ReceiverVaultId field to given value.


### GetSenderVaultId

`func (o *TransferRequest) GetSenderVaultId() int32`

GetSenderVaultId returns the SenderVaultId field if non-nil, zero value otherwise.

### GetSenderVaultIdOk

`func (o *TransferRequest) GetSenderVaultIdOk() (*int32, bool)`

GetSenderVaultIdOk returns a tuple with the SenderVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderVaultId

`func (o *TransferRequest) SetSenderVaultId(v int32)`

SetSenderVaultId sets SenderVaultId field to given value.


### GetStarkSignature

`func (o *TransferRequest) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *TransferRequest) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *TransferRequest) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


