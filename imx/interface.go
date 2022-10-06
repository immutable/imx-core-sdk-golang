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
	tradesAPI            api.TradesApi
	ordersAPI            api.OrdersApi
	transfersAPI         api.TransfersApi
	depositsAPI          api.DepositsApi
	withdrawalsAPI       api.WithdrawalsApi
	mintsAPI             api.MintsApi
	assetsAPI            api.AssetsApi
	usersAPI             api.UsersApi
	metadataAPI          api.MetadataApi
	metadataRefreshesAPI api.MetadataRefreshesApi
	tokensAPI            api.TokensApi
	balancesAPI          api.BalancesApi
	projectsAPI          api.ProjectsApi
	collectionsAPI       api.CollectionsApi
	api.EncodingApi
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

	c.assetsAPI = apiClient.AssetsApi
	c.balancesAPI = apiClient.BalancesApi
	c.collectionsAPI = apiClient.CollectionsApi
	c.depositsAPI = apiClient.DepositsApi
	c.EncodingApi = apiClient.EncodingApi
	c.metadataAPI = apiClient.MetadataApi
	c.mintsAPI = apiClient.MintsApi
	c.ordersAPI = apiClient.OrdersApi
	c.projectsAPI = apiClient.ProjectsApi
	c.tokensAPI = apiClient.TokensApi
	c.tradesAPI = apiClient.TradesApi
	c.transfersAPI = apiClient.TransfersApi
	c.usersAPI = apiClient.UsersApi
	c.withdrawalsAPI = apiClient.WithdrawalsApi

	if err := c.attachCoreContract(context.TODO()); err != nil {
		return nil, err
	}

	if err := c.attachRegistrationContract(context.TODO()); err != nil {
		return nil, err
	}
	return &c, nil
}
