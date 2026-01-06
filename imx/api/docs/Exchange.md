# Exchange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **string** | Time this transaction was created | [optional] 
**Data** | Pointer to [**OptionalExchangeData**](OptionalExchangeData.md) |  | [optional] 
**Id** | Pointer to **int32** | Transaction ID | [optional] 
**Provider** | Pointer to **string** | Provider name | [optional] 
**Status** | Pointer to **string** | Transaction status | [optional] 
**Type** | Pointer to **string** | Transaction type | [optional] 
**UpdatedAt** | Pointer to **string** | Time this transaction was updated | [optional] 
**WalletAddress** | Pointer to **string** | Ethereum address of the user who created transaction | [optional] 

## Methods

### NewExchange

`func NewExchange() *Exchange`

NewExchange instantiates a new Exchange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeWithDefaults

`func NewExchangeWithDefaults() *Exchange`

NewExchangeWithDefaults instantiates a new Exchange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Exchange) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Exchange) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Exchange) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Exchange) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetData

`func (o *Exchange) GetData() OptionalExchangeData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Exchange) GetDataOk() (*OptionalExchangeData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Exchange) SetData(v OptionalExchangeData)`

SetData sets Data field to given value.

### HasData

`func (o *Exchange) HasData() bool`

HasData returns a boolean if a field has been set.

### GetId

`func (o *Exchange) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Exchange) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Exchange) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Exchange) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *Exchange) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Exchange) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Exchange) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Exchange) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *Exchange) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Exchange) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Exchange) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Exchange) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Exchange) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Exchange) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Exchange) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Exchange) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Exchange) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Exchange) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Exchange) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Exchange) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetWalletAddress

`func (o *Exchange) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *Exchange) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *Exchange) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *Exchange) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


