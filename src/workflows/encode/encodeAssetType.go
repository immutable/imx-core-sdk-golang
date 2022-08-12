package encode

import (
	"context"
	"fmt"
	"math/big"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/tokens"
)

const (
	defaultAssetType  = "asset"
	mintableAssetType = "mintable-asset"
)

// GetEncodedAssetTypeForEth returns encoded AssetType for ETH.
func GetEncodedAssetTypeForEth(
	ctx context.Context,
	encodingAPI api.EncodingApi) (*big.Int, error) {
	encodeAssetRequestToken := api.NewEncodeAssetRequestToken()
	encodeAssetRequestToken.SetType(string(tokens.ETHType))
	return GetEncodedAssetType(ctx, encodingAPI, encodeAssetRequestToken, defaultAssetType)
}

// GetEncodedAssetTypeForERC20 returns encoded AssetType for ERC20 token.
func GetEncodedAssetTypeForERC20(
	ctx context.Context,
	encodingAPI api.EncodingApi,
	tokenID *string, tokenAddress string) (*big.Int, error) {
	tokentype := string(tokens.ERC20Type)
	encodeAssetRequestToken := &api.EncodeAssetRequestToken{
		Data: &api.EncodeAssetTokenData{
			TokenAddress: &tokenAddress,
			TokenId:      tokenID,
		},
		Type: &tokentype,
	}
	return GetEncodedAssetType(ctx, encodingAPI, encodeAssetRequestToken, defaultAssetType)
}

// GetEncodedAssetTypeForERC721 returns encoded AssetType for ERC721 token.
func GetEncodedAssetTypeForERC721(
	ctx context.Context,
	encodingAPI api.EncodingApi,
	tokenID, tokenAddress string) (*big.Int, error) {
	tokentype := string(tokens.ERC721Type)
	encodeAssetRequestToken := &api.EncodeAssetRequestToken{
		Data: &api.EncodeAssetTokenData{
			TokenAddress: &tokenAddress,
			TokenId:      &tokenID,
		},
		Type: &tokentype,
	}
	return GetEncodedAssetType(ctx, encodingAPI, encodeAssetRequestToken, defaultAssetType)
}

// GetEncodedMintableAssetTypeForERC721 returns encoded MintableAssetType for ERC721 token.
func GetEncodedMintableAssetTypeForERC721(
	ctx context.Context,
	encodingAPI api.EncodingApi,
	tokenID, tokenAddress, blueprint string) (*big.Int, error) {
	tokentype := string(tokens.ERC721Type)
	encodeAssetRequestToken := &api.EncodeAssetRequestToken{
		Data: &api.EncodeAssetTokenData{
			Id:           &tokenID,
			TokenAddress: &tokenAddress,
			Blueprint:    &blueprint,
		},
		Type: &tokentype,
	}
	return GetEncodedAssetType(ctx, encodingAPI, encodeAssetRequestToken, mintableAssetType)
}

// GetEncodedAssetType performs encoding on asset details to get an AssetType (required for stark contract request)
func GetEncodedAssetType(
	ctx context.Context,
	encodingAPI api.EncodingApi,
	encodeAssetRequestToken *api.EncodeAssetRequestToken,
	assetType string) (*big.Int, error) {
	encodeAssetRequest := api.EncodeAssetRequest{
		Token: *encodeAssetRequestToken,
	}
	encodedAssetResponse, httpResp, err := encodingAPI.EncodeAsset(ctx, assetType).EncodeAssetRequest(encodeAssetRequest).Execute()
	if err != nil {
		return nil, fmt.Errorf("error when calling `Encoding.EncodeAsset`: %v, HTTP response body: %v", err, httpResp.Body)
	}

	encodedAssetType, ok := new(big.Int).SetString(encodedAssetResponse.AssetType, 10)
	if !ok {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v", encodedAssetResponse.AssetType)
	}
	return encodedAssetType, nil
}
