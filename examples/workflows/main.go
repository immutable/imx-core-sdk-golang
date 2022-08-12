package main

import (
	"context"
	"log"
	"strings"

	"github.com/joho/godotenv"
	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/config"
	"immutable.com/imx-core-sdk-golang/examples/workflows/deposits"
	"immutable.com/imx-core-sdk-golang/examples/workflows/minting"
	"immutable.com/imx-core-sdk-golang/examples/workflows/onboarding"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/examples/workflows/withdrawals"
	"immutable.com/imx-core-sdk-golang/signers/stark"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
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
	apiClient := api.NewAPIClient(configuration)

	// Using context value to switch/specify the server before sending request. If nothing is specified, the default server will be used which will be first one in the open api spec list.
	ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Ropsten)

	cfg := config.GetConfig(config.Ropsten, alchemyAPIKey)
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

	onboarding.DemoUserRegistrationWorkflow(ctx, apiClient.UsersApi, l1signer)
	minting.DemoMintingTokens(ctx, apiClient.MintsApi, l1signer, envs["MINT_TOKEN_ID"], envs["MINT_TOKEN_ADDRESS"])

	deposits.DemoDepositWorkflow(ctx, ethClient, apiClient, l1signer)
	// trades.Demo_TradesWorkflow(ctx, apiClient, l1signer, l2signer)
	// burn.Demo_BurnWorkflow(ctx, apiClient, l1signer, l2signer)
	// transfers.Demo_TransferWorkflow(ctx, apiClient, l1signer, l2signer)
	// transfers.Demo_BatchTransferWorkflow(ctx, apiClient, l1signer, l2signer)

	// Withdrawls Demo
	// Will need to prepare withdrawal first and wait for to be ready to complete the withdrawal.
	withdrawals.DemoPrepareEthWithdrawalWorkflow(ctx, apiClient, l1signer, l2signer)
	withdrawals.DemoPrepareERC20WithdrawalWorkflow(ctx, apiClient, envs["WITHDRAWAL_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	withdrawals.DemoPrepareERC721WithdrawalWorkflow(ctx,
		apiClient,
		envs["WITHDRAWAL_ERC721TOKEN_ID"],
		envs["WITHDRAWAL_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// Make sure the tokens are ready for withdrawal before performing complete withdrawal.
	withdrawals.DemoCompleteEthWithdrawalWorkflow(ctx, ethClient, apiClient, l1signer, l2signer)
	withdrawals.DemoCompleteERC20WithdrawalWorkflow(ctx, ethClient, apiClient, envs["WITHDRAWAL_ERC20TOKEN_ADDRESS"], l1signer, l2signer)
	withdrawals.DemoCompleteERC721WithdrawalWorkflow(ctx,
		ethClient,
		apiClient,
		envs["WITHDRAWAL_ERC721TOKEN_ID"],
		envs["WITHDRAWAL_ERC721TOKEN_ADDRESS"],
		l1signer,
		l2signer)

	// orders.Demo_OrdersWorkflow(ctx, apiClient, l1signer, l2signer)
}
