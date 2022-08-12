package withdrawals

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-openapi/runtime"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/mints"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
	helpers "immutable.com/imx-core-sdk-golang/workflows/utils"
)

// CompleteWithdrawal performs the complete withdrawal workflow on ERC721Withdrawal
func (w *ERC721Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	ethClient *ethereum.Client,
	api *client.ImmutableXAPI,
	l1signer signers.L1Signer,
	starkKeyHex string,
) (*eth.Transaction, error) {
	getMintableTokenParams := mints.NewGetMintableTokenDetailsByClientTokenIDParamsWithContext(ctx)
	getMintableTokenParams.SetTokenID(w.TokenID)
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
	mintingBlob := getMintingBlob(w.TokenID, blueprint)
	assetType, err := helpers.GetEncodedMintableAssetTypeForERC721(ctx, api, w.TokenID, w.TokenAddress, blueprint)
	if err != nil {
		return nil, err
	}

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)

	if isRegistered {
		return withdrawAndMintNft(ctx, ethClient, l1signer, starkKey, assetType, mintingBlob)
	}
	return registerAndWithdrawAndMintNft(ctx, ethClient, l1signer, api, starkKeyHex, starkKey, assetType, mintingBlob)
}

func withdrawAndMintNft(ctx context.Context, ethClient *ethereum.Client, l1signer signers.L1Signer, starkKey, assetType *big.Int, mintingBlob []byte) (*eth.Transaction, error) {
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
	tnx, err := ethClient.RegistrationContract.RegsiterAndWithdrawAndMint(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, mintingBlob)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func getMintingBlob(tokenID, blueprint string) []byte {
	return []byte(fmt.Sprintf("{%s}:{%s}", tokenID, blueprint))
}
