# GetMetadataRefreshResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionAddress** | **string** | The collection address | 
**CompletedAt** | Pointer to **NullableString** | When the metadata refresh completed | [optional] 
**RefreshId** | **string** | The metadata refresh ID | 
**StartedAt** | **string** | When the metadata refresh started | 
**Status** | **string** | The metadata refresh status | 
**Summary** | [**[]MetadataRefreshSummary**](MetadataRefreshSummary.md) | The current metadata refresh summary. The summary continue to update until metadata refresh is completed | 

## Methods

### NewGetMetadataRefreshResponse

`func NewGetMetadataRefreshResponse(collectionAddress string, refreshId string, startedAt string, status string, summary []MetadataRefreshSummary, ) *GetMetadataRefreshResponse`

NewGetMetadataRefreshResponse instantiates a new GetMetadataRefreshResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMetadataRefreshResponseWithDefaults

`func NewGetMetadataRefreshResponseWithDefaults() *GetMetadataRefreshResponse`

NewGetMetadataRefreshResponseWithDefaults instantiates a new GetMetadataRefreshResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionAddress

`func (o *GetMetadataRefreshResponse) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *GetMetadataRefreshResponse) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *GetMetadataRefreshResponse) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetCompletedAt

`func (o *GetMetadataRefreshResponse) GetCompletedAt() string`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetMetadataRefreshResponse) GetCompletedAtOk() (*string, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetMetadataRefreshResponse) SetCompletedAt(v string)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetMetadataRefreshResponse) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetMetadataRefreshResponse) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetMetadataRefreshResponse) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetRefreshId

`func (o *GetMetadataRefreshResponse) GetRefreshId() string`

GetRefreshId returns the RefreshId field if non-nil, zero value otherwise.

### GetRefreshIdOk

`func (o *GetMetadataRefreshResponse) GetRefreshIdOk() (*string, bool)`

GetRefreshIdOk returns a tuple with the RefreshId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshId

`func (o *GetMetadataRefreshResponse) SetRefreshId(v string)`

SetRefreshId sets RefreshId field to given value.


### GetStartedAt

`func (o *GetMetadataRefreshResponse) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetMetadataRefreshResponse) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetMetadataRefreshResponse) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetStatus

`func (o *GetMetadataRefreshResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMetadataRefreshResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMetadataRefreshResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSummary

`func (o *GetMetadataRefreshResponse) GetSummary() []MetadataRefreshSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetMetadataRefreshResponse) GetSummaryOk() (*[]MetadataRefreshSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetMetadataRefreshResponse) SetSummary(v []MetadataRefreshSummary)`

SetSummary sets Summary field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


