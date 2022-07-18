package deposits

import (
	"context"

	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	imx "immutable.com/imx-core-sdk-golang/workflows/types"
)

type TokenDeposit interface {
	Execute(ctx context.Context, ethClient *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) (*eth.Transaction, error)
}

type ETHDeposit struct {
	Type   imx.TokenTypeEnum
	Amount string
}

type ERC20Deposit struct {
	Type         imx.TokenTypeEnum
	Amount       string
	TokenAddress string
	Symbol       string
}

type ERC721Deposit struct {
	Type         imx.TokenTypeEnum
	TokenID      string
	TokenAddress string
}

func NewETHDeposit(amount string) *ETHDeposit {
	this := ETHDeposit{}
	this.Type = imx.ETHType
	this.Amount = amount
	return &this
}

func NewERC20Deposit(amount, tokenAddress, symbol string) *ERC20Deposit {
	this := ERC20Deposit{}
	this.Type = imx.ERC20Type
	this.Amount = amount
	this.TokenAddress = tokenAddress
	this.Symbol = symbol
	return &this
}

func NewERC721Deposit(tokenID, tokenAddress string) *ERC721Deposit {
	this := ERC721Deposit{}
	this.Type = imx.ERC721Type
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}
