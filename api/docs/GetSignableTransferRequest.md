# GetSignableTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderEtherKey** | **string** | Ethereum address of the transferring user | 
**SignableRequests** | [**[]SignableTransferDetails**](SignableTransferDetails.md) | List of signable transfer details | 

## Methods

### NewGetSignableTransferRequest

`func NewGetSignableTransferRequest(senderEtherKey string, signableRequests []SignableTransferDetails, ) *GetSignableTransferRequest`

NewGetSignableTransferRequest instantiates a new GetSignableTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTransferRequestWithDefaults

`func NewGetSignableTransferRequestWithDefaults() *GetSignableTransferRequest`

NewGetSignableTransferRequestWithDefaults instantiates a new GetSignableTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderEtherKey

`func (o *GetSignableTransferRequest) GetSenderEtherKey() string`

GetSenderEtherKey returns the SenderEtherKey field if non-nil, zero value otherwise.

### GetSenderEtherKeyOk

`func (o *GetSignableTransferRequest) GetSenderEtherKeyOk() (*string, bool)`

GetSenderEtherKeyOk returns a tuple with the SenderEtherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderEtherKey

`func (o *GetSignableTransferRequest) SetSenderEtherKey(v string)`

SetSenderEtherKey sets SenderEtherKey field to given value.


### GetSignableRequests

`func (o *GetSignableTransferRequest) GetSignableRequests() []SignableTransferDetails`

GetSignableRequests returns the SignableRequests field if non-nil, zero value otherwise.

### GetSignableRequestsOk

`func (o *GetSignableTransferRequest) GetSignableRequestsOk() (*[]SignableTransferDetails, bool)`

GetSignableRequestsOk returns a tuple with the SignableRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableRequests

`func (o *GetSignableTransferRequest) SetSignableRequests(v []SignableTransferDetails)`

SetSignableRequests sets SignableRequests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


