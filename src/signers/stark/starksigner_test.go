package stark

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
)

func TestStarkSigner_VerifySignature(t *testing.T) {
	privateKey := ""
	if privateKey == "" {
		fmt.Println("WARN: You need to provide a private key to run this test.")
		return
	}
	l1Signer, _ := GetL1Signer(privateKey, nil)

	l2Signer, _ := GenerateStarkSigner(l1Signer)
	message := "0x074180eaec7e68712b5a0fbf5d63a70c33940c9b02e60565e36f84d705b669e"
	hash, _ := utils.HexToInt(message)
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
			s := &StarkSigner{
				publicKey: tt.fields.publicKey,
			}
			assert.Equalf(t, tt.want, s.GetAddress(), "GetAddress()")
		})
	}
}

type BaseL1Signer struct {
	privateKey *ecdsa.PrivateKey
	chainId    *big.Int
}

func (b *BaseL1Signer) GetAddress() string {
	address := crypto.PubkeyToAddress(b.privateKey.PublicKey)
	return address.Hex()
}

func (b *BaseL1Signer) SignMessage(message string) ([]byte, error) {
	hashBytes := []byte(message)
	return crypto.Sign(accounts.TextHash(hashBytes), b.privateKey)
}

func (b *BaseL1Signer) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(b.chainId)
	return types.SignTx(tx, signer, b.privateKey)
}

func GetL1Signer(privateKey string, chainId *big.Int) (signers.L1Signer, error) {
	privateKey, err := utils.TrimHexPrefix(privateKey)
	if err != nil {
		return nil, err
	}
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &BaseL1Signer{privateKey: privKey, chainId: chainId}, nil
}
