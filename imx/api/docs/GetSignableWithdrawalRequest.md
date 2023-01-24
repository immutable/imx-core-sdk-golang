# GetSignableWithdrawalRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **string** | Amount of the token to withdraw | 
**Token** | [**GetSignableWithdrawalRequestToken**](GetSignableWithdrawalRequestToken.md) |  | 
**User** | **string** | Ethereum address of the user who is making this withdrawal | 

## Methods

### NewGetSignableWithdrawalRequest

`func NewGetSignableWithdrawalRequest(amount string, token GetSignableWithdrawalRequestToken, user string, ) *GetSignableWithdrawalRequest`

NewGetSignableWithdrawalRequest instantiates a new GetSignableWithdrawalRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignableWithdrawalRequestWithDefaults

`func NewGetSignableWithdrawalRequestWithDefaults() *GetSignableWithdrawalRequest`

NewGetSignableWithdrawalRequestWithDefaults instantiates a new GetSignableWithdrawalRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GetSignableWithdrawalRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GetSignableWithdrawalRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GetSignableWithdrawalRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetToken

`func (o *GetSignableWithdrawalRequest) GetToken() GetSignableWithdrawalRequestToken`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetSignableWithdrawalRequest) GetTokenOk() (*GetSignableWithdrawalRequestToken, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetSignableWithdrawalRequest) SetToken(v GetSignableWithdrawalRequestToken)`

SetToken sets Token field to given value.


### GetUser

`func (o *GetSignableWithdrawalRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetSignableWithdrawalRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetSignableWithdrawalRequest) SetUser(v string)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


