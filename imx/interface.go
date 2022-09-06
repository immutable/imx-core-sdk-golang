package imx

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/internal/contracts"
)

var (
	Sandbox = Environment{
		BaseAPIPath:                 "https://api.ropsten.x.immutable.com",
		EthereumRPC:                 "https://eth-ropsten.alchemyapi.io/v2/",
		RegistrationContractAddress: "0x6C21EC8DE44AE44D0992ec3e2d9f1aBb6207D864",
		CoreContractAddress:         "0x4527BE8f31E2ebFbEF4fCADDb5a17447B27d2aef",
		ChainID:                     big.NewInt(3), // Ropsten
	}
	Mainnet = Environment{
		BaseAPIPath:                 "https://api.x.immutable.com",
		EthereumRPC:                 "https://eth-mainnet.alchemyapi.io/v2/",
		RegistrationContractAddress: "0x72a06bf2a1CE5e39cBA06c0CAb824960B587d64c",
		CoreContractAddress:         "0x5FDCCA53617f4d2b9134B29090C87D01058e27e9",
		ChainID:                     big.NewInt(1),
	}
)

// Config is used to initialise NewClient object.
type Config struct {
	AlchemyAPIKey string
	APIConfig     *api.Configuration
	Environment
}

// Environment holds ethereum network and contract address information. Part of the config.
type Environment struct {
	BaseAPIPath                 string
	EthereumRPC                 string
	RegistrationContractAddress string
	CoreContractAddress         string
	ChainID                     *big.Int
}

// Client implements functions to get the work done with Immutable X API.
// It manages communication with the Immutable X API.
type Client struct {
	Environment          Environment
	EthClient            *ethclient.Client
	registrationContract *contracts.Registration
	coreContract         *contracts.Core
	tradesApi            api.TradesApi
	ordersApi            api.OrdersApi
	transfersApi         api.TransfersApi
	depositsApi          api.DepositsApi
	withdrawalsApi       api.WithdrawalsApi
	mintsApi             api.MintsApi
	api.AssetsApi
	usersApi api.UsersApi
	api.EncodingApi
	api.MetadataApi
	api.ProjectsApi
	tokensApi api.TokensApi
	api.CollectionsApi
	api.BalancesApi
}

// NewClient creates a new Client. Requires config to setup and initialise.
// See examples for usage reference.
func NewClient(cfg *Config) (*Client, error) {
	c := Client{
		Environment: cfg.Environment,
	}

	ethClient, err := ethclient.Dial(cfg.EthereumRPC + cfg.AlchemyAPIKey)
	if err != nil {
		return nil, err
	}
	c.EthClient = ethClient

	cfg.APIConfig.Servers = api.ServerConfigurations{{URL: cfg.BaseAPIPath}}
	apiClient := api.NewAPIClient(cfg.APIConfig)

	c.AssetsApi = apiClient.AssetsApi
	c.BalancesApi = apiClient.BalancesApi
	c.CollectionsApi = apiClient.CollectionsApi
	c.depositsApi = apiClient.DepositsApi
	c.EncodingApi = apiClient.EncodingApi
	c.MetadataApi = apiClient.MetadataApi
	c.mintsApi = apiClient.MintsApi
	c.ordersApi = apiClient.OrdersApi
	c.ProjectsApi = apiClient.ProjectsApi
	c.tokensApi = apiClient.TokensApi
	c.tradesApi = apiClient.TradesApi
	c.transfersApi = apiClient.TransfersApi
	c.usersApi = apiClient.UsersApi
	c.withdrawalsApi = apiClient.WithdrawalsApi

	if err := c.attachCoreContract(context.TODO()); err != nil {
		return nil, err
	}

	if err := c.attachRegistrationContract(context.TODO()); err != nil {
		return nil, err
	}
	return &c, nil
}
