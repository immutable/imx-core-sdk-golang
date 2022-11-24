package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

const (
	ERC20TokenDecimals = 18 // This is a fixed value for almost all known ERC20 tokens.
)

func main() {
	ctx, envs, c, l1signer := common.CommonInitialise(".env")
	l2signer := common.NewStarkSigner(envs["STARK_PRIVATE_KEY_IN_HEX"])

	// Withdrawals Demo
	// For overview of the withdrawal process, please see https://docs.x.immutable.com/docs/guides/asset-management/asset-withdrawals/#withdrawal-process

	// After prepare withdrawal should wait for getWithdrawal endpoint to confirm that the withdrawal is ready to perform complete withdrawal step.
	// See https://docs.x.immutable.com/reference/#/operations/getWithdrawal it should return "rollup_status": "confirmed".

	ethAmountInWei, err := strconv.ParseUint(envs["WITHDRAW_ETH_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}
	erc20AmountInWei, err := strconv.ParseUint(envs["WITHDRAW_ERC20_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}

	DemoPrepareEthWithdrawalWorkflow(ctx, c, ethAmountInWei, l1signer, l2signer)
	DemoPrepareERC20WithdrawalWorkflow(ctx, c, erc20AmountInWei, envs["WITHDRAW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	DemoPrepareERC721WithdrawalWorkflow(ctx,
		c,
		envs["WITHDRAW_ERC721TOKEN_ID"],
		envs["WITHDRAW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// Note: Can only run either Prepare or Complete Withdrawal at a time as we need to wait for the withdrawal to be ready for collection.

	// Make sure the tokens are ready for withdrawal before performing complete withdrawal.
	DemoCompleteEthWithdrawalWorkflow(ctx, c, l1signer, l2signer)
	DemoCompleteERC20WithdrawalWorkflow(ctx, c, envs["WITHDRAW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	DemoCompleteERC721WithdrawalWorkflow(ctx,
		c,
		envs["WITHDRAW_ERC721TOKEN_ID"],
		envs["WITHDRAW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)
}

// DemoPrepareEthWithdrawalWorkflow contains sample code for preparing withdrawal of Eth tokens
func DemoPrepareEthWithdrawalWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	// To declare eth tokens, use imx.SignableETHToken method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: strconv.FormatUint(amount, 10),
		Token:  imx.SignableETHToken(),
	}

	response, err := c.PrepareWithdrawal(ctx, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	getWithdrawalResponse, err := c.GetWithdrawal(ctx, strconv.FormatInt(int64(response.WithdrawalId), 10))
	if err != nil {
		log.Panicf("error calling GetWithdrawal: %v", err)
	}
	val, _ = json.MarshalIndent(getWithdrawalResponse, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoPrepareERC20WithdrawalWorkflow contains sample code for preparing withdrawal of ERC20 tokens
func DemoPrepareERC20WithdrawalWorkflow(ctx context.Context, c *imx.Client, amount imx.Wei, tokenAddress string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

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

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoPrepareERC721WithdrawalWorkflow contains sample code for preparing withdrawal of ERC721 tokens.
func DemoPrepareERC721WithdrawalWorkflow(ctx context.Context, c *imx.Client, tokenID, tokenAddress string, l1signer imx.L1Signer, l2signer imx.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

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

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteEthWithdrawalWorkflow contains sample code for completing withdrawal of Eth tokens
func DemoCompleteEthWithdrawalWorkflow(
	ctx context.Context,
	c *imx.Client,
	l1signer imx.L1Signer,
	l2signer imx.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	ethWithdrawal := imx.NewEthWithdrawal()
	transaction, err := ethWithdrawal.CompleteWithdrawal(ctx, c, l1signer, l2signer.GetPublicKey(), nil)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.CompleteEthWithdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteERC20WithdrawalWorkflow contains sample code for completing withdrawal of ERC20 tokens
func DemoCompleteERC20WithdrawalWorkflow(
	ctx context.Context,
	c *imx.Client,
	tokenAddress string,
	l1signer imx.L1Signer,
	l2signer imx.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	erc20Withdrawal := imx.NewERC20Withdrawal(tokenAddress)
	transaction, err := erc20Withdrawal.CompleteWithdrawal(ctx, c, l1signer, l2signer.GetPublicKey(), nil)

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteERC721WithdrawalWorkflow contains sample code for completing withdrawal of ERC721 tokens
func DemoCompleteERC721WithdrawalWorkflow(
	ctx context.Context,
	c *imx.Client,
	tokenID, tokenAddress string,
	l1signer imx.L1Signer,
	l2signer imx.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	erc721Withdrawal := imx.NewERC721Withdrawal(tokenID, tokenAddress)
	transaction, err := erc721Withdrawal.CompleteWithdrawal(ctx, c, l1signer, l2signer.GetPublicKey(), nil)

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
