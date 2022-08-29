package withdrawals

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/immutable/imx-core-sdk-golang/ethereumutil"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/internal/convert"
	"github.com/immutable/imx-core-sdk-golang/internal/encode"
	"github.com/immutable/imx-core-sdk-golang/signers"
	"github.com/immutable/imx-core-sdk-golang/workflows/registration"
)

// CompleteEthWithdrawal performs the complete withdrawal workflow for ETH
func CompleteEthWithdrawal(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	clientAPI *api.APIClient,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*eth.Transaction, error) {
	assetType, err := encode.GetEncodedAssetTypeForEth(ctx, clientAPI.EncodingApi)
	if err != nil {
		return nil, err
	}

	return completeFungiblesWithdrawal(ctx, ethClient, clientAPI, l1signer, starkKeyHex, assetType)
}

// CompleteWithdrawal performs the complete withdrawal workflow on ERC20Withdrawal
func (w *ERC20Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	clientAPI *api.APIClient,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*eth.Transaction, error) {
	assetType, err := encode.GetEncodedAssetTypeForERC20(ctx, clientAPI.EncodingApi, w.TokenAddress)
	if err != nil {
		return nil, err
	}

	return completeFungiblesWithdrawal(ctx, ethClient, clientAPI, l1signer, starkKeyHex, assetType)
}

func completeFungiblesWithdrawal(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	clientAPI *api.APIClient,
	l1signer signers.L1Signer,
	starkKeyHex string,
	assetType *big.Int,
) (*eth.Transaction, error) {
	starkKey, err := convert.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	if isRegistered {
		return withdrawFungibles(ctx, ethClient, l1signer, starkKey, assetType)
	}
	return registerAndWithdrawFungibles(ctx, ethClient, l1signer, clientAPI.UsersApi, starkKeyHex, starkKey, assetType)
}

func withdrawFungibles(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	l1signer signers.L1Signer,
	starkKey, assetType *big.Int,
) (*eth.Transaction, error) {
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
	ethClient *ethereumutil.Client,
	l1signer signers.L1Signer,
	usersAPI api.UsersApi,
	starkKeyHex string,
	starkKey *big.Int,
	assetType *big.Int,
) (*eth.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := registration.GetSignableRegistrationOnchain(ctx, usersAPI, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := convert.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	transaction, err := ethClient.RegistrationContract.RegisterAndWithdraw(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}
