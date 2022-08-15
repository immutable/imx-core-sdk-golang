package onboarding

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/examples/workflows/utils"
	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/signers/stark"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
)

func DemoUserRegistrationWorkflow(ctx context.Context, usersAPI api.UsersApi, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Printf("Running %s", utils.GetCurrentFunctionName())

	l2signer, err := stark.GenerateStarkSigner(l1signer)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	response, err := registration.RegisterOffchain(ctx, usersAPI, l1signer, l2signer, "")
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("registration success, response: ", string(val))

	// Get the accounts registered on offchain.
	usersResponse, httpResp, err := usersAPI.GetUsers(ctx, l1signer.GetAddress()).Execute()
	if err != nil {
		log.Panicf("error when calling `usersAPI.GetUsers: %v, HTTP response body: %v", err, httpResp.Body)
	}
	log.Println("registered accounts: ", usersResponse.GetAccounts())

	log.Printf("Running %s completed.", utils.GetCurrentFunctionName())
	log.Println("-------------------------------------------------------")
}
