package ethereum

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/internal/convert"
)

// Signer implements L1Signer interface.
type Signer struct {
	privateKey *ecdsa.PrivateKey
	chainID    *big.Int
}

// NewSigner creates a new Signer object with the given user's privateKey and chainID.
func NewSigner(privateKeyInHex string, chainID *big.Int) (imx.L1Signer, error) {
	privateKey, err := convert.TrimHexPrefix(privateKeyInHex)
	if err != nil {
		return nil, err
	}
	privateKeyInEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	return &Signer{privateKey: privateKeyInEcdsa, chainID: chainID}, nil
}

// GetAddress gets the address which is the checksum of the public key associated with this signer.
func (b *Signer) GetAddress() string {
	address := crypto.PubkeyToAddress(b.privateKey.PublicKey)
	return address.Hex()
}

// GetPublicKey gets the public key associated with this signer.
func (b *Signer) GetPublicKey() string {
	pubKey := crypto.FromECDSAPub(&b.privateKey.PublicKey)
	return hexutil.Encode(pubKey)
}

// SignMessage signs the given message using the private key and returns the signature.
func (b *Signer) SignMessage(message string) ([]byte, error) {
	hashBytes := []byte(message)
	return crypto.Sign(accounts.TextHash(hashBytes), b.privateKey)
}

// SignTx signs the given transaction using this signer and private key. Returns the signed transaction.
func (b *Signer) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(b.chainID)
	return types.SignTx(tx, signer, b.privateKey)
}
