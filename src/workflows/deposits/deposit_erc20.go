package deposits

import (
	"context"
	"fmt"
	"math/big"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/encode"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
)

// Deposit performs the deposit workflow on the ERC20Deposit.
func (d *ERC20Deposit) Deposit(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, l1signer signers.L1Signer) (*eth.Transaction, error) {
	// Get decimals for this specific ERC20
	token, httpResp, err := clientAPI.TokensApi.GetToken(ctx, d.TokenAddress).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Tokens.GetToken`: %v, http reponse body: %v", err, httpResp.Body)
	}

	decimals, err := utils.FromStringToDecimal(token.Decimals)
	if err != nil {
		return nil, fmt.Errorf("error parsing token decimals: %v", err)
	}
	amount := utils.ToWei(d.Amount, int(*decimals))

	// Approve whether an amount of token from an account can be spent by a third-party account
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	ierc20Contract, err := ethClient.NewIERC20Contract(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}
	_, err = ierc20Contract.Approve(auth, ethClient.StarkContractAddress, amount)
	if err != nil {
		return nil, err
	}

	// Get signable deposit details
	signableDepositRequest := newSignableDepositRequestForERC20(amount.String(), d.TokenAddress, l1signer.GetAddress(), int(*decimals))
	signableDepositResponse, err := getSignableDeposit(ctx, clientAPI.DepositsApi, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := encode.GetEncodedAssetTypeForERC20(ctx, clientAPI.EncodingApi, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKeyHex := signableDepositResponse.StarkKey
	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	quantizedAmount, _ := new(big.Int).SetString(signableDepositResponse.Amount, 10)

	if isRegistered {
		return depositERC20(ctx, ethClient, l1signer, starkKey, big.NewInt(int64(signableDepositResponse.VaultId)), assetType, quantizedAmount)
	} else {
		return registerAndDepositERC20(ctx, ethClient, l1signer, clientAPI.UsersApi, starkKeyHex, starkKey, big.NewInt(int64(signableDepositResponse.VaultId)), assetType, quantizedAmount)
	}
}

func newSignableDepositRequestForERC20(amount, tokenAddress, user string, decimals int) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  *tokens.NewSignableTokenERC20(decimals, tokenAddress),
		User:   user,
	}
}

func depositERC20(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int) (*eth.Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.DepositERC20(auth, starkKey, assetType, vaultID, quantizedAmount)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC20(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	usersAPI api.UsersApi,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int) (*eth.Transaction, error) {
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

	tnx, err := ethClient.CoreContract.RegisterAndDepositERC20(
		auth,
		common.HexToAddress(etherKey),
		starkKey,
		operatorSignature,
		assetType,
		vaultID,
		quantizedAmount,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
