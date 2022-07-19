package utils

import (
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

// NewSignableTokenEth returns a new ETH type token.
// https://docs.x.immutable.com/docs/token-data-object#type-eth
func NewSignableTokenEth() *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"decimals": 18,
		},
		Type: string(types.ETHType),
	}
}

// NewSignableTokenERC20 returns a new ERC20 type token.
// https://docs.x.immutable.com/docs/token-data-object#type-erc20
func NewSignableTokenERC20(decimals, tokenAddress string) *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"decimals":      decimals,
			"token_address": tokenAddress,
		},
		Type: string(types.ERC20Type),
	}
}

// NewSignableTokenERC721 returns a new ERC721 type token.
// https://docs.x.immutable.com/docs/token-data-object#type-erc721
func NewSignableTokenERC721(tokenID, tokenAddress string) *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"token_id":      tokenID,
			"token_address": tokenAddress,
		},
		Type: string(types.ERC721Type),
	}
}
