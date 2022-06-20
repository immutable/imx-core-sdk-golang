package deposits

import (
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk/api"
	"immutable.com/imx-core-sdk/utils/ethereum"
	. "immutable.com/imx-core-sdk/workflows/types"
)

type TokenDeposit interface {
	Execute(e *ethereum.Client, apis *api.APIClient) (*eth.Transaction, error)
}

type ETHDeposit struct {
	Type   TokenTypeEnum
	Amount string
}

type ERC20Deposit struct {
	Type         TokenTypeEnum
	Amount       string
	TokenAddress string
	Symbol       string
}

type ERC721Deposit struct {
	Type         TokenTypeEnum
	TokenId      string
	TokenAddress string
}

func NewETHDeposit(amount string) *ETHDeposit {
	this := ETHDeposit{}
	this.Type = ETHType
	this.Amount = amount
	return &this
}
func NewERC20Deposit(amount string, tokenAddress string, symbol string) *ERC20Deposit {
	this := ERC20Deposit{}
	this.Type = ERC20Type
	this.Amount = amount
	this.TokenAddress = tokenAddress
	this.Symbol = symbol
	return &this
}
func NewERC721Deposit(tokenId string, tokenAddress string) *ERC721Deposit {
	this := ERC721Deposit{}
	this.Type = ERC721Type
	this.TokenId = tokenId
	this.TokenAddress = tokenAddress
	return &this
}
