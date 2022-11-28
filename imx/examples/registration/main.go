package main

import (
	"encoding/json"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer := common.CommonInitialise(".env")
	l2signer := common.NewStarkSigner(envs["STARK_PRIVATE_KEY_IN_HEX"])

	// Note:	This is the first step to perform for a new user wanting to use IMX.
	// 			After user registration, next step will be to create a project.

	// User registration demo.
	response, err := c.RegisterOffchain(ctx, l1signer, l2signer, "")
	if err != nil {
		log.Panicf("error in RegisterOffchain: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("RegisterOffchain response: ", string(val))

	// Get the accounts registered on offchain.
	usersResponse, err := c.GetUsers(ctx, l1signer.GetAddress())
	if err != nil {
		log.Panicf("error when calling `GetUsers: %v", err)
	}
	log.Println("Registered accounts: ", usersResponse.GetAccounts())
}
