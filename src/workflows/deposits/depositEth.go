package deposits

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	. "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/deposits"
	"immutable.com/imx-core-sdk-golang/api/client/encoding"
	"immutable.com/imx-core-sdk-golang/api/client/users"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

func (d *ETHDeposit) Execute(e *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) (*Transaction, error) {
	if d.Type != types.ETHType {
		return nil, errors.New("invalid token type")
	}
	ctx := context.Background()

	amount, err := utils.ParseEtherToWei(d.Amount)
	if err != nil {
		return nil, fmt.Errorf("Error when parsing deposit amount: %v\n", err)
	}
	amountStr := amount.String()
	user := l1signer.GetAddress().Hex()
	getSignableDepositRequest := &models.GetSignableDepositRequest{
		Amount: &amountStr,
		Token: &models.SignableToken{
			Data: map[string]interface{}{
				"decimals": 18,
			},
			Type: string(types.ETHType),
		},
		User: &user,
	}
	params := deposits.NewGetSignableDepositParams()
	params.SetGetSignableDepositRequest(getSignableDepositRequest)
	signableDepositOK, err := apis.Deposits.GetSignableDeposit(params)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Deposits.GetSignableDeposit`: %v", err)
	}
	signableDeposit := signableDepositOK.GetPayload()

	encodeAssetRequest := &models.EncodeAssetRequest{
		Token: &models.EncodeAssetRequestToken{
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
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v\n", *signableDeposit.StarkKey)
	}
	isRegistered, _ := e.RegistrationContract.IsRegistered(&bind.CallOpts{From: l1signer.GetAddress()}, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return depositEth(
			ctx,
			e,
			l1signer,
			starkKey,
			big.NewInt(*signableDeposit.VaultID),
			assetType,
			amount,
		)
	} else {
		return registerAndDepositEth(
			ctx,
			e,
			l1signer,
			apis.Users,
			starkKey,
			big.NewInt(*signableDeposit.VaultID),
			assetType,
			amount,
		)
	}
}

func registerAndDepositEth(
	ctx context.Context,
	e *ethereum.Client,
	l1signer signers.L1Signer,
	userApi users.ClientService,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	amount *big.Int,
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
		return nil, fmt.Errorf("Error when calling `UserApi.GetSignableRegistration`: %v", err)
	}
	auth, err := e.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(*signableRegistration.GetPayload().OperatorSignature)
	if err != nil {
		return nil, err
	}
	auth.Value = amount
	tnx, err := e.CoreContract.RegisterAndDepositEth(
		auth,
		l1signer.GetAddress(),
		starkPublicKey,
		operatorSignature,
		assetType,
		vaultId,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func depositEth(
	ctx context.Context,
	e *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	amount *big.Int,
) (*Transaction, error) {
	auth, err := e.BuildTransactOpts(ctx, l1signer)
	auth.Value = amount
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.Deposit(auth, starkPublicKey, assetType, vaultId)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
