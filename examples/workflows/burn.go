package main

import (
	"context"
	"encoding/json"
	"log"
	"strconv"

	immutable "github.com/immutable/imx-core-sdk-golang"
)

func demoBurnWorkflow(ctx context.Context, c *immutable.Client, l1signer immutable.L1Signer, l2signer immutable.L2Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	burnRequest := immutable.GetSignableBurnRequest{
		Amount: "100000000",
		Sender: l1signer.GetAddress(),
		Token:  immutable.SignableETHToken(),
	}

	/*
		signableToken := tokens.NewSignableTokenERC20(18, "0x73f99ca65b1a0aef2d4591b1b543d789860851bf")
		ftcToBurn, _ := convert.ToDenomination(".025", 18)
		burnRequest := burnWorkflow.GetSignableBurnRequest{
			 Amount: ftcToBurn.String(),
			 Sender: l1signer.GetAddress(),
			 Token:  signableToken,
		}
	*/

	/*
		signableToken := tokens.NewSignableTokenERC721("token id to burn", "token contract address")
		burnRequest := burnWorkflow.GetSignableBurnRequest{
			 Amount: "1",
			 Sender: l1signer.GetAddress(),
			 Token:  signableToken,
		}
	*/

	burnResponse, err := c.Burn(ctx, l1signer, l2signer, burnRequest)
	if err != nil {
		log.Panicf("error calling burn workflow: %v", err)
	}
	val, _ := json.MarshalIndent(burnResponse, "", "  ")
	log.Printf("Burn response:\n%s\n", val)

	transferID := strconv.FormatInt(int64(burnResponse.TransferId), 10) // GetBurn method requires transferID as a string
	transfer, err := c.GetBurn(ctx, transferID)
	if err != nil {
		log.Panicf("error calling GetBurn workflow: %v", err)
	}
	val, _ = json.MarshalIndent(transfer, "", "  ")
	log.Printf("GetBurn response:\n%s\n", val)

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
