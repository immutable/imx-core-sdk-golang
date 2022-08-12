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

// Deposit performs the deposit workflow on the ERC721Deposit.
func (d *ERC721Deposit) Deposit(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, l1signer signers.L1Signer) (*eth.Transaction, error) {
	// Approve whether an amount of token from an account can be spent by a third-party account
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tokenID, ok := new(big.Int).SetString(d.TokenID, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokenID to bigint: %v", d.TokenID)
	}
	ierc721Contract, err := ethClient.NewIERC721Contract(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}
	_, err = ierc721Contract.Approve(auth, ethClient.StarkContractAddress, tokenID)
	if err != nil {
		return nil, err
	}

	// Get signable deposit details
	signableDepositRequest := newSignableDepositRequestForERC721("1", d.TokenID, d.TokenAddress, l1signer.GetAddress())
	signableDeposit, err := getSignableDeposit(ctx, clientAPI.DepositsApi, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := encode.GetEncodedAssetTypeForERC721(ctx, clientAPI.EncodingApi, d.TokenID, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	// Passing starkKeyHex to register method because it may be padded and converting back from Int loses the padding
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
		return depositERC721(ctx, ethClient, l1signer, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenID)
	} else {
		return registerAndDepositERC721(ctx, ethClient, l1signer, clientAPI.UsersApi, starkKeyHex, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenID)
	}
}

func newSignableDepositRequestForERC721(amount, tokenID, tokenAddress, user string) *api.GetSignableDepositRequest {
	return &api.GetSignableDepositRequest{
		Amount: amount,
		Token:  *tokens.NewSignableTokenERC721(tokenID, tokenAddress),
		User:   user,
	}
}

func depositERC721(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	tokenID *big.Int) (*eth.Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.DepositNft(auth, starkPublicKey, assetType, vaultID, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC721(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	usersAPI api.UsersApi,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	tokenID *big.Int) (*eth.Transaction, error) {
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
	tnx, err := ethClient.RegistrationContract.RegisterAndDepositNft(
		auth,
		common.HexToAddress(etherKey),
		starkKey,
		operatorSignature,
		assetType,
		vaultID,
		tokenID,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
