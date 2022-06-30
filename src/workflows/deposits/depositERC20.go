package deposits

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/deposits"
	"immutable.com/imx-core-sdk-golang/api/client/encoding"
	"immutable.com/imx-core-sdk-golang/api/client/tokens"
	"immutable.com/imx-core-sdk-golang/api/client/users"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

func (d *ERC20Deposit) Execute(ethClient *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) (*Transaction, error) {
	if d.Type != types.ERC20Type {
		return nil, errors.New("invalid token type")
	}
	ctx := context.Background()

	// Get decimals for this specific ERC20
	getTokenParams := tokens.NewGetTokenParams()
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
	amountStr := amount.String()
	user := l1signer.GetAddress().Hex()
	getSignableDepositRequest := &models.GetSignableDepositRequest{
		Amount: &amountStr,
		Token: &models.SignableToken{
			Data: map[string]interface{}{
				"decimals":      *decimals,
				"token_address": d.TokenAddress,
			},
			Type: string(d.Type),
		},
		User: &user,
	}
	signableDepositParams := deposits.NewGetSignableDepositParams()
	signableDepositParams.SetGetSignableDepositRequest(getSignableDepositRequest)
	signableDepositOK, err := apis.Deposits.GetSignableDeposit(signableDepositParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Deposits.GetSignableDeposit`: %v", err)
	}
	signableDeposit := signableDepositOK.GetPayload()

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	encodeAssetRequest := &models.EncodeAssetRequest{
		Token: &models.EncodeAssetRequestToken{
			Data: &models.EncodeAssetTokenData{
				TokenAddress: d.TokenAddress,
			},
			Type: string(d.Type),
		},
	}
	encodeParams := encoding.NewEncodeAssetParams()
	encodeParams.SetAssetType("asset")
	encodeParams.SetEncodeAssetRequest(encodeAssetRequest)
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

	isRegistered, _ := ethClient.RegistrationContract.IsRegistered(nil, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return depositERC20(ctx, ethClient, l1signer, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, amount)
	} else {
		return registerAndDepositERC20(
			ctx,
			ethClient,
			l1signer,
			apis.Users,
			starkKey,
			big.NewInt(*signableDeposit.VaultID),
			assetType,
			amount,
		)
	}
}

func depositERC20(
	ctx context.Context,
	e *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
) (*Transaction, error) {
	auth, err := e.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.DepositERC20(auth, starkPublicKey, assetType, vaultId, quantizedAmount)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC20(
	ctx context.Context,
	e *ethereum.Client,
	l1signer signers.L1Signer,
	userApi users.ClientService,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	quantizedAmount *big.Int,
) (*Transaction, error) {
	etherKey := l1signer.GetAddress().Hex()
	starkKey := hexutil.EncodeBig(starkPublicKey)
	registrationRequest := &models.GetSignableRegistrationRequest{
		EtherKey: &etherKey,
		StarkKey: &starkKey,
	}
	params := users.NewGetSignableRegistrationParams()
	params.SetGetSignableRegistrationRequest(registrationRequest)
	signableRegistration, err := userApi.GetSignableRegistration(params)
	if err != nil {
		return nil, fmt.Errorf("error when calling `UserApi.GetSignableRegistration`: %v", err)
	}

	auth, err := e.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(*signableRegistration.GetPayload().OperatorSignature)
	if err != nil {
		return nil, err
	}

	tnx, err := e.CoreContract.RegisterAndDepositERC20(
		auth,
		l1signer.GetAddress(),
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
