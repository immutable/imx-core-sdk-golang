# OptionalExchangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | Pointer to **string** | Provider transaction ID | [optional] 
**FeesAmount** | Pointer to **float32** | Fees amount | [optional] 
**FromAmount** | Pointer to **float32** | Amount that was exchanged from | [optional] 
**FromCurrency** | Pointer to **string** | Currency that was exchanged from | [optional] 
**ProviderWalletAddress** | Pointer to **string** | Provider wallet address that was used for transferring crypto | [optional] 
**ToAmount** | Pointer to **float32** | Amount that was exchanged to | [optional] 
**ToCurrency** | Pointer to **string** | Currency that was exchanged to | [optional] 
**TransferId** | Pointer to **string** | Transfer ID | [optional] 

## Methods

### NewOptionalExchangeData

`func NewOptionalExchangeData() *OptionalExchangeData`

NewOptionalExchangeData instantiates a new OptionalExchangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionalExchangeDataWithDefaults

`func NewOptionalExchangeDataWithDefaults() *OptionalExchangeData`

NewOptionalExchangeDataWithDefaults instantiates a new OptionalExchangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *OptionalExchangeData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *OptionalExchangeData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *OptionalExchangeData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *OptionalExchangeData) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFeesAmount

`func (o *OptionalExchangeData) GetFeesAmount() float32`

GetFeesAmount returns the FeesAmount field if non-nil, zero value otherwise.

### GetFeesAmountOk

`func (o *OptionalExchangeData) GetFeesAmountOk() (*float32, bool)`

GetFeesAmountOk returns a tuple with the FeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmount

`func (o *OptionalExchangeData) SetFeesAmount(v float32)`

SetFeesAmount sets FeesAmount field to given value.

### HasFeesAmount

`func (o *OptionalExchangeData) HasFeesAmount() bool`

HasFeesAmount returns a boolean if a field has been set.

### GetFromAmount

`func (o *OptionalExchangeData) GetFromAmount() float32`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *OptionalExchangeData) GetFromAmountOk() (*float32, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *OptionalExchangeData) SetFromAmount(v float32)`

SetFromAmount sets FromAmount field to given value.

### HasFromAmount

`func (o *OptionalExchangeData) HasFromAmount() bool`

HasFromAmount returns a boolean if a field has been set.

### GetFromCurrency

`func (o *OptionalExchangeData) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *OptionalExchangeData) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *OptionalExchangeData) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.

### HasFromCurrency

`func (o *OptionalExchangeData) HasFromCurrency() bool`

HasFromCurrency returns a boolean if a field has been set.

### GetProviderWalletAddress

`func (o *OptionalExchangeData) GetProviderWalletAddress() string`

GetProviderWalletAddress returns the ProviderWalletAddress field if non-nil, zero value otherwise.

### GetProviderWalletAddressOk

`func (o *OptionalExchangeData) GetProviderWalletAddressOk() (*string, bool)`

GetProviderWalletAddressOk returns a tuple with the ProviderWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderWalletAddress

`func (o *OptionalExchangeData) SetProviderWalletAddress(v string)`

SetProviderWalletAddress sets ProviderWalletAddress field to given value.

### HasProviderWalletAddress

`func (o *OptionalExchangeData) HasProviderWalletAddress() bool`

HasProviderWalletAddress returns a boolean if a field has been set.

### GetToAmount

`func (o *OptionalExchangeData) GetToAmount() float32`

GetToAmount returns the ToAmount field if non-nil, zero value otherwise.

### GetToAmountOk

`func (o *OptionalExchangeData) GetToAmountOk() (*float32, bool)`

GetToAmountOk returns a tuple with the ToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmount

`func (o *OptionalExchangeData) SetToAmount(v float32)`

SetToAmount sets ToAmount field to given value.

### HasToAmount

`func (o *OptionalExchangeData) HasToAmount() bool`

HasToAmount returns a boolean if a field has been set.

### GetToCurrency

`func (o *OptionalExchangeData) GetToCurrency() string`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *OptionalExchangeData) GetToCurrencyOk() (*string, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *OptionalExchangeData) SetToCurrency(v string)`

SetToCurrency sets ToCurrency field to given value.

### HasToCurrency

`func (o *OptionalExchangeData) HasToCurrency() bool`

HasToCurrency returns a boolean if a field has been set.

### GetTransferId

`func (o *OptionalExchangeData) GetTransferId() string`

GetTransferId returns the TransferId field if non-nil, zero value otherwise.

### GetTransferIdOk

`func (o *OptionalExchangeData) GetTransferIdOk() (*string, bool)`

GetTransferIdOk returns a tuple with the TransferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferId

`func (o *OptionalExchangeData) SetTransferId(v string)`

SetTransferId sets TransferId field to given value.

### HasTransferId

`func (o *OptionalExchangeData) HasTransferId() bool`

HasTransferId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


