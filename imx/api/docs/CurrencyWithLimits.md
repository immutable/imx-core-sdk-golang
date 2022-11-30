# CurrencyWithLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | Provider currency code | [optional] 
**Limits** | Pointer to [**AggregateLimit**](AggregateLimit.md) |  | [optional] 
**Provider** | Pointer to **string** | Provider name (e.g. moonpay) | [optional] 
**Symbol** | Pointer to **string** | Currency Symbol | [optional] 

## Methods

### NewCurrencyWithLimits

`func NewCurrencyWithLimits() *CurrencyWithLimits`

NewCurrencyWithLimits instantiates a new CurrencyWithLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrencyWithLimitsWithDefaults

`func NewCurrencyWithLimitsWithDefaults() *CurrencyWithLimits`

NewCurrencyWithLimitsWithDefaults instantiates a new CurrencyWithLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *CurrencyWithLimits) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CurrencyWithLimits) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CurrencyWithLimits) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *CurrencyWithLimits) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetLimits

`func (o *CurrencyWithLimits) GetLimits() AggregateLimit`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *CurrencyWithLimits) GetLimitsOk() (*AggregateLimit, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *CurrencyWithLimits) SetLimits(v AggregateLimit)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *CurrencyWithLimits) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetProvider

`func (o *CurrencyWithLimits) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CurrencyWithLimits) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CurrencyWithLimits) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CurrencyWithLimits) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSymbol

`func (o *CurrencyWithLimits) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CurrencyWithLimits) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CurrencyWithLimits) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CurrencyWithLimits) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


