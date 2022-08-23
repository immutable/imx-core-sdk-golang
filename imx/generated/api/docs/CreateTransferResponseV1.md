# CreateTransferResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SentSignature** | **string** | [deprecated] Sent signature | 
**Status** | **string** | [deprecated] The status of transfer | 
**Time** | **int32** | [deprecated] Time of the transfer | 
**TransferId** | **int32** | ID of the transfer | 

## Methods

### NewCreateTransferResponseV1

`func NewCreateTransferResponseV1(sentSignature string, status string, time int32, transferId int32, ) *CreateTransferResponseV1`

NewCreateTransferResponseV1 instantiates a new CreateTransferResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferResponseV1WithDefaults

`func NewCreateTransferResponseV1WithDefaults() *CreateTransferResponseV1`

NewCreateTransferResponseV1WithDefaults instantiates a new CreateTransferResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSentSignature

`func (o *CreateTransferResponseV1) GetSentSignature() string`

GetSentSignature returns the SentSignature field if non-nil, zero value otherwise.

### GetSentSignatureOk

`func (o *CreateTransferResponseV1) GetSentSignatureOk() (*string, bool)`

GetSentSignatureOk returns a tuple with the SentSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentSignature

`func (o *CreateTransferResponseV1) SetSentSignature(v string)`

SetSentSignature sets SentSignature field to given value.


### GetStatus

`func (o *CreateTransferResponseV1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTransferResponseV1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTransferResponseV1) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTime

`func (o *CreateTransferResponseV1) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateTransferResponseV1) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateTransferResponseV1) SetTime(v int32)`

SetTime sets Time field to given value.


### GetTransferId

`func (o *CreateTransferResponseV1) GetTransferId() int32`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *CreateTransferResponseV1) GetTransferIdOk() (*int32, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *CreateTransferResponseV1) SetTransferId(v int32)`

SetTransferId sets TransferId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


