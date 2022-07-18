package deposits

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/tokens"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
)

// Execute performs the deposit workflow on the ERC20Deposit.
func (d *ERC20Deposit) Execute(ctx context.Context, ethClient *ethereum.Client, api *client.ImmutableXAPI, l1signer signers.L1Signer) (*eth.Transaction, error) {
	if d.Type != types.ERC20Type {
		return nil, errors.New("invalid token type")
	}

	// Get decimals for this specific ERC20
	getTokenParams := tokens.NewGetTokenParamsWithContext(ctx)
	getTokenParams.SetAddress(d.TokenAddress)
	token, err := api.Tokens.GetToken(getTokenParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Tokens.GetToken`: %v", err)
	}

	decimals, err := utils.FromStringToDecimal(*token.GetPayload().Decimals)
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
	signableDepositRequest := NewSignableDepositRequestForERC20(amount.String(), d.TokenAddress, l1signer.GetAddress(), strconv.Itoa(int(*decimals)))
	signableDeposit, err := GetSignableDeposit(ctx, api.Deposits, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := helpers.GetEncodedAssetTypeForERC20(ctx, api, "", d.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKeyHex := *signableDeposit.StarkKey
	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return depositERC20(ctx, ethClient, l1signer, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, amount)
	} else {
		return registerAndDepositERC20(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, amount)
	}
}

func depositERC20(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
) (*eth.Transaction, error) {
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
	api *client.ImmutableXAPI,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
) (*eth.Transaction, error) {
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
