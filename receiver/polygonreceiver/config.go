package polygonreceiver // import "github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver"

import "go.opentelemetry.io/collector/receiver/scraperhelper"

// Config defines configuration for Collectd receiver.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`

	// Lock down the encoding of the payload, leave empty for attribute based detection
	EthereumJsonRPCEndpoint string `mapstructure:"ethereum_jsonrpc_endpoint"`
	PolygonJsonRPCEndpoint  string `mapstructure:"polygon_jsonrpc_endpoint"`
	EtherscanAPIKey         string `mapstructure:"etherscan_api_key"`
	PolygonscanAPIKey       string `mapstructure:"polygonscan_api_key"`
	Chain                   string `mapstructure:"chain"`
	DatadogAPIKey           string `mapstructure:"datadog_api_key"`
}
