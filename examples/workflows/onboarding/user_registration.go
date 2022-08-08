package onboarding

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/api"

	"immutable.com/imx-core-sdk-golang/signers"
	"immutable.com/imx-core-sdk-golang/signers/stark"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
)

func DemoUserRegistrationWorkflow(ctx context.Context, apiClient api.UsersApi, l1signer signers.L1Signer) {
	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_UserRegistrationWorkflow")

	l2signer, err := stark.GenerateStarkSigner(l1signer)
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	response, err := registration.RegisterOffchain(ctx, apiClient, l1signer, l2signer, "")
	if err != nil {
		log.Panicf("error in creating StarkSigner: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("registration success, response: ", string(val))

	accounts, err := registration.IsRegisteredOffChain(ctx, apiClient, l1signer.GetAddress())
	if err != nil {
		log.Panicf("error in IsRegisteredOffChain function: %v\n", err)
	}
	log.Println("registered accounts: ", accounts)

	log.Println("Running DemoUserRegistrationWorkflow completed.")
	log.Println("-------------------------------------------------------")
}
