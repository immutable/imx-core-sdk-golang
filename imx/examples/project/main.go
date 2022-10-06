package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, _, c, l1signer, _ := common.CommonInitialise(".env")

	// Note: Next step after creating a project would be to create a collection.

	// Create a new project demo.
	response, err := c.CreateProject(ctx, l1signer, "Tinker", "TinkerCompany", "contact@tinker.com.au")
	if err != nil {
		log.Panicf("error in CreateProject: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("Created new project, response: ", string(val))

	// Get the project details we just created.
	projectReponse, err := c.GetProject(ctx, l1signer, strconv.FormatInt(int64(response.Id), 10))
	if err != nil {
		log.Panicf("error when calling `GetProject: %v", err)
	}
	val, err = json.MarshalIndent(projectReponse, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("Project details: ", string(val))
}
