package factories

import (
	"immutable.com/imx-core-sdk/api"
	"immutable.com/imx-core-sdk/config"
)

func NewApiClient(env config.EnvConfig) *api.APIClient {
	apiConfiguration := newConfiguration(env.CoreApiEndpoint)
	apiConfiguration.AddDefaultHeader("x-sdk-version", "imx-core-sdk-golang")
	return api.NewAPIClient(apiConfiguration)
}

func NewApiClientWithHeaders(env config.EnvConfig, headers map[string]string) *api.APIClient {
	apiConfiguration := newConfiguration(env.CoreApiEndpoint)
	apiConfiguration.AddDefaultHeader("x-sdk-version", "imx-core-sdk-golang")
	for key, val := range headers {
		apiConfiguration.AddDefaultHeader(key, val)
	}
	return api.NewAPIClient(apiConfiguration)
}

func newConfiguration(serverUrl string) *api.Configuration {
	cfg := &api.Configuration{
		DefaultHeader: make(map[string]string),
		UserAgent:     "OpenAPI-Generator/1.0.0/go",
		Debug:         false,
		Servers: api.ServerConfigurations{
			{
				URL: serverUrl,
			},
		},
		OperationServers: map[string]api.ServerConfigurations{},
	}
	return cfg
}
