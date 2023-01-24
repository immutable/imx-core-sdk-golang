# GetSignableDepositRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the token the user is depositing | 
**Token** | [**GetSignableDepositRequestToken**](GetSignableDepositRequestToken.md) |  | 
**User** | **string** | User who is depositing | 

## Methods

### NewGetSignableDepositRequest

`func NewGetSignableDepositRequest(amount string, token GetSignableDepositRequestToken, user string, ) *GetSignableDepositRequest`

NewGetSignableDepositRequest instantiates a new GetSignableDepositRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableDepositRequestWithDefaults

`func NewGetSignableDepositRequestWithDefaults() *GetSignableDepositRequest`

NewGetSignableDepositRequestWithDefaults instantiates a new GetSignableDepositRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSignableDepositRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSignableDepositRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSignableDepositRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToken

`func (o *GetSignableDepositRequest) GetToken() GetSignableDepositRequestToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetSignableDepositRequest) GetTokenOk() (*GetSignableDepositRequestToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetSignableDepositRequest) SetToken(v GetSignableDepositRequestToken)`

SetToken sets Token field to given value.


### GetUser

`func (o *GetSignableDepositRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSignableDepositRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSignableDepositRequest) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


