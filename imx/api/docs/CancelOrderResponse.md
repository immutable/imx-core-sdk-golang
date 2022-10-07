# CancelOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | ID of the cancelled order | 
**Status** | **string** | New status of the order | 

## Methods

### NewCancelOrderResponse

`func NewCancelOrderResponse(orderId int32, status string, ) *CancelOrderResponse`

NewCancelOrderResponse instantiates a new CancelOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelOrderResponseWithDefaults

`func NewCancelOrderResponseWithDefaults() *CancelOrderResponse`

NewCancelOrderResponseWithDefaults instantiates a new CancelOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CancelOrderResponse) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CancelOrderResponse) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CancelOrderResponse) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetStatus

`func (o *CancelOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


