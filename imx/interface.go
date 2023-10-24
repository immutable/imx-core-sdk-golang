package imx

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/contracts"
)

var (
	Sandbox = Environment{
		BaseAPIPath:                 "https://api.sandbox.x.immutable.com",
		EthereumRPC:                 "https://eth-sepolia.g.alchemy.com/v2/",
		RegistrationContractAddress: "0xDbA6129C02E69405622fAdc3d5A7f8d23eac3b97",
		CoreContractAddress:         "0x2d5C349fD8464DA06a3f90b4B0E9195F3d1b7F98",
		ChainID:                     big.NewInt(11155111), // Sepolia
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

// Environment holds Ethereum network and contract address information.
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
	Environment           Environment
	EthClient             *ethclient.Client
	RegistrationContract  *contracts.Registration
	CoreContract          *contracts.Core
	TradesAPI             api.TradesApi
	OrdersAPI             api.OrdersApi
	TransfersAPI          api.TransfersApi
	DepositsAPI           api.DepositsApi
	WithdrawalsAPI        api.WithdrawalsApi
	MintsAPI              api.MintsApi
	AssetsAPI             api.AssetsApi
	UsersAPI              api.UsersApi
	MetadataAPI           api.MetadataApi
	MetadataRefreshesAPI  api.MetadataRefreshesApi
	TokensAPI             api.TokensApi
	BalancesAPI           api.BalancesApi
	ProjectsAPI           api.ProjectsApi
	CollectionsAPI        api.CollectionsApi
	EncodingAPI           api.EncodingApi
	ExchangesAPI          api.ExchangesApi
	NftCheckoutPrimaryAPI api.NftCheckoutPrimaryApi
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

	c.AssetsAPI = apiClient.AssetsApi
	c.BalancesAPI = apiClient.BalancesApi
	c.CollectionsAPI = apiClient.CollectionsApi
	c.DepositsAPI = apiClient.DepositsApi
	c.EncodingAPI = apiClient.EncodingApi
	c.ExchangesAPI = apiClient.ExchangesApi
	c.MetadataAPI = apiClient.MetadataApi
	c.MetadataRefreshesAPI = apiClient.MetadataRefreshesApi
	c.MintsAPI = apiClient.MintsApi
	c.NftCheckoutPrimaryAPI = apiClient.NftCheckoutPrimaryApi
	c.OrdersAPI = apiClient.OrdersApi
	c.ProjectsAPI = apiClient.ProjectsApi
	c.TokensAPI = apiClient.TokensApi
	c.TradesAPI = apiClient.TradesApi
	c.TransfersAPI = apiClient.TransfersApi
	c.UsersAPI = apiClient.UsersApi
	c.WithdrawalsAPI = apiClient.WithdrawalsApi

	if err := c.attachCoreContract(context.TODO()); err != nil {
		return nil, err
	}

	if err := c.attachRegistrationContract(context.TODO()); err != nil {
		return nil, err
	}
	return &c, nil
}
