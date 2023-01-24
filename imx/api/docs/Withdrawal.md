# Withdrawal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RollupStatus** | **string** | Status of the on-chain batch confirmation for this withdrawal | 
**Sender** | **string** | Ethereum address of the user who requested this withdrawal | 
**Status** | **string** | Status of this withdrawal | 
**Timestamp** | **string** | Time when this withdrawal was initiated | 
**Token** | [**WithdrawalToken**](WithdrawalToken.md) |  | 
**TransactionId** | **int32** | Sequential ID of this transaction | 
**WithdrawnToWallet** | **bool** | Withdrawal has been transferred to user&#39;s Layer 1 wallet | 

## Methods

### NewWithdrawal

`func NewWithdrawal(rollupStatus string, sender string, status string, timestamp string, token WithdrawalToken, transactionId int32, withdrawnToWallet bool, ) *Withdrawal`

NewWithdrawal instantiates a new Withdrawal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWithdrawalWithDefaults

`func NewWithdrawalWithDefaults() *Withdrawal`

NewWithdrawalWithDefaults instantiates a new Withdrawal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRollupStatus

`func (o *Withdrawal) GetRollupStatus() string`

GetRollupStatus returns the RollupStatus field if non-nil, zero value otherwise.

### GetRollupStatusOk

`func (o *Withdrawal) GetRollupStatusOk() (*string, bool)`

GetRollupStatusOk returns a tuple with the RollupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollupStatus

`func (o *Withdrawal) SetRollupStatus(v string)`

SetRollupStatus sets RollupStatus field to given value.


### GetSender

`func (o *Withdrawal) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *Withdrawal) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *Withdrawal) SetSender(v string)`

SetSender sets Sender field to given value.


### GetStatus

`func (o *Withdrawal) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Withdrawal) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Withdrawal) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Withdrawal) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Withdrawal) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Withdrawal) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetToken

`func (o *Withdrawal) GetToken() WithdrawalToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Withdrawal) GetTokenOk() (*WithdrawalToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Withdrawal) SetToken(v WithdrawalToken)`

SetToken sets Token field to given value.


### GetTransactionId

`func (o *Withdrawal) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Withdrawal) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Withdrawal) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.


### GetWithdrawnToWallet

`func (o *Withdrawal) GetWithdrawnToWallet() bool`

GetWithdrawnToWallet returns the WithdrawnToWallet field if non-nil, zero value otherwise.

### GetWithdrawnToWalletOk

`func (o *Withdrawal) GetWithdrawnToWalletOk() (*bool, bool)`

GetWithdrawnToWalletOk returns a tuple with the WithdrawnToWallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWithdrawnToWallet

`func (o *Withdrawal) SetWithdrawnToWallet(v bool)`

SetWithdrawnToWallet sets WithdrawnToWallet field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


