package imx

import "github.com/ethereum/go-ethereum/core/types"

// L1Signer interface to implement signing functionality using ethereum wallet key.
type L1Signer interface {
	SignMessage(message string) ([]byte, error)
	SignTx(tx *types.Transaction) (*types.Transaction, error)
	GetAddress() string
	GetPublicKey() string
}

// L1Signer interface to implement signing functionality using Stark wallet key.
type L2Signer interface {
	SignMessage(message string) (string, error)
	GetPublicKey() string
}
