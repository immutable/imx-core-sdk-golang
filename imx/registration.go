package imx

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/internal/contracts"
)

// RegisterOffchain performs off chain user registration i.e, on the L2 network.
func (c *Client) RegisterOffchain(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	userEmail string) (*api.RegisterUserResponse, error) {
	if userEmail != "" {
		if !isValidEmail(userEmail) {
			return nil, fmt.Errorf("given userEmail is invalid: %v", userEmail)
		}
	}

	etherKey := l1signer.GetAddress()
	starkKey := l2signer.GetAddress()

	signableRegistrationRequest := api.NewGetSignableRegistrationRequest(etherKey, starkKey)
	signableResponse, httpResponse, err := c.usersAPI.GetSignableRegistrationOffchain(ctx).GetSignableRegistrationRequest(*signableRegistrationRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `UserApi.GetSignableRegistrationOffchain`: %v, HTTP response body: %v", err, httpResponse.Body)
	}

	ethSignature, starkSignature, err := createSignatures(&signableResponse.SignableMessage, &signableResponse.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	registerUserRequest := api.NewRegisterUserRequest(ethSignature, etherKey, starkKey, starkSignature)
	registerUserResponse, httpResp, err := c.usersAPI.RegisterUser(ctx).RegisterUserRequest(*registerUserRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `UserApi.RegisterUser`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return registerUserResponse, nil
}

// IsRegisteredOffChain checks if the given public address is already registered on the offchain (L2 network).
func (c *Client) IsRegisteredOffChain(ctx context.Context, publicAddress string) (*bool, error) {
	usersResponse, httpResp, err := c.usersAPI.GetUsers(ctx, publicAddress).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Users.GetUsers: %v, HTTP response body: %v", err, httpResp.Body)
	}
	isRegistered := len(usersResponse.GetAccounts()) > 0
	return &isRegistered, nil
}

// IsRegisteredOnChain checks if the given public address is already registered on the onchain (L1 network).
func (c *Client) IsRegisteredOnChain(ctx context.Context, contract *contracts.Registration, starkPublicKey string) (*bool, error) {
	starkKey, err := hexutil.DecodeBig(starkPublicKey)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v", starkPublicKey)
	}

	isRegistered, err := contract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	if err != nil {
		isRegistered = false
		return &isRegistered, fmt.Errorf("error: %v", err)
	}
	return &isRegistered, nil
}

/*
GetUsers Get stark keys for a registered user

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param user User
@return GetUsersApiResponse
*/
func (c *Client) GetUsers(ctx context.Context, user string) (*api.GetUsersApiResponse, error) {
	response, httpResponse, err := c.usersAPI.GetUsers(ctx, user).Execute()
	if err != nil {
		return nil, fmt.Errorf("error in getting the users: %v, HTTP response body: %v", err, httpResponse.Body)
	}
	return response, nil
}

// getSignableRegistrationOnchain is a helper function to get the operator signature and payload hash to assist clients in the process of user registration.
func (c *Client) getSignableRegistrationOnchain(ctx context.Context, etherKey, starkKey string) (*api.GetSignableRegistrationResponse, error) {
	signableRegistrationRequest := api.NewGetSignableRegistrationRequest(etherKey, starkKey)
	signableRegistrationResponse, httpResp, err := c.usersAPI.GetSignableRegistration(ctx).
		GetSignableRegistrationRequest(*signableRegistrationRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error in `GetSignableRegistrationRequest.Execute`: %v, HTTP response body: %v", err, httpResp.Body)
	}
	return signableRegistrationResponse, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
