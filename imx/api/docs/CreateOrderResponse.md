# CreateOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **int32** | ID of the created order | 
**RequestId** | Pointer to **string** | Request ID as a reference for an asynchronous order creation request | [optional] 
**Status** | **string** | Status of the created order | 
**Time** | **int32** | Timestamp of the created order | 

## Methods

### NewCreateOrderResponse

`func NewCreateOrderResponse(orderId int32, status string, time int32, ) *CreateOrderResponse`

NewCreateOrderResponse instantiates a new CreateOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderResponseWithDefaults

`func NewCreateOrderResponseWithDefaults() *CreateOrderResponse`

NewCreateOrderResponseWithDefaults instantiates a new CreateOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateOrderResponse) GetOrderId() int32`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateOrderResponse) GetOrderIdOk() (*int32, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateOrderResponse) SetOrderId(v int32)`

SetOrderId sets OrderId field to given value.


### GetRequestId

`func (o *CreateOrderResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *CreateOrderResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *CreateOrderResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *CreateOrderResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetStatus

`func (o *CreateOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTime

`func (o *CreateOrderResponse) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *CreateOrderResponse) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *CreateOrderResponse) SetTime(v int32)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


