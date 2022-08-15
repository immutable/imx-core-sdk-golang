package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"immutable.com/imx-core-sdk-golang/api"
	"immutable.com/imx-core-sdk-golang/config"
)

const (
	AlchemyAPIKey = "Provide your alchemy API Key"
)

func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}

func main() {

	configuration := api.NewConfiguration()
	configuration.Debug = true // Enables debug logging.

	// Multiple ways to configure the server
	// 1. Set default host on config
	// 2. Per request config via context.

	/*
		// Setup default host
		configuration.Host = "api.x.immutable.com"
		configuration.Scheme = "https"
		ctx := context.Background()
	*/

	// or

	/*
		// Alternatively add to configuration servers then it can be selectable with context value as below.
		mainnetConfig := config.GetConfig(config.MainNet, AlchemyAPIKey)
		configuration.Servers = append(configuration.Servers, api.ServerConfiguration{
			URL:         mainnetConfig.CoreAPIEndpoint,
			Description: "Production Server",
		})
	*/

	// Using context value to switch/specify the server before sending request. If nothing is specified, the default server will be used which will be first one in the open api spec list.
	ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Ropsten)

	apiClient := api.NewAPIClient(configuration)
	listAssetsResponse, r, err := apiClient.AssetsApi.ListAssets(ctx).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.ListAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssets`: ListAssetsResponse
	if listAssetsResponse != nil {
		if listAssetsResponse.Result != nil {
			res, err := PrettyStruct(listAssetsResponse.Result)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println(res)
		}
	}
}
