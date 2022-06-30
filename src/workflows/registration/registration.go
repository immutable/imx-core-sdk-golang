package registration

import (
	"context"
	"fmt"
	"net/mail"

	"immutable.com/imx-core-sdk-golang/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/users"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/generated/contracts"
	"immutable.com/imx-core-sdk-golang/signers"
)

// RegisterOffchain performs user registration off chain.
func RegisterOffchain(ctx context.Context, api *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer, userEmail string) (*users.RegisterUserOK, error) {
	if userEmail != "" {
		if !isValidEmail(userEmail) {
			return nil, fmt.Errorf("given userEmail is invalid: %v\n", userEmail)
		}
	}

	signableRegistrationOffchainParams := users.NewGetSignableRegistrationOffchainParamsWithContext(ctx)

	etherKey := l1signer.GetAddress().Hex()
	starkKey := l2signer.GetAddress()
	signableRegistrationOffchainParams.SetGetSignableRegistrationRequest(&models.GetSignableRegistrationRequest{
		EtherKey: &etherKey,
		StarkKey: &starkKey,
	})

	signableRegistrationOffchain, err := api.Users.GetSignableRegistrationOffchain(signableRegistrationOffchainParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Users.GetSignableRegistrationOffchain`: %v\n", err)
	}

	payload := signableRegistrationOffchain.GetPayload()
	ethSignature, err := l1signer.SignMessage(*payload.SignableMessage)
	if err != nil {
		return nil, fmt.Errorf("error in l1signer.SignMessage for generating EthSignature from SignableMessage: %v", err)
	}

	starkSignature, err := l2signer.SignMessage(*payload.PayloadHash)
	if err != nil {
		return nil, fmt.Errorf("error in l2signer.SignMessage for generating StarkSignature from PayloadHash: %v", err)
	}

	ethSignatureEncodedInHex := hexutil.Encode(ethSignature)

	return registerUser(ctx, api, etherKey, ethSignatureEncodedInHex, starkKey, starkSignature, userEmail)
}

func IsRegisteredOffChain(ctx context.Context, api *client.ImmutableXAPI, publicAddress string) ([]string, error) {
	getUserParams := users.NewGetUsersParamsWithContext(ctx)
	getUserParams.SetUser(publicAddress)
	users, err := api.Users.GetUsers(getUserParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `api.Users.GetUsers: %v", err)
	}
	return users.GetPayload().Accounts, nil
}

func IsRegisteredOnChain(contract *contracts.Registration, starkPublicKey string) (*bool, error) {
	starkKey, err := utils.HexToInt(starkPublicKey)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v\n", starkPublicKey)
	}

	isRegistered, err := contract.IsRegistered(nil, starkKey)
	if err != nil {
		isRegistered = false
		return &isRegistered, fmt.Errorf("error: %v\n", err)
	}
	return &isRegistered, nil
}

func GetSignableRegistrationOnchain(api *client.ImmutableXAPI, etherKey string, starkKey string) (*users.GetSignableRegistrationOK, error) {
	ctx := context.Background()
	signableRegistrationParams := users.NewGetSignableRegistrationParamsWithContext(ctx)

	signableRegistrationParams.SetGetSignableRegistrationRequest(&models.GetSignableRegistrationRequest{
		EtherKey: &etherKey,
		StarkKey: &starkKey,
	})

	signableRegistration, err := api.Users.GetSignableRegistration(signableRegistrationParams)
	if err != nil {
		return nil, fmt.Errorf("Error when calling `api.Users.GetSignableRegistration`: %v\n", err)
	}
	return signableRegistration, nil
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func registerUser(ctx context.Context, api *client.ImmutableXAPI, etherKey string, etherSignature string, starkKey string, starkSignature string, userEmail string) (*users.RegisterUserOK, error) {
	registerUserParams := users.NewRegisterUserParamsWithContext(ctx)
	registerUserParams.SetRegisterUserRequest(&models.RegisterUserRequest{
		Email:          userEmail,
		EthSignature:   &etherSignature,
		EtherKey:       &etherKey,
		StarkKey:       &starkKey,
		StarkSignature: &starkSignature,
	})

	registerUserResponse, err := api.Users.RegisterUser(registerUserParams)
	if err != nil {
		return nil, fmt.Errorf("Failed to RegisterUser, error: %v\n", err)
	}
	return registerUserResponse, nil
}
