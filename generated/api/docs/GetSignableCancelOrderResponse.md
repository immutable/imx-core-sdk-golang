# GetSignableCancelOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | ID of the order to be cancelled | 
**PayloadHash** | **string** | Hash of the payload to be signed for cancel order | 
**SignableMessage** | **string** | Message to sign from wallet to confirm cancel order | 

## Methods

### NewGetSignableCancelOrderResponse

`func NewGetSignableCancelOrderResponse(orderId int32, payloadHash string, signableMessage string, ) *GetSignableCancelOrderResponse`

NewGetSignableCancelOrderResponse instantiates a new GetSignableCancelOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableCancelOrderResponseWithDefaults

`func NewGetSignableCancelOrderResponseWithDefaults() *GetSignableCancelOrderResponse`

NewGetSignableCancelOrderResponseWithDefaults instantiates a new GetSignableCancelOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetSignableCancelOrderResponse) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetSignableCancelOrderResponse) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetSignableCancelOrderResponse) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetPayloadHash

`func (o *GetSignableCancelOrderResponse) GetPayloadHash() string`

GetPayloadHash returns the PayloadHash field if non-nil, zero value otherwise.

### GetPayloadHashOk

`func (o *GetSignableCancelOrderResponse) GetPayloadHashOk() (*string, bool)`

GetPayloadHashOk returns a tuple with the PayloadHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadHash

`func (o *GetSignableCancelOrderResponse) SetPayloadHash(v string)`

SetPayloadHash sets PayloadHash field to given value.


### GetSignableMessage

`func (o *GetSignableCancelOrderResponse) GetSignableMessage() string`

GetSignableMessage returns the SignableMessage field if non-nil, zero value otherwise.

### GetSignableMessageOk

`func (o *GetSignableCancelOrderResponse) GetSignableMessageOk() (*string, bool)`

GetSignableMessageOk returns a tuple with the SignableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignableMessage

`func (o *GetSignableCancelOrderResponse) SetSignableMessage(v string)`

SetSignableMessage sets SignableMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


