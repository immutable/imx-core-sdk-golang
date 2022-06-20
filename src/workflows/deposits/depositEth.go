package deposits

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	. "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk/api"
	"immutable.com/imx-core-sdk/utils"
	"immutable.com/imx-core-sdk/utils/ethereum"
	"immutable.com/imx-core-sdk/workflows/types"
	"math/big"
)

func (d *ETHDeposit) Execute(e *ethereum.Client, apis *api.APIClient) (*Transaction, error) {
	if d.Type != types.ETHType {
		return nil, errors.New("invalid token type")
	}
	ctx := context.Background()

	amount, err := utils.ParseEtherToWei(d.Amount)
	if err != nil {
		return nil, fmt.Errorf("Error when parsing deposit amount: %v\n", err)
	}

	getSignableDepositRequest := api.NewGetSignableDepositRequest(amount.String(), *api.NewSignableToken(), e.GetAddress().Hex())
	getSignableDepositRequest.Token.SetType(string(types.ETHType))
	getSignableDepositRequest.Token.SetData(map[string]interface{}{
		"decimals": 18,
	})

	signableDeposit, resp, err := apis.DepositsApi.GetSignableDeposit(ctx).GetSignableDepositRequest(*getSignableDepositRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `DepositsApi.GetSignableDeposit`: %v\nFull HTTP response: %v", err, resp)
	}

	encodeAssetRequest := api.NewEncodeAssetRequest(*api.NewEncodeAssetRequestToken())
	encodeAssetRequest.Token.SetType(string(d.Type))
	encodedAsset, resp, err := apis.EncodingApi.EncodeAsset(ctx, "asset").EncodeAssetRequest(*encodeAssetRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `EncodingApi.EncodeAsset`: %v\nFull HTTP response: %v", err, resp)
	}

	assetType, err := utils.HexToInt(encodedAsset.AssetType)
	if err != nil {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v\n", encodedAsset.AssetType)
	}
	starkKey, err := utils.HexToInt(signableDeposit.StarkKey)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v\n", signableDeposit.StarkKey)
	}
	isRegistered, err := e.RegistrationContract.IsRegistered(&bind.CallOpts{From: e.GetAddress()}, starkKey)
	if err != nil {
		return nil, fmt.Errorf("Error when calling `Registration.IsRegistered`: %v\n", err)
	}

	if isRegistered {
		return depositEth(ctx, e, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount)
	} else {
		return registerAndDepositEth(ctx, e, apis.UsersApi, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount)
	}
}

func registerAndDepositEth(ctx context.Context, e *ethereum.Client, userApi *api.UsersApiService, starkPublicKey *big.Int, vaultId *big.Int, assetType *big.Int, amount *big.Int) (*Transaction, error) {
	registrationRequest := api.NewGetSignableRegistrationRequest(e.GetAddress().Hex(), starkPublicKey.String())
	signableRegistration, resp, err := userApi.GetSignableRegistration(ctx).GetSignableRegistrationRequest(*registrationRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `UserApi.GetSignableRegistration`: %v\nFull HTTP response: %v", err, resp)
	}
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.RegisterAndDepositEth(auth, e.GetAddress(), starkPublicKey, operatorSignature, assetType, vaultId)
	if err != nil {
		return nil, err
	}
	tnx.Value().Set(amount)

	tnx, err = e.SignTransaction(ctx, tnx)
	if err != nil {
		return nil, err
	}

	if err = e.Client.SendTransaction(ctx, tnx); err != nil {
		return nil, err
	}
	return tnx, nil
}

func depositEth(ctx context.Context, e *ethereum.Client, starkPublicKey *big.Int, vaultId *big.Int, assetType *big.Int, amount *big.Int) (*Transaction, error) {
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.Deposit(auth, starkPublicKey, assetType, vaultId)
	if err != nil {
		return nil, err
	}
	tnx.Value().Set(amount)

	tnx, err = e.SignTransaction(ctx, tnx)
	if err != nil {
		return nil, err
	}

	if err = e.Client.SendTransaction(ctx, tnx); err != nil {
		return nil, err
	}
	return tnx, nil
}
