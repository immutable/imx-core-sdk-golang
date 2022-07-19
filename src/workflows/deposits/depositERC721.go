package deposits

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
)

// Deposit performs the deposit workflow on the ERC721Deposit.
func (d *ERC721Deposit) Deposit(ctx context.Context, ethClient *ethereum.Client, api *client.ImmutableXAPI, l1signer signers.L1Signer) (*eth.Transaction, error) {
	if d.Type != types.ERC721Type {
		return nil, errors.New("invalid token type")
	}

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
	signableDepositRequest := NewSignableDepositRequestForERC721("1", d.TokenID, d.TokenAddress, l1signer.GetAddress())
	signableDeposit, err := GetSignableDeposit(ctx, api.Deposits, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	assetType, err := helpers.GetEncodedAssetTypeForERC721(ctx, api, d.TokenID, d.TokenAddress)
	if err != nil {
		return nil, err
	}

	// Passing starkKeyHex to register method because it may be padded and converting back from Int loses the padding
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
		return depositERC721(ctx, ethClient, l1signer, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, tokenID)
	} else {
		return registerAndDepositERC721(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, tokenID)
	}
}

func depositERC721(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	tokenID *big.Int,
) (*eth.Transaction, error) {
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
	api *client.ImmutableXAPI,
	starkKeyHex string,
	starkKey *big.Int,
	vaultID *big.Int,
	assetType *big.Int,
	tokenID *big.Int,
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
