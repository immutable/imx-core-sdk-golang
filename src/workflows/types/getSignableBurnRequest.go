package types

import (
	"immutable.com/imx-core-sdk-golang/api"
)

type GetSignableBurnRequest struct {
	// Amount of the token to burn
	Amount string
	// Ethereum address of the user
	Sender string
	// Token to burn
	Token *api.SignableToken
}
