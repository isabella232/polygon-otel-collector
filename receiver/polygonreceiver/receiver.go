package polygonreceiver

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/maticnetwork/polygon-otel-collector/receiver/polygonreceiver/internal/metadata"
	"github.com/nanmu42/etherscan-api"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
	"github.com/umbracle/ethgo/contract"
	"github.com/umbracle/ethgo/jsonrpc"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/model/pdata"
	"go.opentelemetry.io/collector/receiver/scraperhelper"
	"go.uber.org/zap"
)

// polygonReceiver implements the component.MetricsReceiver for Ethereum protocol.

var (
	prev         pdata.Timestamp
	prevBorBlock uint64
)

type polygonReceiver struct {
	config            *Config
	settings          component.ReceiverCreateSettings
	polygonClient     *jsonrpc.Client
	ethClient         *jsonrpc.Client
	polygonscanClient *etherscan.Client
	etherscanClient   *etherscan.Client
	logger            *zap.SugaredLogger
	mb                *metadata.MetricsBuilder
	ddAPIClient       *datadog.APIClient
	heimdallClient    *HeimdallClient
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
		logger:   set.Logger.Sugar(),
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

	configuration := datadog.NewConfiguration()
	r.ddAPIClient = datadog.NewAPIClient(configuration)

	c := NewHeimdallClient("https://tendermint.api.matic.network", "https://heimdall-api.polygon.technology")
	r.heimdallClient = c

	return nil
}

func (r *polygonReceiver) scrape(ctx context.Context) (pdata.Metrics, error) {
	md := pdata.NewMetrics()
	ilm := md.ResourceMetrics().AppendEmpty().InstrumentationLibraryMetrics().AppendEmpty()
	ilm.InstrumentationLibrary().SetName("otelcol/polygon")

	now := pdata.NewTimestampFromTime(time.Now())
	if prev == 0 {
		prev = now
	}

	r.recordBorBlockMetrics(now, prev)
	r.recordHeimdallBlockMetrics(now, prev)
	r.recordCheckpointMetrics(now)
	r.recordHeimdallUnconfirmedTransactions(now)
	r.recordRootChainStateSyncs(now)
	r.recordSideChainStateSyncs(now)
	r.recordHeimdallEndBlock(now)

	r.mb.Emit(ilm.Metrics())

	prev = now

	return md, nil
}

func (r *polygonReceiver) recordBorBlockMetrics(now pdata.Timestamp, prev pdata.Timestamp) {
	number, err := r.polygonClient.Eth().BlockNumber()
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
		return
	}

	if prev == 0 {
		prevBorBlock = number
		return
	}

	bd := number - prevBorBlock
	td := now.AsTime().Sub(prev.AsTime())

	s := td.Seconds()
	bt := s / float64(bd)
	r.mb.RecordPolygonBorAverageBlockTimeDataPoint(now, bt, "polygon-"+r.config.Chain)
	r.mb.RecordPolygonBorLastBlockDataPoint(now, int64(number), "polygon"+r.config.Chain)
	prevBorBlock = number
}

func (r *polygonReceiver) recordHeimdallBlockMetrics(now pdata.Timestamp, prev pdata.Timestamp) {
	// Get checkpoint signatures
	block, err := r.heimdallClient.Block(nil)
	if err != nil {
		r.logger.Error("failed to get block", zap.Error(err))
		return
	}

	lb, err := strconv.ParseInt(block.Result.Block.Header.Height, 10, 64)
	if err != nil {
		r.logger.Error("failed to parse block height", zap.Error(err))
		return
	}
	pbs := strconv.FormatInt(lb-1, 10)
	prevBlock, err := r.heimdallClient.Block(&pbs)
	if err != nil {
		r.logger.Error("failed to get block", zap.Error(err))
		return
	}

	bt1, err := time.Parse(time.RFC3339Nano, block.Result.Block.Header.Time)
	if err != nil {
		r.logger.Error("failed to parse time", zap.Error(err))
		return
	}
	bt2, err := time.Parse(time.RFC3339Nano, prevBlock.Result.Block.Header.Time)
	if err != nil {
		r.logger.Error("failed to parse time", zap.Error(err))
		return
	}
	bt := bt1.Sub(bt2)

	r.mb.RecordPolygonHeimdallAverageBlockTimeDataPoint(now, bt.Seconds(), "polygon-"+r.config.Chain)
	r.mb.RecordPolygonHeimdallLastBlockDataPoint(now, lb, "polygon-"+r.config.Chain)
}

