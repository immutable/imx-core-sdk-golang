package utils

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/encoding"
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/workflows/types"
	"math/big"
)

const (
	defaultAssetType  = "asset"
	mintabelAssetType = "mintable-asset"
)

// GetEncodedAssetTypeForEth returns encoded AssetType for ETH.
func GetEncodedAssetTypeForEth(
	ctx context.Context,
	api *client.ImmutableXAPI,
) (*big.Int, error) {
	encodeAssetRequestToken := &models.EncodeAssetRequestToken{
		Type: string(types.ETHType),
	}
	return GetEncodedAssetType(ctx, api, encodeAssetRequestToken, defaultAssetType)
}

// GetEncodedAssetTypeForERC20 returns encoded AssetType for ERC20 token.
func GetEncodedAssetTypeForERC20(
	ctx context.Context,
	api *client.ImmutableXAPI,
	tokenId, tokenAddress string,
) (*big.Int, error) {
	encodeAssetRequestToken := &models.EncodeAssetRequestToken{
		Data: &models.EncodeAssetTokenData{
			TokenAddress: tokenAddress,
			TokenID:      tokenId,
		},
		Type: string(types.ERC20Type),
	}
	return GetEncodedAssetType(ctx, api, encodeAssetRequestToken, defaultAssetType)
}

// GetEncodedAssetTypeForERC721 returns encoded AssetType for ERC721 token.
func GetEncodedAssetTypeForERC721(
	ctx context.Context,
	api *client.ImmutableXAPI,
	tokenId, tokenAddress string,
) (*big.Int, error) {
	encodeAssetRequestToken := &models.EncodeAssetRequestToken{
		Data: &models.EncodeAssetTokenData{
			TokenAddress: tokenAddress,
			TokenID:      tokenId,
		},
		Type: string(types.ERC721Type),
	}
	return GetEncodedAssetType(ctx, api, encodeAssetRequestToken, defaultAssetType)
}

// GetEncodedMintableAssetTypeForERC721 returns encoded MintableAssetType for ERC721 token.
func GetEncodedMintableAssetTypeForERC721(
	ctx context.Context,
	api *client.ImmutableXAPI,
	tokenId, tokenAddress, blueprint string,
) (*big.Int, error) {
	encodeAssetRequestToken := &models.EncodeAssetRequestToken{
		Data: &models.EncodeAssetTokenData{
			ID:           tokenId,
			TokenAddress: tokenAddress,
			Blueprint:    blueprint,
		},
		Type: string(types.ERC721Type),
	}
	return GetEncodedAssetType(ctx, api, encodeAssetRequestToken, mintabelAssetType)
}

// GetEncodedAssetType performs encoding on asset details to get an AssetType (required for stark contract request)
func GetEncodedAssetType(
	ctx context.Context,
	api *client.ImmutableXAPI,
	encodeAssetRequestToken *models.EncodeAssetRequestToken,
	assetType string,
) (*big.Int, error) {

	encodeParams := encoding.NewEncodeAssetParamsWithContext(ctx)
	encodeParams.SetAssetType(assetType)
	encodeParams.SetEncodeAssetRequest(&models.EncodeAssetRequest{
		Token: encodeAssetRequestToken,
	})
	encodedAsset, err := api.Encoding.EncodeAsset(encodeParams)
	if err != nil {
		return nil, fmt.Errorf("error when calling `Encoding.EncodeAsset`: %v", err)
	}

	encodedAssetType, ok := new(big.Int).SetString(*encodedAsset.GetPayload().AssetType, 10)
	if !ok {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v\n", *encodedAsset.GetPayload().AssetType)
	}
	return encodedAssetType, nil
}
