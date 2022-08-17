package deposits

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/tokens"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/encode"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
)

// Deposit performs the deposit workflow on the ETHDeposit.
func (d *ETHDeposit) Deposit(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, l1signer signers.L1Signer) (*eth.Transaction, error) {
	amount, err := utils.ToDenomination(d.Amount, utils.EtherDecimals)
	if err != nil {
		return nil, fmt.Errorf("error when parsing deposit amount: %v", err)
	}
	signableDepositRequest := newSignableDepositRequestForEth(amount.String(), l1signer.GetAddress())
	signableDeposit, err := getSignableDeposit(ctx, clientAPI.DepositsApi, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	assetType, err := encode.GetEncodedAssetTypeForEth(ctx, clientAPI.EncodingApi)
	if err != nil {
		return nil, err
	}

	starkKeyHex := signableDeposit.StarkKey
	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %s", starkKeyHex)
	}
	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return depositEth(ctx, ethClient, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount)
	}
	return registerAndDepositEth(ctx, ethClient, l1signer, clientAPI.UsersApi, starkKeyHex, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, amount)
}

func newSignableDepositRequestForEth(amount, user string) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  *tokens.NewSignableTokenEth(),
		User:   user,
	}
}

func registerAndDepositEth(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	usersAPI api.UsersApi,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	amount *big.Int) (*eth.Transaction, error) {
	etherKey := l1signer.GetAddress()
	signableRegistration, err := registration.GetSignableRegistrationOnchain(ctx, usersAPI, etherKey, starkKeyHex)
	if err != nil {
		return nil, err
	}

	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	auth.Value = amount
	tnx, err := ethClient.CoreContract.RegisterAndDepositEth(
		auth,
		common.HexToAddress(etherKey),
		starkKey,
		operatorSignature,
		assetType,
		vaultID,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func depositEth(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	amount *big.Int) (*eth.Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	auth.Value = amount
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.Deposit(auth, starkPublicKey, assetType, vaultID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
