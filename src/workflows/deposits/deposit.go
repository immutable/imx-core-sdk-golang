package deposits

import (
	"context"

	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
)

type TokenDeposit interface {
	Deposit(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, l1signer signers.L1Signer) (*eth.Transaction, error)
}

// ETHDeposit implements TokenDeposit. Used to deposit Eth Tokens.
type ETHDeposit struct {
	Amount string
}

// ERC20Deposit implements TokenDeposit. Used to deposit ERC20 Tokens.
type ERC20Deposit struct {
	Amount       string
	TokenAddress string
}

// ERC721Deposit implements TokenDeposit. Used to deposit ERC721 Tokens.
type ERC721Deposit struct {
	TokenID      string
	TokenAddress string
}

// NewETHDeposit instantiates a new ETHDeposit object with the given amount.
func NewETHDeposit(amount string) *ETHDeposit {
	this := ETHDeposit{}
	this.Amount = amount
	return &this
}

// NewERC20Deposit instantiates a new ERC20Deposit object with given amount and tokenAddress.
func NewERC20Deposit(amount, tokenAddress string) *ERC20Deposit {
	this := ERC20Deposit{}
	this.Amount = amount
	this.TokenAddress = tokenAddress
	return &this
}

// NewERC721Deposit instantiates a new ERC721Deposit object with given tokenID and tokenAddress.
func NewERC721Deposit(tokenID, tokenAddress string) *ERC721Deposit {
	this := ERC721Deposit{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}
