package deposits

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/deposits"
)

func DemoDepositWorkflow(ctx context.Context, ethClient *ethereum.Client, clientAPI *api.APIClient, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	depositRequest := deposits.NewETHDeposit("0.03")
	transaction, err := depositRequest.Deposit(ctx, ethClient, clientAPI, l1signer)

	/*
		// Uncomment for ERC20
		amount := "Amount to deposit"
		tokenAddress := "Token address here"
		depositERC20Request := deposits.NewERC20Deposit(amount, tokenAddress)
		transaction, err := depositERC20Request.Deposit(ctx, ethClient, clientAPI, l1signer)
	*/

	/*
		// Uncomment for ERC721
		depositERC721Request := deposits.NewERC721Deposit("7", "0x0fB969a08c7C39BA99c1628b59c0B7E5611BD396")
		transaction, err := depositERC721Request.Deposit(ctx, ethClient, clientAPI, l1signer)
	*/

	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())
	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
