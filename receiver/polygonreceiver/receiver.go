package polygonreceiver

import (
	"context"
	"time"

	"github.com/umbracle/ethgo/jsonrpc"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
	conventions "go.opentelemetry.io/collector/model/semconv/v1.5.0"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
)

const (
	metricPrefix = "polygon."
)

// polygonReceiver implements the component.MetricsReceiver for Ethereum protocol.
type polygonReceiver struct {
	config   *Config
	settings component.ReceiverCreateSettings
	client   *jsonrpc.Client
}

// newPolygonReceiver creates the Polygon receiver with the given parameters.
func newPolygonReceiver(
	_ context.Context,
	set component.ReceiverCreateSettings,
	config *Config,
	nextConsumer consumer.Metrics,
) (component.MetricsReceiver, error) {
	recv := polygonReceiver{
		config:   config,
		settings: set,
	}

	scrp, err := scraperhelper.NewScraper(typeStr, recv.scrape, scraperhelper.WithStart(recv.start))
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(&recv.config.ScraperControllerSettings, set, nextConsumer, scraperhelper.AddScraper(scrp))
}

func (r *polygonReceiver) start(ctx context.Context, _ component.Host) error {
	client, err := jsonrpc.NewClient("https://polygon-mainnet.g.alchemy.com/v2/vvAT2dHbNSp3KYrdG_curFSVgHGl3Jd5")
	if err != nil {
		panic(err)
	}
	r.client = client

	return nil
}

func (r *polygonReceiver) scrape(ctx context.Context) (pdata.Metrics, error) {
	now := pdata.NewTimestampFromTime(time.Now())
	md := pdata.NewMetrics()
	rs := md.ResourceMetrics().AppendEmpty()
	rs.SetSchemaUrl(conventions.SchemaURL)
	ilm := rs.InstrumentationLibraryMetrics().AppendEmpty()

	number, err := r.client.Eth().BlockNumber()
	if err != nil {
		panic(err)
	}

	dest := ilm.Metrics().AppendEmpty()
	dest.SetName(metricPrefix + "eth_block_number")
	dest.SetDescription("The current block number.")
	dest.SetDataType(pdata.MetricDataTypeSum)
	sum := dest.Sum()
	sum.SetIsMonotonic(true)
	sum.SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)

	dp := sum.DataPoints().AppendEmpty()
	dp.SetIntVal(int64(number))
	dp.SetTimestamp(now)

	return md, nil
}
