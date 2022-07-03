package main

import (
	"context"
	"fmt"
	"immutable.com/imx-core-sdk-golang/config"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/workflows"
	"immutable.com/imx-core-sdk-golang/workflows/deposits"
	"os"
)

func main() {
	ctx := context.Background()
	privateKey := "l1 private key here"
	alchemyApiKey := "alchemy api key here"

	devConfig := config.GetConfig(config.Ropsten, alchemyApiKey)
	workflow, err := workflows.NewWorkflow(devConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating new workflow: %v", err)
	}
	defer workflow.CloseEthClient()

	chainId, _ := workflow.EthClient.Client.ChainID(ctx)
	l1signer, _ := utils.GetL1Signer(privateKey, chainId)

	depositRequest := deposits.NewETHDeposit("0.03")
	transaction, err := workflow.Deposit(depositRequest, l1signer)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error calling deposit workflow: %v", err)
	}

	fmt.Println("transaction hash:", transaction.Hash())
}
