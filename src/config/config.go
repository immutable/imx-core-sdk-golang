package config

import (
	"net/url"

	"immutable.com/imx-core-sdk-golang/generated/client"
)

const (
	EthereumDefaultGasLimit uint64  = 30000        // Ethereum default gas limit in gas units
	EthereumDefaultGasPrice uint64  = 160000000000 // Ethereum default gas price in wei
	EthereumEGSAPIKey       string  = ""           // Ethereum eth gas station (EGS) API key
	EthereumEGSSpeed        string  = "fast"       // Ethereum eth gas station (EGS) speed
	EthereumGasMultiplier   float64 = 1.5          // Ethereum default gas multiplier
	EthereumMaxGasPrice     uint64  = 300000000000 // Ethereum max gas price in wei
)

type Config struct {
	EthereumClientEndpoint      url.URL
	CoreApiEndpoint             url.URL
	RegistrationContractAddress string
	StarkContractAddress        string
}

func NewDevConfig() (*Config, error) {
	ethereumClientEndpoint, err := url.Parse("https://eth-ropsten.alchemyapi.io/v2/")
	if err != nil {
		return nil, err
	}
	coreApiEndpoint, err := url.Parse("https://api.dev.x.immutable.com")
	if err != nil {
		return nil, err
	}

	return &Config{
		EthereumClientEndpoint:      *ethereumClientEndpoint,
		CoreApiEndpoint:             *coreApiEndpoint,
		RegistrationContractAddress: "0x7EB840223a3b1E0e8D54bF8A6cd83df5AFfC88B2",
		StarkContractAddress:        "0xd05323731807A35599BF9798a1DE15e89d6D6eF1",
	}, nil
}

func NewRopstenConfig() (*Config, error) {
	ethereumClientEndpoint, err := url.Parse("https://eth-ropsten.alchemyapi.io/v2/")
	if err != nil {
		return nil, err
	}
	coreApiEndpoint, err := url.Parse("https://api.ropsten.x.immutable.com")
	if err != nil {
		return nil, err
	}

	return &Config{
		EthereumClientEndpoint:      *ethereumClientEndpoint,
		CoreApiEndpoint:             *coreApiEndpoint,
		RegistrationContractAddress: "0x6C21EC8DE44AE44D0992ec3e2d9f1aBb6207D864",
		StarkContractAddress:        "0x4527BE8f31E2ebFbEF4fCADDb5a17447B27d2aef",
	}, nil
}

func NewMainNetConfig() (*Config, error) {
	ethereumClientEndpoint, err := url.Parse("https://eth-mainnet.alchemyapi.io/v2/")
	if err != nil {
		return nil, err
	}
	coreApiEndpoint, err := url.Parse("https://api.x.immutable.com")
	if err != nil {
		return nil, err
	}

	return &Config{
		EthereumClientEndpoint:      *ethereumClientEndpoint,
		CoreApiEndpoint:             *coreApiEndpoint,
		RegistrationContractAddress: "0x72a06bf2a1CE5e39cBA06c0CAb824960B587d64c",
		StarkContractAddress:        "0x5FDCCA53617f4d2b9134B29090C87D01058e27e9",
	}, nil
}

func NewTransportConfig(coreEndPoint url.URL) *client.TransportConfig {
	return &client.TransportConfig{
		Host:     coreEndPoint.Host,
		BasePath: coreEndPoint.Path,
		Schemes:  []string{coreEndPoint.Scheme},
	}
}
