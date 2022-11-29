# ExchangeCreateExchangeAndURLResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Created transaction ID | [optional] 
**Provider** | Pointer to **string** | Provider name (e.g. moonpay) | [optional] 
**Type** | Pointer to **string** | Transaction type | [optional] 
**Url** | Pointer to **string** | Widget URL | [optional] 
**WalletAddress** | Pointer to **string** | Ethereum address of the user who created transaction | [optional] 

## Methods

### NewExchangeCreateExchangeAndURLResponse

`func NewExchangeCreateExchangeAndURLResponse() *ExchangeCreateExchangeAndURLResponse`

NewExchangeCreateExchangeAndURLResponse instantiates a new ExchangeCreateExchangeAndURLResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExchangeCreateExchangeAndURLResponseWithDefaults

`func NewExchangeCreateExchangeAndURLResponseWithDefaults() *ExchangeCreateExchangeAndURLResponse`

NewExchangeCreateExchangeAndURLResponseWithDefaults instantiates a new ExchangeCreateExchangeAndURLResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExchangeCreateExchangeAndURLResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExchangeCreateExchangeAndURLResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExchangeCreateExchangeAndURLResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ExchangeCreateExchangeAndURLResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *ExchangeCreateExchangeAndURLResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ExchangeCreateExchangeAndURLResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ExchangeCreateExchangeAndURLResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ExchangeCreateExchangeAndURLResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetType

`func (o *ExchangeCreateExchangeAndURLResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExchangeCreateExchangeAndURLResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExchangeCreateExchangeAndURLResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExchangeCreateExchangeAndURLResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *ExchangeCreateExchangeAndURLResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExchangeCreateExchangeAndURLResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExchangeCreateExchangeAndURLResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ExchangeCreateExchangeAndURLResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetWalletAddress

`func (o *ExchangeCreateExchangeAndURLResponse) GetWalletAddress() string`

GetWalletAddress returns the WalletAddress field if non-nil, zero value otherwise.

### GetWalletAddressOk

`func (o *ExchangeCreateExchangeAndURLResponse) GetWalletAddressOk() (*string, bool)`

GetWalletAddressOk returns a tuple with the WalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddress

`func (o *ExchangeCreateExchangeAndURLResponse) SetWalletAddress(v string)`

SetWalletAddress sets WalletAddress field to given value.

### HasWalletAddress

`func (o *ExchangeCreateExchangeAndURLResponse) HasWalletAddress() bool`

HasWalletAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


