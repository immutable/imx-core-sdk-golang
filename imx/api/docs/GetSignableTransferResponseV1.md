# GetSignableTransferResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the asset being transferred | 
**AssetId** | **string** | ID of the asset being transferred | 
**ExpirationTimestamp** | **int32** | Token in request to match in SDK implementation | 
**Nonce** | **int32** | Nonce of the transfer | 
**PayloadHash** | **string** | Hash of the payload | 
**ReceiverStarkKey** | **string** | Receiver of the transfer | 
**ReceiverVaultId** | **int32** | ID of the vault being transferred to | 
**SenderStarkKey** | Pointer to **string** | Sender of the transfer | [optional] 
**SenderVaultId** | **int32** | ID of the vault being transferred from | 
**SignableMessage** | **string** | Message to sign with L1 wallet to confirm transfer request | 

## Methods

### NewGetSignableTransferResponseV1

`func NewGetSignableTransferResponseV1(amount string, assetId string, expirationTimestamp int32, nonce int32, payloadHash string, receiverStarkKey string, receiverVaultId int32, senderVaultId int32, signableMessage string, ) *GetSignableTransferResponseV1`

NewGetSignableTransferResponseV1 instantiates a new GetSignableTransferResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTransferResponseV1WithDefaults

`func NewGetSignableTransferResponseV1WithDefaults() *GetSignableTransferResponseV1`

NewGetSignableTransferResponseV1WithDefaults instantiates a new GetSignableTransferResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSignableTransferResponseV1) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSignableTransferResponseV1) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSignableTransferResponseV1) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *GetSignableTransferResponseV1) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *GetSignableTransferResponseV1) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *GetSignableTransferResponseV1) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetExpirationTimestamp

`func (o *GetSignableTransferResponseV1) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *GetSignableTransferResponseV1) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *GetSignableTransferResponseV1) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetNonce

`func (o *GetSignableTransferResponseV1) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *GetSignableTransferResponseV1) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *GetSignableTransferResponseV1) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetPayloadHash

`func (o *GetSignableTransferResponseV1) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableTransferResponseV1) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableTransferResponseV1) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.


### GetReceiverStarkKey

`func (o *GetSignableTransferResponseV1) GetReceiverStarkKey() string`

GetReceiverStarkKey returns the ReceiverStarkKey field if non-nil, zero value otherwise.

### GetReceiverStarkKeyOk

`func (o *GetSignableTransferResponseV1) GetReceiverStarkKeyOk() (*string, bool)`

GetReceiverStarkKeyOk returns a tuple with the ReceiverStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverStarkKey

`func (o *GetSignableTransferResponseV1) SetReceiverStarkKey(v string)`

SetReceiverStarkKey sets ReceiverStarkKey field to given value.


### GetReceiverVaultId

`func (o *GetSignableTransferResponseV1) GetReceiverVaultId() int32`

GetReceiverVaultId returns the ReceiverVaultId field if non-nil, zero value otherwise.

### GetReceiverVaultIdOk

`func (o *GetSignableTransferResponseV1) GetReceiverVaultIdOk() (*int32, bool)`

GetReceiverVaultIdOk returns a tuple with the ReceiverVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverVaultId

`func (o *GetSignableTransferResponseV1) SetReceiverVaultId(v int32)`

SetReceiverVaultId sets ReceiverVaultId field to given value.


### GetSenderStarkKey

`func (o *GetSignableTransferResponseV1) GetSenderStarkKey() string`

GetSenderStarkKey returns the SenderStarkKey field if non-nil, zero value otherwise.

### GetSenderStarkKeyOk

`func (o *GetSignableTransferResponseV1) GetSenderStarkKeyOk() (*string, bool)`

GetSenderStarkKeyOk returns a tuple with the SenderStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderStarkKey

`func (o *GetSignableTransferResponseV1) SetSenderStarkKey(v string)`

SetSenderStarkKey sets SenderStarkKey field to given value.

### HasSenderStarkKey

`func (o *GetSignableTransferResponseV1) HasSenderStarkKey() bool`

HasSenderStarkKey returns a boolean if a field has been set.

### GetSenderVaultId

`func (o *GetSignableTransferResponseV1) GetSenderVaultId() int32`

GetSenderVaultId returns the SenderVaultId field if non-nil, zero value otherwise.

### GetSenderVaultIdOk

`func (o *GetSignableTransferResponseV1) GetSenderVaultIdOk() (*int32, bool)`

GetSenderVaultIdOk returns a tuple with the SenderVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderVaultId

`func (o *GetSignableTransferResponseV1) SetSenderVaultId(v int32)`

SetSenderVaultId sets SenderVaultId field to given value.


### GetSignableMessage

`func (o *GetSignableTransferResponseV1) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableTransferResponseV1) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableTransferResponseV1) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


