# ListTokensResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cursor** | **string** | Generated cursor returned by previous query | 
**Result** | [**[]TokenDetails**](TokenDetails.md) | Tokens matching query parameters | 

## Methods

### NewListTokensResponse

`func NewListTokensResponse(cursor string, result []TokenDetails, ) *ListTokensResponse`

NewListTokensResponse instantiates a new ListTokensResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTokensResponseWithDefaults

`func NewListTokensResponseWithDefaults() *ListTokensResponse`

NewListTokensResponseWithDefaults instantiates a new ListTokensResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCursor

`func (o *ListTokensResponse) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *ListTokensResponse) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *ListTokensResponse) SetCursor(v string)`

SetCursor sets Cursor field to given value.


### GetResult

`func (o *ListTokensResponse) GetResult() []TokenDetails`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *ListTokensResponse) GetResultOk() (*[]TokenDetails, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *ListTokensResponse) SetResult(v []TokenDetails)`

SetResult sets Result field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


