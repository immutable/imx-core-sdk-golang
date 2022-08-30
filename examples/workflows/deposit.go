package main

import (
	"context"
	"log"

	immutable "github.com/immutable/imx-core-sdk-golang"
)

func demoDepositEthWorkflow(ctx context.Context, c *immutable.Client, amount string, l1signer immutable.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	depositRequest := immutable.NewETHDeposit(amount)
	transaction, err := depositRequest.Deposit(ctx, c, l1signer)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func demoDepositERC20Workflow(ctx context.Context, c *immutable.Client, amount, tokenAddress string, l1signer immutable.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	transaction, err := immutable.NewERC20Deposit(amount, tokenAddress).Deposit(ctx, c, l1signer)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func demoDepositERC721Workflow(ctx context.Context, c *immutable.Client, tokenID, tokenAddress string, l1signer immutable.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	depositERC721Request := immutable.NewERC721Deposit(tokenID, tokenAddress)
	transaction, err := depositERC721Request.Deposit(ctx, c, l1signer)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
