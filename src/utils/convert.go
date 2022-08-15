package utils

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

func HexToInt(hexString string) (*big.Int, error) {
	addressInt := new(big.Int)
	hexString, err := TrimHexPrefix(hexString)
	if err != nil {
		return nil, err
	}
	addressInt.SetString(hexString, 16)
	return addressInt, nil
}

func HexToByteArray(hexString string) ([]byte, error) {
	hexString, err := TrimHexPrefix(hexString)
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(hexString)
}

// ToUnquantized converts amount to denomination based on decimals.
// e.g. set decimals to 18 to convert amount in eth to wei.
func ToUnquantized(amount string, decimals int) (*big.Int, error) {
	amountInDecimals, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amountInDecimals.Mul(mul)

	unquantizedValue := new(big.Int)
	unquantizedValue, ok := unquantizedValue.SetString(result.String(), 10)
	if !ok {
		return nil, fmt.Errorf("failed to convert amount to big.Int")
	}

	return unquantizedValue, nil
}
