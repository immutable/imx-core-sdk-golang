package stark

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math/big"
	"path/filepath"
	"runtime"

	"github.com/dontpanicdao/caigo"
)

// package-global variable to represent our standard stark curve
var curve *caigo.StarkCurve

// GenerateKey generates a random key that can be used to create StarkSigner.
// On creation save this key for future usage as this key will be required to reuse your stark signer.
// @return Randomly generated private key.
func GenerateKey() (*big.Int, error) {
	if curve == nil {
		var err error
		curve, err = loadCurve()
		if err != nil {
			return nil, err
		}
	}
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}

	return grind(privKey.D), nil
}

// Initialises the Stark Elliptic Curve.
func loadCurve() (*caigo.StarkCurve, error) {
	dir, err := currentDirname()
	if err != nil {
		return nil, err
	}
	pedersenParamsFilePath := filepath.Join(dir, "pedersen_params.json")
	sc, err := caigo.SC(caigo.WithConstants(pedersenParamsFilePath))
	if err != nil {
		return nil, err
	}
	return &sc, nil
}

// Create a hash from a key + an index
func hashKeyWithIndex(key *big.Int, index byte) *big.Int {
	bytes := sha256.Sum256(append(key.Bytes(), index))
	return new(big.Int).SetBytes(bytes[:])
}

/*
grind function receives a key seed and produces an appropriate StarkEx key from a uniform distribution.

Although it is possible to define a StarkEx key as a residue (mod) between the StarkEx EC order and a
random 256bit digest value, the result would be a biased key. In order to prevent this bias, we
deterministically search (by applying more hashes, AKA grinding) for a value lower than the largest
256bit multiple of StarkEx EC order.

https://github.com/starkware-libs/starkware-crypto-utils/blob/dev/src/js/key_derivation.js#L119
*/
func grind(key *big.Int) *big.Int {
	sha256EcMaxDigest, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	starkEcOrder, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)

	upperBound := new(big.Int).Sub(sha256EcMaxDigest, new(big.Int).Rem(sha256EcMaxDigest, starkEcOrder))

	var i byte = 0
	key = hashKeyWithIndex(key, i)
	for key.Cmp(upperBound) >= 0 {
		key = hashKeyWithIndex(key, i)
		i += 1
	}
	return new(big.Int).Rem(key, starkEcOrder)
}

// currentDirname gets the full directory path of the caller of this function.
func currentDirname() (string, error) {
	_, filename, _, ok := runtime.Caller(1) // Caller 1 will get this function callers path.
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filepath.Dir(filename), nil
}
