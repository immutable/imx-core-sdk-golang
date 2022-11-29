package stark

import (
	"math/big"
	"testing"

	"github.com/immutable/imx-core-sdk-golang/imx/signers/ethereum"

	"github.com/stretchr/testify/assert"
)

func TestStarkSignerFactory_GenerateRandomKey(t *testing.T) {
	key1, err := GenerateKey()
	assert.NoError(t, err)
	key2, err := GenerateKey()
	assert.NoError(t, err)

	assert.NotEqualf(t, key1, key2, "Generated random keys are not same")
}

func TestStarkSignerFactory_Grinding(t *testing.T) {
	privateKey, ok := new(big.Int).SetString("86F3E7293141F20A8BAFF320E8EE4ACCB9D4A4BF2B4D295E8CEE784DB46E0519", 16)
	assert.True(t, ok)
	expectedKeyValueAfterGrinding, ok := new(big.Int).SetString("5c8c8683596c732541a59e03007b2d30dbbbb873556fe65b5fb63c16688f941", 16)
	assert.True(t, ok)

	assert.Equalf(t, expectedKeyValueAfterGrinding, grind(privateKey), "Verify grinding logic")
}

func TestStarkSignerFactory_GenerateLegacyKey(t *testing.T) {
	l1Signer, err := ethereum.NewSigner("5c7b4b5cad9a3fc7b1ba235a49cd74e615488a18b0d6a531739fd1062935104d", big.NewInt(5))
	assert.NoError(t, err)
	key1, err := GenerateLegacyKey(l1Signer)
	assert.NoError(t, err)
	assert.Equalf(t, "0x556413893a023efd75f62cd4eca825f2be7e918b5188f1db06cbec12d7d1b88", key1, "Check the generated key matches")

	key2, err := GenerateLegacyKey(l1Signer)
	assert.NoError(t, err)

	assert.Equalf(t, key1, key2, "Generated keys are same")
}
