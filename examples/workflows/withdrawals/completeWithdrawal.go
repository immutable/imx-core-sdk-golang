package withdrawals

import (
	"context"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/utils/ethereum"
	withdrawalsWorkflow "immutable.com/imx-core-sdk-golang/workflows/withdrawals"
)

func Demo_CompleteWithdrawalWorkflow(ctx context.Context, ethClient *ethereum.Client, api *client.ImmutableXAPI, l1signer signers.L1Signer, l2signer signers.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_CompleteWithdrawalWorkflow")

	transaction, err := withdrawalsWorkflow.CompleteEthWithdrawal(ctx, ethClient, api, l1signer, l2signer.GetAddress())

	/*
		// Uncomment for ERC20
		tokenId := "ERC20 Token Id Here"
		tokenAddress := "ERC20 Token Address Here"
		erc20Withdrawal := withdrawalsWorkflow.NewERC20Withdrawal(tokenId, tokenAddress)
		transaction, err = erc20Withdrawal.CompleteWithdrawal(ctx, ethClient, api, l1signer, l2signer.GetAddress())
	*/

	/*
		// Uncomment for ERC721
		tokenId := "ERC721 Token Id Here"
		tokenAddress := "ERC721 Token Address Here"
		erc721Withdrawal := withdrawalsWorkflow.NewERC721Withdrawal(tokenId, tokenAddress)
		transaction, err = erc721Withdrawal.CompleteWithdrawal(ctx, ethClient, api, l1signer, l2signer.GetAddress())
	*/

	if err != nil {
		log.Panicf("error calling complete withdrawal workflow: %v", err)
	}
	log.Println("transaction hash:", transaction.Hash())

	log.Println("Running Demo_CompleteWithdrawalWorkflow completed")
	log.Println("-------------------------------------------------------")
}
