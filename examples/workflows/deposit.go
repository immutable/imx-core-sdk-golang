package main

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
)

func demoDepositEthWorkflow(ctx context.Context, c *imx.Client, amount string, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	depositRequest := imx.NewETHDeposit(amount)
	transaction, err := depositRequest.Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func demoDepositERC20Workflow(ctx context.Context, c *imx.Client, amount, tokenAddress string, l1signer imx.L1Signer) {
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

func demoDepositERC721Workflow(ctx context.Context, c *imx.Client, tokenID, tokenAddress string, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	depositERC721Request := imx.NewERC721Deposit(tokenID, tokenAddress)
	transaction, err := depositERC721Request.Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
