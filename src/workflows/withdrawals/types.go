package withdrawals

import (
	"context"
	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	. "immutable.com/imx-core-sdk-golang/workflows/types"
)

type TokenWithdrawal interface {
	CompleteWithdrawal(ctx context.Context, ethClient *ethereum.Client, api *client.ImmutableXAPI, l1signer signers.L1Signer, starkPublicKey string) (*eth.Transaction, error)
}

type ETHWithdrawal struct {
	Type TokenTypeEnum
}

type ERC20Withdrawal struct {
	Type         TokenTypeEnum
	TokenId      string
	TokenAddress string
}

type ERC721Withdrawal struct {
	Type         TokenTypeEnum
	TokenId      string
	TokenAddress string
}

func NewETHWithdrawal() *ETHWithdrawal {
	this := ETHWithdrawal{}
	this.Type = ETHType
	return &this
}

func NewERC20Withdrawal(tokenId, tokenAddress string) *ERC20Withdrawal {
	this := ERC20Withdrawal{}
	this.Type = ERC20Type
	this.TokenId = tokenId
	this.TokenAddress = tokenAddress
	return &this
}

func NewERC721Withdrawal(tokenId, tokenAddress string) *ERC721Withdrawal {
	this := ERC721Withdrawal{}
	this.Type = ERC721Type
	this.TokenId = tokenId
	this.TokenAddress = tokenAddress
	return &this
}
