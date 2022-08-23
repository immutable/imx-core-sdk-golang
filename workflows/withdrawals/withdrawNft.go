package withdrawals

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	eth "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-openapi/runtime"
	"github.com/immutable/imx-core-sdk-golang/ethereum"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/internal/encode"
	"github.com/immutable/imx-core-sdk-golang/internal/utils"
	"github.com/immutable/imx-core-sdk-golang/signers"
	"github.com/immutable/imx-core-sdk-golang/workflows/registration"
)

// CompleteWithdrawal performs the completion step of the withdrawal process for ERC721 token.
func (w *ERC721Withdrawal) CompleteWithdrawal(
	ctx context.Context,
	ethClient *ethereum.Client,
	clientAPI *api.APIClient,
	l1signer signers.L1Signer,
	starkKeyHex string) (*eth.Transaction, error) {
	mintableTokenResponse, httpResp, err := clientAPI.MintsApi.GetMintableTokenDetailsByClientTokenId(ctx, w.TokenAddress, w.TokenID).Execute()
	if err != nil {
		if err.(*runtime.APIError).IsCode(404) {
			// Token is already minted on L1
			return w.withdrawMintedNft(ctx, ethClient, clientAPI, l1signer, starkKeyHex)
		}
		return nil, fmt.Errorf("error when calling `clientAPI.MintsApi.GetMintableTokenDetailsByClientTokenId.Execute`: %v, HTTP response body: %v", err, httpResp.Body)
	}

	blueprint := mintableTokenResponse.Blueprint
	mintingBlob := getMintingBlob(w.TokenID, blueprint)
	assetType, err := encode.GetEncodedMintableAssetTypeForERC721(ctx, clientAPI.EncodingApi, w.TokenID, w.TokenAddress, blueprint)
	if err != nil {
		return nil, err
	}

	starkKey, err := utils.HexToInt(starkKeyHex)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKeyHex to bigint: %s", starkKeyHex)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and withdraw flow to execute.

	if isRegistered {
		return withdrawAndMintNft(ctx, ethClient, l1signer, starkKey, assetType, mintingBlob)
	}
	return registerAndWithdrawAndMintNft(ctx, ethClient, l1signer, clientAPI.UsersApi, starkKeyHex, starkKey, assetType, mintingBlob)
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
	usersAPI api.UsersApi,
	starkKeyHex string,
	starkKey *big.Int,
	assetType *big.Int,
	mintingBlob []byte) (*eth.Transaction, error) {
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
	tnx, err := ethClient.RegistrationContract.RegsiterAndWithdrawAndMint(auth, common.HexToAddress(etherKey), starkKey, operatorSignature, assetType, mintingBlob)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func getMintingBlob(tokenID, blueprint string) []byte {
	return []byte(fmt.Sprintf("{%s}:{%s}", tokenID, blueprint))
}
