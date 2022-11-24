package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer := common.CommonInitialise(".env")

	// For transfer to work the user need to supply the same stark private key associated with their account.
	// Ethereum wallet and stark wallet are 1:1 mapping.
	if envs["STARK_PRIVATE_KEY_IN_HEX"] == "" {
		log.Fatalf("No STARK_PRIVATE_KEY_IN_HEX supplied, a registered user required to perform transfer. See Registration example to register user")
	}
	l2signer := common.NewStarkSigner(envs["STARK_PRIVATE_KEY_IN_HEX"])

	// Transfer workflow demo
	DemoTransferWorkflow(ctx, c, l1signer, l2signer, envs)
	//DemoBatchNftTransferWorkflow(ctx, c, l1signer, l2signer, envs)
}

func DemoTransferWorkflow(ctx context.Context, c *imx.Client, l1signer imx.L1Signer, l2signer imx.L2Signer, envs map[string]string) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	// To declare tokens, use utils.NewSignableToken[type] method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object

	transferRequest := api.GetSignableTransferRequestV1{
		Amount:   envs["TRANSFER_NUMBER_OF_TOKENS"],
		Sender:   l1signer.GetAddress(),
		Token:    imx.SignableETHToken(),
		Receiver: envs["RECEIVER_ETHEREUM_ACCOUNT_ADDRESS"],
	}

	response, err := c.Transfer(ctx, l1signer, l2signer, transferRequest)
	if err != nil {
		log.Panicf("error calling transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}

func DemoBatchNftTransferWorkflow(ctx context.Context, c *imx.Client, l1signer imx.L1Signer, l2signer imx.L2Signer, envs map[string]string) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", common.GetCurrentFunctionName())

	// To declare tokens, use tokens.NewSignableTokenERC721 method.
	// For more information about ETH, ERC20, and ERC721 tokens see https://docs.x.immutable.com/docs/token-data-object
	signableToken1 := imx.SignableERC721Token(envs["TRANSFER_NFT_TOKEN_ID_1"], envs["TRANSFER_NFT_TOKEN_ADDRESS_1"])
	signableToken2 := imx.SignableERC721Token(envs["TRANSFER_NFT_TOKEN_ID_2"], envs["TRANSFER_NFT_TOKEN_ADDRESS_2"])
	amount := "1"

	transferRequest1 := api.SignableTransferDetails{
		Amount:   amount,
		Receiver: envs["RECEIVER_ETHEREUM_ACCOUNT_ADDRESS"],
		Token:    signableToken1,
	}
	transferRequest2 := api.SignableTransferDetails{
		Amount:   amount,
		Receiver: envs["RECEIVER_ETHEREUM_ACCOUNT_ADDRESS"],
		Token:    signableToken2,
	}

	batchTransferRequest := api.GetSignableTransferRequest{
		SenderEtherKey: l1signer.GetAddress(),
		SignableRequests: []api.SignableTransferDetails{
			transferRequest1,
			transferRequest2,
		},
	}

	response, err := c.BatchNftTransfer(ctx, l1signer, l2signer, batchTransferRequest)
	if err != nil {
		log.Panicf("error calling batch transfer workflow: %v", err)
	}
	val, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("response:\n%s\n", val)

	log.Printf("Running %s completed.", common.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
