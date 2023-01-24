# SignableTransferDetailsToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Token data. See https://docs.x.immutable.com/docs/token-data-object | [optional] 
**Type** | Pointer to **string** | Type of token | [optional] 

## Methods

### NewSignableTransferDetailsToken

`func NewSignableTransferDetailsToken() *SignableTransferDetailsToken`

NewSignableTransferDetailsToken instantiates a new SignableTransferDetailsToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignableTransferDetailsTokenWithDefaults

`func NewSignableTransferDetailsTokenWithDefaults() *SignableTransferDetailsToken`

NewSignableTransferDetailsTokenWithDefaults instantiates a new SignableTransferDetailsToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SignableTransferDetailsToken) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignableTransferDetailsToken) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignableTransferDetailsToken) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SignableTransferDetailsToken) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *SignableTransferDetailsToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignableTransferDetailsToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignableTransferDetailsToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SignableTransferDetailsToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


