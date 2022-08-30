package stark

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"path"
	"path/filepath"
	"runtime"

	"github.com/aarbt/hdkeys"
	"github.com/dontpanicdao/caigo"
	"github.com/ethereum/go-ethereum/common/hexutil"
	immutable "github.com/immutable/imx-core-sdk-golang"
)

const (
	seedMessage = "Only sign this request if you’ve initiated an action with Immutable X."
)

// GenerateStarkSigner factory method to create StarkSigner.
func GenerateSigner(signer immutable.L1Signer) (*Signer, error) {
	seed, err := generateSeed(signer, seedMessage)
	if err != nil {
		return nil, fmt.Errorf("failed to generate seed using l1signer: %v", err)
	}

	starkPath := generateStarkPath([]byte("starkex"), []byte("immutablex"), hexutil.MustDecode(signer.GetAddress()), "1")
	childKey, err := hdkeys.NewMasterKey(seed).Chain(starkPath)
	if err != nil {
		return nil, err
	}

	// Last 32 bits
	childPrivateKey := childKey.Serialize()[46:]
	privateStarkKey := grind(new(big.Int).SetBytes(childPrivateKey))

	dir, err := currentDirname()
	if err != nil {
		return nil, err
	}
	pedersenParamsFilePath := path.Join(dir, "pedersen_params.json")
	starkCurve, err := caigo.SC(caigo.WithConstants(pedersenParamsFilePath))
	if err != nil {
		return nil, fmt.Errorf("error generating stark curve: %v", err)
	}

	starkX, _, err := starkCurve.PrivateToPoint(privateStarkKey)
	if err != nil {
		return nil, fmt.Errorf("error generating stark public key: %v", err)
	}

	return NewSigner(privateStarkKey, starkX, &starkCurve), nil
}

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

func generateSeed(signer immutable.L1Signer, seedMessage string) ([]byte, error) {
	signature, err := signer.SignMessage(seedMessage)
	if err != nil {
		return nil, err
	}
	return signature[32:64], nil
}

func generateStarkPath(layer, application, eth []byte, index string) string {
	layerHash := sha256.Sum256(layer)
	appHash := sha256.Sum256(application)

	a := binary.BigEndian.Uint32(layerHash[28:]) & (1<<31 - 1)
	b := binary.BigEndian.Uint32(appHash[28:]) & (1<<31 - 1)
	c := binary.BigEndian.Uint32(eth[16:]) & (1<<31 - 1)
	// Mask of 31 binary 1 digits, at the position 32 from the end (counting from 1)
	d := (binary.BigEndian.Uint64(eth[12:]) & ((1<<31 - 1) << 31)) >> 31

	return fmt.Sprintf("m/2645'/%d'/%d'/%d'/%d'/%s", a, b, c, d, index)
}

// currentDirname gets the full directory path of the caller of this function.
func currentDirname() (string, error) {
	_, filename, _, ok := runtime.Caller(1) // Caller 1 will get this function callers path.
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filepath.Dir(filename), nil
}
