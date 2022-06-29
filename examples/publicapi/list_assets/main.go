package main

import (
	"fmt"

	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/api/client/assets"
	"immutable.com/imx-core-sdk-golang/config"
)

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
			for _, asset := range (*listAssetsResponse.Payload).Result {
				if asset != nil {
					fmt.Printf("%+v\n", *asset)
				}
			}
		}
	}
}
