package withdrawals

import (
	"context"

	eth "github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
)

type TokenWithdrawal interface {
	CompleteWithdrawal(ctx context.Context, ethClient *ethereum.Client, api *client.ImmutableXAPI, l1signer signers.L1Signer, starkPublicKey string) (*eth.Transaction, error)
}

type ERC20Withdrawal struct {
	TokenID      string
	TokenAddress string
}

type ERC721Withdrawal struct {
	TokenID      string
	TokenAddress string
}

func NewERC20Withdrawal(tokenID, tokenAddress string) *ERC20Withdrawal {
	this := ERC20Withdrawal{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}

func NewERC721Withdrawal(tokenID, tokenAddress string) *ERC721Withdrawal {
	this := ERC721Withdrawal{}
	this.TokenID = tokenID
	this.TokenAddress = tokenAddress
	return &this
}
