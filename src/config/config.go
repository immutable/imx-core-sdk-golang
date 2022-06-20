package config

var (
	EthereumDefaultGasLimit uint64  // Ethereum default gas limit in gas units
	EthereumDefaultGasPrice uint64  // Ethereum default gas price in wei
	EthereumEGSAPIKey       string  // Ethereum eth gas station (EGS) API key
	EthereumEGSSpeed        string  // Ethereum eth gas station (EGS) speed
	EthereumGasMultiplier   float64 // Ethereum default gas multiplier
	EthereumMaxGasPrice     uint64  // Ethereum max gas price in wei
)

type EnvConfig struct {
	EthereumClientEndpoint      string
	CoreApiEndpoint             string
	RegistrationContractAddress string
	StarkContractAddress        string
}

func init() {
	EthereumDefaultGasLimit = 300000
	EthereumDefaultGasPrice = 160000000000
	EthereumGasMultiplier = 1.5
	EthereumMaxGasPrice = 300000000000
	EthereumEGSAPIKey = ""
	EthereumEGSSpeed = "fast"
}

func DevConfig(alchemyApiKey string) EnvConfig {
	return EnvConfig{
		EthereumClientEndpoint:      "https://eth-ropsten.alchemyapi.io/v2/" + alchemyApiKey,
		CoreApiEndpoint:             "https://api.dev.x.immutable.com",
		RegistrationContractAddress: "0x7EB840223a3b1E0e8D54bF8A6cd83df5AFfC88B2",
		StarkContractAddress:        "0xd05323731807A35599BF9798a1DE15e89d6D6eF1",
	}
}

func RopstenConfig(alchemyApiKey string) EnvConfig {
	return EnvConfig{
		EthereumClientEndpoint:      "https://eth-ropsten.alchemyapi.io/v2/" + alchemyApiKey,
		CoreApiEndpoint:             "https://api.ropsten.x.immutable.com",
		RegistrationContractAddress: "0x6C21EC8DE44AE44D0992ec3e2d9f1aBb6207D864",
		StarkContractAddress:        "0x4527BE8f31E2ebFbEF4fCADDb5a17447B27d2aef",
	}
}

func MainNetConfig(alchemyApiKey string) EnvConfig {
	return EnvConfig{
		EthereumClientEndpoint:      "https://eth-mainnet.alchemyapi.io/v2/" + alchemyApiKey,
		CoreApiEndpoint:             "https://api.x.immutable.com",
		RegistrationContractAddress: "0x72a06bf2a1CE5e39cBA06c0CAb824960B587d64c",
		StarkContractAddress:        "0x5FDCCA53617f4d2b9134B29090C87D01058e27e9",
	}
}
