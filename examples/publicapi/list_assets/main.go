package main

import (
	"encoding/json"
	"fmt"
	"log"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/assets"
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

	apiUrl := config.GetApiUrl(config.Dev)

	httpClient := client.NewHTTPClientWithConfig(nil, config.NewTransportConfig(apiUrl))
	listAssetsResponse, err := httpClient.Assets.ListAssets(assets.NewListAssetsParams())
	if err != nil {
		fmt.Errorf("error: %v", err.Error())
	}

	if listAssetsResponse != nil {
		if listAssetsResponse.Payload != nil {
			res, err := PrettyStruct((*listAssetsResponse.Payload).Result)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res)
		}
	}
}
