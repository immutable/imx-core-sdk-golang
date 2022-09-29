package imx

import (
	"context"
	"fmt"
	"math/big"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

const (
	defaultAssetType  = "asset"
	mintableAssetType = "mintable-asset"
)

// ETHAsset returns encoded AssetType for ETH.
func (c *Client) encodeETHAsset(ctx context.Context) (*big.Int, error) {
	encodeAssetRequestToken := api.NewEncodeAssetRequestToken()
	encodeAssetRequestToken.SetType(string(ETHTokenType))
	return c.encodedAssetType(ctx, encodeAssetRequestToken, defaultAssetType)
}

// ERC20Asset returns encoded AssetType for ERC20 token.
func (c *Client) encodeERC20Asset(ctx context.Context, tokenAddress string) (*big.Int, error) {
	encodeAssetRequestToken := &api.EncodeAssetRequestToken{
		Data: &api.EncodeAssetTokenData{
			TokenAddress: &tokenAddress,
		},
		Type: s(ERC20TokenType),
	}
	return c.encodedAssetType(ctx, encodeAssetRequestToken, defaultAssetType)
}

// ERC721Asset returns encoded AssetType for ERC721 token.
func (c *Client) encodeERC721Asset(ctx context.Context, tokenID, tokenAddress string) (*big.Int, error) {
	encodeAssetRequestToken := &api.EncodeAssetRequestToken{
		Data: &api.EncodeAssetTokenData{
			TokenAddress: &tokenAddress,
			TokenId:      &tokenID,
		},
		Type: s(ERC721TokenType),
	}
	return c.encodedAssetType(ctx, encodeAssetRequestToken, defaultAssetType)
}

// MintableERC721Asset returns encoded MintableAssetType for ERC721 token.
func (c *Client) encodeMintableERC721Asset(ctx context.Context, tokenID, tokenAddress, blueprint string) (*big.Int, error) {
	encodeAssetRequestToken := &api.EncodeAssetRequestToken{
		Data: &api.EncodeAssetTokenData{
			Id:           &tokenID,
			TokenAddress: &tokenAddress,
			Blueprint:    &blueprint,
		},
		Type: s(ERC721TokenType),
	}
	return c.encodedAssetType(ctx, encodeAssetRequestToken, mintableAssetType)
}

// encodedAssetType performs encoding on asset details to get an AssetType (required for stark contract request)
func (c *Client) encodedAssetType(
	ctx context.Context,
	encodeAssetRequestToken *api.EncodeAssetRequestToken,
	assetType string) (*big.Int, error) {
	encodeAssetRequest := api.EncodeAssetRequest{
		Token: *encodeAssetRequestToken,
	}
	encodedAssetResponse, httpResponse, err := c.EncodeAsset(ctx, assetType).EncodeAssetRequest(encodeAssetRequest).Execute()
	if err != nil {
		return nil, NewIMXError(httpResponse, err)
	}

	encodedAssetType, ok := new(big.Int).SetString(encodedAssetResponse.AssetType, 10)
	if !ok {
		return nil, fmt.Errorf("error converting encoded asset type to bigint: %v", encodedAssetResponse.AssetType)
	}
	return encodedAssetType, nil
}
