package main

import (
	"context"
	"encoding/json"
	"log"

	immutable "github.com/immutable/imx-core-sdk-golang"
	"github.com/immutable/imx-core-sdk-golang/api"
)

func demoTransferWorkflow(ctx context.Context, c *immutable.Client, l1signer immutable.L1Signer, l2signer immutable.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	amount := "100000000"
	sender := l1signer.GetAddress()
	receiver := "Set receiver address here"
	transferRequest := api.GetSignableTransferRequestV1{
		Amount:   amount,
		Sender:   sender,
		Token:    immutable.SignableETHToken(),
		Receiver: receiver,
	}

	response, err := c.CreateTransfer(ctx, l1signer, l2signer, transferRequest)
	if err != nil {
		log.Panicf("error calling transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func demoBatchNftTransferWorkflow(ctx context.Context, c *immutable.Client, l1signer immutable.L1Signer, l2signer immutable.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	// To declare tokens, use tokens.NewSignableTokenERC721 method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken1 := immutable.SignableERC721Token("Token ID Here", "Token Address Hex Here")
	signableToken2 := immutable.SignableERC721Token("Token ID Here", "Token Address Hex Here")
	amount := "1"
	sender := l1signer.GetAddress()
	receiver := "Receiver Address Here"

	transferRequest1 := api.SignableTransferDetails{
		Amount:   amount,
		Receiver: receiver,
		Token:    signableToken1,
	}
	transferRequest2 := api.SignableTransferDetails{
		Amount:   amount,
		Receiver: receiver,
		Token:    signableToken2,
	}

	batchTransferRequest := api.GetSignableTransferRequest{
		SenderEtherKey: sender,
		SignableRequests: []api.SignableTransferDetails{
			transferRequest1,
			transferRequest2,
		},
	}

	response, err := c.CreateBatchNftTransfer(ctx, l1signer, l2signer, batchTransferRequest)
	if err != nil {
		log.Panicf("error calling batch transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}