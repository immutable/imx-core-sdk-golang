package main

import (
	"context"
	"log"
	"strconv"
	"strings"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/workflow"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/ethereum"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/stark"
	"github.com/joho/godotenv"
)

func main() {
	var envs map[string]string
	envs, err := godotenv.Read(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	alchemyAPIKey := envs["ALCHEMY_API_KEY"]
	signerPrivateKey := envs["OWNER_ACCOUNT_PRIVATE_KEY"]
	enableDebugLogging := strings.EqualFold(envs["DEBUG_LOGGING"], "true")

	apiConfiguration := api.NewConfiguration()
	// Enable debug logging.
	if enableDebugLogging {
		apiConfiguration.Debug = true
	}

	// Using context value to switch/specify the server before sending request.
	// If nothing is specified, the default server will be used which will be first one in the open api spec list.
	ctx := context.TODO()

	cfg := imx.Config{
		APIConfig:     apiConfiguration,
		AlchemyAPIKey: alchemyAPIKey,
		Environment:   imx.Sandbox,
	}

	c, err := imx.NewClient(&cfg)
	if err != nil {
		log.Panicf("error on NewClient: %v\n", err)
	}
	defer c.EthClient.Close()

	l1signer, err := ethereum.NewSigner(signerPrivateKey, cfg.ChainID)
	if err != nil {
		log.Panicf("error in creating L1Signer: %v\n", err)
	}

	key, err := stark.GenerateKey()
	if err != nil {
		log.Panicf("error in Generating Stark Private Key: %v\n", err)
	}

	l2signer, err := stark.NewSigner(key)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	// User registration demo.
	workflow.DemoUserRegistrationWorkflow(ctx, c, l1signer)

	// Minting tokens demo
	workflow.DemoMintingTokens(ctx, c, l1signer, envs["MINT_TOKEN_ID"], envs["MINT_TOKEN_ADDRESS"])

	// Deposit tokens demo.
	ethAmountInWei, err := strconv.ParseUint(envs["DW_ETH_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}
	workflow.DemoDepositEthWorkflow(ctx, c, ethAmountInWei, l1signer)

	erc20AmountInWei, err := strconv.ParseUint(envs["DW_ERC20_AMOUNT_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}
	workflow.DemoDepositERC20Workflow(ctx, c, erc20AmountInWei, envs["DW_ERC20TOKEN_ADDRESS"], l1signer)
	workflow.DemoDepositERC721Workflow(ctx, c, envs["DW_ERC721TOKEN_ID"], envs["DW_ERC721TOKEN_ADDRESS"], l1signer)

	workflow.DemoTradesWorkflow(ctx, c, l1signer, l2signer)
	workflow.DemoTransferWorkflow(ctx, c, l1signer, l2signer)
	workflow.DemoBatchNftTransferWorkflow(ctx, c, l1signer, l2signer)

	// Withdrawals Demo
	// After prepare withdrawal workflow is performed. Must wait for getWithdrawal endpoint
	// https://docs.x.immutable.com/reference/#/operations/getWithdrawal to return "rollup_status": "confirmed"
	// before calling complete withdrawal workflow.
	workflow.DemoPrepareEthWithdrawalWorkflow(ctx, c, ethAmountInWei, l1signer, l2signer)
	workflow.DemoPrepareERC20WithdrawalWorkflow(ctx, c, erc20AmountInWei, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	workflow.DemoPrepareERC721WithdrawalWorkflow(ctx,
		c,
		envs["DW_ERC721TOKEN_ID"],
		envs["DW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// Make sure the tokens are ready for withdrawal before performing complete withdrawal.
	workflow.DemoCompleteEthWithdrawalWorkflow(ctx, c, l1signer, l2signer)
	workflow.DemoCompleteERC20WithdrawalWorkflow(ctx, c, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	workflow.DemoCompleteERC721WithdrawalWorkflow(ctx,
		c,
		envs["DW_ERC721TOKEN_ID"],
		envs["DW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// Orders Demo
	listingPriceInWei, err := strconv.ParseUint(envs["CREATE_ORDER_LISTING_PRICE_IN_WEI"], 10, 64)
	if err != nil {
		log.Panicf("error in converting ethAmountInWei from string to int: %v\n", err)
	}
	workflow.DemoOrdersWorkflow(ctx, c, listingPriceInWei, l1signer, l2signer)
}
