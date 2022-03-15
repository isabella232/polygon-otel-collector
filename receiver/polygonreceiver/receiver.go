package polygonreceiver

import (
	"context"
	"time"

	"github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver/internal/metadata"
	"github.com/nanmu42/etherscan-api"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/jsonrpc"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"
)

// polygonReceiver implements the component.MetricsReceiver for Ethereum protocol.
type polygonReceiver struct {
	config            *Config
	settings          component.ReceiverCreateSettings
	client            *jsonrpc.Client
	polygonscanClient *etherscan.Client
	etherscanClient   *etherscan.Client
	logger            *zap.SugaredLogger
	mb                *metadata.MetricsBuilder
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
		logger:   zap.L().Sugar(),
		mb:       metadata.NewMetricsBuilder(metadata.DefaultMetricsSettings()),
	}

	scrp, err := scraperhelper.NewScraper(typeStr, recv.scrape, scraperhelper.WithStart(recv.start))
	if err != nil {
		return nil, err
	}
	return scraperhelper.NewScraperControllerReceiver(&recv.config.ScraperControllerSettings, set, nextConsumer, scraperhelper.AddScraper(scrp))
}

func (r *polygonReceiver) start(ctx context.Context, _ component.Host) error {
	client, err := jsonrpc.NewClient(r.config.JsonRPCEndpoint)
	if err != nil {
		panic(err)
	}
	r.client = client
	r.polygonscanClient = etherscan.NewCustomized(etherscan.Customization{
		Key:     r.config.PolygonscanAPIKey,
		BaseURL: "https://api.polygonscan.com/api?",
	})
	r.etherscanClient = etherscan.New(etherscan.Mainnet, r.config.EtherscanAPIKey)

	return nil
}

func (r *polygonReceiver) scrape(ctx context.Context) (pdata.Metrics, error) {
	md := pdata.NewMetrics()
	ilm := md.ResourceMetrics().AppendEmpty().InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/polygon")
	now := pdata.NewTimestampFromTime(time.Now())

	number, err := r.client.Eth().BlockNumber()
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
	}
	block, err := r.client.Eth().GetBlockByNumber(ethgo.BlockNumber(number), true)
	if err != nil {
		r.logger.Error("failed to get block", zap.Error(err))
	}
	prevBlock, err := r.client.Eth().GetBlockByNumber(ethgo.BlockNumber(number-1), true)
	if err != nil {
		r.logger.Error("failed to get previous block", zap.Error(err))
	}
	if block != nil && prevBlock != nil {
		bd := time.Unix(int64(block.Timestamp), 0).Sub(time.Unix(int64(prevBlock.Timestamp), 0))
		r.mb.RecordPolygonLastBlockTimeDataPoint(now, bd.Milliseconds(), "polygon-mainnet")
		r.mb.RecordPolygonLastBlockDataPoint(pdata.Timestamp(block.Timestamp), int64(number), "polygon-mainnet")
	}

	// Get latest checkpoint transaction
	bn, err := r.etherscanClient.BlockNumber(now.AsTime().Unix(), "before")
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
	}

	sb := bn - 1000
	txl, err := r.etherscanClient.NormalTxByAddress("0x86E4Dc95c7FBdBf52e33D563BbDB00823894C287", &sb, &bn, 1, 0, true)
	if err != nil {
		r.logger.Error("failed to get transaction", zap.Error(err))
	}

	for _, tx := range txl {
		r.mb.RecordPolygonSubmitCheckpointDataPoint(now, int64(tx.BlockNumber), "ethereum-mainnet")
	}

	r.mb.Emit(ilm.Metrics())

	return md, nil
}
