package imx

import "github.com/ethereum/go-ethereum/core/types"

type L1Signer interface {
	SignMessage(message string) ([]byte, error)
	SignTx(tx *types.Transaction) (*types.Transaction, error)
	GetAddress() string
}

type L2Signer interface {
	SignMessage(message string) (string, error)
	GetAddress() string
}