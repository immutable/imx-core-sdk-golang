package withdrawals

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/tokens"
	withdrawalsWorkflow "immutable.com/imx-core-sdk-golang/workflows/withdrawals"
)

const (
	ERC20TokenDecimals = 18 // This is a fixed value for almost all known ERC20 tokens.
)

// DemoPrepareEthWithdrawalWorkflow contains sample code for preparing withdrawal of Eth tokens
func DemoPrepareEthWithdrawalWorkflow(ctx context.Context, clientAPI *api.APIClient, amount string, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken := tokens.NewSignableTokenEth()
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: amount,
		Token:  *signableToken,
	}

	response, err := withdrawalsWorkflow.PrepareEthWithdrawal(ctx, clientAPI.WithdrawalsApi, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareEthWithdrawal workflow: %v", err)
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
	signableToken := tokens.NewSignableTokenERC20(ERC20TokenDecimals, tokenAddress)
	amount := "5"
	withdrawalRequest := api.GetSignableWithdrawalRequest{
		Amount: amount,
		Token:  *signableToken,
	}

	response, err := withdrawalsWorkflow.PrepareERC20Withdrawal(ctx, clientAPI.WithdrawalsApi, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareERC20Withdrawal workflow: %v", err)
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

	response, err := withdrawalsWorkflow.PrepareERC721Withdrawal(ctx, clientAPI.WithdrawalsApi, l1signer, l2signer, withdrawalRequest)
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.PrepareERC721Withdrawal workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
