# CreateTradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Current status of trade | 
**TradeId** | **int32** | ID of trade within Immutable X | 

## Methods

### NewCreateTradeResponse

`func NewCreateTradeResponse(status string, tradeId int32, ) *CreateTradeResponse`

NewCreateTradeResponse instantiates a new CreateTradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTradeResponseWithDefaults

`func NewCreateTradeResponseWithDefaults() *CreateTradeResponse`

NewCreateTradeResponseWithDefaults instantiates a new CreateTradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateTradeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTradeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTradeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTradeId

`func (o *CreateTradeResponse) GetTradeId() int32`

GetTradeId returns the TradeId field if non-nil, zero value otherwise.

### GetTradeIdOk

`func (o *CreateTradeResponse) GetTradeIdOk() (*int32, bool)`

GetTradeIdOk returns a tuple with the TradeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeId

`func (o *CreateTradeResponse) SetTradeId(v int32)`

SetTradeId sets TradeId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


