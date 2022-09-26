package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer, _ := common.CommonInitialise()

	// Note:	A project should be created before creating a new collection.
	//
	// 			A smart contract should be deployed and we use that contract address in creation of our collection.
	// 			Please see for how to deploy your smart contract, steps 7,8 and 9.
	//			https://docs.x.immutable.com/docs/zero-to-hero-nft-minting#step-7-create-nft-contract

	// Create a new collection
	projectID, err := strconv.ParseInt(envs["COLLECTION_PROJECT_ID"], 10, 32)
	if err != nil {
		log.Panicf("error in converting project id from string to int: %v", err)
	}

	createCollectionRequest := api.NewCreateCollectionRequest(envs["COLLECTION_CONTRACT_ADDRESS"],
		"TinkerCollection",
		l1signer.GetAddress(),
		int32(projectID))

	response, err := c.CreateCollection(ctx, l1signer, createCollectionRequest)
	if err != nil {
		log.Panicf("error in CreateCollection: %v\n", err)
	}

	val, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		log.Panicf("error in json marshaling: %v\n", err)
	}
	log.Println("Created new collection, response: ", string(val))

	// Get the collection details we just created.
	collectionReponse, err := c.GetCollection(ctx, envs["COLLECTION_CONTRACT_ADDRESS"])
	if err != nil {
		log.Panicf("error when calling `GetCollection: %v", err)
	}
	log.Println("Created Collection Name: ", collectionReponse.Name)
}
