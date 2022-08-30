# ListOrdersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Generated cursor returned by previous query | 
**Remaining** | **int32** | Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results | 
**Result** | [**[]Order**](Order.md) | Orders matching query parameters | 

## Methods

### NewListOrdersResponse

`func NewListOrdersResponse(cursor string, remaining int32, result []Order, ) *ListOrdersResponse`

NewListOrdersResponse instantiates a new ListOrdersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrdersResponseWithDefaults

`func NewListOrdersResponseWithDefaults() *ListOrdersResponse`

NewListOrdersResponseWithDefaults instantiates a new ListOrdersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListOrdersResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListOrdersResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListOrdersResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetRemaining

`func (o *ListOrdersResponse) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *ListOrdersResponse) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *ListOrdersResponse) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### GetResult

`func (o *ListOrdersResponse) GetResult() []Order`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListOrdersResponse) GetResultOk() (*[]Order, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListOrdersResponse) SetResult(v []Order)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


