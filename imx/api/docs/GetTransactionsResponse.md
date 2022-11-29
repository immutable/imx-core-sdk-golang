# GetTransactionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | Pointer to **string** | Generated cursor returned by previous query | [optional] 
**Remaining** | Pointer to **int32** | Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results | [optional] 
**Result** | Pointer to [**[]Exchange**](Exchange.md) | Transactions matching query parameters | [optional] 

## Methods

### NewGetTransactionsResponse

`func NewGetTransactionsResponse() *GetTransactionsResponse`

NewGetTransactionsResponse instantiates a new GetTransactionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTransactionsResponseWithDefaults

`func NewGetTransactionsResponseWithDefaults() *GetTransactionsResponse`

NewGetTransactionsResponseWithDefaults instantiates a new GetTransactionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *GetTransactionsResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetTransactionsResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetTransactionsResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetTransactionsResponse) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetRemaining

`func (o *GetTransactionsResponse) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *GetTransactionsResponse) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *GetTransactionsResponse) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.

### HasRemaining

`func (o *GetTransactionsResponse) HasRemaining() bool`

HasRemaining returns a boolean if a field has been set.

### GetResult

`func (o *GetTransactionsResponse) GetResult() []Exchange`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetTransactionsResponse) GetResultOk() (*[]Exchange, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetTransactionsResponse) SetResult(v []Exchange)`

SetResult sets Result field to given value.

### HasResult

`func (o *GetTransactionsResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


