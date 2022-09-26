package convert

import (
	"encoding/hex"
	"fmt"
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
