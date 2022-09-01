# ListCollectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Generated cursor returned by previous query | 
**Remaining** | **int32** | Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results | 
**Result** | [**[]Collection**](Collection.md) | Collections matching query parameters | 

## Methods

### NewListCollectionsResponse

`func NewListCollectionsResponse(cursor string, remaining int32, result []Collection, ) *ListCollectionsResponse`

NewListCollectionsResponse instantiates a new ListCollectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectionsResponseWithDefaults

`func NewListCollectionsResponseWithDefaults() *ListCollectionsResponse`

NewListCollectionsResponseWithDefaults instantiates a new ListCollectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListCollectionsResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListCollectionsResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListCollectionsResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetRemaining

`func (o *ListCollectionsResponse) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *ListCollectionsResponse) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *ListCollectionsResponse) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### GetResult

`func (o *ListCollectionsResponse) GetResult() []Collection`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListCollectionsResponse) GetResultOk() (*[]Collection, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListCollectionsResponse) SetResult(v []Collection)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


