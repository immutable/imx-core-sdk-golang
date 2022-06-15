package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"immutable.com/imx-core-sdk/api"
	"immutable.com/imx-core-sdk/servers"
	"testing"
)

func TestListAssets(t *testing.T) {
	cfg := api.NewConfiguration()
	client := api.NewAPIClient(cfg)
	ctx := context.WithValue(context.Background(), api.ContextServerIndex, servers.Dev)
	assets, resp, err := client.AssetsApi.ListAssets(ctx).Execute()
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.NotEmpty(t, assets)
	//fmt.Fprintf(os.Stdout, "Response from `AssetsApi.ListAssets`: %v\n", assets)
}
