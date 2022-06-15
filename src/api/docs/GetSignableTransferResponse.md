# GetSignableTransferResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SenderStarkKey** | **string** | Sender of the transfer | 
**SignableMessage** | **string** | Message to sign with L1 wallet to confirm transfer request | 
**SignableResponses** | [**[]SignableTransferResponseDetails**](SignableTransferResponseDetails.md) | List of transfer responses without the sender stark key | 

## Methods

### NewGetSignableTransferResponse

`func NewGetSignableTransferResponse(senderStarkKey string, signableMessage string, signableResponses []SignableTransferResponseDetails, ) *GetSignableTransferResponse`

NewGetSignableTransferResponse instantiates a new GetSignableTransferResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableTransferResponseWithDefaults

`func NewGetSignableTransferResponseWithDefaults() *GetSignableTransferResponse`

NewGetSignableTransferResponseWithDefaults instantiates a new GetSignableTransferResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSenderStarkKey

`func (o *GetSignableTransferResponse) GetSenderStarkKey() string`

GetSenderStarkKey returns the SenderStarkKey field if non-nil, zero value otherwise.

### GetSenderStarkKeyOk

`func (o *GetSignableTransferResponse) GetSenderStarkKeyOk() (*string, bool)`

GetSenderStarkKeyOk returns a tuple with the SenderStarkKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderStarkKey

`func (o *GetSignableTransferResponse) SetSenderStarkKey(v string)`

SetSenderStarkKey sets SenderStarkKey field to given value.


### GetSignableMessage

`func (o *GetSignableTransferResponse) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableTransferResponse) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableTransferResponse) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.


### GetSignableResponses

`func (o *GetSignableTransferResponse) GetSignableResponses() []SignableTransferResponseDetails`

GetSignableResponses returns the SignableResponses field if non-nil, zero value otherwise.

### GetSignableResponsesOk

`func (o *GetSignableTransferResponse) GetSignableResponsesOk() (*[]SignableTransferResponseDetails, bool)`

GetSignableResponsesOk returns a tuple with the SignableResponses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableResponses

`func (o *GetSignableTransferResponse) SetSignableResponses(v []SignableTransferResponseDetails)`

SetSignableResponses sets SignableResponses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


