package deposits

import (
	"fmt"
	types2 "github.com/ethereum/go-ethereum/core/types"
)

func (d *ERC20Deposit) Deposit() types2.Transaction {
	fmt.Println(d.Amount)
	return types2.Transaction{}
}
