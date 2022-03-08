# GethInfluxDB Receiver

This receiver accepts metrics data as [InfluxDB v2 Line Protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/) and it's specially adapted to accept and convert go-ethereum client InfluxDB v2 metrics to be correctly parsed into `prometheus-v1` schema expected by [influx2otel](https://github.com/influxdata/influxdb-observability/tree/main/influx2otel) library.

Supported pipeline types: metrics

Write endpoints exist at `/api/v2/write` (InfluxDB 2.x compatibility).
Write query parameters `org`/`bucket` (InfluxDB 2.x) are ignored.
Write query parameter `precision` is optional, defaults to `ns`.

Write responses:
- 204: success, no further response needed (no content)
- 400: permanent failure; check response body for details
- 500: retryable error; check response body for details

## Configuration

The following configuration options are supported:

* `endpoint` (default = 0.0.0.0:8086) HTTP service endpoint for the line protocol receiver
* `token` (default = "") Authentication token to control write access

The full list of settings exposed for this receiver are documented in [config.go](config.go).

Example:
```yaml
receivers:
  influxdb:
    endpoint: 0.0.0.0:8080
    token: myToken
```

## Definitions

[InfluxDB](https://www.influxdata.com/products/influxdb/) is an open-source time series database.

[Telegraf](https://www.influxdata.com/time-series-platform/telegraf/) is an open-source metrics agent, similar to the OpenTelemetry Collector.
Telegraf has [hundreds of plugins](https://www.influxdata.com/products/integrations/?_integrations_dropdown=telegraf-plugins).

[Line protocol](https://docs.influxdata.com/influxdb/v2.0/reference/syntax/line-protocol/) is a textual HTTP payload format used to move metrics between Telegraf agents and InfluxDB instances.

[Ethereum Client]([https://github.com/ethereum/go-ethereum) Ethereum Go reference implementation.
## Schema

The InfluxDB->OpenTelemetry conversion [schema](https://github.com/influxdata/influxdb-observability/blob/main/docs/index.md) and [implementation](https://github.com/influxdata/influxdb-observability/tree/main/influx2otel) are hosted at https://github.com/influxdata/influxdb-observability .
This receiver automatically detects schema at parse time.

### Example: Metrics - `prometheus-v1`
```
cpu_temp,foo=bar gauge=87.332
http_requests_total,method=post,code=200 counter=1027
http_requests_total,method=post,code=400 counter=3
http_request_duration_seconds 0.05=24054,0.1=33444,0.2=100392,0.5=129389,1=133988,sum=53423,count=144320
rpc_duration_seconds 0.01=3102,0.05=3272,0.5=4773,0.9=9001,0.99=76656,sum=1.7560473e+07,count=2693
```

### Example: Metrics - `prometheus-v2`
```
prometheus,foo=bar cpu_temp=87.332
prometheus,method=post,code=200 http_requests_total=1027
prometheus,method=post,code=400 http_requests_total=3
prometheus,le=0.05 http_request_duration_seconds_bucket=24054
prometheus,le=0.1  http_request_duration_seconds_bucket=33444
prometheus,le=0.2  http_request_duration_seconds_bucket=100392
prometheus,le=0.5  http_request_duration_seconds_bucket=129389
prometheus,le=1    http_request_duration_seconds_bucket=133988
prometheus         http_request_duration_seconds_count=144320,http_request_duration_seconds_sum=53423
prometheus,quantile=0.01 rpc_duration_seconds=3102
prometheus,quantile=0.05 rpc_duration_seconds=3272
prometheus,quantile=0.5  rpc_duration_seconds=4773
prometheus,quantile=0.9  rpc_duration_seconds=9001
prometheus,quantile=0.99 rpc_duration_seconds=76656
prometheus               rpc_duration_seconds_count=1.7560473e+07,rpc_duration_seconds_sum=2693
```
