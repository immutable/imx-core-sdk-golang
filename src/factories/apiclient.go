package factories

import (
	"immutable.com/imx-core-sdk-golang/api/client"
	"immutable.com/imx-core-sdk-golang/config"
)

func NewAPIClient(env *config.Config) *client.ImmutableXAPI {
	return client.NewHTTPClientWithConfig(nil, config.NewTransportConfig(&env.CoreAPIEndpoint))
}
