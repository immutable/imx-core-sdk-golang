# WidgetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | Amount that will be prefilled in the widget | [optional] 
**SupportedCurrencies** | Pointer to **[]string** | Currencies that will be available in the widget. If not defined all available currencies will be shown | [optional] 
**Theme** | Pointer to **string** | Widget theme dark by default | [optional] 

## Methods

### NewWidgetParams

`func NewWidgetParams() *WidgetParams`

NewWidgetParams instantiates a new WidgetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWidgetParamsWithDefaults

`func NewWidgetParamsWithDefaults() *WidgetParams`

NewWidgetParamsWithDefaults instantiates a new WidgetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WidgetParams) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WidgetParams) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WidgetParams) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *WidgetParams) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetSupportedCurrencies

`func (o *WidgetParams) GetSupportedCurrencies() []string`

GetSupportedCurrencies returns the SupportedCurrencies field if non-nil, zero value otherwise.

### GetSupportedCurrenciesOk

`func (o *WidgetParams) GetSupportedCurrenciesOk() (*[]string, bool)`

GetSupportedCurrenciesOk returns a tuple with the SupportedCurrencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCurrencies

`func (o *WidgetParams) SetSupportedCurrencies(v []string)`

SetSupportedCurrencies sets SupportedCurrencies field to given value.

### HasSupportedCurrencies

`func (o *WidgetParams) HasSupportedCurrencies() bool`

HasSupportedCurrencies returns a boolean if a field has been set.

### GetTheme

`func (o *WidgetParams) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *WidgetParams) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *WidgetParams) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *WidgetParams) HasTheme() bool`

HasTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


