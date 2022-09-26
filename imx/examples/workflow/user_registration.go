package workflow

import (
	"context"
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx"
	"github.com/immutable/imx-core-sdk-golang/imx/signers/stark"
)

func DemoUserRegistrationWorkflow(ctx context.Context, c *imx.Client, l1signer imx.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", getCurrentFunctionName())

	key, err := stark.GenerateKey()
	if err != nil {
		log.Panicf("error in generating private key for stark signer: %v\n", err)
	}
	l2signer, err := stark.NewSigner(key)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	response, err := c.RegisterOffchain(ctx, l1signer, l2signer, "")
	if err != nil {
		log.Panicf("error in RegisterOffchain: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("registration success, response: ", string(val))

	// Get the accounts registered on offchain.
	usersResponse, err := c.GetUsers(ctx, l1signer.GetAddress())
	if err != nil {
		log.Panicf("error when calling `usersAPI.GetUsers: %v", err)
	}
	log.Println("registered accounts: ", usersResponse.GetAccounts())

	log.Printf("Running %s completed.", getCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
