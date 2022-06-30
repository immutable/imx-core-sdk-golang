package utils

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
)

// BaseL1Signer is an L1Signer implementation to assist for demo code.
type BaseL1Signer struct {
	privateKey *ecdsa.PrivateKey
	chainId    *big.Int
}

// NewBaseL1Signer creates a new BaseL1Signer object with the given user's privateKey and chainId.
func NewBaseL1Signer(privateKeyInHex string, chainId *big.Int) (signers.L1Signer, error) {
	privateKey, err := utils.TrimHexPrefix(privateKeyInHex)
	if err != nil {
		return nil, err
	}
	privateKeyInEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return &BaseL1Signer{privateKey: privateKeyInEcdsa, chainId: chainId}, nil
}

// GetAddress gets the address which is the public key associated with this signer.
func (b *BaseL1Signer) GetAddress() common.Address {
	address := crypto.PubkeyToAddress(b.privateKey.PublicKey)
	return address
}

// SignMessage signs the given message using the private key and returns the signature.
func (b *BaseL1Signer) SignMessage(message string) ([]byte, error) {
	bytes, err := common.ParseHexOrString(message)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(accounts.TextHash(bytes), b.privateKey)
}

// SignTx signs the given transaction using this signer and private key. Returns the signed transaction.
func (b *BaseL1Signer) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(b.chainId)
	return types.SignTx(tx, signer, b.privateKey)
}
