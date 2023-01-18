# NftprimarytransactionTransactionData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractAddress** | Pointer to **string** | Contract address of the asset | [optional] 
**CreatedAt** | Pointer to **string** | Timestamp when the transaction was created | [optional] 
**ExternalId** | Pointer to **string** | External transaction id | [optional] 
**FeesAmount** | Pointer to **float32** | Fees to pay on this transaction | [optional] 
**FromAmount** | Pointer to **float32** | Amount of the currency specified in &#x60;from_currency&#x60; that the buyer paid for the transaction | [optional] 
**FromCurrency** | Pointer to **string** | Currency that the buyer used for the transaction | [optional] 
**MintId** | Pointer to **string** | Minting transaction ID - see mintTokens response | [optional] 
**MintStatus** | Pointer to **string** | Mint status | [optional] 
**OfferId** | Pointer to **string** | Temporary asset id. Might be a token id if the token id is known or a generic description if it&#39;s not | [optional] 
**Provider** | Pointer to **string** | Provider name | [optional] 
**SellerWalletAddress** | Pointer to **string** | Ethereum address of the seller | [optional] 
**Status** | Pointer to **string** | Transaction status enums(created, waitingPayment, pending, completed, failed) | [optional] 
**ToAmount** | Pointer to **float32** | Amount of the currency specified in &#x60;to_currency&#x60; that the seller received from the checkout provider for the transaction | [optional] 
**ToCurrency** | Pointer to **string** | Currency (crypto) that the checkout provider sent to the seller for the transaction | [optional] 
**TokenId** | Pointer to **string** | ID of the token that has been successfully minted - should be the same as &#x60;offer_id&#x60; | [optional] 
**TransactionId** | Pointer to **string** | Transaction id | [optional] 
**UpdatedAt** | Pointer to **string** | Timestamp when the transaction was updated | [optional] 
**UserWalletAddress** | Pointer to **string** | Wallet address that receives the minted NFT | [optional] 

## Methods

### NewNftprimarytransactionTransactionData

`func NewNftprimarytransactionTransactionData() *NftprimarytransactionTransactionData`

NewNftprimarytransactionTransactionData instantiates a new NftprimarytransactionTransactionData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNftprimarytransactionTransactionDataWithDefaults

`func NewNftprimarytransactionTransactionDataWithDefaults() *NftprimarytransactionTransactionData`

NewNftprimarytransactionTransactionDataWithDefaults instantiates a new NftprimarytransactionTransactionData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractAddress

`func (o *NftprimarytransactionTransactionData) GetContractAddress() string`

GetContractAddress returns the ContractAddress field if non-nil, zero value otherwise.

### GetContractAddressOk

`func (o *NftprimarytransactionTransactionData) GetContractAddressOk() (*string, bool)`

GetContractAddressOk returns a tuple with the ContractAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractAddress

`func (o *NftprimarytransactionTransactionData) SetContractAddress(v string)`

SetContractAddress sets ContractAddress field to given value.

### HasContractAddress

`func (o *NftprimarytransactionTransactionData) HasContractAddress() bool`

HasContractAddress returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NftprimarytransactionTransactionData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NftprimarytransactionTransactionData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NftprimarytransactionTransactionData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NftprimarytransactionTransactionData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *NftprimarytransactionTransactionData) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NftprimarytransactionTransactionData) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NftprimarytransactionTransactionData) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NftprimarytransactionTransactionData) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetFeesAmount

`func (o *NftprimarytransactionTransactionData) GetFeesAmount() float32`

GetFeesAmount returns the FeesAmount field if non-nil, zero value otherwise.

### GetFeesAmountOk

`func (o *NftprimarytransactionTransactionData) GetFeesAmountOk() (*float32, bool)`

GetFeesAmountOk returns a tuple with the FeesAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeesAmount

`func (o *NftprimarytransactionTransactionData) SetFeesAmount(v float32)`

SetFeesAmount sets FeesAmount field to given value.

### HasFeesAmount

`func (o *NftprimarytransactionTransactionData) HasFeesAmount() bool`

HasFeesAmount returns a boolean if a field has been set.

### GetFromAmount

`func (o *NftprimarytransactionTransactionData) GetFromAmount() float32`

GetFromAmount returns the FromAmount field if non-nil, zero value otherwise.

### GetFromAmountOk

`func (o *NftprimarytransactionTransactionData) GetFromAmountOk() (*float32, bool)`

GetFromAmountOk returns a tuple with the FromAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAmount

`func (o *NftprimarytransactionTransactionData) SetFromAmount(v float32)`

SetFromAmount sets FromAmount field to given value.

### HasFromAmount

`func (o *NftprimarytransactionTransactionData) HasFromAmount() bool`

HasFromAmount returns a boolean if a field has been set.

### GetFromCurrency

`func (o *NftprimarytransactionTransactionData) GetFromCurrency() string`

GetFromCurrency returns the FromCurrency field if non-nil, zero value otherwise.

### GetFromCurrencyOk

`func (o *NftprimarytransactionTransactionData) GetFromCurrencyOk() (*string, bool)`

GetFromCurrencyOk returns a tuple with the FromCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCurrency

`func (o *NftprimarytransactionTransactionData) SetFromCurrency(v string)`

SetFromCurrency sets FromCurrency field to given value.

### HasFromCurrency

`func (o *NftprimarytransactionTransactionData) HasFromCurrency() bool`

HasFromCurrency returns a boolean if a field has been set.

### GetMintId

`func (o *NftprimarytransactionTransactionData) GetMintId() string`

GetMintId returns the MintId field if non-nil, zero value otherwise.

