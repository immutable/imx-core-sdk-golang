# NftprimarytransactionCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | Pointer to **string** | Contract address of the asset to be created | [optional] 
**OfferId** | Pointer to **string** | Temporary asset id | [optional] 
**Provider** | Pointer to **string** | Provider name | [optional] 
**SellerWalletAddress** | Pointer to **string** | Ethereum address of the seller | [optional] 
**TokenId** | Pointer to **string** | Asset id | [optional] 
**TransactionId** | Pointer to **string** | Transaction id | [optional] 
**Url** | Pointer to **string** | Widget Url signed by provider | [optional] 
**UserWalletAddress** | Pointer to **string** | Ethereum address of the user who wants to create transaction | [optional] 

## Methods

### NewNftprimarytransactionCreateResponse

`func NewNftprimarytransactionCreateResponse() *NftprimarytransactionCreateResponse`

NewNftprimarytransactionCreateResponse instantiates a new NftprimarytransactionCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftprimarytransactionCreateResponseWithDefaults

`func NewNftprimarytransactionCreateResponseWithDefaults() *NftprimarytransactionCreateResponse`

NewNftprimarytransactionCreateResponseWithDefaults instantiates a new NftprimarytransactionCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *NftprimarytransactionCreateResponse) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *NftprimarytransactionCreateResponse) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *NftprimarytransactionCreateResponse) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *NftprimarytransactionCreateResponse) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetOfferId

`func (o *NftprimarytransactionCreateResponse) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *NftprimarytransactionCreateResponse) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *NftprimarytransactionCreateResponse) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *NftprimarytransactionCreateResponse) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetProvider

`func (o *NftprimarytransactionCreateResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NftprimarytransactionCreateResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NftprimarytransactionCreateResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NftprimarytransactionCreateResponse) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSellerWalletAddress

`func (o *NftprimarytransactionCreateResponse) GetSellerWalletAddress() string`

GetSellerWalletAddress returns the SellerWalletAddress field if non-nil, zero value otherwise.

### GetSellerWalletAddressOk

`func (o *NftprimarytransactionCreateResponse) GetSellerWalletAddressOk() (*string, bool)`

GetSellerWalletAddressOk returns a tuple with the SellerWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerWalletAddress

`func (o *NftprimarytransactionCreateResponse) SetSellerWalletAddress(v string)`

SetSellerWalletAddress sets SellerWalletAddress field to given value.

### HasSellerWalletAddress

`func (o *NftprimarytransactionCreateResponse) HasSellerWalletAddress() bool`

HasSellerWalletAddress returns a boolean if a field has been set.

### GetTokenId

`func (o *NftprimarytransactionCreateResponse) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *NftprimarytransactionCreateResponse) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *NftprimarytransactionCreateResponse) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *NftprimarytransactionCreateResponse) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTransactionId

`func (o *NftprimarytransactionCreateResponse) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *NftprimarytransactionCreateResponse) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *NftprimarytransactionCreateResponse) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *NftprimarytransactionCreateResponse) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUrl

`func (o *NftprimarytransactionCreateResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *NftprimarytransactionCreateResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *NftprimarytransactionCreateResponse) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *NftprimarytransactionCreateResponse) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUserWalletAddress

`func (o *NftprimarytransactionCreateResponse) GetUserWalletAddress() string`

GetUserWalletAddress returns the UserWalletAddress field if non-nil, zero value otherwise.

### GetUserWalletAddressOk

`func (o *NftprimarytransactionCreateResponse) GetUserWalletAddressOk() (*string, bool)`

GetUserWalletAddressOk returns a tuple with the UserWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserWalletAddress

`func (o *NftprimarytransactionCreateResponse) SetUserWalletAddress(v string)`

SetUserWalletAddress sets UserWalletAddress field to given value.

### HasUserWalletAddress

`func (o *NftprimarytransactionCreateResponse) HasUserWalletAddress() bool`

HasUserWalletAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


