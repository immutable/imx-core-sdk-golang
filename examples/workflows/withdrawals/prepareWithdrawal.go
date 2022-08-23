package withdrawals

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	converter "github.com/immutable/imx-core-sdk-golang/internal/utils"
	"github.com/immutable/imx-core-sdk-golang/signers"
	"github.com/immutable/imx-core-sdk-golang/tokens"
	withdrawalsWorkflow "github.com/immutable/imx-core-sdk-golang/workflows/withdrawals"
)

const (
	ERC20TokenDecimals = 18 // This is a fixed value for almost all known ERC20 tokens.
)

// DemoPrepareEthWithdrawalWorkflow contains sample code for preparing withdrawal of Eth tokens
func DemoPrepareEthWithdrawalWorkflow(ctx context.Context, clientAPI *api.APIClient, amount string, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// Convert Eth Amount to its denomination.
	ethTokenValue, err := converter.ToDenomination(amount, converter.EtherDecimals)
	if err != nil {
		log.Panicf("error converting Eth amount: %v", err)
	}

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken := tokens.NewSignableTokenEth()
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: ethTokenValue.String(),
		Token:  *signableToken,
	}

	response, err := withdrawalsWorkflow.PrepareWithdrawal(ctx, clientAPI.WithdrawalsApi, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoPrepareERC20WithdrawalWorkflow contains sample code for preparing withdrawal of ERC20 tokens
func DemoPrepareERC20WithdrawalWorkflow(ctx context.Context, clientAPI *api.APIClient, tokenAddress string, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object

	amount := "5"
	// Convert ERC20 token amount to its denomination.
	erc20TokenValue, err := converter.ToDenomination(amount, ERC20TokenDecimals)
	if err != nil {
		log.Panicf("error converting ERC20 amount: %v", err)
	}

	signableToken := tokens.NewSignableTokenERC20(ERC20TokenDecimals, tokenAddress)

	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: erc20TokenValue.String(),
		Token:  *signableToken,
	}

	response, err := withdrawalsWorkflow.PrepareWithdrawal(ctx, clientAPI.WithdrawalsApi, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoPrepareERC721WithdrawalWorkflow contains sample code for preparing withdrawal of ERC721 tokens.
func DemoPrepareERC721WithdrawalWorkflow(ctx context.Context, clientAPI *api.APIClient, tokenID, tokenAddress string, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken := tokens.NewSignableTokenERC721(tokenID, tokenAddress)
	amount := "1"
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: amount,
		Token:  *signableToken,
	}

	response, err := withdrawalsWorkflow.PrepareWithdrawal(ctx, clientAPI.WithdrawalsApi, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareWithdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
