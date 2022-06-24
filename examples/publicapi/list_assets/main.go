package main

import (
	"fmt"
	"immutable.com/imx-core-sdk-golang/generated/client"
	"immutable.com/imx-core-sdk-golang/generated/client/assets"
)

func main() {
	listAssetsResponse, err := client.Default.Assets.ListAssets(assets.NewListAssetsParams())
	if err != nil {
		fmt.Errorf("error: %v", err.Error())
	}
	for _, asset := range listAssetsResponse.Payload.Result {
		fmt.Printf("%+v\n", asset)
	}
}
