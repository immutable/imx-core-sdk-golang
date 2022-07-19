package withdrawals

import (
	"context"

	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	imx "immutable.com/imx-core-sdk-golang/workflows/types"
)

type TokenWithdrawal interface {
	CompleteWithdrawal(ctx context.Context, ethClient *ethereum.Client, api *client.ImmutableXAPI, l1signer signers.L1Signer, starkPublicKey string) (*eth.Transaction, error)
}

type ETHWithdrawal struct {
	Type imx.TokenTypeEnum
}

type ERC20Withdrawal struct {
	Type         imx.TokenTypeEnum
	TokenID      string
	TokenAddress string
}

type ERC721Withdrawal struct {
	Type         imx.TokenTypeEnum
	TokenID      string
	TokenAddress string
}

func NewETHWithdrawal() *ETHWithdrawal {
	this := ETHWithdrawal{}
	this.Type = imx.ETHType
	return &this
}

func NewERC20Withdrawal(tokenID, tokenAddress string) *ERC20Withdrawal {
	this := ERC20Withdrawal{}
	this.Type = imx.ERC20Type
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}

func NewERC721Withdrawal(tokenID, tokenAddress string) *ERC721Withdrawal {
	this := ERC721Withdrawal{}
	this.Type = imx.ERC721Type
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}
