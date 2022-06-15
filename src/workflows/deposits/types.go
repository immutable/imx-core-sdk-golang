package deposits

import (
	eth "github.com/ethereum/go-ethereum/core/types"
	. "immutable.com/imx-core-sdk/workflows/types"
)

type TokenDeposit interface {
	Deposit() eth.Transaction
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
