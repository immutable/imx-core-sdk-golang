# FeeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | **string** | ID of the asset these fees relate to | 
**FeeLimit** | **string** | Fee limit | 
**SourceVaultId** | **int32** | ID of vault the asset belong to | 

## Methods

### NewFeeInfo

`func NewFeeInfo(assetId string, feeLimit string, sourceVaultId int32, ) *FeeInfo`

NewFeeInfo instantiates a new FeeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeInfoWithDefaults

`func NewFeeInfoWithDefaults() *FeeInfo`

NewFeeInfoWithDefaults instantiates a new FeeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *FeeInfo) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *FeeInfo) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *FeeInfo) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.


### GetFeeLimit

`func (o *FeeInfo) GetFeeLimit() string`

GetFeeLimit returns the FeeLimit field if non-nil, zero value otherwise.

### GetFeeLimitOk

`func (o *FeeInfo) GetFeeLimitOk() (*string, bool)`

GetFeeLimitOk returns a tuple with the FeeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeLimit

`func (o *FeeInfo) SetFeeLimit(v string)`

SetFeeLimit sets FeeLimit field to given value.


### GetSourceVaultId

`func (o *FeeInfo) GetSourceVaultId() int32`

GetSourceVaultId returns the SourceVaultId field if non-nil, zero value otherwise.

### GetSourceVaultIdOk

`func (o *FeeInfo) GetSourceVaultIdOk() (*int32, bool)`

GetSourceVaultIdOk returns a tuple with the SourceVaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVaultId

`func (o *FeeInfo) SetSourceVaultId(v int32)`

SetSourceVaultId sets SourceVaultId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


