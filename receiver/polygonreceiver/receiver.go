package polygonreceiver

import (
	"context"
	"net/http"
	"sort"
	"time"

	"github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver/internal/metadata"
	"github.com/nanmu42/etherscan-api"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
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
	polygonClient     *jsonrpc.Client
	ethClient         *jsonrpc.Client
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
	ec, err := jsonrpc.NewClient(r.config.EthereumJsonRPCEndpoint)
	if err != nil {
		panic(err)
	}
	r.ethClient = ec

	pc, err := jsonrpc.NewClient(r.config.PolygonJsonRPCEndpoint)
	if err != nil {
		panic(err)
	}
	r.polygonClient = pc

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

	number, err := r.polygonClient.Eth().BlockNumber()
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
	}
	block, err := r.polygonClient.Eth().GetBlockByNumber(ethgo.BlockNumber(number), true)
	if err != nil {
		r.logger.Error("failed to get block", zap.Error(err))
	}
	prevBlock, err := r.polygonClient.Eth().GetBlockByNumber(ethgo.BlockNumber(number-1), true)
	if err != nil {
		r.logger.Error("failed to get previous block", zap.Error(err))
	}

	if block != nil && prevBlock != nil {
		bd := time.Unix(int64(block.Timestamp), 0).Sub(time.Unix(int64(prevBlock.Timestamp), 0))
		r.mb.RecordPolygonLastBlockTimeDataPoint(now, bd.Milliseconds(), "polygon-mainnet")
	}
	r.mb.RecordPolygonLastBlockDataPoint(now, int64(number), "polygon-mainnet")

	// Get latest checkpoint transaction
	bn, err := r.ethClient.Eth().BlockNumber()
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
	}

	txd := now.AsTime().Sub(time.Time(checkpoint.TimeStamp))
	r.mb.RecordPolygonSubmitCheckpointTimeDataPoint(now, txd.Seconds(), "ethereum-mainnet")

	checkpointEventSig := abi.MustNewEvent("event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)")

	bnp := ethgo.BlockNumber(bn - 1000)
	lbp := ethgo.Latest

	logs, err := r.ethClient.Eth().GetLogs(&ethgo.LogFilter{
		From:    &bnp,
		To:      &lbp,
		Address: []ethgo.Address{ethgo.BytesToAddress([]byte("0x86E4Dc95c7FBdBf52e33D563BbDB00823894C287"))},
		Topics:  [][]*ethgo.Hash{checkpointEventSig.ID()},
	})
	if err != nil {
		r.logger.Error("failed to get logs", zap.Error(err))
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(logs, func(i, j int) bool {
		return l[i].BlockNumber > l[j].BlockNumber
	})

	for _, l := range logs {

	}

	// Get checkpoint signatures
	_, err = http.Get("https://sentinel.matic.network/api/v2/monitor/checkpoint-signatures/checkpoint/")
	if err != nil {
		r.logger.Error("failed to get checkpoint signatures", zap.Error(err))
	}

	r.mb.Emit(ilm.Metrics())

	return md, nil
}
