package deposits

import (
	"context"

	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
)

type TokenDeposit interface {
	Deposit(ctx context.Context, ethClient *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) (*eth.Transaction, error)
}

type ETHDeposit struct {
	Amount string
}

type ERC20Deposit struct {
	Amount       string
	TokenAddress string
}

type ERC721Deposit struct {
	TokenID      string
	TokenAddress string
}

func NewETHDeposit(amount string) *ETHDeposit {
	this := ETHDeposit{}
	this.Amount = amount
	return &this
}

func NewERC20Deposit(amount, tokenAddress string) *ERC20Deposit {
	this := ERC20Deposit{}
	this.Amount = amount
	this.TokenAddress = tokenAddress
	return &this
}

func NewERC721Deposit(tokenID, tokenAddress string) *ERC721Deposit {
	this := ERC721Deposit{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}
