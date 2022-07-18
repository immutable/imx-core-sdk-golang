package withdrawals

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
	"math/big"
)

func (w *ERC721Withdrawal) withdrawMintedNft(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*Transaction, error) {

	assetType, err := helpers.GetEncodedAssetTypeForERC721(ctx, api, w.TokenId, w.TokenAddress)
	if err != nil {
		return nil, err
	}

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s\n", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	tokenId, ok := new(big.Int).SetString(w.TokenId, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokenId to bigint: %v\n", w.TokenId)
	}

	if isRegistered {
		return withdrawMintedNft(ctx, ethClient, l1signer, starkKey, assetType, tokenId)
	} else {
		return registerAndWithdrawMintedNft(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, assetType, tokenId)
	}
}

func withdrawMintedNft(ctx context.Context, ethClient *ethereum.Client, l1signer signers.L1Signer, starkKey, assetType, tokenId *big.Int) (*Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.WithdrawNft(auth, starkKey, assetType, tokenId)
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
	tokenId *big.Int,
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
	tnx, err := ethClient.RegistrationContract.RegisterAndWithdrawNft(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, tokenId)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
