package workflow

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
)

func DemoDepositEthWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	transaction, err := imx.NewETHDeposit(amount).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoDepositERC20Workflow(ctx context.Context, c *imx.Client, amount imx.Wei, tokenAddress string, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	transaction, err := imx.NewERC20Deposit(amount, tokenAddress).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoDepositERC721Workflow(ctx context.Context, c *imx.Client, tokenID, tokenAddress string, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	transaction, err := imx.NewERC721Deposit(tokenID, tokenAddress).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
