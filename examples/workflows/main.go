package main

import (
	"context"
	"log"
	"strings"

	immutable "github.com/immutable/imx-core-sdk-golang"
	"github.com/immutable/imx-core-sdk-golang/api"
	"github.com/immutable/imx-core-sdk-golang/signers/ethereum"
	"github.com/immutable/imx-core-sdk-golang/signers/stark"
	"github.com/joho/godotenv"
)

func main() {
	var envs map[string]string
	envs, err := godotenv.Read("../.env")
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
	// ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Sandbox)
	ctx := context.TODO()

	cfg := immutable.Config{
		APIConfig:   configuration,
		EthereumRPC: "https://eth-ropsten.alchemyapi.io/v2/" + alchemyAPIKey,
	}

	c, err := immutable.NewClient(immutable.Sandbox, cfg)
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
		log.Panicf("error in creating BaseL1Signer: %v\n", err)
	}

	l2signer, err := stark.GenerateSigner(l1signer)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	// User registration demo.
	demoUserRegistrationWorkflow(ctx, c, l1signer)

	// Minting tokens demo
	demoMintingTokens(ctx, c, l1signer, envs["MINT_TOKEN_ID"], envs["MINT_TOKEN_ADDRESS"])

	// Deposit tokens demo.
	demoDepositEthWorkflow(ctx, c, envs["DW_ETH_AMOUNT"], l1signer)
	demoDepositERC20Workflow(ctx, c, envs["DW_ERC20_AMOUNT"], envs["DW_ERC20TOKEN_ADDRESS"], l1signer)
	demoDepositERC721Workflow(ctx, c, envs["DW_ERC721TOKEN_ID"], envs["DW_ERC721TOKEN_ADDRESS"], l1signer)

	demoTradesWorkflow(ctx, c, l1signer, l2signer)
	demoBurnWorkflow(ctx, c, l1signer, l2signer)
	demoTransferWorkflow(ctx, c, l1signer, l2signer)
	demoBatchNftTransferWorkflow(ctx, c, l1signer, l2signer)

	// Withdrawals Demo
	// After prepare withdrawal workflow is performed. Must wait for getWithdrawal endpoint
	// https://docs.x.immutable.com/reference/#/operations/getWithdrawal to return "rollup_status": "confirmed"
	// before calling complete withdrawal workflow.
	demoPrepareEthWithdrawalWorkflow(ctx, c, envs["DW_ETH_AMOUNT"], l1signer, l2signer)
	demoPrepareERC20WithdrawalWorkflow(ctx, c, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	demoPrepareERC721WithdrawalWorkflow(ctx,
		c,
		envs["DW_ERC721TOKEN_ID"],
		envs["DW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// Make sure the tokens are ready for withdrawal before performing complete withdrawal.
	demoCompleteEthWithdrawalWorkflow(ctx, c, l1signer, l2signer)
	demoCompleteERC20WithdrawalWorkflow(ctx, c, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	demoCompleteERC721WithdrawalWorkflow(ctx,
		c,
		envs["DW_ERC721TOKEN_ID"],
		envs["DW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	demoOrdersWorkflow(ctx, c, l1signer, l2signer)
}
