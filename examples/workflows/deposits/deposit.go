package deposits

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	"immutable.com/imx-core-sdk-golang/workflows/deposits"
)

func Demo_DepositWorkflow(ctx context.Context, ethClient *ethereum.Client, apis *client.ImmutableXAPI, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_DepositWorkflow")

	depositRequest := deposits.NewETHDeposit("0.03")
	transaction, err := depositRequest.Deposit(ctx, ethClient, apis, l1signer)

	/*
		// Uncomment for ERC20
		amount := "Amount to deposit"
		tokenAddress := "Token address here"
		depositERC20Request := deposits.NewERC20Deposit(amount, tokenAddress)
		transaction, err := depositERC20Request.Deposit(ctx, ethClient, apis, l1signer)
	*/

	/*
		// Uncomment for ERC721
		depositERC721Request := deposits.NewERC721Deposit("7", "0x0fB969a08c7C39BA99c1628b59c0B7E5611BD396")
		transaction, err := depositERC721Request.Deposit(ctx, ethClient, apis, l1signer)
	*/

	if err != nil {
		log.Panicf("error calling deposit workflow: %v", err)
	}

	log.Println("transaction hash:", transaction.Hash())

	log.Println("Running Demo_DepositWorkflow completed")
	log.Println("-------------------------------------------------------")
}
