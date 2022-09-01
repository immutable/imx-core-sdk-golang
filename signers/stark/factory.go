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
// generally bad practice, but has no state + we never configure it
// unexported so can be updated in future if necessary
// TODO: should we do this in init() to allow us to handle the eaten error here
var curve, _ = loadCurve()

// GenerateStarkSigner factory method to create StarkSigner.
func GenerateKey() (*big.Int, error) {
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
