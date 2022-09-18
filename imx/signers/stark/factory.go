package stark

import (
	"errors"
	"math/big"
	"path"
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
	return curve.GetRandomPrivateKey()
}

func loadCurve() (*caigo.StarkCurve, error) {
	dir, err := currentDirname()
	if err != nil {
		return nil, err
	}
	pedersenParamsFilePath := path.Join(dir, "pedersen_params.json")
	sc, err := caigo.SC(caigo.WithConstants(pedersenParamsFilePath))
	if err != nil {
		return nil, err
	}
	return &sc, nil
}

// currentDirname gets the full directory path of the caller of this function.
func currentDirname() (string, error) {
	_, filename, _, ok := runtime.Caller(1) // Caller 1 will get this function callers path.
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filepath.Dir(filename), nil
}
