package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"immutable.com/imx-core-sdk-golang/api"
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
	apiClient := api.NewAPIClient(configuration)
	listAssetsResponse, r, err := apiClient.AssetsApi.ListAssets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AssetsApi.ListAssets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAssets`: ListAssetsResponse
	fmt.Fprintf(os.Stdout, "Response from `AssetsApi.ListAssets`: %v\n", listAssetsResponse)

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
