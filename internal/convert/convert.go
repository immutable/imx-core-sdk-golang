package convert

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/shopspring/decimal"
)

// EtherDecimals for Ether is expressed in 18 decimals
const EtherDecimals = 18

// TrimHexPrefix removes the prefix `0x` from the given hex string.
func TrimHexPrefix(hexString string) (string, error) {
	if len(hexString) < 2 {
		return "", fmt.Errorf("invalid hex string %s", hexString)
	}
	if hexString[:2] == "0x" {
		return hexString[2:], nil
	}
	return hexString, nil
}

func HexToByteArray(hexString string) ([]byte, error) {
	hexString, err := TrimHexPrefix(hexString)
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(hexString)
}

// ToDenomination converts amount to denomination based on decimals.
// e.g. set decimals to 18 to convert amount in eth to wei.
func ToDenomination(amount string, decimals int) (*big.Int, error) {
	amountInDecimals, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amountInDecimals.Mul(mul)

	denomination := new(big.Int)
	denomination, ok := denomination.SetString(result.String(), 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert amount to big.Int")
	}

	return denomination, nil
}
