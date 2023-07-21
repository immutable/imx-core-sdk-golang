package stark

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"embed"
	"encoding/binary"
	"fmt"
	"math/big"

	"github.com/immutable/imx-core-sdk-golang/imx"

	"github.com/aarbt/hdkeys"
	"github.com/dontpanicdao/caigo"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

var (
	// package-global variable to represent our standard stark curve
	curve *caigo.StarkCurve

	//go:embed pedersen_params.json
	pedersenParamsBytes []byte

	//go:embed *.json
	_ embed.FS // Unused but required to import the embed module
)

// GenerateKey generates a random key that can be used to create StarkSigner.
// On creation save this key for future usage as this key will be required to reuse your stark signer.
// @return Randomly generated private key.
func GenerateKey() (string, error) {
	if curve == nil {
		var err error
		curve, err = loadCurve()
		if err != nil {
			return "", err
		}
	}
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return "", err
	}
	return hexutil.EncodeBig(grind(privKey.D)), nil
}

const (
	DefaultSeedMessage = "Only sign this request if you’ve initiated an action with Immutable X."
	ApplicationName    = "immutablex"
	LayerName          = "starkex"
	Index              = "1"
)

// GenerateLegacyKey Generates a deterministic Stark private key from the provided signer.
// @return Deterministically generated private key.
func GenerateLegacyKey(signer imx.L1Signer) (string, error) {
	seed, err := generateSeed(signer, DefaultSeedMessage)
	if err != nil {
		return "", fmt.Errorf("failed to generate seed using l1signer: %v", err)
	}

	starkPath := getStarkPath(LayerName, ApplicationName, signer.GetAddress(), Index)
	childKey, err := hdkeys.NewMasterKey(seed).Chain(starkPath)
	if err != nil {
		return "", err
	}

	// Last 32 bits
	childPrivateKey := childKey.Serialize()[46:]
	return hexutil.EncodeBig(grind(new(big.Int).SetBytes(childPrivateKey))), nil
}

// Initialises the Stark Elliptic Curve.
func loadCurve() (*caigo.StarkCurve, error) {
	caigo.PedersenParamsRaw = pedersenParamsBytes
	return &caigo.Curve, nil
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

	//index is 0, 0, 1, 2...
	var i byte = 0
	key = hashKeyWithIndex(key, i)
	for key.Cmp(upperBound) >= 0 {
		key = hashKeyWithIndex(key, i)
		i += 1
	}
	return new(big.Int).Rem(key, starkEcOrder)
}

func grindCorrect(key *big.Int) *big.Int {
	sha256EcMaxDigest, _ := new(big.Int).SetString("10000000000000000000000000000000000000000000000000000000000000000", 16)
	starkEcOrder, _ := new(big.Int).SetString("0800000000000010ffffffffffffffffb781126dcae7b2321e66a241adc64d2f", 16)

	upperBound := new(big.Int).Sub(sha256EcMaxDigest, new(big.Int).Rem(sha256EcMaxDigest, starkEcOrder))

	//index is 0, 1, 2, ...
	var i byte = 0
	key = hashKeyWithIndex(key, i)
	for key.Cmp(upperBound) >= 0 {
		i += 1
		key = hashKeyWithIndex(key, i)
	}
	return new(big.Int).Rem(key, starkEcOrder)
}

// generateSeed generates the seed value for the given seed message.
func generateSeed(signer imx.L1Signer, seedMessage string) ([]byte, error) {
	signature, err := signer.SignMessage(seedMessage)
	if err != nil {
		return nil, err
	}
	return signature[32:64], nil
}

func getStarkPath(layerName, applicationName, ethereumAddress, index string) string {
	// Starkware keys are derived with the following BIP43-compatible derivation path, with direct inspiration from BIP44:
	//
	// m / purpose' / layer' / application' / eth_address_1' / eth_address_2' / index
	// where:
	//
	// m 			- the seed.
	// purpose 		- 2645 (the number of this EIP).
	// layer 		- the 31 lowest bits of sha256 on the layer name. Serve as a domain separator between different technologies. In the context of starkex, the value would be 579218131.
	// application 	- the 31 lowest bits of sha256 of the application name. Serve as a domain separator between different applications.
	// 					In the context of DeversiFi in June 2020, it is the 31 lowest bits of sha256(starkexdvf) and the value would be 1393043894.
	// eth_address_1 / eth_address_2 - the first and second 31 lowest bits of the corresponding eth_address.
	// index 		- to allow multiple keys per eth_address.
	//
	// See for more info regarding path derivation https://docs.starkware.co/starkex/key-derivation.html and https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2645.md

	layerHash := sha256.Sum256([]byte(layerName))
	appHash := sha256.Sum256([]byte(applicationName))
	layer := binary.BigEndian.Uint32(layerHash[28:]) & (1<<31 - 1)
	application := binary.BigEndian.Uint32(appHash[28:]) & (1<<31 - 1)

	ethereumAddressInBytes := hexutil.MustDecode(ethereumAddress)
	ethAddress1 := binary.BigEndian.Uint32(ethereumAddressInBytes[16:]) & (1<<31 - 1)
	// Mask of 31 binary 1 digits, at the position 32 from the end (counting from 1)
	ethAddress2 := (binary.BigEndian.Uint64(ethereumAddressInBytes[12:]) & ((1<<31 - 1) << 31)) >> 31

	return fmt.Sprintf("m/2645'/%d'/%d'/%d'/%d'/%s", layer, application, ethAddress1, ethAddress2, index)
}
