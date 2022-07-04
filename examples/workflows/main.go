package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"immutable.com/imx-core-sdk-golang/examples/workflows/deposits"

	"immutable.com/imx-core-sdk-golang/config"

	"immutable.com/imx-core-sdk-golang/examples/workflows/onboarding"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/workflows"
)

const (
	signerPrivateKey string = "" // Provide your wallet private key here
	alchemyApiKey    string = "" // Provide your alchemy key here
)

func main() {
	ctx := context.Background()
	cfg := config.GetConfig(config.Ropsten, alchemyApiKey)
	workflow, err := workflows.NewWorkflow(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating new workflow: %v", err)
	}
	defer workflow.CloseEthClient()

	chainId, err := workflow.EthClient.Client.ChainID(ctx)
	if err != nil {
		log.Fatalf("error obtaining chain id: %v\n", err)
	}

	l1signer, err := utils.NewBaseL1Signer(signerPrivateKey, chainId)
	if err != nil {
		log.Fatalf("error in creating BaseL1Signer: %v\n", err)
	}

	onboarding.Demo_UserRegistrationWorkflow(ctx, workflow, l1signer)
	deposits.Demo_DepositWorkflow(ctx, workflow, l1signer)
}
