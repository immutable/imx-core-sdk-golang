# Deposit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Status of this deposit in Immutable X | 
**Timestamp** | **string** | Timestamp of the deposit | 
**Token** | [**DepositToken**](DepositToken.md) |  | 
**TransactionId** | **int32** | Sequential ID of this transaction within Immutable X | 
**User** | **string** | Ethereum address of the user making this deposit | 

## Methods

### NewDeposit

`func NewDeposit(status string, timestamp string, token DepositToken, transactionId int32, user string, ) *Deposit`

NewDeposit instantiates a new Deposit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDepositWithDefaults

`func NewDepositWithDefaults() *Deposit`

NewDepositWithDefaults instantiates a new Deposit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Deposit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Deposit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Deposit) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Deposit) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Deposit) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Deposit) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetToken

`func (o *Deposit) GetToken() DepositToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Deposit) GetTokenOk() (*DepositToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Deposit) SetToken(v DepositToken)`

SetToken sets Token field to given value.


### GetTransactionId

`func (o *Deposit) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Deposit) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Deposit) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.


### GetUser

`func (o *Deposit) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Deposit) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Deposit) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


