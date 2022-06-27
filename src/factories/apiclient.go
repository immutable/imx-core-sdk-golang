package factories

import (
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/config"
)

func NewApiClient(env config.Config) *client.ImmutableXAPI {
	return client.NewHTTPClientWithConfig(nil, config.NewTransportConfig(env.CoreApiEndpoint))
}