func (r *polygonReceiver) recordCheckpointMetrics(now pdata.Timestamp) {
	// Get logs for the latest checkpoint transaction
	bn, err := r.ethClient.Eth().BlockNumber()
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
		return
	}

	checkpointEventSig := abi.MustNewEvent("event NewHeaderBlock(address indexed proposer, uint256 indexed headerBlockId, uint256 indexed reward, uint256 start, uint256 end, bytes32 root)")

	bnp := ethgo.BlockNumber(bn - 10000)
	lbp := ethgo.BlockNumber(bn)
	h := checkpointEventSig.ID()
	topics := []*ethgo.Hash{&h}
	logs, err := r.ethClient.Eth().GetLogs(&ethgo.LogFilter{
		From:    &bnp,
		To:      &lbp,
		Address: []ethgo.Address{ethgo.HexToAddress("0x86E4Dc95c7FBdBf52e33D563BbDB00823894C287")},
		Topics:  [][]*ethgo.Hash{topics},
	})
	if err == nil && len(logs) > 0 {
		// Exit if there are no logs.
		log := logs[len(logs)-1]

		event, err := checkpointEventSig.ParseLog(log)
		if err != nil {
			r.logger.Error("failed to parse log", zap.Error(err))
			return
		}

		hbi := event["headerBlockId"].(*big.Int)
		hbiTrim := strings.TrimRight(fmt.Sprintf("%v", hbi), "0000")

		////
		b, err := r.ethClient.Eth().GetBlockByNumber(ethgo.BlockNumber(log.BlockNumber), true)
		if err != nil {
			r.logger.Error("failed to get block", zap.Error(err))
			return
		}
		txd := now.AsTime().Sub(time.Unix(int64(b.Timestamp), 0))
		r.mb.RecordPolygonEthSubmitCheckpointTimeDataPoint(now, txd.Seconds(), "ethereum-"+r.config.Chain)
		////

		// Get checkpoint signatures
		res, err := http.Get(fmt.Sprintf("https://sentinel.matic.network/api/v2/monitor/checkpoint-signatures/checkpoint/%s", hbiTrim))
		if err != nil {
			r.logger.Error("failed to get checkpoint signatures", zap.Error(err))
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			r.logger.Error("failed to read checkpoint signatures", zap.Error(err))
			return
		}

		signatures := &CheckpointSignatures{}
		err = json.Unmarshal(body, &signatures)
		if err != nil {
			r.logger.Error("failed to unmarshal checkpoint signatures", zap.Error(err))
			return
		}

		for _, signature := range signatures.Result {
			if signature.HasSigned {
				r.mb.RecordPolygonHeimdallCheckpointValidatorsSignedDataPoint(now, 1, "polygon-"+r.config.Chain, signature.SignerAddress)
			} else {
				r.mb.RecordPolygonHeimdallCheckpointValidatorsSignedDataPoint(now, 0, "polygon-"+r.config.Chain, signature.SignerAddress)
			}
		}
	} else {
		r.logger.Error("failed to get logs", zap.Error(err))
	}
}

func (r *polygonReceiver) checkpointSignaturesDatadogEvent(signedCount int64) error {
	body := datadog.EventCreateRequest{
		Title: "The number of validators who have signed the checkpoint is less than 90",
		Text:  "Check the validators.",
		Tags: &[]string{
			"test:ExamplePostaneventreturnsOKresponse",
		},
	}
	ctx := context.WithValue(context.Background(), datadog.ContextAPIKeys, map[string]datadog.APIKey{
		"apiKeyAuth": {
			Key: r.config.DatadogAPIKey,
		},
	})

	_, _, err := r.ddAPIClient.EventsApi.CreateEvent(ctx, body)
	if err != nil {
		return err
	}

	return nil
}

