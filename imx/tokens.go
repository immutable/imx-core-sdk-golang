package imx

import (
	"context"

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

/*
GetToken Get details of a token with the given ID

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@param id Token ID
@return TokenDetails
*/
func (c *Client) GetToken(ctx context.Context, id string) (*api.TokenDetails, error) {
	response, httpResponse, err := c.tokensAPI.GetToken(ctx, id).Execute()
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}

/*
NewListTokensRequest Creates a new ApiListTokensRequest object with required params.

@param ctx context.Context - for cancellation, deadlines, tracing, etc or context.Background().
@return ApiListTokensRequest
*/
func (c *Client) NewListTokensRequest(ctx context.Context) api.ApiListTokensRequest {
	return c.tokensAPI.ListTokens(ctx)
}

/*
ListTokens Gets a list of tokens

@param req the api request struct with all params populated to make the request
@return ListTokensResponse
*/
func (c *Client) ListTokens(req *api.ApiListTokensRequest) (*api.ListTokensResponse, error) {
	response, httpResponse, err := c.tokensAPI.ListTokensExecute(*req)
	if err != nil {
		return nil, NewAPIError(httpResponse, err)
	}
	return response, nil
}
