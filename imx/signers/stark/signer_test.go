package stark

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestStarkSigner_VerifySignature(t *testing.T) {
	key, err := GenerateKey()
	assert.NoError(t, err)

	l2Signer, err := NewSigner(key)
	assert.NoError(t, err)

	message := "0x074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e"
	hash, ok := new(big.Int).SetString(message, 0)
	assert.True(t, ok)

	signature, err := l2Signer.SignMessage(message)
	assert.NoError(t, err)

	starkPubKey := l2Signer.GetPublicKey()
	err = l2Signer.VerifySignature(hash, signature, starkPubKey)
	assert.NoError(t, err)
}

func TestStarkSigner_ShouldReturnCorrectAddress(t *testing.T) {
	mockPrivateKey := "019b5de47e5fc8f2e8c3415b42a126aadb462637f2feca1df3733fe3f37cf50f"
	expectedPublicKey := "0x0790436373c1d5b7a88ce4fd7ac96591a385b2b6392d1ea44a165f75115b82ac"
	l2Signer, err := NewSigner(mockPrivateKey)
	assert.NoError(t, err)

	assert.Equalf(t, expectedPublicKey, l2Signer.GetPublicKey(), "Check the generated public key is as expected")
}

func TestStarkSigner_GetAddress(t *testing.T) {
	type fields struct {
		publicKey *big.Int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "returns correct length given 3 digits",
			fields: fields{publicKey: hexutil.MustDecodeBig("0x123")},
			want:   "0x0000000000000000000000000000000000000000000000000000000000000123",
		},
		{
			name:   "returns correct length given 63 digits",
			fields: fields{publicKey: hexutil.MustDecodeBig("0x163dac0ba9e48ca45aee143a7cd495e0e858b52640f14f28f047033829757d8")},
			want:   "0x0163dac0ba9e48ca45aee143a7cd495e0e858b52640f14f28f047033829757d8",
		},
		{
			name:   "returns correct length given 64 digits",
			fields: fields{publicKey: hexutil.MustDecodeBig("0x163dac0ba9e48ca45aee143a7cd495e0e858b52640f14f28f047033829757d80")},
			want:   "0x163dac0ba9e48ca45aee143a7cd495e0e858b52640f14f28f047033829757d80",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Signer{
				publicKey: tt.fields.publicKey,
			}
			assert.Equalf(t, tt.want, s.GetPublicKey(), "GetPublicKey()")
		})
	}
}
