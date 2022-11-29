# AggregateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAmount** | Pointer to **float32** | Max transaction amount | [optional] 
**MinAmount** | Pointer to **float32** | Min transaction amount | [optional] 

## Methods

### NewAggregateLimit

`func NewAggregateLimit() *AggregateLimit`

NewAggregateLimit instantiates a new AggregateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateLimitWithDefaults

`func NewAggregateLimitWithDefaults() *AggregateLimit`

NewAggregateLimitWithDefaults instantiates a new AggregateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAmount

`func (o *AggregateLimit) GetMaxAmount() float32`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *AggregateLimit) GetMaxAmountOk() (*float32, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *AggregateLimit) SetMaxAmount(v float32)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *AggregateLimit) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetMinAmount

`func (o *AggregateLimit) GetMinAmount() float32`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *AggregateLimit) GetMinAmountOk() (*float32, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *AggregateLimit) SetMinAmount(v float32)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *AggregateLimit) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


