package workflow

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
)

const (
	ERC20TokenDecimals = 18 // This is a fixed value for almost all known ERC20 tokens.
)

// DemoPrepareEthWithdrawalWorkflow contains sample code for preparing withdrawal of Eth tokens
func DemoPrepareEthWithdrawalWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare eth tokens, use imx.SignableETHToken method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: strconv.FormatUint(amount, 10),
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
func DemoPrepareERC20WithdrawalWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, tokenAddress string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare ERC20 token, use imx.SignableERC20Token method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: strconv.FormatUint(amount, 10),
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
func DemoPrepareERC721WithdrawalWorkflow(ctx context.Context, c *imx.Client, tokenID, tokenAddress string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare ERC721 token, use imx.SignableERC721Token method.
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
