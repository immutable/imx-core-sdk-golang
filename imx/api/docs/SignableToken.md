# SignableToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Token data. See https://docs.x.immutable.com/docs/token-data-object | [optional] 
**Type** | Pointer to **string** | Type of token | [optional] 

## Methods

### NewSignableToken

`func NewSignableToken() *SignableToken`

NewSignableToken instantiates a new SignableToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignableTokenWithDefaults

`func NewSignableTokenWithDefaults() *SignableToken`

NewSignableTokenWithDefaults instantiates a new SignableToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SignableToken) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SignableToken) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SignableToken) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SignableToken) HasData() bool`

HasData returns a boolean if a field has been set.

### GetType

`func (o *SignableToken) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignableToken) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignableToken) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SignableToken) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


