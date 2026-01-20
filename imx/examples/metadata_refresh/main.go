package main

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer := common.CommonInitialise(".env")

	tokens := strings.Split(envs["TOKEN_IDS"], ",")
	createMetadataRequest := api.NewCreateMetadataRefreshRequest(envs["COLLECTION_CONTRACT_ADDRESS"], tokens)

	response, err := c.CreateMetadataRefresh(ctx, l1signer, createMetadataRequest)
	if err != nil {
		log.Panicf("error in CreateMetadataRefresh: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("Created new metadata refresh, response: ", string(val))

	// Get the metadata refresh details we just created.
	resultResponse, err := c.GetMetadataRefreshResults(ctx, l1signer, response.RefreshId)
	if err != nil {
		log.Println("error when calling `GetMetadataRefreshResults: %v", err)
	}
	val, err = json.MarshalIndent(resultResponse, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("GetMetadataRefreshResults, response: ", string(val))

	errorsReponse, err := c.GetMetadataRefreshErrors(ctx, l1signer, response.RefreshId, nil, nil)
	if err != nil {
		log.Println("error when calling `GetMetadataRefreshErrors: %v", err)
	}
	val, err = json.MarshalIndent(errorsReponse, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("GetMetadataRefreshErrors, response: ", string(val))
}
