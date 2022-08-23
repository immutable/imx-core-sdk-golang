package tokens

import (
	"github.com/immutable/imx-core-sdk-golang/generated/api"
)

type TokenType string

const (
	ETHType    TokenType = "ETH"
	ERC20Type  TokenType = "ERC20"
	ERC721Type TokenType = "ERC721"
)

// NewSignableTokenEth returns a new ETH type token.
// https://docs.x.github.com/immutable/docs/token-data-object#type-eth
func NewSignableTokenEth() *api.SignableToken {
	var tokenType = string(ETHType)
	return &api.SignableToken{
		Data: map[string]interface{}{
			"decimals": 18,
		},
		Type: &tokenType,
	}
}

// NewSignableTokenERC20 returns a new ERC20 type token.
// https://docs.x.github.com/immutable/docs/token-data-object#type-erc20
func NewSignableTokenERC20(decimals int, tokenAddress string) *api.SignableToken {
	var tokenType = string(ERC20Type)
	return &api.SignableToken{
		Data: map[string]interface{}{
			"decimals":      decimals,
			"token_address": tokenAddress,
		},
		Type: &tokenType,
	}
}

// NewSignableTokenERC721 returns a new ERC721 type token.
// https://docs.x.github.com/immutable/docs/token-data-object#type-erc721
func NewSignableTokenERC721(tokenID, tokenAddress string) *api.SignableToken {
	var tokenType = string(ERC721Type)
	return &api.SignableToken{
		Data: map[string]interface{}{
			"token_id":      tokenID,
			"token_address": tokenAddress,
		},
		Type: &tokenType,
	}
}
