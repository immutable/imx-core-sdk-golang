package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/internal/convert"
)

const (
	ERC20TokenDecimals = 18 // This is a fixed value for almost all known ERC20 tokens.
)

// DemoPrepareEthWithdrawalWorkflow contains sample code for preparing withdrawal of Eth tokens
func demoPrepareEthWithdrawalWorkflow(ctx context.Context, c *imx.Client, amount string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// Convert Eth Amount to its denomination.
	ethTokenValue, err := convert.ToDenomination(amount, convert.EtherDecimals)
	if err != nil {
		log.Panicf("error converting Eth amount: %v", err)
	}

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: ethTokenValue.String(),
		Token:  imx.SignableETHToken(),
	}

	response, err := c.PrepareWithdrawal(ctx, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoPrepareERC20WithdrawalWorkflow contains sample code for preparing withdrawal of ERC20 tokens
func demoPrepareERC20WithdrawalWorkflow(ctx context.Context, c *imx.Client, tokenAddress string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object

	amount := "5"
	// Convert ERC20 token amount to its denomination.
	erc20TokenValue, err := convert.ToDenomination(amount, ERC20TokenDecimals)
	if err != nil {
		log.Panicf("error converting ERC20 amount: %v", err)
	}

	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: erc20TokenValue.String(),
		Token:  imx.SignableERC20Token(ERC20TokenDecimals, tokenAddress),
	}

	response, err := c.PrepareWithdrawal(ctx, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoPrepareERC721WithdrawalWorkflow contains sample code for preparing withdrawal of ERC721 tokens.
func demoPrepareERC721WithdrawalWorkflow(ctx context.Context, c *imx.Client, tokenID, tokenAddress string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	amount := "1"
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: amount,
		Token:  imx.SignableERC721Token(tokenID, tokenAddress),
	}

	response, err := c.PrepareWithdrawal(ctx, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
