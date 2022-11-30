# ProviderGetMintStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ProviderMintStatus**](ProviderMintStatus.md) | Mint status for requested transactions | [optional] 

## Methods

### NewProviderGetMintStatusResponse

`func NewProviderGetMintStatusResponse() *ProviderGetMintStatusResponse`

NewProviderGetMintStatusResponse instantiates a new ProviderGetMintStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGetMintStatusResponseWithDefaults

`func NewProviderGetMintStatusResponseWithDefaults() *ProviderGetMintStatusResponse`

NewProviderGetMintStatusResponseWithDefaults instantiates a new ProviderGetMintStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ProviderGetMintStatusResponse) GetData() []ProviderMintStatus`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ProviderGetMintStatusResponse) GetDataOk() (*[]ProviderMintStatus, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ProviderGetMintStatusResponse) SetData(v []ProviderMintStatus)`

SetData sets Data field to given value.

### HasData

`func (o *ProviderGetMintStatusResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


