# ProviderMintStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailureReason** | Pointer to **string** | reason for mint failure | [optional] 
**Id** | Pointer to **string** | External transaction id | [optional] 
**Status** | Pointer to **string** | Mint status | [optional] 
**StatusChangedAt** | Pointer to **string** | status updated-at timestamp | [optional] 
**TokenId** | Pointer to **string** | Asset id | [optional] 
**TransactionHash** | Pointer to **[]string** | Transaction hash | [optional] 

## Methods

### NewProviderMintStatus

`func NewProviderMintStatus() *ProviderMintStatus`

NewProviderMintStatus instantiates a new ProviderMintStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderMintStatusWithDefaults

`func NewProviderMintStatusWithDefaults() *ProviderMintStatus`

NewProviderMintStatusWithDefaults instantiates a new ProviderMintStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailureReason

`func (o *ProviderMintStatus) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *ProviderMintStatus) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *ProviderMintStatus) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *ProviderMintStatus) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### GetId

`func (o *ProviderMintStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderMintStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderMintStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderMintStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ProviderMintStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ProviderMintStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ProviderMintStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ProviderMintStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusChangedAt

`func (o *ProviderMintStatus) GetStatusChangedAt() string`

GetStatusChangedAt returns the StatusChangedAt field if non-nil, zero value otherwise.

### GetStatusChangedAtOk

`func (o *ProviderMintStatus) GetStatusChangedAtOk() (*string, bool)`

GetStatusChangedAtOk returns a tuple with the StatusChangedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusChangedAt

`func (o *ProviderMintStatus) SetStatusChangedAt(v string)`

SetStatusChangedAt sets StatusChangedAt field to given value.

### HasStatusChangedAt

`func (o *ProviderMintStatus) HasStatusChangedAt() bool`

HasStatusChangedAt returns a boolean if a field has been set.

### GetTokenId

`func (o *ProviderMintStatus) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *ProviderMintStatus) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *ProviderMintStatus) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *ProviderMintStatus) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTransactionHash

`func (o *ProviderMintStatus) GetTransactionHash() []string`

GetTransactionHash returns the TransactionHash field if non-nil, zero value otherwise.

### GetTransactionHashOk

`func (o *ProviderMintStatus) GetTransactionHashOk() (*[]string, bool)`

GetTransactionHashOk returns a tuple with the TransactionHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionHash

`func (o *ProviderMintStatus) SetTransactionHash(v []string)`

SetTransactionHash sets TransactionHash field to given value.

### HasTransactionHash

`func (o *ProviderMintStatus) HasTransactionHash() bool`

HasTransactionHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


