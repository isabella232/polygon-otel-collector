package polygonreceiver // import "github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver"

import "go.opentelemetry.io/collector/receiver/scraperhelper"

// Config defines configuration for Collectd receiver.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`

	EthereumJsonRPCEndpoint string `mapstructure:"ethereum_jsonrpc_endpoint"`
	PolygonJsonRPCEndpoint  string `mapstructure:"polygon_jsonrpc_endpoint"`
	EtherscanAPIKey         string `mapstructure:"etherscan_api_key"`
	PolygonscanAPIKey       string `mapstructure:"polygonscan_api_key"`
	Chain                   string `mapstructure:"chain"`
	DatadogAPIKey           string `mapstructure:"datadog_api_key"`
	TendermintEndpoint      string `mapstructure:"tendermint_endpoint"`
	HeimdalEndpoint         string `mapstructure:"heimdal_endpoint"`
	SentinelEndpoint        string `mapstructure:"sentinel_endpoint"`
	RootChainProxyContract  string `mapstructure:"root_chain_proxy_contract"`
	StateSenderContract     string `mapstructure:"state_sender_contract"`
	StateReceiverContract   string `mapstructure:"state_receiver_contract"`
}
