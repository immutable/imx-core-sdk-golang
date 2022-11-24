package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/immutable/imx-core-sdk-golang/imx/api"
	"github.com/immutable/imx-core-sdk-golang/imx/examples/common"
)

func main() {
	ctx, envs, c, l1signer := common.CommonInitialise(".env")

	// Note:	A project should be created before creating a new collection.
	//
	//          Another pre-requisite to creating a collection step is to deploy a smart contract.
	//          The address of the deployed smart contract is required for creation of our collection.
	// 			For how to deploy your smart contract, please see steps 7,8 and 9 of
	//			https://docs.x.immutable.com/docs/zero-to-hero-nft-minting#step-7-create-nft-contract

	// Create a new collection
	var projectID int32
	fmt.Sscan(envs["COLLECTION_PROJECT_ID"], &projectID)

	createCollectionRequest := api.NewCreateCollectionRequest(envs["COLLECTION_CONTRACT_ADDRESS"],
		"TinkerCollection",
		l1signer.GetPublicKey(),
		projectID)

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
