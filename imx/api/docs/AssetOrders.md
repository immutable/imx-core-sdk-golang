# AssetOrders

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuyOrders** | Pointer to **[]map[string]interface{}** | Buy orders for this asset | [optional] 
**SellOrders** | Pointer to **[]map[string]interface{}** | Sell orders for this asset | [optional] 

## Methods

### NewAssetOrders

`func NewAssetOrders() *AssetOrders`

NewAssetOrders instantiates a new AssetOrders object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetOrdersWithDefaults

`func NewAssetOrdersWithDefaults() *AssetOrders`

NewAssetOrdersWithDefaults instantiates a new AssetOrders object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuyOrders

`func (o *AssetOrders) GetBuyOrders() []map[string]interface{}`

GetBuyOrders returns the BuyOrders field if non-nil, zero value otherwise.

### GetBuyOrdersOk

`func (o *AssetOrders) GetBuyOrdersOk() (*[]map[string]interface{}, bool)`

GetBuyOrdersOk returns a tuple with the BuyOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyOrders

`func (o *AssetOrders) SetBuyOrders(v []map[string]interface{})`

SetBuyOrders sets BuyOrders field to given value.

### HasBuyOrders

`func (o *AssetOrders) HasBuyOrders() bool`

HasBuyOrders returns a boolean if a field has been set.

### GetSellOrders

`func (o *AssetOrders) GetSellOrders() []map[string]interface{}`

GetSellOrders returns the SellOrders field if non-nil, zero value otherwise.

### GetSellOrdersOk

`func (o *AssetOrders) GetSellOrdersOk() (*[]map[string]interface{}, bool)`

GetSellOrdersOk returns a tuple with the SellOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellOrders

`func (o *AssetOrders) SetSellOrders(v []map[string]interface{})`

SetSellOrders sets SellOrders field to given value.

### HasSellOrders

`func (o *AssetOrders) HasSellOrders() bool`

HasSellOrders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


