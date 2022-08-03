# CreateTransferRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount to transfer | 
**AssetId** | **string** | ID of the asset to transfer | 
**ExpirationTimestamp** | **int32** | Expiration timestamp for this transfer | 
**Nonce** | **int32** | Nonce of the transfer | 
**ReceiverStarkKey** | **string** | Public stark key of the user receiving the transfer | 
**ReceiverVaultId** | **int32** | ID of the vault into which the asset will be transferred to | 
**SenderStarkKey** | **string** | Public stark key of the user sending the transfer | 
**SenderVaultId** | **int32** | ID of the vault into which the asset is from | 
**StarkSignature** | **string** | Transfer payload signature | 

## Methods

### NewCreateTransferRequestV1

`func NewCreateTransferRequestV1(amount string, assetId string, expirationTimestamp int32, nonce int32, receiverStarkKey string, receiverVaultId int32, senderStarkKey string, senderVaultId int32, starkSignature string, ) *CreateTransferRequestV1`

NewCreateTransferRequestV1 instantiates a new CreateTransferRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestV1WithDefaults

`func NewCreateTransferRequestV1WithDefaults() *CreateTransferRequestV1`

NewCreateTransferRequestV1WithDefaults instantiates a new CreateTransferRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *CreateTransferRequestV1) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateTransferRequestV1) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateTransferRequestV1) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *CreateTransferRequestV1) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *CreateTransferRequestV1) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *CreateTransferRequestV1) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetExpirationTimestamp

`func (o *CreateTransferRequestV1) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *CreateTransferRequestV1) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *CreateTransferRequestV1) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetNonce

`func (o *CreateTransferRequestV1) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateTransferRequestV1) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateTransferRequestV1) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetReceiverStarkKey

`func (o *CreateTransferRequestV1) GetReceiverStarkKey() string`

GetReceiverStarkKey returns the ReceiverStarkKey field if non-nil, zero value otherwise.

### GetReceiverStarkKeyOk

`func (o *CreateTransferRequestV1) GetReceiverStarkKeyOk() (*string, bool)`

GetReceiverStarkKeyOk returns a tuple with the ReceiverStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverStarkKey

`func (o *CreateTransferRequestV1) SetReceiverStarkKey(v string)`

SetReceiverStarkKey sets ReceiverStarkKey field to given value.


### GetReceiverVaultId

`func (o *CreateTransferRequestV1) GetReceiverVaultId() int32`

GetReceiverVaultId returns the ReceiverVaultId field if non-nil, zero value otherwise.

### GetReceiverVaultIdOk

`func (o *CreateTransferRequestV1) GetReceiverVaultIdOk() (*int32, bool)`

GetReceiverVaultIdOk returns a tuple with the ReceiverVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverVaultId

`func (o *CreateTransferRequestV1) SetReceiverVaultId(v int32)`

SetReceiverVaultId sets ReceiverVaultId field to given value.


### GetSenderStarkKey

`func (o *CreateTransferRequestV1) GetSenderStarkKey() string`

GetSenderStarkKey returns the SenderStarkKey field if non-nil, zero value otherwise.

### GetSenderStarkKeyOk

`func (o *CreateTransferRequestV1) GetSenderStarkKeyOk() (*string, bool)`

GetSenderStarkKeyOk returns a tuple with the SenderStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderStarkKey

`func (o *CreateTransferRequestV1) SetSenderStarkKey(v string)`

SetSenderStarkKey sets SenderStarkKey field to given value.


### GetSenderVaultId

`func (o *CreateTransferRequestV1) GetSenderVaultId() int32`

GetSenderVaultId returns the SenderVaultId field if non-nil, zero value otherwise.

### GetSenderVaultIdOk

`func (o *CreateTransferRequestV1) GetSenderVaultIdOk() (*int32, bool)`

GetSenderVaultIdOk returns a tuple with the SenderVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderVaultId

`func (o *CreateTransferRequestV1) SetSenderVaultId(v int32)`

SetSenderVaultId sets SenderVaultId field to given value.


### GetStarkSignature

`func (o *CreateTransferRequestV1) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *CreateTransferRequestV1) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *CreateTransferRequestV1) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


