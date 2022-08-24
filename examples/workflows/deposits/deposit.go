package deposits

import (
	"context"
	"log"

	"github.com/immutable/imx-core-sdk-golang/ethereum"
	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
	"github.com/immutable/imx-core-sdk-golang/workflows/deposits"
)

func DemoDepositEthWorkflow(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, amount string, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	depositRequest := deposits.NewETHDeposit(amount)
	transaction, err := depositRequest.Deposit(ctx, ethClient, clientAPI, l1signer)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoDepositERC20Workflow(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, amount, tokenAddress string, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	transaction, err := deposits.NewERC20Deposit(amount, tokenAddress).Deposit(ctx, ethClient, clientAPI, l1signer)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoDepositERC721Workflow(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, tokenID, tokenAddress string, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	depositERC721Request := deposits.NewERC721Deposit(tokenID, tokenAddress)
	transaction, err := depositERC721Request.Deposit(ctx, ethClient, clientAPI, l1signer)
	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
