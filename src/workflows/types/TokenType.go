package types

type TokenTypeEnum string

const (
	ETHType    = TokenTypeEnum("ETH")
	ERC20Type  = TokenTypeEnum("ERC20")
	ERC721Type = TokenTypeEnum("ERC721")
)
