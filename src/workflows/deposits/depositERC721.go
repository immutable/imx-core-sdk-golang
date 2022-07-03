package deposits

import (
	"context"
	"errors"
	"fmt"
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
	"math/big"
)

func (d *ERC721Deposit) Execute(e *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) (*Transaction, error) {
	if d.Type != types.ERC721Type {
		return nil, errors.New("invalid token type")
	}
	ctx := context.Background()

	// Approve whether an amount of token from an account can be spent by a third-party account
	auth, err := e.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tokenId, ok := new(big.Int).SetString(d.TokenId, 10)
	if !ok {
		return nil, fmt.Errorf("error converting tokentId to bigint: %v\n", d.TokenId)
	}
	ierc721Contract, err := e.NewIERC721Contract(ctx, d.TokenAddress)
	if err != nil {
		return nil, err
	}
	_, err = ierc721Contract.Approve(auth, e.StarkContractAddress, tokenId)
	if err != nil {
		return nil, err
	}

	// Get signable deposit details
	var amount = "1"
	user := l1signer.GetAddress().Hex()
	getSignableDepositRequest := &models.GetSignableDepositRequest{
		Amount: &amount,
		Token: &models.SignableToken{
			Data: map[string]interface{}{
				"token_id":      d.TokenId,
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
				TokenID:      d.TokenId,
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

	isRegistered, _ := e.RegistrationContract.IsRegistered(nil, starkKey)
	// Note: if we reach here, it means we are registered off-chain.
	// Above call will return an error user is not registered but this is for on-chain
	// we should swallow this error to allow the register and deposit flow to execute.

	if isRegistered {
		return depositERC721(ctx, e, l1signer, starkKey, big.NewInt(*signableDeposit.VaultID), assetType, tokenId)
	} else {
		return registerAndDepositERC721(
			ctx,
			e,
			l1signer,
			apis.Users,
			starkKey,
			big.NewInt(*signableDeposit.VaultID),
			assetType,
			tokenId,
		)
	}
}

func depositERC721(
	ctx context.Context,
	e *ethereum.Client,
	l1signer signers.L1Signer,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	tokenId *big.Int,
) (*Transaction, error) {
	auth, err := e.BuildTransactOpts(ctx, l1signer)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.DepositNft(auth, starkPublicKey, assetType, vaultId, tokenId)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC721(
	ctx context.Context,
	e *ethereum.Client,
	l1signer signers.L1Signer,
	userApi users.ClientService,
	starkPublicKey *big.Int,
	vaultId *big.Int,
	assetType *big.Int,
	tokenId *big.Int,
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
	tnx, err := e.RegistrationContract.RegisterAndDepositNft(
		auth,
		l1signer.GetAddress(),
		starkPublicKey,
		operatorSignature,
		assetType,
		vaultId,
		tokenId,
	)
	if err != nil {
		return nil, err
	}
	return tnx, nil
}
