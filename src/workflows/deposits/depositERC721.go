package deposits

import (
	"context"
	"errors"
	"fmt"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk/api"
	"immutable.com/imx-core-sdk/utils"
	"immutable.com/imx-core-sdk/utils/ethereum"
	"immutable.com/imx-core-sdk/workflows/types"
	"math/big"
)

func (d *ERC721Deposit) Execute(e *ethereum.Client, apis *api.APIClient) (*ethtypes.Transaction, error) {
	if d.Type != types.ERC721Type {
		return nil, errors.New("invalid token type")
	}
	ctx := context.Background()

	// Approve whether an amount of token from an account can be spent by a third-party account
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tokenId, ok := utils.StringToBigInt(d.TokenId)
	if !ok {
		return nil, fmt.Errorf("error converting tokentId to bigint: %v\n", d.TokenId)
	}
	approve, err := e.IERC721Contract.Approve(auth, e.StarkContractAddress, tokenId)
	if err != nil {
		return nil, err
	}
	// TODO: check if this needs to be signed before sending also populate gas fields?
	e.Client.SendTransaction(ctx, approve)

	// Get signable deposit details
	const amount = "1"
	getSignableDepositRequest := api.NewGetSignableDepositRequest(amount, *api.NewSignableToken(), e.GetAddress().Hex())
	getSignableDepositRequest.Token.SetType(string(d.Type))
	getSignableDepositRequest.Token.SetData(map[string]interface{}{
		"token_id":      d.TokenId,
		"token_address": d.TokenAddress,
	})

	signableDeposit, resp, err := apis.DepositsApi.GetSignableDeposit(ctx).GetSignableDepositRequest(*getSignableDepositRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `DepositsApi.GetSignableDeposit`: %v\nFull HTTP response: %v", err, resp)
	}

	// Perform encoding on asset details to get an assetType (required for stark contract request)
	encodeAssetRequest := api.NewEncodeAssetRequest(*api.NewEncodeAssetRequestToken())
	encodeAssetRequest.Token.SetType(string(d.Type))
	encodeAssetRequest.Token.SetData(api.EncodeAssetTokenData{
		TokenAddress: &d.TokenAddress,
		TokenId:      &d.TokenId,
	})
	encodingResult, resp, err := apis.EncodingApi.EncodeAsset(ctx, "asset").EncodeAssetRequest(*encodeAssetRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `EncodingApi.EncodeAsset`: %v\nFull HTTP response: %v", err, resp)
	}

	// TODO: verify this is the correct type conversion
	assetType, err := utils.HexToInt(encodingResult.AssetType)
	if err != nil {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v\n", encodingResult.AssetType)
	}
	starkKey, err := utils.HexToInt(signableDeposit.StarkKey)
	if err != nil {
		return nil, fmt.Errorf("error converting StarkKey to bigint: %v\n", signableDeposit.StarkKey)
	}
	// TODO: verify if it's ok to set CallOpts to nil
	isRegistered, err := e.RegistrationContract.IsRegistered(nil, starkKey)
	if err != nil {
		return nil, fmt.Errorf("Error when calling `Registration.IsRegistered`: %v\n", err)
	}

	// TODO: verify if big.NewInt(int64(signableDeposit.VaultId)) is the correct conversion
	if isRegistered {
		return depositERC721(ctx, e, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenId)
	} else {
		return registerAndDepositERC721(ctx, e, apis.UsersApi, starkKey, big.NewInt(int64(signableDeposit.VaultId)), assetType, tokenId)
	}
}

func depositERC721(ctx context.Context, e *ethereum.Client, starkPublicKey *big.Int, vaultId *big.Int, assetType *big.Int, tokenId *big.Int) (*ethtypes.Transaction, error) {
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}
	tnx, err := e.CoreContract.DepositNft(auth, starkPublicKey, assetType, vaultId, tokenId)
	if err != nil {
		return nil, err
	}

	tnx, err = e.SignTransaction(ctx, tnx)
	if err != nil {
		return nil, err
	}

	// TODO: verify if tnx needs to be populated with gas/nounce etc and signed before sending
	if err = e.Client.SendTransaction(ctx, tnx); err != nil {
		return nil, err
	}
	return tnx, nil
}

func registerAndDepositERC721(ctx context.Context, e *ethereum.Client, userApi *api.UsersApiService, starkPublicKey *big.Int, vaultId *big.Int, assetType *big.Int, tokenId *big.Int) (*ethtypes.Transaction, error) {
	registrationRequest := api.NewGetSignableRegistrationRequest(e.GetAddress().Hex(), starkPublicKey.String())
	signableRegistration, resp, err := userApi.GetSignableRegistration(ctx).GetSignableRegistrationRequest(*registrationRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("Error when calling `UserApi.GetSignableRegistration`: %v\nFull HTTP response: %v", err, resp)
	}
	auth, err := e.BuildTransactOpts(ctx)
	if err != nil {
		return nil, err
	}

	operatorSignature, err := utils.HexToByteArray(signableRegistration.OperatorSignature)
	if err != nil {
		return nil, err
	}
	tnx, err := e.RegistrationContract.RegisterAndDepositNft(auth, e.GetAddress(), starkPublicKey, operatorSignature, assetType, vaultId, tokenId)
	if err != nil {
		return nil, err
	}

	tnx, err = e.SignTransaction(ctx, tnx)
	if err != nil {
		return nil, err
	}

	if err = e.Client.SendTransaction(ctx, tnx); err != nil {
		return nil, err
	}
	return tnx, nil
}
