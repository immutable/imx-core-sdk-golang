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

	cfg, err := config.NewDevConfig()
	if err != nil || cfg == nil {
		fmt.Errorf("failed to get the config. error: %v", err)
	}

	httpClient := client.NewHTTPClientWithConfig(nil, config.NewTransportConfig(cfg.CoreApiEndpoint))
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
