package registration

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/generated/contracts"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
)

// RegisterOffchain performs user registration off chain.
func RegisterOffchain(ctx context.Context,
	userAPI api.UsersApi,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
	userEmail string) (*api.RegisterUserResponse, error) {
	if userEmail != "" {
		if !isValidEmail(userEmail) {
			return nil, fmt.Errorf("given userEmail is invalid: %v", userEmail)
		}
	}

	etherKey := l1signer.GetAddress()
	starkKey := l2signer.GetAddress()

	signableRegistrationRequest := api.NewGetSignableRegistrationRequest(etherKey, starkKey)
	signableRegistrationOffchainResponse, httpResponse, err := userAPI.GetSignableRegistrationOffchain(ctx).GetSignableRegistrationRequest(*signableRegistrationRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `UserApi.GetSignableRegistrationOffchain`: %v, full HTTP response: %v", err, httpResponse)
	}

	ethSignature, err := l1signer.SignMessage(signableRegistrationOffchainResponse.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error in l1signer.SignMessage for generating EthSignature from SignableMessage: %v", err)
	}

	starkSignature, err := l2signer.SignMessage(signableRegistrationOffchainResponse.PayloadHash)
	if err != nil {
		return nil, fmt.Errorf("error in l2signer.SignMessage for generating StarkSignature from PayloadHash: %v", err)
	}

	ethSignatureEncodedInHex := hexutil.Encode(ethSignature)

	registerUserRequest := api.NewRegisterUserRequest(ethSignatureEncodedInHex, etherKey, starkKey, starkSignature)
	registerUserResponse, httpResp, err := userAPI.RegisterUser(ctx).RegisterUserRequest(*registerUserRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `UserApi.RegisterUser`: %v, full HTTP response: %v", err, httpResp)
	}
	return registerUserResponse, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsRegisteredOffChain(ctx context.Context, usersAPI api.UsersApi, publicAddress string) ([]string, error) {
	getUsersRequest := usersAPI.GetUsers(ctx, publicAddress)
	usersResponse, httpResp, err := usersAPI.GetUsersExecute(getUsersRequest)
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Users.GetUsers: %v, full HTTP response %v", err, httpResp)
	}
	return usersResponse.GetAccounts(), nil
}

func IsRegisteredOnChain(ctx context.Context, contract *contracts.Registration, starkPublicKey string) (*bool, error) {
	starkKey, err := utils.HexToInt(starkPublicKey)
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

func GetSignableRegistrationOnchain(ctx context.Context, usersAPI api.UsersApi, etherKey, starkKey string) (*api.GetSignableRegistrationResponse, error) {
	signableRegistrationRequest := api.NewGetSignableRegistrationRequest(etherKey, starkKey)
	apiGetSignableRegistrationRequest := usersAPI.GetSignableRegistration(ctx)
	apiGetSignableRegistrationRequest.GetSignableRegistrationRequest(*signableRegistrationRequest)
	signableRegistrationResponse, httpResp, err := usersAPI.GetSignableRegistrationExecute(apiGetSignableRegistrationRequest)
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Users.GetSignableRegistration`: %v, full HTTP response %v", err, httpResp)
	}
	return signableRegistrationResponse, nil
}
