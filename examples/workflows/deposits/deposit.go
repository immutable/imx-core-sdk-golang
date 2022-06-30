package deposits

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/workflows"
	"immutable.com/imx-core-sdk-golang/workflows/deposits"
)

func Demo_DepositWorkflow(ctx context.Context, workflow *workflows.Workflows, l1signer signers.L1Signer) {

	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_DepositWorkflow")

	depositRequest := deposits.NewETHDeposit("0.03")
	transaction, err := workflow.Deposit(depositRequest, l1signer)
	if err != nil {
		log.Fatalf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())

	log.Println("Running Demo_DepositWorkflow completed")
	log.Println("-------------------------------------------------------")
}
