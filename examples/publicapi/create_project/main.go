package main

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/immutable/imx-core-sdk-golang/config"
	"github.com/immutable/imx-core-sdk-golang/examples/workflows/utils"
	"github.com/immutable/imx-core-sdk-golang/generated/api"
	"github.com/immutable/imx-core-sdk-golang/signers"
)

func main() {
	environment := config.GetConfig(config.Sandbox, "alchemy api key")
	configuration := api.NewConfiguration()
	apiClient := api.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Sandbox)
	l1signer, err := utils.NewBaseL1Signer("Project Owner Private Key", environment.ChainID)
	if err != nil {
		log.Panicf("error in creating BaseL1Signer: %v\n", err)
	}

	timestamp, signature, err := GetProjectOwnerAuthorisationHeaders(l1signer)
	if err != nil {
		log.Panicf("error in GetProjectOwnerAuthorisationHeaders: %v\n", err)
	}

	createProjectRequest := api.NewCreateProjectRequest("Company Name", "contact@email.com", "Project Name")
	createProjectResponse, httpResp, err := apiClient.ProjectsApi.CreateProject(ctx).
		CreateProjectRequest(*createProjectRequest).
		IMXTimestamp(timestamp).
		IMXSignature(signature).
		Execute()
	if err != nil {
		log.Panicf("error when calling `CreateProject`: %v, HTTP response body: %v", err, httpResp.Body)
	}

	log.Printf("Created: %+v", createProjectResponse)
}

// GetProjectOwnerAuthorisationHeaders generates authorisation headers for api calls which require project owner authorisation in the form of a signed timestamp.
func GetProjectOwnerAuthorisationHeaders(l1signer signers.L1Signer) (timestamp, signature string, err error) {
	timestamp = strconv.FormatInt(time.Now().Unix(), 10)
	signedTimestamp, err := l1signer.SignMessage(timestamp)
	if err != nil {
		return "", "", err
	}
	signature = hexutil.Encode(signedTimestamp)
	return timestamp, signature, nil
}
