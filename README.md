# OpenTelemetry Polygon distro 

This OpenTelemetry distro is aimed to monitor the Polygon infrastructure, it's focused on serve the specific needs of the Polygon team but it can be used by anyone that runs a node or several nodes and monitor several blockchain and infrastructure related metrics. It's compatible with Ethereum clients too so it can be used to monitor a set of Ethereum nodes.

## Setup

The application can be build using go > v1.17:

`go build .`

Or using Docker with the included Dockerfile

`docker build -t polygon-otel-builder .`

## Components

* [Geth Influxdb receiver](./receiver/gethinfluxdbreceiver/README.md)

