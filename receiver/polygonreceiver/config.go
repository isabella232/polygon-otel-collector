package polygonreceiver // import "github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver"

import "go.opentelemetry.io/collector/receiver/scraperhelper"

// Config defines configuration for Collectd receiver.
type Config struct {
	scraperhelper.ScraperControllerSettings `mapstructure:",squash"`

	// Lock down the encoding of the payload, leave empty for attribute based detection
	Encoding string `mapstructure:"encoding"`
}
