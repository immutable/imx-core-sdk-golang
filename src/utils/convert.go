package utils

import (
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	"github.com/shopspring/decimal"
	"math/big"
	"strconv"
	"strings"
)

func ParseBigFloat(value string) (*big.Float, error) {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	_, err := fmt.Sscan(value, f)
	return f, err
}

// ParseEtherToWei accepts a string and converts to wei.
// https://github.com/ethereum/go-ethereum/issues/21221
func ParseEtherToWei(value string) (*big.Int, error) {
	eth, err := ParseBigFloat(value)
	if err != nil {
		return nil, err
	}
	return EtherToWei(eth), nil
}

func EtherToWei(eth *big.Float) *big.Int {
	truncInt, _ := eth.Int(nil)
	truncInt = new(big.Int).Mul(truncInt, big.NewInt(params.Ether))
	fracStr := strings.Split(fmt.Sprintf("%.18f", eth), ".")[1]
	fracStr += strings.Repeat("0", 18-len(fracStr))
	fracInt, _ := new(big.Int).SetString(fracStr, 10)
	wei := new(big.Int).Add(truncInt, fracInt)
	return wei
}

func RemoveHex(hexString string) (string, error) {
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
	hexString, err := RemoveHex(hexString)
	if err != nil {
		return nil, err
	}
	addressInt.SetString(hexString, 16)
	return addressInt, nil
}

func HexToByteArray(hexString string) ([]byte, error) {
	hexString, err := RemoveHex(hexString)
	if err != nil {
		return nil, err
	}
	return hex.DecodeString(hexString)
}

// ToWei decimals to wei
func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}

func FromStringToDecimal(value string) (*uint8, error) {
	decimals, err := strconv.ParseUint(value, 10, 8)
	if err != nil {
		return nil, err
	}
	result := uint8(decimals)
	return &result, nil
}
