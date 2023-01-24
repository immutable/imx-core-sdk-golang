# SignableTransferDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the token to transfer | 
**Receiver** | **string** | Ethereum address of the receiving user | 
**Token** | [**SignableTransferDetailsToken**](SignableTransferDetailsToken.md) |  | 

## Methods

### NewSignableTransferDetails

`func NewSignableTransferDetails(amount string, receiver string, token SignableTransferDetailsToken, ) *SignableTransferDetails`

NewSignableTransferDetails instantiates a new SignableTransferDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignableTransferDetailsWithDefaults

`func NewSignableTransferDetailsWithDefaults() *SignableTransferDetails`

NewSignableTransferDetailsWithDefaults instantiates a new SignableTransferDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SignableTransferDetails) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SignableTransferDetails) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SignableTransferDetails) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetReceiver

`func (o *SignableTransferDetails) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *SignableTransferDetails) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *SignableTransferDetails) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetToken

`func (o *SignableTransferDetails) GetToken() SignableTransferDetailsToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SignableTransferDetails) GetTokenOk() (*SignableTransferDetailsToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SignableTransferDetails) SetToken(v SignableTransferDetailsToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


