# GetMetadataRefreshErrorsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Generated cursor returned by previous query | 
**Remaining** | **int32** | Remaining results flag. 1: there are remaining results matching this query, 0: no remaining results | 
**Result** | [**[]MetadataRefreshErrors**](MetadataRefreshErrors.md) | Metadata refresh errors matching query parameters | 

## Methods

### NewGetMetadataRefreshErrorsResponse

`func NewGetMetadataRefreshErrorsResponse(cursor string, remaining int32, result []MetadataRefreshErrors, ) *GetMetadataRefreshErrorsResponse`

NewGetMetadataRefreshErrorsResponse instantiates a new GetMetadataRefreshErrorsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetadataRefreshErrorsResponseWithDefaults

`func NewGetMetadataRefreshErrorsResponseWithDefaults() *GetMetadataRefreshErrorsResponse`

NewGetMetadataRefreshErrorsResponseWithDefaults instantiates a new GetMetadataRefreshErrorsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *GetMetadataRefreshErrorsResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetMetadataRefreshErrorsResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetMetadataRefreshErrorsResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetRemaining

`func (o *GetMetadataRefreshErrorsResponse) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *GetMetadataRefreshErrorsResponse) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *GetMetadataRefreshErrorsResponse) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.


### GetResult

`func (o *GetMetadataRefreshErrorsResponse) GetResult() []MetadataRefreshErrors`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *GetMetadataRefreshErrorsResponse) GetResultOk() (*[]MetadataRefreshErrors, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *GetMetadataRefreshErrorsResponse) SetResult(v []MetadataRefreshErrors)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


