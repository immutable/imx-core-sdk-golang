package deposits

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/encoding"
	"immutable.com/imx-core-sdk-golang/api/client/tokens"
	"immutable.com/imx-core-sdk-golang/api/client/users"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

func (d *ERC20Deposit) Execute(ctx context.Context, ethClient *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) (*Transaction, error) {
	if d.Type != types.ERC20Type {
		return nil, errors.New("invalid token type")
	}

	// Get decimals for this specific ERC20
	getTokenParams := tokens.NewGetTokenParamsWithContext(ctx)
	getTokenParams.SetAddress(d.TokenAddress)
	token, err := apis.Tokens.GetToken(getTokenParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Tokens.GetToken`: %v", err)
	}

	decimals, err := utils.FromStringToDecimal(*token.GetPayload().Decimals)
	if err != nil {
		return nil, fmt.Errorf("error parsing token decimals: %v\n", err)
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
	signableDeposit, err := GetSignableDeposit(ctx, apis.Deposits, signableDepositRequest)
	if err != nil {
		return nil, err
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	encodeParams := encoding.NewEncodeAssetParamsWithContext(ctx)
	encodeParams.SetAssetType("asset")
	encodeParams.SetEncodeAssetRequest(&models.EncodeAssetRequest{
		Token: &models.EncodeAssetRequestToken{
			Data: &models.EncodeAssetTokenData{
				TokenAddress: d.TokenAddress,
			},
			Type: string(d.Type),
		},
	})
	encodedAsset, err := apis.Encoding.EncodeAsset(encodeParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Encoding.EncodeAsset`: %v", err)
	}

	assetType, ok := new(big.Int).SetString(*encodedAsset.GetPayload().AssetType, 10)
	if !ok {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v\n", *encodedAsset.GetPayload().AssetType)
	}
	starkKey, err := utils.HexToInt(*signableDeposit.StarkKey)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v\n", signableDeposit.StarkKey)
	}

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(&bind.CallOpts{Context: ctx}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return depositERC20(ctx, ethClient, l1signer, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, amount)
	} else {
		return registerAndDepositERC20(ctx, ethClient, l1signer, apis.Users, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, amount)
	}
}

func depositERC20(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
) (*Transaction, error) {
	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := ethClient.CoreContract.DepositERC20(auth, starkPublicKey, assetType, vaultId, quantizedAmount)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC20(
	ctx context.Context,
	ethClient *ethereum.Client,
	l1signer signers.L1Signer,
	userApi users.ClientService,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
) (*Transaction, error) {
	etherKey := l1signer.GetAddress()
	starkKey := hexutil.EncodeBig(starkPublicKey)
	params := users.NewGetSignableRegistrationParamsWithContext(ctx)
	params.SetGetSignableRegistrationRequest(&models.GetSignableRegistrationRequest{
		EtherKey: &etherKey,
		StarkKey: &starkKey,
	})
	signableRegistration, err := userApi.GetSignableRegistration(params)
	if err != nil {
		return nil, fmt.Errorf("error when calling `UserApi.GetSignableRegistration`: %v", err)
	}

	auth, err := ethClient.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(*signableRegistration.GetPayload().OperatorSignature)
	if err != nil {
		return nil, err
	}

	tnx, err := ethClient.CoreContract.RegisterAndDepositERC20(
		auth,
		common.HexToAddress(l1signer.GetAddress()),
		starkPublicKey,
		operatorSignature,
		assetType,
		vaultId,
		quantizedAmount,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
