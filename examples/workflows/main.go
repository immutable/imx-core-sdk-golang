package main

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/config"
	"immutable.com/imx-core-sdk-golang/examples/workflows/burn"
	"immutable.com/imx-core-sdk-golang/examples/workflows/deposits"
	"immutable.com/imx-core-sdk-golang/examples/workflows/onboarding"
	"immutable.com/imx-core-sdk-golang/examples/workflows/trades"
	"immutable.com/imx-core-sdk-golang/examples/workflows/transfers"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/factories"
	"immutable.com/imx-core-sdk-golang/signers/stark"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
)

const (
	signerPrivateKey string = "" // Provide your wallet private key here
	alchemyApiKey    string = "" // Provide your alchemy key here
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig(config.Ropsten, alchemyApiKey)
	apiClient := factories.NewApiClient(cfg)
	ethClient, err := factories.NewEthereumClient(context.Background(), cfg, ethereum.DefaultGasParams)
	if err != nil {
		log.Fatalf("error dialing ethereum client: %v\n", err)
	}
	defer ethClient.Client.Close()

	chainId, err := ethClient.Client.ChainID(ctx)
	if err != nil {
		log.Fatalf("error obtaining chain id: %v\n", err)
	}

	l1signer, err := utils.NewBaseL1Signer(signerPrivateKey, chainId)
	if err != nil {
		log.Fatalf("error in creating BaseL1Signer: %v\n", err)
	}

	l2signer, err := stark.GenerateStarkSigner(l1signer)
	if err != nil {
		log.Fatalf("error in creating StarkSigner: %v\n", err)
	}

	onboarding.Demo_UserRegistrationWorkflow(ctx, apiClient, l1signer)
	deposits.Demo_DepositWorkflow(ctx, ethClient, apiClient, l1signer)
	trades.Demo_TradesWorkflow(ctx, apiClient, l1signer, l2signer)
	burn.Demo_BurnWorkflow(ctx, apiClient, l1signer, l2signer)
	transfers.Demo_TransferWorkflow(ctx, apiClient, l1signer, l2signer)
	transfers.Demo_BatchTransferWorkflow(ctx, apiClient, l1signer, l2signer)
}
