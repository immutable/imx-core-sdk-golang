package onboarding

import (
	"context"
	"encoding/json"
	"log"

	"immutable.com/imx-core-sdk-golang/signers"

	"immutable.com/imx-core-sdk-golang/signers/stark"
	"immutable.com/imx-core-sdk-golang/workflows"
	"immutable.com/imx-core-sdk-golang/workflows/registration"
)

func Demo_UserRegistrationWorkflow(ctx context.Context, workflow *workflows.Workflows, l1signer signers.L1Signer) {

	log.Println("-------------------------------------------------------")
	log.Println("Running Demo_UserRegistrationWorkflow")

	l2signer, err := stark.GenerateStarkSigner(l1signer)
	if err != nil {
		log.Fatalf("error in creating StarkSigner: %v\n", err)
	}

	response, err := registration.RegisterOffchain(ctx, workflow.ApiClient, l1signer, l2signer, "")
	if err != nil {
		log.Fatalf("error in creating StarkSigner: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Fatalf("error in json marshaling: %v\n", err)
	}
	log.Println("registration success, response: ", string(val))

	accounts, err := registration.IsRegisteredOffChain(ctx, workflow.ApiClient, l1signer.GetAddress())
	if err != nil {
		log.Fatalf("error in IsRegisteredOffChain function: %v\n", err)
	}
	log.Println("registered accounts: ", accounts)

	log.Println("Running Demo_UserRegistrationWorkflow completed.")
	log.Println("-------------------------------------------------------")
}
