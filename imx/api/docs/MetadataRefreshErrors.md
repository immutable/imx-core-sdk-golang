# MetadataRefreshErrors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientResponseBody** | **NullableString** | Metadata API response for the token | 
**ClientResponseStatusCode** | **NullableString** | Metadata API response code for the token | 
**ClientTokenMetadataUrl** | **string** | Requested metadata url for the token | 
**CollectionAddress** | **string** | The collection contract address | 
**CreatedAt** | **string** | When the error was created | 
**ErrorCode** | **string** | Metadata refresh error code | 
**TokenId** | **string** | The token ID | 

## Methods

### NewMetadataRefreshErrors

`func NewMetadataRefreshErrors(clientResponseBody NullableString, clientResponseStatusCode NullableString, clientTokenMetadataUrl string, collectionAddress string, createdAt string, errorCode string, tokenId string, ) *MetadataRefreshErrors`

NewMetadataRefreshErrors instantiates a new MetadataRefreshErrors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataRefreshErrorsWithDefaults

`func NewMetadataRefreshErrorsWithDefaults() *MetadataRefreshErrors`

NewMetadataRefreshErrorsWithDefaults instantiates a new MetadataRefreshErrors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientResponseBody

`func (o *MetadataRefreshErrors) GetClientResponseBody() string`

GetClientResponseBody returns the ClientResponseBody field if non-nil, zero value otherwise.

### GetClientResponseBodyOk

`func (o *MetadataRefreshErrors) GetClientResponseBodyOk() (*string, bool)`

GetClientResponseBodyOk returns a tuple with the ClientResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientResponseBody

`func (o *MetadataRefreshErrors) SetClientResponseBody(v string)`

SetClientResponseBody sets ClientResponseBody field to given value.


### SetClientResponseBodyNil

`func (o *MetadataRefreshErrors) SetClientResponseBodyNil(b bool)`

 SetClientResponseBodyNil sets the value for ClientResponseBody to be an explicit nil

### UnsetClientResponseBody
`func (o *MetadataRefreshErrors) UnsetClientResponseBody()`

UnsetClientResponseBody ensures that no value is present for ClientResponseBody, not even an explicit nil
### GetClientResponseStatusCode

`func (o *MetadataRefreshErrors) GetClientResponseStatusCode() string`

GetClientResponseStatusCode returns the ClientResponseStatusCode field if non-nil, zero value otherwise.

### GetClientResponseStatusCodeOk

`func (o *MetadataRefreshErrors) GetClientResponseStatusCodeOk() (*string, bool)`

GetClientResponseStatusCodeOk returns a tuple with the ClientResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientResponseStatusCode

`func (o *MetadataRefreshErrors) SetClientResponseStatusCode(v string)`

SetClientResponseStatusCode sets ClientResponseStatusCode field to given value.


### SetClientResponseStatusCodeNil

`func (o *MetadataRefreshErrors) SetClientResponseStatusCodeNil(b bool)`

 SetClientResponseStatusCodeNil sets the value for ClientResponseStatusCode to be an explicit nil

### UnsetClientResponseStatusCode
`func (o *MetadataRefreshErrors) UnsetClientResponseStatusCode()`

UnsetClientResponseStatusCode ensures that no value is present for ClientResponseStatusCode, not even an explicit nil
### GetClientTokenMetadataUrl

`func (o *MetadataRefreshErrors) GetClientTokenMetadataUrl() string`

GetClientTokenMetadataUrl returns the ClientTokenMetadataUrl field if non-nil, zero value otherwise.

### GetClientTokenMetadataUrlOk

`func (o *MetadataRefreshErrors) GetClientTokenMetadataUrlOk() (*string, bool)`

GetClientTokenMetadataUrlOk returns a tuple with the ClientTokenMetadataUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTokenMetadataUrl

`func (o *MetadataRefreshErrors) SetClientTokenMetadataUrl(v string)`

SetClientTokenMetadataUrl sets ClientTokenMetadataUrl field to given value.


### GetCollectionAddress

`func (o *MetadataRefreshErrors) GetCollectionAddress() string`

GetCollectionAddress returns the CollectionAddress field if non-nil, zero value otherwise.

### GetCollectionAddressOk

`func (o *MetadataRefreshErrors) GetCollectionAddressOk() (*string, bool)`

GetCollectionAddressOk returns a tuple with the CollectionAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionAddress

`func (o *MetadataRefreshErrors) SetCollectionAddress(v string)`

SetCollectionAddress sets CollectionAddress field to given value.


### GetCreatedAt

`func (o *MetadataRefreshErrors) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetadataRefreshErrors) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetadataRefreshErrors) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetErrorCode

`func (o *MetadataRefreshErrors) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *MetadataRefreshErrors) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *MetadataRefreshErrors) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetTokenId

`func (o *MetadataRefreshErrors) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MetadataRefreshErrors) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MetadataRefreshErrors) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


