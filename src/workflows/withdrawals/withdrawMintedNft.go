package withdrawals

import (
	"context"
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
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
)

func (w *ERC721Withdrawal) withdrawMintedNft(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*eth.Transaction, error) {
	assetType, err := helpers.GetEncodedAssetTypeForERC721(ctx, api, w.TokenID, w.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	tokenID, ok := new(big.Int).SetString(w.TokenID, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokenID to bigint: %v", w.TokenID)
	}

	if isRegistered {
		return withdrawMintedNft(ctx, ethClient, l1signer, starkKey, assetType, tokenID)
	} else {
		return registerAndWithdrawMintedNft(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, assetType, tokenID)
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
	api *client.ImmutableXAPI,
	starkKeyHex string,
	starkKey *big.Int,
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
	tnx, err := ethClient.RegistrationContract.RegisterAndWithdrawNft(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, tokenID)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
