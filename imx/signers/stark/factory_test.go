package stark

import (
	"context"
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
	type test struct {
		name                 string
		privateKey           string
		wantKeyAfterGrinding string
	}
	tests := []test{
		{
			name:                 "correct grind key",
			privateKey:           "86F3E7293141F20A8BAFF320E8EE4ACCB9D4A4BF2B4D295E8CEE784DB46E0519",
			wantKeyAfterGrinding: "5c8c8683596c732541a59e03007b2d30dbbbb873556fe65b5fb63c16688f941",
		},
		{
			name:                 "private key is above the key limit",
			privateKey:           "a978531943ad2e2a8af34e0e2a7d306dc99516d489be16e4ea2ee74c90a9d88f",
			wantKeyAfterGrinding: "1e8108d99e74b769d6b998a5a41ff2745f0607496f2eed39abfd161837408e7",
		},
		{
			name:                 "private key starts with zero",
			privateKey:           "086F3E7293141F20A8BAFF320E8EE4ACCB9D4A4BF2B4D295E8CEE784DB46E051",
			wantKeyAfterGrinding: "2b2c6db790a95ce05426c3d67247547f1a72d104fd5af24553d42b7557ab082",
		},
	}
	for _, tt := range tests {
		privateKey, ok := new(big.Int).SetString(tt.privateKey, 16)
		assert.True(t, ok)
		wantKeyAfterGrinding, ok := new(big.Int).SetString(tt.wantKeyAfterGrinding, 16)
		assert.True(t, ok)
		assert.Equalf(t, wantKeyAfterGrinding, grindKeyV100Beta1(privateKey), "Verify grinding logic")
	}
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

func TestStarkSignerFactory_GenerateLegacyKey_GrindkeyFix(t *testing.T) {
	type test struct {
		name          string
		privateKey    string
		wantPublicKey string
	}
	tests := []test{
		{
			name:          "should generate Legacy Stark public key",
			privateKey:    "5c7b4b5cad9a3fc7b1ba235a49cd74e615488a18b0d6a531739fd1062935104d",
			wantPublicKey: "0x0579f97e8084dfbbead9bffd750df780e06d8c09a3ba7f40ebe51d46b47df043",
		},
		{
			name:       "Random Eth Private Key hashed with no index, should generate correct legacy (imx-sdk-js) stark key",
			privateKey: "0x4aee6bb7807a684341e39303fb1292f1c66b0f0bcac81407d10e80ba9202fc4a",
			// Compatible with grindKeyLegacy, With Expected Index of 0,0,0,... // Hash Iterated 0 times
			wantPublicKey: "0x3c35bb183cc613d57e6c823a6c7a17a14fba9dc65a5ece97a87020e32cbbd9",
		},
		{
			name:       "Eth Private Key hashed with index once, should generate correct legacy (imx-sdk-js) stark key",
			privateKey: "0xabb730db6bd22a773263c9266b0e157e6a9985c75ff68c68fdebdef3e500892c",
			// Compatible with grindKeyLegacy, With Expected Index of 0,0,0,... // Hash Iterated 1 time
			wantPublicKey: "0x0743a17b8a566fe39cfebc41db631210bf328b91ae6c3968f1f9a96ac48a02aa",
		},
		{
			name:       "Eth Private Key hashed with indexed more than once, should generate correct legacy (imx-sdk-js) stark key",
			privateKey: "0x2b9b32d0ed8863016001ffb15af781dca725b1581a3d4d61c0e2c4c0c03ede8a",
			// Compatible with grindKeyLegacy, With Expected Index of 0,0,0,... // Hash Iterated 2 times
			wantPublicKey: "0x07d14076863573a2310a948472512331416ceb655a304b1811a127e6c710e350",
		},
		{
			name:       "Eth Private Key hashed with index once, should generate Legacy Stark key backwards compatible with versions between 1.0.0-beta.3 to 2.0.0",
			privateKey: "0x1a245f2fa7c4f04a65d45a3877ad00b1423d081490dcc1a7050c8d7c11ec5c8f",
			// Compatible with grindKey, with Expected Index of 0,1,2,... // Hash Iterated 1 time
			wantPublicKey: "0x07ca61905954bdd858ae63704b111da5ca52b30a21fa9aaeee9aef9b24e89607",
		},
		{
			name:       "Eth Private Key hashed with index more than once, should generate Legacy Stark key compatible with versions between 1.0.0-beta.3 to 2.0.0, 2.0.2 and above",
			privateKey: "0xe516a5b715dc53ba4bf06ed98cbe50897921f45f76d4138f8a7d4ba02a89e10d",
			// Compatible with grindKey, with Expected Index of 0,1,2,... // Hash Iterated 2 times
			wantPublicKey: "0x0157c6efcb2c69604c1b46ffeb568b3b10ea7093da762e1daf66b59d9b63b5f9",
		},
		{
			name:       "should generate Legacy Stark key compatible with versions between 1.0.0-beta.3 to 2.0.0, 2.0.2 and above",
			privateKey: "ba3c969f4957e6bf24e5cf8a931bdba4f90d27c01bb7dff738e4593142826db7",
			// Compatible with GrindKey, with Expected Index of 0,1,2,... // Hash Iterated 0 times
			wantPublicKey: "0x04eb684f7318d2b90ef8703e6cd77e362414f4997934ee3c99ee50db3411dfef",
		},
		{
			name:       "Eth Private Key hashed with index once, should generate Legacy Stark key backwards compatible with version 1.0.0-beta.1",
			privateKey: "0xa25022c521d884d195c119b705ac515fabb48971d3cf2369f188c96b0e6404ef",
			// Compatible with grindKeyV100Beta1, with Expected Index of 0,0,1,... // Hash Iterated 1 time
			wantPublicKey: "0x03aa9edf4a5b33f623d3dd37a6df62e5590f172fd4d1a9483532464aa13bcbb9",
		},
		{
			name:       "Eth Private Key hashed with index more than once, should generate Legacy Stark key backwards compatible with version 1.0.0-beta.1",
			privateKey: "0x719b531dfdbf5e327646ccd992923c9d8846d60767b49d0589d631c1f54d1f12",
			// Compatible with core-sdk GrindKey, with Expected Index of 0,0,1,... // Hash Iterated 2 times
			wantPublicKey: "0x057a8278637277ebaaa63b000185507eb847eec8f69c7758fc10a490d0226e7d",
		},
		{
			name:       "Eth Private Key hashed with index once, should generate correct stark key",
			privateKey: "0xe7ecb8f91175446248a6cfa45a9526c85bc4cd7cbd9427e3e65a82d3f5fb8cdc",
			// Hash Iterated 1 time
			wantPublicKey: "0x05f8f00b03e896cb8c02a936bab73623fc3651fddabfe0d28fbc309c63642c10",
		},
		{
			name:       "Eth Private Key hashed with index more than once, should generate correct stark key",
			privateKey: "0x82930dfa052b198828f70529148b02178b03515e910217dab32bf9fc046bc9d9",
			// Hash Iterated 2 times
			wantPublicKey: "0x04d34c9f519b2a3538b02c4e939188b591065eb8c2f210277c6ed1de85b16fab",
		},
		{
			name:       "Random Eth Private Key hashed with no index, should generate correct legacy stark key",
			privateKey: "0x0ac63e2143a49a14a87d865c8d7993806ff16ac1c3288ff97101569881f0d306",
			// Hash Iterated 0 times
			wantPublicKey: "0x06e6ac4bb44f3295b881532452f90eccb5601314fafe306db17684b47aa388bd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1Signer, err := ethereum.NewSigner(tt.privateKey, big.NewInt(5))
			assert.NoError(t, err)
			starkKeyInHex, err := GenerateLegacyKey(l1Signer)
			if err != nil {
				t.Fatal(err)
			}
			starkSigner, err := NewSigner(starkKeyInHex)
			assert.NoError(t, err)
			expectedStarkPubKey, success := new(big.Int).SetString(tt.wantPublicKey, 0)
			assert.True(t, success)
			assert.Equal(t, expectedStarkPubKey, starkSigner.PublicKey())
		})
	}
}

