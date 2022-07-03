package workflows

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/config"
	"immutable.com/imx-core-sdk-golang/factories"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/deposits"
)

type Workflows struct {
	ApiClient *client.ImmutableXAPI
	EthClient *ethereum.Client
	Config    config.Config
}

func NewWorkflow(config config.Config) (*Workflows, error) {
	ethereumClient, err := factories.NewEthereumClient(context.Background(), config, ethereum.DefaultGasParams)
	if err != nil {
		return nil, fmt.Errorf("Error dialing ethereum client: %v\n", err)
	}
	apiClient := factories.NewApiClient(config)
	workflows := &Workflows{
		ApiClient: apiClient,
		EthClient: ethereumClient,
		Config:    config,
	}
	return workflows, nil
}

func NewWorkflowWithGasParams(config config.Config, gasParams ethereum.GasParams) (*Workflows, error) {
	ethereumClient, err := factories.NewEthereumClient(context.Background(), config, gasParams)
	if err != nil {
		return nil, fmt.Errorf("Error dialing ethereum client: %v\n", err)
	}
	apiClient := factories.NewApiClient(config)
	workflows := &Workflows{
		ApiClient: apiClient,
		EthClient: ethereumClient,
		Config:    config,
	}
	return workflows, nil
}

func (w *Workflows) Deposit(depositRequest deposits.TokenDeposit, l1signer signers.L1Signer) (*types.Transaction, error) {
	return depositRequest.Execute(w.EthClient, w.ApiClient, l1signer)
}

func (w *Workflows) CloseEthClient() {
	w.EthClient.Client.Close()
}
