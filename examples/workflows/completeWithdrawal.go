package main

import (
	"context"
	"log"

	immutable "github.com/immutable/imx-core-sdk-golang"
)

// DemoCompleteEthWithdrawalWorkflow contains sample code for completing withdrawal of Eth tokens
func demoCompleteEthWithdrawalWorkflow(
	ctx context.Context,
	c *immutable.Client,
	l1signer immutable.L1Signer,
	l2signer immutable.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	transaction, err := c.CompleteEthWithdrawal(ctx, l1signer, l2signer.GetAddress(), nil)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.CompleteEthWithdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteERC20WithdrawalWorkflow contains sample code for completing withdrawal of ERC20 tokens
func demoCompleteERC20WithdrawalWorkflow(
	ctx context.Context,
	c *immutable.Client,
	tokenAddress string,
	l1signer immutable.L1Signer,
	l2signer immutable.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	erc20Withdrawal := immutable.NewERC20Withdrawal(tokenAddress)
	transaction, err := erc20Withdrawal.CompleteWithdrawal(ctx, c, l1signer, l2signer.GetAddress(), nil)

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteERC721WithdrawalWorkflow contains sample code for completing withdrawal of ERC721 tokens
func demoCompleteERC721WithdrawalWorkflow(
	ctx context.Context,
	c *immutable.Client,
	tokenID, tokenAddress string,
	l1signer immutable.L1Signer,
	l2signer immutable.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	erc721Withdrawal := immutable.NewERC721Withdrawal(tokenID, tokenAddress)
	transaction, err := erc721Withdrawal.CompleteWithdrawal(ctx, c, l1signer, l2signer.GetAddress(), nil)

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
