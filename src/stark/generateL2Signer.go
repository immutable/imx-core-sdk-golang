// Package stark implements utilities for generating signers.L2Signer
package stark

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/aarbt/hdkeys"
	"github.com/dontpanicdao/caigo"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"math/big"
)

const (
	SignableMessageString = "Only sign this request if you’ve initiated an action with Immutable X."
)

func hashKeyWithIndex(key *big.Int, index byte) *big.Int {
	bytes := sha256.Sum256(append(key.Bytes(), index))
	return new(big.Int).SetBytes(bytes[:])
}

func grind(key *big.Int) *big.Int {
	var i byte = 0

	key = hashKeyWithIndex(key, i)

	secpOrder, _ := new(big.Int).SetString("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141", 16)
	order, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)

	upperBound := new(big.Int).Sub(secpOrder, new(big.Int).Rem(secpOrder, order))

	for key.Cmp(upperBound) >= 0 {
		key = hashKeyWithIndex(key, i)
		i += 1
	}

	return new(big.Int).Rem(key, order)
}

func prefixMessageDigest(hashBytes []byte) []byte {
	ext := []byte("\x19Ethereum Signed Message:\n")
	len := []byte(fmt.Sprintf("%d", len(hashBytes)))
	return append(ext, append(len, hashBytes...)...)
}

func generateSeed(signer signers.L1Signer) ([]byte, error) {
	signature, err := signer.SignMessage(SignableMessageString)
	if err != nil {
		return nil, err
	}

	return signature[32:64], nil
}

func generatePath(layer, application, eth []byte, index string) string {
	layerHash := sha256.Sum256(layer)
	appHash := sha256.Sum256(application)

	a := binary.BigEndian.Uint32(layerHash[28:]) & (1<<31 - 1)
	b := binary.BigEndian.Uint32(appHash[28:]) & (1<<31 - 1)
	c := binary.BigEndian.Uint32(eth[16:]) & (1<<31 - 1)
	// Mask of 31 binary 1 digits, at the position 32 from the end (counting from 1)
	d := (binary.BigEndian.Uint64(eth[12:]) & ((1<<31 - 1) << 31)) >> 31

	return fmt.Sprintf("m/2645'/%d'/%d'/%d'/%d'/%s", a, b, c, d, index)
}

// GenerateSigner creates a signers.L2Signer from provided signer.
// pedersenParamsPath should be a local path to 'pedersen_params.json' file.
// If pedersenParamsPath is empty string the file will be pulled from Starkware
// official github repository. For production deployments it is recommended
// to have the file stored locally.
func GenerateSigner(signer signers.L1Signer, pedersenParamsPath string) (*StarkSigner, error) {
	seed, err := generateSeed(signer)
	if err != nil {
		return nil, fmt.Errorf("error getting seed from l1signer: %v", err)
	}

	masterKey := hdkeys.NewMasterKey(seed)

	path := generatePath([]byte("starkex"), []byte("immutablex"), signer.GetAddress().Bytes(), "1")

	childKey, err := masterKey.Chain(path)
	if err != nil {
		return nil, err
	}

	// Last 32 bits
	childPrivateKey := childKey.Serialize()[46:]

	privateStarkKey := grind(new(big.Int).SetBytes(childPrivateKey))

	starkCurve, err := caigo.SCWithConstants(pedersenParamsPath)
	if err != nil {
		return nil, fmt.Errorf("error generating stark curve: %v", err)
	}

	starkX, _, err := starkCurve.PrivateToPoint(privateStarkKey)
	if err != nil {
		return nil, fmt.Errorf("error generating stark public key: %v", err)
	}

	return NewStarkSigner(privateStarkKey, starkX, &starkCurve), nil
}

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
	hash, err := utils.HexToInt(message)
	if err != nil {
		return "", err
	}
	r, s, err := base.curve.Sign(hash, base.privateKey)
	if err != nil {
		return "", err
	}
	return serializeSignature(r, s), nil
}

func (base *StarkSigner) VerifySignature(hash *big.Int, signature string, requiredSigner string) error {
	// All signatures must be 130 characters hex encoded 0x + 64 bytes with 2 characters each
	requiredSigLength := len("0x") + 64*2
	if len(signature) != requiredSigLength {
		return fmt.Errorf("invalid signature")
	}

	pubKey, _ := utils.HexToInt(requiredSigner)
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

func rsFromSig(signature string) (*big.Int, *big.Int) {
	sig := signature[2:]
	size := 64
	r := new(big.Int)
	r.SetString(sig[0:size], 16)
	s := new(big.Int)
	s.SetString(sig[size:size*2], 16)
	return r, s
}

// GetAddress returns the stark public key of the StarkSigner as a 64 digit hex string prefixed with 0x.
func (s *StarkSigner) GetAddress() string {
	return fmt.Sprintf("0x%064x", s.publicKey)
}

func serializeSignature(r, s *big.Int) string {
	return fmt.Sprintf("0x%064s%064s", r.Text(16), s.Text(16))
}
