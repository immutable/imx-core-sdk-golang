package factories

import (
	"context"
	"immutable.com/imx-core-sdk/config"
	"immutable.com/imx-core-sdk/utils/ethereum"
	"math/big"
)

func NewEthereumClient(ctx context.Context, privateKey string, env config.EnvConfig) (*ethereum.Client, error) {
	ethClient, err := ethereum.NewEthereumClient(ethereum.ClientParams{
		URL:           env.EthereumClientEndpoint,
		GasLimit:      config.EthereumDefaultGasLimit,
		GasPrice:      big.NewInt(int64(config.EthereumDefaultGasPrice)),
		GasMultiplier: config.EthereumGasMultiplier,
		MaxGasPrice:   big.NewInt(int64(config.EthereumMaxGasPrice)),
		EGSAPIKey:     config.EthereumEGSAPIKey,
		EGSSpeed:      config.EthereumEGSSpeed,
	})
	if err != nil {
		return nil, err
	}
	if err := ethClient.AttachRegistrationContract(ctx, env.RegistrationContractAddress); err != nil {
		return nil, err
	}
	if err := ethClient.AttachCoreContract(ctx, env.StarkContractAddress); err != nil {
		return nil, err
	}
	if err := ethClient.LoadPrivateKey(privateKey); err != nil {
		return nil, err
	}
	return ethClient, nil
}
