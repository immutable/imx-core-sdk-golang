# SignableTransferResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the asset being transferred | 
**AssetId** | **string** | ID of the asset being transferred | 
**ExpirationTimestamp** | **int32** | Timestamp when this transfer will expire | 
**Nonce** | **int32** | Nonce of the transfer | 
**PayloadHash** | **string** | Hash of the payload to be signed for transfer | 
**ReceiverStarkKey** | **string** | Receiver of the transfer | 
**ReceiverVaultId** | **int32** | ID of the vault being transferred to | 
**SenderVaultId** | **int32** | ID of the vault being transferred from | 
**Token** | [**SignableToken**](SignableToken.md) |  | 

## Methods

### NewSignableTransferResponseDetails

`func NewSignableTransferResponseDetails(amount string, assetId string, expirationTimestamp int32, nonce int32, payloadHash string, receiverStarkKey string, receiverVaultId int32, senderVaultId int32, token SignableToken, ) *SignableTransferResponseDetails`

NewSignableTransferResponseDetails instantiates a new SignableTransferResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignableTransferResponseDetailsWithDefaults

`func NewSignableTransferResponseDetailsWithDefaults() *SignableTransferResponseDetails`

NewSignableTransferResponseDetailsWithDefaults instantiates a new SignableTransferResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SignableTransferResponseDetails) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SignableTransferResponseDetails) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SignableTransferResponseDetails) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAssetId

`func (o *SignableTransferResponseDetails) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *SignableTransferResponseDetails) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *SignableTransferResponseDetails) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetExpirationTimestamp

`func (o *SignableTransferResponseDetails) GetExpirationTimestamp() int32`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *SignableTransferResponseDetails) GetExpirationTimestampOk() (*int32, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *SignableTransferResponseDetails) SetExpirationTimestamp(v int32)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.


### GetNonce

`func (o *SignableTransferResponseDetails) GetNonce() int32`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *SignableTransferResponseDetails) GetNonceOk() (*int32, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *SignableTransferResponseDetails) SetNonce(v int32)`

SetNonce sets Nonce field to given value.


### GetPayloadHash

`func (o *SignableTransferResponseDetails) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *SignableTransferResponseDetails) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *SignableTransferResponseDetails) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.


### GetReceiverStarkKey

`func (o *SignableTransferResponseDetails) GetReceiverStarkKey() string`

GetReceiverStarkKey returns the ReceiverStarkKey field if non-nil, zero value otherwise.

### GetReceiverStarkKeyOk

`func (o *SignableTransferResponseDetails) GetReceiverStarkKeyOk() (*string, bool)`

GetReceiverStarkKeyOk returns a tuple with the ReceiverStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverStarkKey

`func (o *SignableTransferResponseDetails) SetReceiverStarkKey(v string)`

SetReceiverStarkKey sets ReceiverStarkKey field to given value.


### GetReceiverVaultId

`func (o *SignableTransferResponseDetails) GetReceiverVaultId() int32`

GetReceiverVaultId returns the ReceiverVaultId field if non-nil, zero value otherwise.

### GetReceiverVaultIdOk

`func (o *SignableTransferResponseDetails) GetReceiverVaultIdOk() (*int32, bool)`

GetReceiverVaultIdOk returns a tuple with the ReceiverVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiverVaultId

`func (o *SignableTransferResponseDetails) SetReceiverVaultId(v int32)`

SetReceiverVaultId sets ReceiverVaultId field to given value.


### GetSenderVaultId

`func (o *SignableTransferResponseDetails) GetSenderVaultId() int32`

GetSenderVaultId returns the SenderVaultId field if non-nil, zero value otherwise.

### GetSenderVaultIdOk

`func (o *SignableTransferResponseDetails) GetSenderVaultIdOk() (*int32, bool)`

GetSenderVaultIdOk returns a tuple with the SenderVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderVaultId

`func (o *SignableTransferResponseDetails) SetSenderVaultId(v int32)`

SetSenderVaultId sets SenderVaultId field to given value.


### GetToken

`func (o *SignableTransferResponseDetails) GetToken() SignableToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignableTransferResponseDetails) GetTokenOk() (*SignableToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignableTransferResponseDetails) SetToken(v SignableToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


