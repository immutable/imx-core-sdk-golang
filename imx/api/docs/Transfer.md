# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Receiver** | **string** | Ethereum address of the user who received this transfer | 
**Status** | **string** | Status of the transaction | 
**Timestamp** | **NullableString** | Timestamp of the transfer | 
**Token** | [**TransferToken**](TransferToken.md) |  | 
**TransactionId** | **int32** | Sequential transaction ID | 
**User** | **string** | Ethereum address of the user  who submitted this transfer | 

## Methods

### NewTransfer

`func NewTransfer(receiver string, status string, timestamp NullableString, token TransferToken, transactionId int32, user string, ) *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceiver

`func (o *Transfer) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *Transfer) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *Transfer) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetStatus

`func (o *Transfer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transfer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transfer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Transfer) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Transfer) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Transfer) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### SetTimestampNil

`func (o *Transfer) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *Transfer) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil
### GetToken

`func (o *Transfer) GetToken() TransferToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Transfer) GetTokenOk() (*TransferToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Transfer) SetToken(v TransferToken)`

SetToken sets Token field to given value.


### GetTransactionId

`func (o *Transfer) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Transfer) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Transfer) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.


### GetUser

`func (o *Transfer) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Transfer) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Transfer) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


