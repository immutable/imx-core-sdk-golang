package config

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

// The order of this enum should match with the order specified in the OpenApi Spec document.
const (
	Sandbox = iota // Sandbox network
	Mainnet        // Production network
)

type Config struct {
	EthereumClientEndpoint      string
	CoreAPIEndpoint             string
	RegistrationContractAddress string
	StarkContractAddress        string
	ChainID                     *big.Int
}

type Network int

func GetAPIURL(network Network) string {
	return getBaseConfigs()[network].CoreAPIEndpoint
}

func GetConfig(network Network, alchemyKey string) Config {
	baseConfig := getBaseConfigs()[network]
	baseConfig.EthereumClientEndpoint += alchemyKey

	return baseConfig
}

func getBaseConfigs() map[Network]Config {
	return map[Network]Config{
		Sandbox: {
			EthereumClientEndpoint:      "https://eth-ropsten.alchemyapi.io/v2/",
			CoreAPIEndpoint:             "https://api.ropsten.x.immutable.com",
			RegistrationContractAddress: "0x6C21EC8DE44AE44D0992ec3e2d9f1aBb6207D864",
			StarkContractAddress:        "0x4527BE8f31E2ebFbEF4fCADDb5a17447B27d2aef",
			ChainID:                     params.RopstenChainConfig.ChainID,
		},
		Mainnet: {
			EthereumClientEndpoint:      "https://eth-mainnet.alchemyapi.io/v2/",
			CoreAPIEndpoint:             "https://api.x.immutable.com",
			RegistrationContractAddress: "0x72a06bf2a1CE5e39cBA06c0CAb824960B587d64c",
			StarkContractAddress:        "0x5FDCCA53617f4d2b9134B29090C87D01058e27e9",
			ChainID:                     params.MainnetChainConfig.ChainID,
		},
	}
}
