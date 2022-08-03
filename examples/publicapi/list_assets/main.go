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
	apiClient := api.NewAPIClient(configuration)

	// Using context value to switch/specify the server before sending request. If nothig is specified default server will be used which will be first one in the open api spec list.
	ctx := context.WithValue(context.Background(), api.ContextServerIndex, config.Ropsten)
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
