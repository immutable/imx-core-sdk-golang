package main

import (
	"context"
	"log"
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

	configuration := api.NewConfiguration()
	// Enable debug logging.
	if enableDebugLogging {
		configuration.Debug = true
	}

	// Using context value to switch/specify the server before sending request.
	// If nothing is specified, the default server will be used which will be first one in the open api spec list.
	ctx := context.TODO()

	cfg := imx.Config{
		APIConfig:     configuration,
		AlchemyAPIKey: alchemyAPIKey,
		Environment:   imx.Sandbox,
	}

	c, err := imx.NewClient(cfg)
	if err != nil {
		log.Panicf("error on NewClient: %v\n", err)
	}
	defer c.EthClient.Close()

	chainID, err := c.EthClient.ChainID(ctx)
	if err != nil {
		log.Panicf("error obtaining chain id: %v\n", err)
	}

	l1signer, err := ethereum.NewSigner(signerPrivateKey, chainID)
	if err != nil {
		log.Panicf("error in creating L1Signer: %v\n", err)
	}

	key, _ := stark.GenerateKey()
	l2signer, err := stark.NewSigner(key)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	// User registration demo.
	workflow.DemoUserRegistrationWorkflow(ctx, c, l1signer)

	// Minting tokens demo
	workflow.DemoMintingTokens(ctx, c, l1signer, envs["MINT_TOKEN_ID"], envs["MINT_TOKEN_ADDRESS"])

	// Deposit tokens demo.
	workflow.DemoDepositEthWorkflow(ctx, c, envs["DW_ETH_AMOUNT"], l1signer)
	workflow.DemoDepositERC20Workflow(ctx, c, envs["DW_ERC20_AMOUNT"], envs["DW_ERC20TOKEN_ADDRESS"], l1signer)
	workflow.DemoDepositERC721Workflow(ctx, c, envs["DW_ERC721TOKEN_ID"], envs["DW_ERC721TOKEN_ADDRESS"], l1signer)

	workflow.DemoTradesWorkflow(ctx, c, l1signer, l2signer)
	workflow.DemoBurnWorkflow(ctx, c, l1signer, l2signer)
	workflow.DemoTransferWorkflow(ctx, c, l1signer, l2signer)
	workflow.DemoBatchNftTransferWorkflow(ctx, c, l1signer, l2signer)

	// Withdrawals Demo
	// After prepare withdrawal workflow is performed. Must wait for getWithdrawal endpoint
	// https://docs.x.immutable.com/reference/#/operations/getWithdrawal to return "rollup_status": "confirmed"
	// before calling complete withdrawal workflow.
	workflow.DemoPrepareEthWithdrawalWorkflow(ctx, c, envs["DW_ETH_AMOUNT"], l1signer, l2signer)
	workflow.DemoPrepareERC20WithdrawalWorkflow(ctx, c, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
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

	workflow.DemoOrdersWorkflow(ctx, c, l1signer, l2signer)
}
