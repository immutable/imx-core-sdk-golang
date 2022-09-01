package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/stark"
)

func demoUserRegistrationWorkflow(ctx context.Context, c *imx.Client, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	key, _ := stark.GenerateKey()
	l2signer, err := stark.NewSigner(key)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	response, err := c.RegisterOffchain(ctx, l1signer, l2signer, "")
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("registration success, response: ", string(val))

	// Get the accounts registered on offchain.
	usersResponse, httpResp, err := c.GetUsers(ctx, l1signer.GetAddress()).Execute()
	if err != nil {
		log.Panicf("error when calling `usersAPI.GetUsers: %v, HTTP response body: %v", err, httpResp.Body)
	}
	log.Println("registered accounts: ", usersResponse.GetAccounts())

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
