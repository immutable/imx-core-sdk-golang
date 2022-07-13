package withdrawals

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
	"math/big"
)

// CompleteWithdrawal performs the complete withdrawal workflow on ETHWithdrawal
func (w *ETHWithdrawal) CompleteWithdrawal(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*Transaction, error) {
	if w.Type != types.ETHType {
		return nil, errors.New("invalid token type")
	}

	assetType, err := helpers.GetEncodedAssetTypeForEth(ctx, api)
	if err != nil {
		return nil, err
	}

	return completeFungiblesWithdrawal(ctx, ethClient, api, l1signer, starkKeyHex, assetType)
}

// CompleteWithdrawal performs the complete withdrawal workflow on ERC20Withdrawal
func (w *ERC20Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*Transaction, error) {
	if w.Type != types.ERC20Type {
		return nil, errors.New("invalid token type")
	}

	assetType, err := helpers.GetEncodedAssetTypeForERC20(ctx, api, w.TokenId, w.TokenAddress)
	if err != nil {
		return nil, err
	}

	return completeFungiblesWithdrawal(ctx, ethClient, api, l1signer, starkKeyHex, assetType)
}

func completeFungiblesWithdrawal(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
	assetType *big.Int,
) (*Transaction, error) {

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s\n", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	if isRegistered {
		return withdrawFungibles(ctx, ethClient, l1signer, starkKey, assetType)
	} else {
		return registerAndWithdrawFungibles(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, assetType)
	}
}

func withdrawFungibles(ctx context.Context, ethClient *ethereum.Client, l1signer signers.L1Signer, starkKey, assetType *big.Int) (*Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.Withdraw(auth, starkKey, assetType)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndWithdrawFungibles(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	api *client.ImmutableXAPI,
	starkKeyHex string,
	starkKey *big.Int,
	assetType *big.Int,
) (*Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := registration.GetSignableRegistrationOnchain(ctx, api, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(*signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.RegistrationContract.RegisterAndWithdraw(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
