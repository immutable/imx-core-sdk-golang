// Package stark implements utilities for generating immutable.L2Signer
package stark

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// StarkSigner implements L2Signer interface.
type Signer struct {
	privateKey *big.Int
	publicKey  *big.Int
}

func NewSigner(privateKey *big.Int) (*Signer, error) {
	x, _, err := curve.PrivateToPoint(privateKey)
	if err != nil {
		return nil, err
	}

	return &Signer{
		privateKey: privateKey,
		publicKey:  x, // TODO: is this right?
	}, nil
}

// has0xPrefix checks if input contains 0x prefix. This method is copied from github.com/ethereum/go-ethereum/common/hexutil
func has0xPrefix(input string) bool {
	return len(input) >= 2 && input[0] == '0' && (input[1] == 'x' || input[1] == 'X')
}

func (base *Signer) SignMessage(message string) (string, error) {
	if !has0xPrefix(message) {
		message = "0x" + message
	}
	hash, err := hexutil.DecodeBig(message)
	if err != nil {
		return "", err
	}
	r, s, err := curve.Sign(hash, base.privateKey)
	if err != nil {
		return "", err
	}
	return serializeSignature(r, s), nil
}

// VerifySignature validates the given signature is signed by the requiredSigner or not.
func (base *Signer) VerifySignature(hash *big.Int, signature, requiredSigner string) error {
	// All signatures must be 130 characters hex encoded 0x + 64 bytes with 2 characters each
	requiredSigLength := len("0x") + 64*2
	if len(signature) != requiredSigLength {
		return fmt.Errorf("invalid signature")
	}

	pubKey, _ := hexutil.DecodeBig(requiredSigner)
	y := curve.GetYCoordinate(pubKey)

	r, s := rsFromSig(signature)

	// verification
	ok := curve.Verify(hash, r, s, pubKey, y)
	if ok {
		return nil
	}
	negY := big.NewInt(0).Mod(big.NewInt(0).Neg(y), curve.P)
	ok = curve.Verify(hash, r, s, pubKey, negY)
	if !ok {
		return fmt.Errorf("verification failed")
	}
	return nil
}

// GetAddress returns the stark public key of the StarkSigner as a 64 digit hex string prefixed with 0x.
func (base *Signer) GetAddress() string {
	return fmt.Sprintf("0x%064x", base.publicKey)
}

func rsFromSig(signature string) (r, s *big.Int) {
	sig := signature[2:]
	size := 64
	r = new(big.Int)
	r.SetString(sig[0:size], 16)
	s = new(big.Int)
	s.SetString(sig[size:size*2], 16)
	return r, s
}

func serializeSignature(r, s *big.Int) string {
	return fmt.Sprintf("0x%064s%064s", r.Text(16), s.Text(16))
}
