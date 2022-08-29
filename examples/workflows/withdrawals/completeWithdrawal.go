package withdrawals

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/ethereumutil"
	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
	withdrawalsWorkflow "github.com/immutable/imx-core-sdk-golang/workflows/withdrawals"
)

// DemoCompleteEthWithdrawalWorkflow contains sample code for completing withdrawal of Eth tokens
func DemoCompleteEthWithdrawalWorkflow(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	clientAPI *api.APIClient,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	transaction, err := withdrawalsWorkflow.CompleteEthWithdrawal(ctx, ethClient, clientAPI, l1signer, l2signer.GetAddress())
	if err != nil {
		log.Panicf("error calling withdrawalsWorkflow.CompleteEthWithdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteERC20WithdrawalWorkflow contains sample code for completing withdrawal of ERC20 tokens
func DemoCompleteERC20WithdrawalWorkflow(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	clientAPI *api.APIClient,
	tokenAddress string,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	erc20Withdrawal := withdrawalsWorkflow.NewERC20Withdrawal(tokenAddress)
	transaction, err := erc20Withdrawal.CompleteWithdrawal(ctx, ethClient, clientAPI, l1signer, l2signer.GetAddress())

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

// DemoCompleteERC721WithdrawalWorkflow contains sample code for completing withdrawal of ERC721 tokens
func DemoCompleteERC721WithdrawalWorkflow(
	ctx context.Context,
	ethClient *ethereumutil.Client,
	clientAPI *api.APIClient,
	tokenID, tokenAddress string,
	l1signer signers.L1Signer,
	l2signer signers.L2Signer,
) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	erc721Withdrawal := withdrawalsWorkflow.NewERC721Withdrawal(tokenID, tokenAddress)
	transaction, err := erc721Withdrawal.CompleteWithdrawal(ctx, ethClient, clientAPI, l1signer, l2signer.GetAddress())

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
