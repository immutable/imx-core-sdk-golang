package stark

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
)

func TestStarkSigner_VerifySignature(t *testing.T) {
	key, _ := GenerateKey()
	l2Signer, _ := NewSigner(key)
	message := "0x074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e"
	hash, _ := hexutil.DecodeBig(message)
	signature, _ := l2Signer.SignMessage(message)
	starkPubKey := l2Signer.GetAddress()
	err := l2Signer.VerifySignature(hash, signature, starkPubKey)
	assert.Nil(t, err)
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
			assert.Equalf(t, tt.want, s.GetAddress(), "GetAddress()")
		})
	}
}
