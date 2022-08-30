# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectionLimitExpiresAt** | **string** | The current period expiry date for collection limit | 
**CollectionMonthlyLimit** | **int32** | The total monthly collection limit | 
**CollectionRemaining** | **int32** | The number of collection remaining in the current period | 
**CompanyName** | **string** | The company name | 
**ContactEmail** | **string** | The project contact email | 
**Id** | **int32** | The project ID | 
**MintLimitExpiresAt** | **string** | The current period expiry date for mint operation limit | 
**MintMonthlyLimit** | **int32** | The total monthly mint operation limit | 
**MintRemaining** | **int32** | The number of mint operation remaining in the current period | 
**Name** | **string** | The project name | 

## Methods

### NewProject

`func NewProject(collectionLimitExpiresAt string, collectionMonthlyLimit int32, collectionRemaining int32, companyName string, contactEmail string, id int32, mintLimitExpiresAt string, mintMonthlyLimit int32, mintRemaining int32, name string, ) *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectionLimitExpiresAt

`func (o *Project) GetCollectionLimitExpiresAt() string`

GetCollectionLimitExpiresAt returns the CollectionLimitExpiresAt field if non-nil, zero value otherwise.

### GetCollectionLimitExpiresAtOk

`func (o *Project) GetCollectionLimitExpiresAtOk() (*string, bool)`

GetCollectionLimitExpiresAtOk returns a tuple with the CollectionLimitExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionLimitExpiresAt

`func (o *Project) SetCollectionLimitExpiresAt(v string)`

SetCollectionLimitExpiresAt sets CollectionLimitExpiresAt field to given value.


### GetCollectionMonthlyLimit

`func (o *Project) GetCollectionMonthlyLimit() int32`

GetCollectionMonthlyLimit returns the CollectionMonthlyLimit field if non-nil, zero value otherwise.

### GetCollectionMonthlyLimitOk

`func (o *Project) GetCollectionMonthlyLimitOk() (*int32, bool)`

GetCollectionMonthlyLimitOk returns a tuple with the CollectionMonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionMonthlyLimit

`func (o *Project) SetCollectionMonthlyLimit(v int32)`

SetCollectionMonthlyLimit sets CollectionMonthlyLimit field to given value.


### GetCollectionRemaining

`func (o *Project) GetCollectionRemaining() int32`

GetCollectionRemaining returns the CollectionRemaining field if non-nil, zero value otherwise.

### GetCollectionRemainingOk

`func (o *Project) GetCollectionRemainingOk() (*int32, bool)`

GetCollectionRemainingOk returns a tuple with the CollectionRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionRemaining

`func (o *Project) SetCollectionRemaining(v int32)`

SetCollectionRemaining sets CollectionRemaining field to given value.


### GetCompanyName

`func (o *Project) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *Project) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *Project) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.


### GetContactEmail

`func (o *Project) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *Project) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *Project) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetId

`func (o *Project) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Project) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Project) SetId(v int32)`

SetId sets Id field to given value.


### GetMintLimitExpiresAt

`func (o *Project) GetMintLimitExpiresAt() string`

GetMintLimitExpiresAt returns the MintLimitExpiresAt field if non-nil, zero value otherwise.

### GetMintLimitExpiresAtOk

`func (o *Project) GetMintLimitExpiresAtOk() (*string, bool)`

GetMintLimitExpiresAtOk returns a tuple with the MintLimitExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintLimitExpiresAt

`func (o *Project) SetMintLimitExpiresAt(v string)`

SetMintLimitExpiresAt sets MintLimitExpiresAt field to given value.


### GetMintMonthlyLimit

`func (o *Project) GetMintMonthlyLimit() int32`

GetMintMonthlyLimit returns the MintMonthlyLimit field if non-nil, zero value otherwise.

### GetMintMonthlyLimitOk

`func (o *Project) GetMintMonthlyLimitOk() (*int32, bool)`

GetMintMonthlyLimitOk returns a tuple with the MintMonthlyLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintMonthlyLimit

`func (o *Project) SetMintMonthlyLimit(v int32)`

SetMintMonthlyLimit sets MintMonthlyLimit field to given value.


### GetMintRemaining

`func (o *Project) GetMintRemaining() int32`

GetMintRemaining returns the MintRemaining field if non-nil, zero value otherwise.

### GetMintRemainingOk

`func (o *Project) GetMintRemainingOk() (*int32, bool)`

GetMintRemainingOk returns a tuple with the MintRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintRemaining

`func (o *Project) SetMintRemaining(v int32)`

SetMintRemaining sets MintRemaining field to given value.


### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


