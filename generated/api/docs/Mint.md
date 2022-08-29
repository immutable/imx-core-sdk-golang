# Mint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fees** | Pointer to [**[]Fee**](Fee.md) | Fee details | [optional] 
**Status** | **string** | Status of this mint | 
**Timestamp** | **string** | Timestamp this mint was initiated | 
**Token** | [**Token**](Token.md) |  | 
**TransactionId** | **int32** | Sequential ID of transaction in Immutable X | 
**User** | **string** | Ethereum address of the user to whom the asset has been minted | 

## Methods

### NewMint

`func NewMint(status string, timestamp string, token Token, transactionId int32, user string, ) *Mint`

NewMint instantiates a new Mint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintWithDefaults

`func NewMintWithDefaults() *Mint`

NewMintWithDefaults instantiates a new Mint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFees

`func (o *Mint) GetFees() []Fee`

GetFees returns the Fees field if non-nil, zero value otherwise.

### GetFeesOk

`func (o *Mint) GetFeesOk() (*[]Fee, bool)`

GetFeesOk returns a tuple with the Fees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFees

`func (o *Mint) SetFees(v []Fee)`

SetFees sets Fees field to given value.

### HasFees

`func (o *Mint) HasFees() bool`

HasFees returns a boolean if a field has been set.

### GetStatus

`func (o *Mint) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Mint) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Mint) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *Mint) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Mint) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Mint) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetToken

`func (o *Mint) GetToken() Token`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *Mint) GetTokenOk() (*Token, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *Mint) SetToken(v Token)`

SetToken sets Token field to given value.


### GetTransactionId

`func (o *Mint) GetTransactionId() int32`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Mint) GetTransactionIdOk() (*int32, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Mint) SetTransactionId(v int32)`

SetTransactionId sets TransactionId field to given value.


### GetUser

`func (o *Mint) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Mint) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Mint) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


