package main

import (
	"context"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/config"
	"immutable.com/imx-core-sdk-golang/examples/workflows/burn"
	"immutable.com/imx-core-sdk-golang/examples/workflows/deposits"
	"immutable.com/imx-core-sdk-golang/examples/workflows/minting"
	"immutable.com/imx-core-sdk-golang/examples/workflows/onboarding"
	"immutable.com/imx-core-sdk-golang/examples/workflows/orders"
	"immutable.com/imx-core-sdk-golang/examples/workflows/trades"
	"immutable.com/imx-core-sdk-golang/examples/workflows/transfers"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/examples/workflows/withdrawals"
	"immutable.com/imx-core-sdk-golang/signers/stark"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
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
	apiClient := api.NewAPIClient(configuration)

	// Using context value to switch/specify the server before sending request. If nothing is specified, the default server will be used which will be first one in the open api spec list.
	ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Sandbox)

	cfg := config.GetConfig(config.Sandbox, alchemyAPIKey)
	ethClient, err := ethereum.NewEthereumClientAndAttachContracts(ctx, &cfg, ethereum.DefaultGasParams)
	if err != nil {
		log.Panicf("error dialing ethereum client: %v\n", err)
	}
	defer ethClient.Client.Close()

	chainID, err := ethClient.Client.ChainID(ctx)
	if err != nil {
		log.Panicf("error obtaining chain id: %v\n", err)
	}

	l1signer, err := utils.NewBaseL1Signer(signerPrivateKey, chainID)
	if err != nil {
		log.Panicf("error in creating BaseL1Signer: %v\n", err)
	}

	l2signer, err := stark.GenerateStarkSigner(l1signer)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	// User registration demo.
	onboarding.DemoUserRegistrationWorkflow(ctx, apiClient.UsersApi, l1signer)

	// Minting tokens demo
	minting.DemoMintingTokens(ctx, apiClient.MintsApi, l1signer, envs["MINT_TOKEN_ID"], envs["MINT_TOKEN_ADDRESS"])

	// Deposit tokens demo.
	deposits.DemoDepositEthWorkflow(ctx, ethClient, apiClient, envs["DW_ETH_AMOUNT"], l1signer)
	deposits.DemoDepositERC20Workflow(ctx, ethClient, apiClient, envs["DW_ERC20_AMOUNT"], envs["DW_ERC20TOKEN_ADDRESS"], l1signer)
	deposits.DemoDepositERC721Workflow(ctx, ethClient, apiClient, envs["DW_ERC721TOKEN_ID"], envs["DW_ERC721TOKEN_ADDRESS"], l1signer)

	trades.DemoTradesWorkflow(ctx, apiClient, l1signer, l2signer)
	burn.DemoBurnWorkflow(ctx, apiClient, l1signer, l2signer)
	transfers.DemoTransferWorkflow(ctx, apiClient, l1signer, l2signer)
	transfers.DemoBatchTransferWorkflow(ctx, apiClient, l1signer, l2signer)

	// Withdrawals Demo
	// After prepare withdrawal workflow is performed. Must wait for getWithdrawal endpoint
	// https://docs.x.immutable.com/reference/#/operations/getWithdrawal to return "rollup_status": "confirmed"
	// before calling complete withdrawal workflow.
	withdrawals.DemoPrepareEthWithdrawalWorkflow(ctx, apiClient, envs["DW_ETH_AMOUNT"], l1signer, l2signer)
	withdrawals.DemoPrepareERC20WithdrawalWorkflow(ctx, apiClient, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	withdrawals.DemoPrepareERC721WithdrawalWorkflow(ctx,
		apiClient,
		envs["DW_ERC721TOKEN_ID"],
		envs["DW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// Make sure the tokens are ready for withdrawal before performing complete withdrawal.
	withdrawals.DemoCompleteEthWithdrawalWorkflow(ctx, ethClient, apiClient, l1signer, l2signer)
	withdrawals.DemoCompleteERC20WithdrawalWorkflow(ctx, ethClient, apiClient, envs["DW_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	withdrawals.DemoCompleteERC721WithdrawalWorkflow(ctx,
		ethClient,
		apiClient,
		envs["DW_ERC721TOKEN_ID"],
		envs["DW_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	orders.DemoOrdersWorkflow(ctx, apiClient, l1signer, l2signer)
}
