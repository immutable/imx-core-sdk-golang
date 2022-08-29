// Package stark implements utilities for generating signers.L2Signer
package stark

import (
	"fmt"
	"math/big"

	"github.com/dontpanicdao/caigo"
	"github.com/immutable/imx-core-sdk-golang/internal/convert"
)

// StarkSigner implements L2Signer interface.
type StarkSigner struct {
	privateKey *big.Int
	publicKey  *big.Int
	curve      *caigo.StarkCurve
}

func NewStarkSigner(privateKey, publicKey *big.Int, curve *caigo.StarkCurve) *StarkSigner {
	return &StarkSigner{
		privateKey: privateKey,
		publicKey:  publicKey,
		curve:      curve,
	}
}

func (base *StarkSigner) SignMessage(message string) (string, error) {
	hash, err := convert.HexToInt(message)
	if err != nil {
		return "", err
	}
	r, s, err := base.curve.Sign(hash, base.privateKey)
	if err != nil {
		return "", err
	}
	return serializeSignature(r, s), nil
}

// VerifySignature validates the given signature is signed by the requiredSigner or not.
func (base *StarkSigner) VerifySignature(hash *big.Int, signature, requiredSigner string) error {
	// All signatures must be 130 characters hex encoded 0x + 64 bytes with 2 characters each
	requiredSigLength := len("0x") + 64*2
	if len(signature) != requiredSigLength {
		return fmt.Errorf("invalid signature")
	}

	pubKey, _ := convert.HexToInt(requiredSigner)
	y := base.curve.GetYCoordinate(pubKey)

	r, s := rsFromSig(signature)

	// verification
	ok := base.curve.Verify(hash, r, s, pubKey, y)
	if ok {
		return nil
	}
	negY := big.NewInt(0).Mod(big.NewInt(0).Neg(y), base.curve.P)
	ok = base.curve.Verify(hash, r, s, pubKey, negY)
	if !ok {
		return fmt.Errorf("verification failed")
	}
	return nil
}

// GetAddress returns the stark public key of the StarkSigner as a 64 digit hex string prefixed with 0x.
func (base *StarkSigner) GetAddress() string {
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
