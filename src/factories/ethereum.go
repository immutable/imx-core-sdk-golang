package factories

import (
	"context"

	"immutable.com/imx-core-sdk-golang/config"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
)

func NewEthereumClient(ctx context.Context, env *config.Config, gasParams ethereum.GasParams) (*ethereum.Client, error) {
	ethClient, err := ethereum.NewEthereumClient(env.EthereumClientEndpoint, gasParams)
	if err != nil {
		return nil, err
	}
	if err := ethClient.AttachRegistrationContract(ctx, env.RegistrationContractAddress); err != nil {
		return nil, err
	}
	if err := ethClient.AttachCoreContract(ctx, env.StarkContractAddress); err != nil {
		return nil, err
	}
	return ethClient, nil
}
