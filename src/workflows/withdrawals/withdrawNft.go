package withdrawals

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-openapi/runtime"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/mints"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
	"math/big"
)

// CompleteWithdrawal performs the complete withdrawal workflow on ERC721Withdrawal
func (w *ERC721Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*Transaction, error) {
	if w.Type != types.ERC721Type {
		return nil, errors.New("invalid token type")
	}

	getMintableTokenParams := mints.NewGetMintableTokenDetailsByClientTokenIDParamsWithContext(ctx)
	getMintableTokenParams.SetTokenID(w.TokenId)
	getMintableTokenParams.SetTokenAddress(w.TokenAddress)
	mintableTokenResponse, err := api.Mints.GetMintableTokenDetailsByClientTokenID(getMintableTokenParams)
	if err != nil {
		if err.(*runtime.APIError).IsCode(404) {
			// Token is already minted on L1
			return w.withdrawMintedNft(ctx, ethClient, api, l1signer, starkKeyHex)
		}
		return nil, fmt.Errorf("error when calling `Mints.GetMintableTokenDetailsByClientTokenID`: %v", err)
	}

	blueprint := *mintableTokenResponse.GetPayload().Blueprint
	mintingBlob := getMintingBlob(w.TokenId, blueprint)
	assetType, err := helpers.GetEncodedMintableAssetTypeForERC721(ctx, api, w.TokenId, w.TokenAddress, blueprint)
	if err != nil {
		return nil, err
	}

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s\n", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	if isRegistered {
		return withdrawAndMintNft(ctx, ethClient, l1signer, starkKey, assetType, mintingBlob)
	} else {
		return registerAndWithdrawAndMintNft(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, assetType, mintingBlob)
	}
}

func withdrawAndMintNft(ctx context.Context, ethClient *ethereum.Client, l1signer signers.L1Signer, starkKey, assetType *big.Int, mintingBlob []byte) (*Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.WithdrawAndMint(auth, starkKey, assetType, mintingBlob)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndWithdrawAndMintNft(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	api *client.ImmutableXAPI,
	starkKeyHex string,
	starkKey *big.Int,
	assetType *big.Int,
	mintingBlob []byte,
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
	tnx, err := ethClient.RegistrationContract.RegsiterAndWithdrawAndMint(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, mintingBlob)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func getMintingBlob(tokenId, blueprint string) []byte {
	return []byte(fmt.Sprintf("{%s}:{%s}", tokenId, blueprint))
}
