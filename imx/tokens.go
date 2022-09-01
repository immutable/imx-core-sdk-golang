package imx

import (
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

const (
	ETHTokenType    = "ETH"
	ERC20TokenType  = "ERC20"
	ERC721TokenType = "ERC721"
)

func s(constant string) *string {
	return &constant
}

// SignableETHToken returns a new ETH type token.
// https://docs.x.immutable.com/docs/token-data-object#type-eth
func SignableETHToken() api.SignableToken {
	return api.SignableToken{
		Data: map[string]interface{}{
			"decimals": 18,
		},
		Type: s(ETHTokenType),
	}
}

// SignableERC20Token returns a new ERC20 type token.
// https://docs.x.immutable.com/docs/token-data-object#type-erc20
func SignableERC20Token(decimals int, tokenAddress string) api.SignableToken {
	return api.SignableToken{
		Data: map[string]interface{}{
			"decimals":      decimals,
			"token_address": tokenAddress,
		},
		Type: s(ERC20TokenType),
	}
}

// SignableERC721Token returns a new ERC721 type token.
// https://docs.x.immutable.com/docs/token-data-object#type-erc721
func SignableERC721Token(tokenID, tokenAddress string) api.SignableToken {
	return api.SignableToken{
		Data: map[string]interface{}{
			"token_id":      tokenID,
			"token_address": tokenAddress,
		},
		Type: s(ERC721TokenType),
	}
}
