package deposits

import (
	"context"
	"errors"
	"fmt"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk/api"
	"immutable.com/imx-core-sdk/utils"
	"immutable.com/imx-core-sdk/utils/ethereum"
	"immutable.com/imx-core-sdk/workflows/types"
	"math/big"
)

func (d *ERC20Deposit) Execute(e *ethereum.Client, apis *api.APIClient) (*ethtypes.Transaction, error) {
	if d.Type != types.ERC20Type {
		return nil, errors.New("invalid token type")
	}
	ctx := context.Background()

	// Get decimals for this specific ERC20
	token, resp, err := apis.TokensApi.GetToken(ctx, d.TokenAddress).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `TokensApi.GetToken`: %v\nFull HTTP response: %v", err, resp)
	}

	decimals, err := utils.FromStringToDecimal(token.Decimals)
	if err != nil {
		return nil, fmt.Errorf("Error parsing token decimals: %v\n", err)
	}
	amount := utils.ToWei(d.Amount, int(*decimals))

	// Approve whether an amount of token from an account can be spent by a third-party account
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	approve, err := e.IERC20Contract.Approve(auth, e.StarkContractAddress, amount)
	if err != nil {
		return nil, err
	}
	// TODO: check if this needs to be signed before sending also populate gas fields?
	e.Client.SendTransaction(ctx, approve)

	// Get signable deposit details
	getSignableDepositRequest := api.NewGetSignableDepositRequest(amount.String(), *api.NewSignableToken(), e.GetAddress().Hex())
	getSignableDepositRequest.Token.SetType(string(d.Type))
	getSignableDepositRequest.Token.SetData(map[string]interface{}{
		"decimals":      int(*decimals),
		"token_address": d.TokenAddress,
	})

	signableDeposit, resp, err := apis.DepositsApi.GetSignableDeposit(ctx).GetSignableDepositRequest(*getSignableDepositRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `DepositsApi.GetSignableDeposit`: %v\nFull HTTP response: %v", err, resp)
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	encodeAssetRequest := api.NewEncodeAssetRequest(*api.NewEncodeAssetRequestToken())
	encodeAssetRequest.Token.SetType(string(d.Type))
	encodeAssetRequest.Token.SetData(api.EncodeAssetTokenData{
		TokenAddress: &d.TokenAddress,
	})
	encodingResult, resp, err := apis.EncodingApi.EncodeAsset(ctx, "asset").EncodeAssetRequest(*encodeAssetRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `EncodingApi.EncodeAsset`: %v\nFull HTTP response: %v", err, resp)
	}

	// TODO: verify this is the correct type conversion
	assetType, err := utils.HexToInt(encodingResult.AssetType)
	if err != nil {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v\n", encodingResult.AssetType)
	}
	starkKey, err := utils.HexToInt(signableDeposit.StarkKey)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v\n", signableDeposit.StarkKey)
	}
	// TODO: verify if it's ok to set CallOpts to nil
	isRegistered, err := e.RegistrationContract.IsRegistered(nil, starkKey)
	if err != nil {
		return nil, fmt.Errorf("Error when calling `Registration.IsRegistered`: %v\n", err)
	}

	// TODO: verify if big.NewInt(int64(signableDeposit.VaultId)) is the correct conversion
	if isRegistered {
		return depositERC20(ctx, e, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount)
	} else {
		return registerAndDepositERC20(ctx, e, apis.UsersApi, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount)
	}
}

func depositERC20(ctx context.Context, e *ethereum.Client, starkPublicKey *big.Int, vaultId *big.Int, assetType *big.Int, quantizedAmount *big.Int) (*ethtypes.Transaction, error) {
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.DepositERC20(auth, starkPublicKey, assetType, vaultId, quantizedAmount)
	if err != nil {
		return nil, err
	}

	tnx, err = e.SignTransaction(ctx, tnx)
	if err != nil {
		return nil, err
	}

	// TODO: verify if tnx needs to be populated with gas/nounce etc and signed before sending
	if err = e.Client.SendTransaction(ctx, tnx); err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC20(ctx context.Context, e *ethereum.Client, userApi *api.UsersApiService, starkPublicKey *big.Int, vaultId *big.Int, assetType *big.Int, quantizedAmount *big.Int) (*ethtypes.Transaction, error) {
	registrationRequest := api.NewGetSignableRegistrationRequest(e.GetAddress().Hex(), starkPublicKey.String())
	signableRegistration, resp, err := userApi.GetSignableRegistration(ctx).GetSignableRegistrationRequest(*registrationRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `UserApi.GetSignableRegistration`: %v\nFull HTTP response: %v", err, resp)
	}
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	// TODO: verify if it is correct to do []byte conversion on OperatorSignature as below
	operatorSignature, err := utils.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.RegisterAndDepositERC20(auth, e.GetAddress(), starkPublicKey, operatorSignature, assetType, vaultId, quantizedAmount)
	if err != nil {
		return nil, err
	}

	tnx, err = e.SignTransaction(ctx, tnx)
	if err != nil {
		return nil, err
	}

	if err = e.Client.SendTransaction(ctx, tnx); err != nil {
		return nil, err
	}
	return tnx, nil
}