func (r *polygonReceiver) recordHeimdallUnconfirmedTransactions(now pdata.Timestamp) {
	// Get heimdall unconfirmed transactions
	res, err := http.Get("https://tendermint.api.matic.network/num_unconfirmed_txs")
	if err != nil {
		r.logger.Error("failed to get unconfirmed transactions", zap.Error(err))
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		r.logger.Error("failed to read unconfirmed transactions", zap.Error(err))
		return
	}

	transactions := &HeimdallUnconfirmedTransactions{}
	err = json.Unmarshal(body, &transactions)
	if err != nil {
		r.logger.Error("failed to unmarshal unconfirmed transactions", zap.Error(err))
		return
	}

	utxs, err := strconv.ParseInt(transactions.Result.Ntxs, 10, 64)
	if err != nil {
		r.logger.Error("failed to parse unconfirmed transactions", zap.Error(err))
		return
	}
	ttxs, err := strconv.ParseInt(transactions.Result.Total, 10, 64)
	if err != nil {
		r.logger.Error("failed to parse total transactions", zap.Error(err))
		return
	}

	r.mb.RecordPolygonHeimdallUnconfirmedTxsDataPoint(now, utxs, "polygon-"+r.config.Chain)
	r.mb.RecordPolygonHeimdallTotalTxsDataPoint(now, ttxs, "polygon-"+r.config.Chain)
}

func (r *polygonReceiver) recordRootChainStateSyncs(now pdata.Timestamp) {
	// Get logs for the latest checkpoint transaction
	bn, err := r.ethClient.Eth().BlockNumber()
	if err != nil {
		r.logger.Error("failed to get block number", zap.Error(err))
		return
	}
	stateSyncedEventSig := abi.MustNewEvent("event StateSynced(uint256 indexed id, address indexed contractAddress, bytes data)")

	bnp := ethgo.BlockNumber(bn - 1000)
	lbp := ethgo.BlockNumber(bn)
	logs, err := r.ethClient.Eth().GetLogs(&ethgo.LogFilter{
		From:    &bnp,
		To:      &lbp,
		Address: []ethgo.Address{ethgo.HexToAddress("0x28e4F3a7f651294B9564800b2D01f35189A5bFbE")},
	})

	if err == nil && len(logs) > 0 {
		log := logs[len(logs)-1]

		event, err := stateSyncedEventSig.ParseLog(log)
		if err != nil {
			r.logger.Error("failed to parse log", zap.Error(err))
			return
		}
		id := event["id"].(*big.Int)
		r.mb.RecordPolygonEthStateSyncDataPoint(now, id.Int64(), "ethereum-"+r.config.Chain)
	}
}

func (r *polygonReceiver) recordSideChainStateSyncs(now pdata.Timestamp) {
	var functions = []string{
		"function lastStateId() public view returns (uint256)",
	}

	abiContract, err := abi.NewABIFromList(functions)
	if err != nil {
		r.logger.Error("failed to create abi", zap.Error(err))
		return
	}

	// Matic token
	addr := ethgo.HexToAddress("0x0000000000000000000000000000000000001001")

	c := contract.NewContract(addr, abiContract, r.polygonClient)
	res, err := c.Call("lastStateId", ethgo.Latest)
	if err != nil {
		r.logger.Error("failed to get last state id", zap.Error(err))
		return
	}

	id, ok := res["0"].(*big.Int)
	if !ok {
		r.logger.Error("failed to parse last state id", zap.Error(err))
		return
	}
	r.mb.RecordPolygonPolygonStateSyncDataPoint(now, id.Int64(), "ethereum-"+r.config.Chain)
}

func (r *polygonReceiver) recordHeimdallEndBlock(now pdata.Timestamp) {
	// Get heimdall end block
	ls, err := r.heimdallClient.LatestSpan()
	if err != nil {
		r.logger.Error("failed to get latest span", zap.Error(err))
		return
	}

	r.mb.RecordPolygonHeimdallCurrentSpanEndBlockDataPoint(now, ls.Result.EndBlock, r.config.Chain)
}
