// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"time"

	"go.opentelemetry.io/collector/model/pdata"
)

// MetricSettings provides common settings for a particular metric.
type MetricSettings struct {
	Enabled bool `mapstructure:"enabled"`
}

// MetricsSettings provides settings for polygonreceiver metrics.
type MetricsSettings struct {
	PolygonBorAverageBlockTime                MetricSettings `mapstructure:"polygon.bor.average_block_time"`
	PolygonBorLastBlock                       MetricSettings `mapstructure:"polygon.bor.last_block"`
	PolygonEthStateSync                       MetricSettings `mapstructure:"polygon.eth.state_sync"`
	PolygonEthSubmitCheckpointTime            MetricSettings `mapstructure:"polygon.eth.submit_checkpoint_time"`
	PolygonHeimdallAverageBlockTime           MetricSettings `mapstructure:"polygon.heimdall.average_block_time"`
	PolygonHeimdallCheckpointValidatorsSigned MetricSettings `mapstructure:"polygon.heimdall.checkpoint_validators_signed"`
	PolygonHeimdallCurrentSpanEndBlock        MetricSettings `mapstructure:"polygon.heimdall.current_span_end_block"`
	PolygonHeimdallLastBlock                  MetricSettings `mapstructure:"polygon.heimdall.last_block"`
	PolygonHeimdallTotalTxs                   MetricSettings `mapstructure:"polygon.heimdall.total_txs"`
	PolygonHeimdallUnconfirmedTxs             MetricSettings `mapstructure:"polygon.heimdall.unconfirmed_txs"`
	PolygonPolygonStateSync                   MetricSettings `mapstructure:"polygon.polygon.state_sync"`
}

func DefaultMetricsSettings() MetricsSettings {
	return MetricsSettings{
		PolygonBorAverageBlockTime: MetricSettings{
			Enabled: true,
		},
		PolygonBorLastBlock: MetricSettings{
			Enabled: true,
		},
		PolygonEthStateSync: MetricSettings{
			Enabled: true,
		},
		PolygonEthSubmitCheckpointTime: MetricSettings{
			Enabled: true,
		},
		PolygonHeimdallAverageBlockTime: MetricSettings{
			Enabled: true,
		},
		PolygonHeimdallCheckpointValidatorsSigned: MetricSettings{
			Enabled: true,
		},
		PolygonHeimdallCurrentSpanEndBlock: MetricSettings{
			Enabled: true,
		},
		PolygonHeimdallLastBlock: MetricSettings{
			Enabled: true,
		},
		PolygonHeimdallTotalTxs: MetricSettings{
			Enabled: true,
		},
		PolygonHeimdallUnconfirmedTxs: MetricSettings{
			Enabled: true,
		},
		PolygonPolygonStateSync: MetricSettings{
			Enabled: true,
		},
	}
}

type metricPolygonBorAverageBlockTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.bor.average_block_time metric with initial data.
func (m *metricPolygonBorAverageBlockTime) init() {
	m.data.SetName("polygon.bor.average_block_time")
	m.data.SetDescription("The average block time.")
	m.data.SetUnit("seconds")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonBorAverageBlockTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonBorAverageBlockTime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonBorAverageBlockTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonBorAverageBlockTime(settings MetricSettings) metricPolygonBorAverageBlockTime {
	m := metricPolygonBorAverageBlockTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonBorLastBlock struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.bor.last_block metric with initial data.
func (m *metricPolygonBorLastBlock) init() {
	m.data.SetName("polygon.bor.last_block")
	m.data.SetDescription("The current block number.")
	m.data.SetUnit("block")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonBorLastBlock) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonBorLastBlock) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonBorLastBlock) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonBorLastBlock(settings MetricSettings) metricPolygonBorLastBlock {
	m := metricPolygonBorLastBlock{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonEthStateSync struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.eth.state_sync metric with initial data.
func (m *metricPolygonEthStateSync) init() {
	m.data.SetName("polygon.eth.state_sync")
	m.data.SetDescription("Total number of StateSync transactions emited.")
	m.data.SetUnit("txs")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonEthStateSync) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonEthStateSync) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonEthStateSync) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonEthStateSync(settings MetricSettings) metricPolygonEthStateSync {
	m := metricPolygonEthStateSync{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonEthSubmitCheckpointTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.eth.submit_checkpoint_time metric with initial data.
func (m *metricPolygonEthSubmitCheckpointTime) init() {
	m.data.SetName("polygon.eth.submit_checkpoint_time")
	m.data.SetDescription("Latest checkpoint transaction time.")
	m.data.SetUnit("seconds")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonEthSubmitCheckpointTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonEthSubmitCheckpointTime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonEthSubmitCheckpointTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonEthSubmitCheckpointTime(settings MetricSettings) metricPolygonEthSubmitCheckpointTime {
	m := metricPolygonEthSubmitCheckpointTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonHeimdallAverageBlockTime struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.heimdall.average_block_time metric with initial data.
func (m *metricPolygonHeimdallAverageBlockTime) init() {
	m.data.SetName("polygon.heimdall.average_block_time")
	m.data.SetDescription("The average block time.")
	m.data.SetUnit("seconds")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonHeimdallAverageBlockTime) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val float64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetDoubleVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonHeimdallAverageBlockTime) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonHeimdallAverageBlockTime) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonHeimdallAverageBlockTime(settings MetricSettings) metricPolygonHeimdallAverageBlockTime {
	m := metricPolygonHeimdallAverageBlockTime{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonHeimdallCheckpointValidatorsSigned struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.heimdall.checkpoint_validators_signed metric with initial data.
func (m *metricPolygonHeimdallCheckpointValidatorsSigned) init() {
	m.data.SetName("polygon.heimdall.checkpoint_validators_signed")
	m.data.SetDescription("Number of validators who signed last checkpoint.")
	m.data.SetUnit("")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonHeimdallCheckpointValidatorsSigned) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonHeimdallCheckpointValidatorsSigned) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonHeimdallCheckpointValidatorsSigned) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonHeimdallCheckpointValidatorsSigned(settings MetricSettings) metricPolygonHeimdallCheckpointValidatorsSigned {
	m := metricPolygonHeimdallCheckpointValidatorsSigned{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonHeimdallCurrentSpanEndBlock struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.heimdall.current_span_end_block metric with initial data.
func (m *metricPolygonHeimdallCurrentSpanEndBlock) init() {
	m.data.SetName("polygon.heimdall.current_span_end_block")
	m.data.SetDescription("The end block of the current span.")
	m.data.SetUnit("block")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonHeimdallCurrentSpanEndBlock) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonHeimdallCurrentSpanEndBlock) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonHeimdallCurrentSpanEndBlock) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonHeimdallCurrentSpanEndBlock(settings MetricSettings) metricPolygonHeimdallCurrentSpanEndBlock {
	m := metricPolygonHeimdallCurrentSpanEndBlock{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonHeimdallLastBlock struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.heimdall.last_block metric with initial data.
func (m *metricPolygonHeimdallLastBlock) init() {
	m.data.SetName("polygon.heimdall.last_block")
	m.data.SetDescription("The current block number.")
	m.data.SetUnit("block")
	m.data.SetDataType(pdata.MetricDataTypeSum)
	m.data.Sum().SetIsMonotonic(false)
	m.data.Sum().SetAggregationTemporality(pdata.MetricAggregationTemporalityCumulative)
	m.data.Sum().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonHeimdallLastBlock) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Sum().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonHeimdallLastBlock) updateCapacity() {
	if m.data.Sum().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Sum().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonHeimdallLastBlock) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Sum().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonHeimdallLastBlock(settings MetricSettings) metricPolygonHeimdallLastBlock {
	m := metricPolygonHeimdallLastBlock{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonHeimdallTotalTxs struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.heimdall.total_txs metric with initial data.
func (m *metricPolygonHeimdallTotalTxs) init() {
	m.data.SetName("polygon.heimdall.total_txs")
	m.data.SetDescription("Total number of transactions.")
	m.data.SetUnit("txs")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonHeimdallTotalTxs) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonHeimdallTotalTxs) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonHeimdallTotalTxs) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonHeimdallTotalTxs(settings MetricSettings) metricPolygonHeimdallTotalTxs {
	m := metricPolygonHeimdallTotalTxs{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonHeimdallUnconfirmedTxs struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.heimdall.unconfirmed_txs metric with initial data.
func (m *metricPolygonHeimdallUnconfirmedTxs) init() {
	m.data.SetName("polygon.heimdall.unconfirmed_txs")
	m.data.SetDescription("Number of unconfirmed transactions.")
	m.data.SetUnit("txs")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonHeimdallUnconfirmedTxs) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonHeimdallUnconfirmedTxs) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonHeimdallUnconfirmedTxs) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonHeimdallUnconfirmedTxs(settings MetricSettings) metricPolygonHeimdallUnconfirmedTxs {
	m := metricPolygonHeimdallUnconfirmedTxs{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

type metricPolygonPolygonStateSync struct {
	data     pdata.Metric   // data buffer for generated metric.
	settings MetricSettings // metric settings provided by user.
	capacity int            // max observed number of data points added to the metric.
}

// init fills polygon.polygon.state_sync metric with initial data.
func (m *metricPolygonPolygonStateSync) init() {
	m.data.SetName("polygon.polygon.state_sync")
	m.data.SetDescription("Total number of StateSync transactions received.")
	m.data.SetUnit("txs")
	m.data.SetDataType(pdata.MetricDataTypeGauge)
	m.data.Gauge().DataPoints().EnsureCapacity(m.capacity)
}

func (m *metricPolygonPolygonStateSync) recordDataPoint(start pdata.Timestamp, ts pdata.Timestamp, val int64, chainAttributeValue string) {
	if !m.settings.Enabled {
		return
	}
	dp := m.data.Gauge().DataPoints().AppendEmpty()
	dp.SetStartTimestamp(start)
	dp.SetTimestamp(ts)
	dp.SetIntVal(val)
	dp.Attributes().Insert(A.Chain, pdata.NewAttributeValueString(chainAttributeValue))
}

// updateCapacity saves max length of data point slices that will be used for the slice capacity.
func (m *metricPolygonPolygonStateSync) updateCapacity() {
	if m.data.Gauge().DataPoints().Len() > m.capacity {
		m.capacity = m.data.Gauge().DataPoints().Len()
	}
}

// emit appends recorded metric data to a metrics slice and prepares it for recording another set of data points.
func (m *metricPolygonPolygonStateSync) emit(metrics pdata.MetricSlice) {
	if m.settings.Enabled && m.data.Gauge().DataPoints().Len() > 0 {
		m.updateCapacity()
		m.data.MoveTo(metrics.AppendEmpty())
		m.init()
	}
}

func newMetricPolygonPolygonStateSync(settings MetricSettings) metricPolygonPolygonStateSync {
	m := metricPolygonPolygonStateSync{settings: settings}
	if settings.Enabled {
		m.data = pdata.NewMetric()
		m.init()
	}
	return m
}

// MetricsBuilder provides an interface for scrapers to report metrics while taking care of all the transformations
// required to produce metric representation defined in metadata and user settings.
type MetricsBuilder struct {
	startTime                                       pdata.Timestamp
	metricPolygonBorAverageBlockTime                metricPolygonBorAverageBlockTime
	metricPolygonBorLastBlock                       metricPolygonBorLastBlock
	metricPolygonEthStateSync                       metricPolygonEthStateSync
	metricPolygonEthSubmitCheckpointTime            metricPolygonEthSubmitCheckpointTime
	metricPolygonHeimdallAverageBlockTime           metricPolygonHeimdallAverageBlockTime
	metricPolygonHeimdallCheckpointValidatorsSigned metricPolygonHeimdallCheckpointValidatorsSigned
	metricPolygonHeimdallCurrentSpanEndBlock        metricPolygonHeimdallCurrentSpanEndBlock
	metricPolygonHeimdallLastBlock                  metricPolygonHeimdallLastBlock
	metricPolygonHeimdallTotalTxs                   metricPolygonHeimdallTotalTxs
	metricPolygonHeimdallUnconfirmedTxs             metricPolygonHeimdallUnconfirmedTxs
	metricPolygonPolygonStateSync                   metricPolygonPolygonStateSync
}

// metricBuilderOption applies changes to default metrics builder.
type metricBuilderOption func(*MetricsBuilder)

// WithStartTime sets startTime on the metrics builder.
func WithStartTime(startTime pdata.Timestamp) metricBuilderOption {
	return func(mb *MetricsBuilder) {
		mb.startTime = startTime
	}
}

func NewMetricsBuilder(settings MetricsSettings, options ...metricBuilderOption) *MetricsBuilder {
	mb := &MetricsBuilder{
		startTime:                                       pdata.NewTimestampFromTime(time.Now()),
		metricPolygonBorAverageBlockTime:                newMetricPolygonBorAverageBlockTime(settings.PolygonBorAverageBlockTime),
		metricPolygonBorLastBlock:                       newMetricPolygonBorLastBlock(settings.PolygonBorLastBlock),
		metricPolygonEthStateSync:                       newMetricPolygonEthStateSync(settings.PolygonEthStateSync),
		metricPolygonEthSubmitCheckpointTime:            newMetricPolygonEthSubmitCheckpointTime(settings.PolygonEthSubmitCheckpointTime),
		metricPolygonHeimdallAverageBlockTime:           newMetricPolygonHeimdallAverageBlockTime(settings.PolygonHeimdallAverageBlockTime),
		metricPolygonHeimdallCheckpointValidatorsSigned: newMetricPolygonHeimdallCheckpointValidatorsSigned(settings.PolygonHeimdallCheckpointValidatorsSigned),
		metricPolygonHeimdallCurrentSpanEndBlock:        newMetricPolygonHeimdallCurrentSpanEndBlock(settings.PolygonHeimdallCurrentSpanEndBlock),
		metricPolygonHeimdallLastBlock:                  newMetricPolygonHeimdallLastBlock(settings.PolygonHeimdallLastBlock),
		metricPolygonHeimdallTotalTxs:                   newMetricPolygonHeimdallTotalTxs(settings.PolygonHeimdallTotalTxs),
		metricPolygonHeimdallUnconfirmedTxs:             newMetricPolygonHeimdallUnconfirmedTxs(settings.PolygonHeimdallUnconfirmedTxs),
		metricPolygonPolygonStateSync:                   newMetricPolygonPolygonStateSync(settings.PolygonPolygonStateSync),
	}
	for _, op := range options {
		op(mb)
	}
	return mb
}

// Emit appends generated metrics to a pdata.MetricsSlice and updates the internal state to be ready for recording
// another set of data points. This function will be doing all transformations required to produce metric representation
// defined in metadata and user settings, e.g. delta/cumulative translation.
func (mb *MetricsBuilder) Emit(metrics pdata.MetricSlice) {
	mb.metricPolygonBorAverageBlockTime.emit(metrics)
	mb.metricPolygonBorLastBlock.emit(metrics)
	mb.metricPolygonEthStateSync.emit(metrics)
	mb.metricPolygonEthSubmitCheckpointTime.emit(metrics)
	mb.metricPolygonHeimdallAverageBlockTime.emit(metrics)
	mb.metricPolygonHeimdallCheckpointValidatorsSigned.emit(metrics)
	mb.metricPolygonHeimdallCurrentSpanEndBlock.emit(metrics)
	mb.metricPolygonHeimdallLastBlock.emit(metrics)
	mb.metricPolygonHeimdallTotalTxs.emit(metrics)
	mb.metricPolygonHeimdallUnconfirmedTxs.emit(metrics)
	mb.metricPolygonPolygonStateSync.emit(metrics)
}

// RecordPolygonBorAverageBlockTimeDataPoint adds a data point to polygon.bor.average_block_time metric.
func (mb *MetricsBuilder) RecordPolygonBorAverageBlockTimeDataPoint(ts pdata.Timestamp, val float64, chainAttributeValue string) {
	mb.metricPolygonBorAverageBlockTime.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonBorLastBlockDataPoint adds a data point to polygon.bor.last_block metric.
func (mb *MetricsBuilder) RecordPolygonBorLastBlockDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonBorLastBlock.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonEthStateSyncDataPoint adds a data point to polygon.eth.state_sync metric.
func (mb *MetricsBuilder) RecordPolygonEthStateSyncDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonEthStateSync.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonEthSubmitCheckpointTimeDataPoint adds a data point to polygon.eth.submit_checkpoint_time metric.
func (mb *MetricsBuilder) RecordPolygonEthSubmitCheckpointTimeDataPoint(ts pdata.Timestamp, val float64, chainAttributeValue string) {
	mb.metricPolygonEthSubmitCheckpointTime.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonHeimdallAverageBlockTimeDataPoint adds a data point to polygon.heimdall.average_block_time metric.
func (mb *MetricsBuilder) RecordPolygonHeimdallAverageBlockTimeDataPoint(ts pdata.Timestamp, val float64, chainAttributeValue string) {
	mb.metricPolygonHeimdallAverageBlockTime.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonHeimdallCheckpointValidatorsSignedDataPoint adds a data point to polygon.heimdall.checkpoint_validators_signed metric.
func (mb *MetricsBuilder) RecordPolygonHeimdallCheckpointValidatorsSignedDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonHeimdallCheckpointValidatorsSigned.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonHeimdallCurrentSpanEndBlockDataPoint adds a data point to polygon.heimdall.current_span_end_block metric.
func (mb *MetricsBuilder) RecordPolygonHeimdallCurrentSpanEndBlockDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonHeimdallCurrentSpanEndBlock.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonHeimdallLastBlockDataPoint adds a data point to polygon.heimdall.last_block metric.
func (mb *MetricsBuilder) RecordPolygonHeimdallLastBlockDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonHeimdallLastBlock.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonHeimdallTotalTxsDataPoint adds a data point to polygon.heimdall.total_txs metric.
func (mb *MetricsBuilder) RecordPolygonHeimdallTotalTxsDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonHeimdallTotalTxs.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonHeimdallUnconfirmedTxsDataPoint adds a data point to polygon.heimdall.unconfirmed_txs metric.
func (mb *MetricsBuilder) RecordPolygonHeimdallUnconfirmedTxsDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonHeimdallUnconfirmedTxs.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// RecordPolygonPolygonStateSyncDataPoint adds a data point to polygon.polygon.state_sync metric.
func (mb *MetricsBuilder) RecordPolygonPolygonStateSyncDataPoint(ts pdata.Timestamp, val int64, chainAttributeValue string) {
	mb.metricPolygonPolygonStateSync.recordDataPoint(mb.startTime, ts, val, chainAttributeValue)
}

// Reset resets metrics builder to its initial state. It should be used when external metrics source is restarted,
// and metrics builder should update its startTime and reset it's internal state accordingly.
func (mb *MetricsBuilder) Reset(options ...metricBuilderOption) {
	mb.startTime = pdata.NewTimestampFromTime(time.Now())
	for _, op := range options {
		op(mb)
	}
}

// Attributes contains the possible metric attributes that can be used.
var Attributes = struct {
	// Chain (The name of a chain.)
	Chain string
}{
	"chain",
}

// A is an alias for Attributes.
var A = Attributes
