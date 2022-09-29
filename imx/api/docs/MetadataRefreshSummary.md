# MetadataRefreshSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failed** | Pointer to **int32** | The number of tokens with failed metadata refreshes | [optional] 
**Pending** | Pointer to **int32** | The number of tokens that has not been refreshed yet | [optional] 
**Succeeded** | Pointer to **int32** | The number of tokens with successful metadata refreshes | [optional] 

## Methods

### NewMetadataRefreshSummary

`func NewMetadataRefreshSummary() *MetadataRefreshSummary`

NewMetadataRefreshSummary instantiates a new MetadataRefreshSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataRefreshSummaryWithDefaults

`func NewMetadataRefreshSummaryWithDefaults() *MetadataRefreshSummary`

NewMetadataRefreshSummaryWithDefaults instantiates a new MetadataRefreshSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailed

`func (o *MetadataRefreshSummary) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *MetadataRefreshSummary) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *MetadataRefreshSummary) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *MetadataRefreshSummary) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetPending

`func (o *MetadataRefreshSummary) GetPending() int32`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *MetadataRefreshSummary) GetPendingOk() (*int32, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *MetadataRefreshSummary) SetPending(v int32)`

SetPending sets Pending field to given value.

### HasPending

`func (o *MetadataRefreshSummary) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetSucceeded

`func (o *MetadataRefreshSummary) GetSucceeded() int32`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *MetadataRefreshSummary) GetSucceededOk() (*int32, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *MetadataRefreshSummary) SetSucceeded(v int32)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *MetadataRefreshSummary) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


