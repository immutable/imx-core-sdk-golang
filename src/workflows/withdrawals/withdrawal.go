package withdrawals

import (
	"context"

	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
)

type TokenWithdrawal interface {
	CompleteWithdrawal(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, l1signer signers.L1Signer, starkPublicKey string) (*eth.Transaction, error)
}

// ERC20Withdrawal implements TokenWithdrawal. Used for withdrawal of ERC20 Tokens.
type ERC20Withdrawal struct {
	TokenAddress string
}

// ERC721Withdrawal implements TokenWithdrawal. Used for withdrawal of ERC721 Tokens.
type ERC721Withdrawal struct {
	TokenID      string
	TokenAddress string
}

// NewERC20Withdrawal instantiates a new ERC20Withdrawal object with the given tokenAddress.
func NewERC20Withdrawal(tokenAddress string) *ERC20Withdrawal {
	this := ERC20Withdrawal{}
	this.TokenAddress = tokenAddress
	return &this
}

// NewERC721Withdrawal instantiates a new ERC721Withdrawal object with the given tokenId and tokenAddress.
func NewERC721Withdrawal(tokenID, tokenAddress string) *ERC721Withdrawal {
	this := ERC721Withdrawal{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}
