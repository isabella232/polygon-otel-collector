package polygonreceiver // import "github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver"

import (
	"context"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	// The value of "type" key in configuration.
	typeStr = "polygon"
)

// NewFactory creates a factory for Elastic exporter.
func NewFactory() component.ReceiverFactory {
	return receiverhelper.NewFactory(
		typeStr,
		createDefaultConfig,
		receiverhelper.WithMetrics(createMetricsReceiver))
}

func createDefaultConfig() config.Receiver {
	scs := scraperhelper.DefaultScraperControllerSettings(typeStr)
	scs.CollectionInterval = 10 * time.Second
	return &Config{
		ScraperControllerSettings: scs,
	}
}

func createMetricsReceiver(
	ctx context.Context,
	params component.ReceiverCreateSettings,
	cfg config.Receiver,
	nextConsumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	config := cfg.(*Config)

	return newPolygonReceiver(ctx, params, config, nextConsumer)
}
