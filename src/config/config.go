package config

import (
	"net/url"

	"immutable.com/imx-core-sdk-golang/api/client"
)

type Config struct {
	EthereumClientEndpoint      string
	CoreAPIEndpoint             url.URL
	RegistrationContractAddress string
	StarkContractAddress        string
}

type Network string

const (
	Dev     = Network("dev")
	Ropsten = Network("ropsten")
	MainNet = Network("mainnet")
)

func getBaseConfigs() map[Network]Config {
	return map[Network]Config{
		Dev: {
			EthereumClientEndpoint: "https://eth-ropsten.alchemyapi.io/v2/",
			CoreAPIEndpoint: url.URL{
				Scheme: "https",
				Host:   "api.dev.x.immutable.com",
			},
			RegistrationContractAddress: "0x7EB840223a3b1E0e8D54bF8A6cd83df5AFfC88B2",
			StarkContractAddress:        "0xd05323731807A35599BF9798a1DE15e89d6D6eF1",
		},
		Ropsten: {
			EthereumClientEndpoint: "https://eth-ropsten.alchemyapi.io/v2/",
			CoreAPIEndpoint: url.URL{
				Scheme: "https",
				Host:   "api.ropsten.x.immutable.com",
			},
			RegistrationContractAddress: "0x6C21EC8DE44AE44D0992ec3e2d9f1aBb6207D864",
			StarkContractAddress:        "0x4527BE8f31E2ebFbEF4fCADDb5a17447B27d2aef",
		},
		MainNet: {
			EthereumClientEndpoint: "https://eth-mainnet.alchemyapi.io/v2/",
			CoreAPIEndpoint: url.URL{
				Scheme: "https",
				Host:   "api.x.immutable.com",
			},
			RegistrationContractAddress: "0x72a06bf2a1CE5e39cBA06c0CAb824960B587d64c",
			StarkContractAddress:        "0x5FDCCA53617f4d2b9134B29090C87D01058e27e9",
		},
	}
}

func GetAPIURL(network Network) url.URL {
	return getBaseConfigs()[network].CoreAPIEndpoint
}

func GetConfig(network Network, alchemyKey string) Config {
	baseConfig := getBaseConfigs()[network]
	baseConfig.EthereumClientEndpoint += alchemyKey

	return baseConfig
}

func NewTransportConfig(coreEndPoint *url.URL) *client.TransportConfig {
	return &client.TransportConfig{
		Host:     coreEndPoint.Host,
		BasePath: coreEndPoint.Path,
		Schemes:  []string{coreEndPoint.Scheme},
	}
}
