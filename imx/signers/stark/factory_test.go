package stark

import (
	"math/big"
	"testing"

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
