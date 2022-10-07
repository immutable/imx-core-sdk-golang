# CancelOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | ID of the order | 
**StarkSignature** | **string** | Payload signature | 

## Methods

### NewCancelOrderRequest

`func NewCancelOrderRequest(orderId int32, starkSignature string, ) *CancelOrderRequest`

NewCancelOrderRequest instantiates a new CancelOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderRequestWithDefaults

`func NewCancelOrderRequestWithDefaults() *CancelOrderRequest`

NewCancelOrderRequestWithDefaults instantiates a new CancelOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CancelOrderRequest) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelOrderRequest) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelOrderRequest) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetStarkSignature

`func (o *CancelOrderRequest) GetStarkSignature() string`

GetStarkSignature returns the StarkSignature field if non-nil, zero value otherwise.

### GetStarkSignatureOk

`func (o *CancelOrderRequest) GetStarkSignatureOk() (*string, bool)`

GetStarkSignatureOk returns a tuple with the StarkSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarkSignature

`func (o *CancelOrderRequest) SetStarkSignature(v string)`

SetStarkSignature sets StarkSignature field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


