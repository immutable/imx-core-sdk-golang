package main

import (
	"context"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer, _ := common.CommonInitialise()

	// Deposit tokens demo.
	ethAmountInWei, err := strconv.ParseUint(envs["DW_ETH_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}
	DemoDepositEthWorkflow(ctx, c, ethAmountInWei, l1signer)

	// erc20AmountInWei, err := strconv.ParseUint(envs["DW_ERC20_AMOUNT_IN_WEI"], 10, 64)
	// if err != nil {
	// 	log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	// }
	//DemoDepositERC20Workflow(ctx, c, erc20AmountInWei, envs["DW_ERC20TOKEN_ADDRESS"], l1signer)
	//DemoDepositERC721Workflow(ctx, c, envs["DW_ERC721TOKEN_ID"], envs["DW_ERC721TOKEN_ADDRESS"], l1signer)
}

func DemoDepositEthWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	transaction, err := imx.NewETHDeposit(amount).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoDepositERC20Workflow(ctx context.Context, c *imx.Client, unDecimalisedAmount uint64, tokenAddress string, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	transaction, err := imx.NewERC20Deposit(unDecimalisedAmount, tokenAddress).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoDepositERC721Workflow(ctx context.Context, c *imx.Client, tokenID, tokenAddress string, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	transaction, err := imx.NewERC721Deposit(tokenID, tokenAddress).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
