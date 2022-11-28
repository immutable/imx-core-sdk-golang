package convert

import (
	"encoding/hex"
	"fmt"
	"math/big"
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

func HexToBigInt(hexString string) (*big.Int, error) {
	hexString, err := TrimHexPrefix(hexString)
	if err != nil {
		return nil, err
	}
	bigIntValue, ok := new(big.Int).SetString(hexString, 16)
	if !ok {
		return nil, fmt.Errorf("error in converting stark private key value from string to big.Int")
	}
	return bigIntValue, nil
}
