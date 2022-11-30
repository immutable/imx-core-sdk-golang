# CreateExchangeAndURLAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | Pointer to **string** | Provider name | [optional] 
**Type** | Pointer to **string** | Transaction type | [optional] 
**WalletAddress** | Pointer to **string** | Ethereum address of the user who wants to create transaction | [optional] 
**Widget** | Pointer to [**WidgetParams**](WidgetParams.md) |  | [optional] 

## Methods

### NewCreateExchangeAndURLAPIRequest

`func NewCreateExchangeAndURLAPIRequest() *CreateExchangeAndURLAPIRequest`

NewCreateExchangeAndURLAPIRequest instantiates a new CreateExchangeAndURLAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateExchangeAndURLAPIRequestWithDefaults

`func NewCreateExchangeAndURLAPIRequestWithDefaults() *CreateExchangeAndURLAPIRequest`

NewCreateExchangeAndURLAPIRequestWithDefaults instantiates a new CreateExchangeAndURLAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *CreateExchangeAndURLAPIRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CreateExchangeAndURLAPIRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CreateExchangeAndURLAPIRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CreateExchangeAndURLAPIRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetType

`func (o *CreateExchangeAndURLAPIRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateExchangeAndURLAPIRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateExchangeAndURLAPIRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateExchangeAndURLAPIRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetWalletAddress

`func (o *CreateExchangeAndURLAPIRequest) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *CreateExchangeAndURLAPIRequest) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *CreateExchangeAndURLAPIRequest) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *CreateExchangeAndURLAPIRequest) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.

### GetWidget

`func (o *CreateExchangeAndURLAPIRequest) GetWidget() WidgetParams`

GetWidget returns the Widget field if non-nil, zero value otherwise.

### GetWidgetOk

`func (o *CreateExchangeAndURLAPIRequest) GetWidgetOk() (*WidgetParams, bool)`

GetWidgetOk returns a tuple with the Widget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidget

`func (o *CreateExchangeAndURLAPIRequest) SetWidget(v WidgetParams)`

SetWidget sets Widget field to given value.

### HasWidget

`func (o *CreateExchangeAndURLAPIRequest) HasWidget() bool`

HasWidget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


