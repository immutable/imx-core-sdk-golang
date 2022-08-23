# GetSignableTransferRequestV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the token to transfer | 
**Receiver** | **string** | Ethereum address of the receiving user | 
**Sender** | **string** | Ethereum address of the transferring user | 
**Token** | [**SignableToken**](SignableToken.md) |  | 

## Methods

### NewGetSignableTransferRequestV1

`func NewGetSignableTransferRequestV1(amount string, receiver string, sender string, token SignableToken, ) *GetSignableTransferRequestV1`

NewGetSignableTransferRequestV1 instantiates a new GetSignableTransferRequestV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTransferRequestV1WithDefaults

`func NewGetSignableTransferRequestV1WithDefaults() *GetSignableTransferRequestV1`

NewGetSignableTransferRequestV1WithDefaults instantiates a new GetSignableTransferRequestV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSignableTransferRequestV1) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSignableTransferRequestV1) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSignableTransferRequestV1) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetReceiver

`func (o *GetSignableTransferRequestV1) GetReceiver() string`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *GetSignableTransferRequestV1) GetReceiverOk() (*string, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *GetSignableTransferRequestV1) SetReceiver(v string)`

SetReceiver sets Receiver field to given value.


### GetSender

`func (o *GetSignableTransferRequestV1) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GetSignableTransferRequestV1) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GetSignableTransferRequestV1) SetSender(v string)`

SetSender sets Sender field to given value.


### GetToken

`func (o *GetSignableTransferRequestV1) GetToken() SignableToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetSignableTransferRequestV1) GetTokenOk() (*SignableToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetSignableTransferRequestV1) SetToken(v SignableToken)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


