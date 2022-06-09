package deposits

import (
	"github.com/ethereum/go-ethereum/core/types"
	types2 "immutable.com/imx-core-sdk/workflows/types"
)

type TokenDeposit interface {
	Deposit() types.Transaction
}

type ETHDeposit struct {
	Type   types2.TokenTypeEnum
	Amount string
}

type ERC20Deposit struct {
	Type         types2.TokenTypeEnum
	Amount       string
	TokenAddress string
	Symbol       string
}

type ERC721Deposit struct {
	Type         types2.TokenTypeEnum
	TokenId      string
	TokenAddress string
}
