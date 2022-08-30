# CreateTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Requests** | [**[]TransferRequest**](TransferRequest.md) | List of transfers | 
**SenderStarkKey** | **string** | Public stark key of the user sending the transfer | 

## Methods

### NewCreateTransferRequest

`func NewCreateTransferRequest(requests []TransferRequest, senderStarkKey string, ) *CreateTransferRequest`

NewCreateTransferRequest instantiates a new CreateTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTransferRequestWithDefaults

`func NewCreateTransferRequestWithDefaults() *CreateTransferRequest`

NewCreateTransferRequestWithDefaults instantiates a new CreateTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequests

`func (o *CreateTransferRequest) GetRequests() []TransferRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *CreateTransferRequest) GetRequestsOk() (*[]TransferRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *CreateTransferRequest) SetRequests(v []TransferRequest)`

SetRequests sets Requests field to given value.


### GetSenderStarkKey

`func (o *CreateTransferRequest) GetSenderStarkKey() string`

GetSenderStarkKey returns the SenderStarkKey field if non-nil, zero value otherwise.

### GetSenderStarkKeyOk

`func (o *CreateTransferRequest) GetSenderStarkKeyOk() (*string, bool)`

GetSenderStarkKeyOk returns a tuple with the SenderStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderStarkKey

`func (o *CreateTransferRequest) SetSenderStarkKey(v string)`

SetSenderStarkKey sets SenderStarkKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


