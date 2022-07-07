package utils

import (
	"immutable.com/imx-core-sdk-golang/api/models"
	"immutable.com/imx-core-sdk-golang/workflows/types"
)

func NewSignableTokenEth() *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"decimals": 18,
		},
		Type: string(types.ETHType),
	}
}

func NewSignableTokenERC20(decimals, tokenAddress string) *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"decimals":      decimals,
			"token_address": tokenAddress,
		},
		Type: string(types.ERC20Type),
	}
}

func NewSignableTokenERC721(tokenId, tokenAddress string) *models.SignableToken {
	return &models.SignableToken{
		Data: map[string]interface{}{
			"token_id":      tokenId,
			"token_address": tokenAddress,
		},
		Type: string(types.ERC721Type),
	}
}
