package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils"
	"math/big"
)

type BaseL1Signer struct {
	privateKey *ecdsa.PrivateKey
	chainId    *big.Int
}

func (b *BaseL1Signer) GetAddress() common.Address {
	address := crypto.PubkeyToAddress(b.privateKey.PublicKey)
	return address
}

func (b *BaseL1Signer) SignMessage(message string) ([]byte, error) {
	bytes, err := common.ParseHexOrString(message)
	if err != nil {
		return nil, err
	}
	return crypto.Sign(accounts.TextHash(bytes), b.privateKey)
}

func (b *BaseL1Signer) SignTx(tx *types.Transaction) (*types.Transaction, error) {
	signer := types.LatestSignerForChainID(b.chainId)
	return types.SignTx(tx, signer, b.privateKey)
}

func GetL1Signer(privateKey string, chainId *big.Int) (signers.L1Signer, error) {
	privateKey, err := utils.RemoveHex(privateKey)
	if err != nil {
		return nil, err
	}
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}

	return &BaseL1Signer{privateKey: privKey, chainId: chainId}, nil
}
