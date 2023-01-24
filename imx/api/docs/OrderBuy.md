# OrderBuy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TokenDataOAIGen**](TokenDataOAIGen.md) |  | 
**Type** | **string** | Type of this asset (ETH/ERC20/ERC721) | 

## Methods

### NewOrderBuy

`func NewOrderBuy(data TokenDataOAIGen, type_ string, ) *OrderBuy`

NewOrderBuy instantiates a new OrderBuy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderBuyWithDefaults

`func NewOrderBuyWithDefaults() *OrderBuy`

NewOrderBuyWithDefaults instantiates a new OrderBuy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *OrderBuy) GetData() TokenDataOAIGen`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OrderBuy) GetDataOk() (*TokenDataOAIGen, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OrderBuy) SetData(v TokenDataOAIGen)`

SetData sets Data field to given value.


### GetType

`func (o *OrderBuy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OrderBuy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OrderBuy) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