func Test_getStarkPublicKeyFromImx(t *testing.T) {
	type args struct {
		ctx        context.Context
		privateKey string
	}
	tests := []struct {
		name                string
		args                args
		wantStarkPublicKey  string
		wantAccountNotFound bool
		wantErr             bool
	}{
		{"Account Not Found",
			args{
				ctx:        context.Background(),
				privateKey: "0x0ac63e2143a49a14a87d865c8d7993806ff16ac1c3288ff97101569881f0d306",
			},
			"", true, false,
		},
		{"Valid account on Prod",
			args{
				ctx:        context.Background(),
				privateKey: "0x719b531dfdbf5e327646ccd992923c9d8846d60767b49d0589d631c1f54d1f12",
			},
			"0x057a8278637277ebaaa63b000185507eb847eec8f69c7758fc10a490d0226e7d", false, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1Signer, err := ethereum.NewSigner(tt.args.privateKey, big.NewInt(5))
			assert.NoError(t, err)
			ethAddress := l1Signer.GetAddress()
			gotStarkPublicKey, gotAccountNotFound, err := getStarkPublicKeyFromImx(tt.args.ctx, ethAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("getStarkPublicKeyFromImx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.wantAccountNotFound, gotAccountNotFound)
			assert.Equal(t, tt.wantStarkPublicKey, gotStarkPublicKey)
		})
	}
}
