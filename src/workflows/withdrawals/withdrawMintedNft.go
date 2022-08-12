package withdrawals

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/encode"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
)

func (w *ERC721Withdrawal) withdrawMintedNft(
	ctx context.Context,
	ethClient *ethereum.Client,
	clientAPI *api.APIClient,
	l1signer signers.L1Signer,
	starkKeyHex string) (*eth.Transaction, error) {
	assetType, err := encode.GetEncodedAssetTypeForERC721(ctx, clientAPI.EncodingApi, w.TokenID, w.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s", starkKeyHex)
	}

	isRegistered, err := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	if err != nil {
		return nil, fmt.Errorf("error when calling 'ethClient.RegistrationContract.IsRegistered': %v", err)
	}

	tokenID, ok := new(big.Int).SetString(w.TokenID, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokenID to bigint: %v", w.TokenID)
	}

	if isRegistered {
		return withdrawMintedNft(ctx, ethClient, l1signer, starkKey, assetType, tokenID)
	} else {
		return registerAndWithdrawMintedNft(ctx, ethClient, l1signer, clientAPI.UsersApi, starkKeyHex, starkKey, assetType, tokenID)
	}
}

func withdrawMintedNft(ctx context.Context, ethClient *ethereum.Client, l1signer signers.L1Signer, starkKey, assetType, tokenID *big.Int) (*eth.Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.WithdrawNft(auth, starkKey, assetType, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndWithdrawMintedNft(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	usersAPI api.UsersApi,
	starkKeyHex string,
	starkKey *big.Int,
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
	tnx, err := ethClient.RegistrationContract.RegisterAndWithdrawNft(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
