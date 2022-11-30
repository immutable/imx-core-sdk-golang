# NftprimarytransactionCreateAPIRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | Pointer to **string** | Contract address of the asset to be created | [optional] 
**OfferId** | Pointer to **string** | Temporary asset id | [optional] 
**Provider** | Pointer to **string** | Provider name | [optional] 
**UserWalletAddress** | Pointer to **string** | Ethereum address of the user who wants to create transaction | [optional] 
**Widget** | Pointer to [**NftprimarytransactionWidgetParams**](NftprimarytransactionWidgetParams.md) |  | [optional] 

## Methods

### NewNftprimarytransactionCreateAPIRequest

`func NewNftprimarytransactionCreateAPIRequest() *NftprimarytransactionCreateAPIRequest`

NewNftprimarytransactionCreateAPIRequest instantiates a new NftprimarytransactionCreateAPIRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftprimarytransactionCreateAPIRequestWithDefaults

`func NewNftprimarytransactionCreateAPIRequestWithDefaults() *NftprimarytransactionCreateAPIRequest`

NewNftprimarytransactionCreateAPIRequestWithDefaults instantiates a new NftprimarytransactionCreateAPIRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *NftprimarytransactionCreateAPIRequest) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *NftprimarytransactionCreateAPIRequest) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *NftprimarytransactionCreateAPIRequest) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *NftprimarytransactionCreateAPIRequest) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetOfferId

`func (o *NftprimarytransactionCreateAPIRequest) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *NftprimarytransactionCreateAPIRequest) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *NftprimarytransactionCreateAPIRequest) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *NftprimarytransactionCreateAPIRequest) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetProvider

`func (o *NftprimarytransactionCreateAPIRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NftprimarytransactionCreateAPIRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NftprimarytransactionCreateAPIRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NftprimarytransactionCreateAPIRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetUserWalletAddress

`func (o *NftprimarytransactionCreateAPIRequest) GetUserWalletAddress() string`

GetUserWalletAddress returns the UserWalletAddress field if non-nil, zero value otherwise.

### GetUserWalletAddressOk

`func (o *NftprimarytransactionCreateAPIRequest) GetUserWalletAddressOk() (*string, bool)`

GetUserWalletAddressOk returns a tuple with the UserWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserWalletAddress

`func (o *NftprimarytransactionCreateAPIRequest) SetUserWalletAddress(v string)`

SetUserWalletAddress sets UserWalletAddress field to given value.

### HasUserWalletAddress

`func (o *NftprimarytransactionCreateAPIRequest) HasUserWalletAddress() bool`

HasUserWalletAddress returns a boolean if a field has been set.

### GetWidget

`func (o *NftprimarytransactionCreateAPIRequest) GetWidget() NftprimarytransactionWidgetParams`

GetWidget returns the Widget field if non-nil, zero value otherwise.

### GetWidgetOk

`func (o *NftprimarytransactionCreateAPIRequest) GetWidgetOk() (*NftprimarytransactionWidgetParams, bool)`

GetWidgetOk returns a tuple with the Widget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidget

`func (o *NftprimarytransactionCreateAPIRequest) SetWidget(v NftprimarytransactionWidgetParams)`

SetWidget sets Widget field to given value.

### HasWidget

`func (o *NftprimarytransactionCreateAPIRequest) HasWidget() bool`

HasWidget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


