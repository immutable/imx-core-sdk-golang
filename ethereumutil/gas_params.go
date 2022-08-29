package ethereumutil

import "math/big"

const (
	DefaultGasLimit      uint64  = 300000       // Ethereum default gas limit in gas units
	DefaultGasPrice      uint64  = 160000000000 // Ethereum default gas price in wei
	DefaultGasMultiplier float64 = 1.5          // Ethereum default gas multiplier
	DefaultMaxGasPrice   uint64  = 300000000000 // Ethereum max gas price in wei
)

// GasParams are used to set the gas limit and gas price.
// When EgsAPIKey is provided, Client.BuildTransactOpts will attempt to set the gas price returned from eth gas station.
// If the api call to eth gas station fails, the gas price will be set at the (Network suggested price * GasMultiplier)
// unless MaxGasPrice is lower, then gas price will be set at MaxGasPrice.
// When EgsAPIKey is an empty string or if the api call to eth gas station fails, the gas price will be set at
// the (Network suggested price * GasMultiplier) unless MaxGasPrice is lower, then gas price will be set at MaxGasPrice.
type GasParams struct {
	GasLimit uint64
	GasPrice *big.Int
	// GasMultiplier is used to adjust the GasPrice returned by the method SuggestGasPrice (eth_gasPrice)
	// to control the speed of the ethereum transaction. E.g. a multiplier of 1.5 would set the gas price
	// at 1.5 times the suggested gas price and should process faster than simply using the suggested gas price.
	GasMultiplier float64
	MaxGasPrice   *big.Int
	// Eth gas station (EGS) API key https://ethgasstation.info
	EgsAPIKey string
	// Eth gas station (EGS) speed https://ethgasstation.info
	EGSSpeed string
}

var DefaultGasParams = NewDefaultGasParams()

func NewDefaultGasParams() GasParams {
	return GasParams{
		GasLimit:      DefaultGasLimit,
		GasPrice:      big.NewInt(int64(DefaultGasPrice)),
		GasMultiplier: DefaultGasMultiplier,
		MaxGasPrice:   big.NewInt(int64(DefaultMaxGasPrice)),
	}
}