### GetMintIdOk

`func (o *NftprimarytransactionTransactionData) GetMintIdOk() (*string, bool)`

GetMintIdOk returns a tuple with the MintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintId

`func (o *NftprimarytransactionTransactionData) SetMintId(v string)`

SetMintId sets MintId field to given value.

### HasMintId

`func (o *NftprimarytransactionTransactionData) HasMintId() bool`

HasMintId returns a boolean if a field has been set.

### GetMintStatus

`func (o *NftprimarytransactionTransactionData) GetMintStatus() string`

GetMintStatus returns the MintStatus field if non-nil, zero value otherwise.

### GetMintStatusOk

`func (o *NftprimarytransactionTransactionData) GetMintStatusOk() (*string, bool)`

GetMintStatusOk returns a tuple with the MintStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMintStatus

`func (o *NftprimarytransactionTransactionData) SetMintStatus(v string)`

SetMintStatus sets MintStatus field to given value.

### HasMintStatus

`func (o *NftprimarytransactionTransactionData) HasMintStatus() bool`

HasMintStatus returns a boolean if a field has been set.

### GetOfferId

`func (o *NftprimarytransactionTransactionData) GetOfferId() string`

GetOfferId returns the OfferId field if non-nil, zero value otherwise.

### GetOfferIdOk

`func (o *NftprimarytransactionTransactionData) GetOfferIdOk() (*string, bool)`

GetOfferIdOk returns a tuple with the OfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferId

`func (o *NftprimarytransactionTransactionData) SetOfferId(v string)`

SetOfferId sets OfferId field to given value.

### HasOfferId

`func (o *NftprimarytransactionTransactionData) HasOfferId() bool`

HasOfferId returns a boolean if a field has been set.

### GetProvider

`func (o *NftprimarytransactionTransactionData) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NftprimarytransactionTransactionData) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NftprimarytransactionTransactionData) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NftprimarytransactionTransactionData) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSellerWalletAddress

`func (o *NftprimarytransactionTransactionData) GetSellerWalletAddress() string`

GetSellerWalletAddress returns the SellerWalletAddress field if non-nil, zero value otherwise.

### GetSellerWalletAddressOk

`func (o *NftprimarytransactionTransactionData) GetSellerWalletAddressOk() (*string, bool)`

GetSellerWalletAddressOk returns a tuple with the SellerWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerWalletAddress

`func (o *NftprimarytransactionTransactionData) SetSellerWalletAddress(v string)`

SetSellerWalletAddress sets SellerWalletAddress field to given value.

### HasSellerWalletAddress

`func (o *NftprimarytransactionTransactionData) HasSellerWalletAddress() bool`

HasSellerWalletAddress returns a boolean if a field has been set.

### GetStatus

`func (o *NftprimarytransactionTransactionData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NftprimarytransactionTransactionData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NftprimarytransactionTransactionData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NftprimarytransactionTransactionData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToAmount

`func (o *NftprimarytransactionTransactionData) GetToAmount() float32`

GetToAmount returns the ToAmount field if non-nil, zero value otherwise.

### GetToAmountOk

`func (o *NftprimarytransactionTransactionData) GetToAmountOk() (*float32, bool)`

GetToAmountOk returns a tuple with the ToAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAmount

`func (o *NftprimarytransactionTransactionData) SetToAmount(v float32)`

SetToAmount sets ToAmount field to given value.

### HasToAmount

`func (o *NftprimarytransactionTransactionData) HasToAmount() bool`

HasToAmount returns a boolean if a field has been set.

### GetToCurrency

`func (o *NftprimarytransactionTransactionData) GetToCurrency() string`

GetToCurrency returns the ToCurrency field if non-nil, zero value otherwise.

### GetToCurrencyOk

`func (o *NftprimarytransactionTransactionData) GetToCurrencyOk() (*string, bool)`

GetToCurrencyOk returns a tuple with the ToCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCurrency

`func (o *NftprimarytransactionTransactionData) SetToCurrency(v string)`

SetToCurrency sets ToCurrency field to given value.

### HasToCurrency

`func (o *NftprimarytransactionTransactionData) HasToCurrency() bool`

HasToCurrency returns a boolean if a field has been set.

### GetTokenId

`func (o *NftprimarytransactionTransactionData) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *NftprimarytransactionTransactionData) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *NftprimarytransactionTransactionData) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.

### HasTokenId

`func (o *NftprimarytransactionTransactionData) HasTokenId() bool`

HasTokenId returns a boolean if a field has been set.

### GetTransactionId

`func (o *NftprimarytransactionTransactionData) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *NftprimarytransactionTransactionData) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *NftprimarytransactionTransactionData) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *NftprimarytransactionTransactionData) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NftprimarytransactionTransactionData) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NftprimarytransactionTransactionData) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NftprimarytransactionTransactionData) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NftprimarytransactionTransactionData) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUserWalletAddress

`func (o *NftprimarytransactionTransactionData) GetUserWalletAddress() string`

GetUserWalletAddress returns the UserWalletAddress field if non-nil, zero value otherwise.

### GetUserWalletAddressOk

`func (o *NftprimarytransactionTransactionData) GetUserWalletAddressOk() (*string, bool)`

GetUserWalletAddressOk returns a tuple with the UserWalletAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserWalletAddress

`func (o *NftprimarytransactionTransactionData) SetUserWalletAddress(v string)`

SetUserWalletAddress sets UserWalletAddress field to given value.

### HasUserWalletAddress

`func (o *NftprimarytransactionTransactionData) HasUserWalletAddress() bool`

HasUserWalletAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


