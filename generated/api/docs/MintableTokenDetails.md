# MintableTokenDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprint** | **string** | Blueprint of this token | 
**ClientTokenId** | **string** | ID provided by the client for this token | 
**TokenId** | **string** | IMX Id of this token | 

## Methods

### NewMintableTokenDetails

`func NewMintableTokenDetails(blueprint string, clientTokenId string, tokenId string, ) *MintableTokenDetails`

NewMintableTokenDetails instantiates a new MintableTokenDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMintableTokenDetailsWithDefaults

`func NewMintableTokenDetailsWithDefaults() *MintableTokenDetails`

NewMintableTokenDetailsWithDefaults instantiates a new MintableTokenDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprint

`func (o *MintableTokenDetails) GetBlueprint() string`

GetBlueprint returns the Blueprint field if non-nil, zero value otherwise.

### GetBlueprintOk

`func (o *MintableTokenDetails) GetBlueprintOk() (*string, bool)`

GetBlueprintOk returns a tuple with the Blueprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprint

`func (o *MintableTokenDetails) SetBlueprint(v string)`

SetBlueprint sets Blueprint field to given value.


### GetClientTokenId

`func (o *MintableTokenDetails) GetClientTokenId() string`

GetClientTokenId returns the ClientTokenId field if non-nil, zero value otherwise.

### GetClientTokenIdOk

`func (o *MintableTokenDetails) GetClientTokenIdOk() (*string, bool)`

GetClientTokenIdOk returns a tuple with the ClientTokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTokenId

`func (o *MintableTokenDetails) SetClientTokenId(v string)`

SetClientTokenId sets ClientTokenId field to given value.


### GetTokenId

`func (o *MintableTokenDetails) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *MintableTokenDetails) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *MintableTokenDetails) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


