package imx

import (
	"context"
	"fmt"
	"math/big"
	"net/mail"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

/*
RegisterOffchain performs off chain user registration i.e, on the L2 network.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param l1Signer Ethereum signer to sign message.
@param l2signer Stark signer to sign the payload hash.
@param userEmail A valid user email.
@return RegisterUserResponse
*/
func (c *Client) RegisterOffchain(ctx context.Context,
	l1signer L1Signer,
	l2signer L2Signer,
	userEmail string,
) (*api.RegisterUserResponse, error) {
	if userEmail != "" {
		if !isValidEmail(userEmail) {
			return nil, fmt.Errorf("given userEmail is invalid: %v", userEmail)
		}
	}

	etherKey := l1signer.GetAddress()
	starkKey := l2signer.GetPublicKey()

	signableRegistrationRequest := api.NewGetSignableRegistrationRequest(etherKey, starkKey)
	signableResponse, httpResponse, err := c.UsersAPI.GetSignableRegistrationOffchain(ctx).GetSignableRegistrationRequest(*signableRegistrationRequest).Execute()
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}

	ethSignature, starkSignature, err := createSignatures(&signableResponse.SignableMessage, &signableResponse.PayloadHash, l1signer, l2signer)
	if err != nil {
		return nil, err
	}

	registerUserRequest := api.NewRegisterUserRequest(ethSignature, etherKey, starkKey, starkSignature)
	registerUserResponse, httpResponse, err := c.UsersAPI.RegisterUser(ctx).RegisterUserRequest(*registerUserRequest).Execute()
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return registerUserResponse, nil
}

/*
IsRegisteredOnChain checks if the given public address is already registered on the onchain (L1 network).

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param starkPublicKey The stark wallet public address.
@return true if registered or false. Nil on error.
*/
func (c *Client) IsRegisteredOnChain(ctx context.Context, starkPublicKey string) (*bool, error) {
	starkKey, ok := new(big.Int).SetString(starkPublicKey, 0)
	if !ok {
		return nil, fmt.Errorf("error converting StarkKey to bigint")
	}

	isRegistered, err := c.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	if err != nil {
		isRegistered = false
		return &isRegistered, fmt.Errorf("error: %v", err)
	}
	return &isRegistered, nil
}

/*
GetUsers Get stark keys for a registered user.
Can also be used to check if the user is registered or not when it returns an error.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param user User
@return GetUsersApiResponse
*/
func (c *Client) GetUsers(ctx context.Context, user string) (*api.GetUsersApiResponse, error) {
	response, httpResponse, err := c.UsersAPI.GetUsers(ctx, user).Execute()
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return response, nil
}

// getSignableRegistrationOnchain is a helper function to get the operator signature and payload hash to assist clients in the process of user registration.
func (c *Client) getSignableRegistrationOnchain(
	ctx context.Context,
	etherKey, starkKey string,
) (*api.GetSignableRegistrationResponse, error) {
	signableRegistrationRequest := api.NewGetSignableRegistrationRequest(etherKey, starkKey)
	signableRegistrationResponse, httpResponse, err := c.UsersAPI.GetSignableRegistration(ctx).
		GetSignableRegistrationRequest(*signableRegistrationRequest).Execute()
	defer httpResponse.Body.Close()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}
	return signableRegistrationResponse, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
