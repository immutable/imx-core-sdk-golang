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
			EthereumClientEndpoint: "https://eth-goerli.g.alchemy.com/v2/",
			CoreAPIEndpoint: url.URL{
				Scheme: "https",
				Host:   "api.dev.x.immutable.com",
			},
			RegistrationContractAddress: "0x2F76E4e48A5f9e517765B70a4DEc67781d35A199",
			StarkContractAddress:        "0x3e6e01355bB66925a65D372bf9c9f3835d9964fA",
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
