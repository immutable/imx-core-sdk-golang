package main

import (
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer, _ := common.CommonInitialise(".env")

	// Deposit tokens demo.

	// Eth Deposit
	ethAmountInWei, err := strconv.ParseUint(envs["DEPOSIT_ETH_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}

	transaction, err := imx.NewETHDeposit(ethAmountInWei).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("Eth deposit: %v", err)
	}
	log.Println("Eth Deposit transaction hash:", transaction.Hash())

	// ERC20 Deposit
	erc20AmountInWei, err := strconv.ParseUint(envs["DEPOSIT_ERC20_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}
	transaction, err = imx.NewERC20Deposit(erc20AmountInWei, envs["DEPOSIT_ERC20TOKEN_ADDRESS"]).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("ERC20 deposit workflow: %v", err)
	}
	log.Println("ERC20 deposit transaction hash:", transaction.Hash())

	// ERC721 Deposit
	transaction, err = imx.NewERC721Deposit(envs["DEPOSIT_ERC721TOKEN_ID"], envs["DEPOSIT_ERC721TOKEN_ADDRESS"]).Deposit(ctx, c, l1signer, nil)
	if err != nil {
		log.Panicf("ERC721 deposit workflow: %v", err)
	}
	log.Println("ERC721 deposit transaction hash:", transaction.Hash())
}
