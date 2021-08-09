# Prometheus Metrics

## Instructions
### Prerequisites
- Golang

### Getting Started
1. Run `go get -v ./...`
2. Run the server with `go run cmd/apiserver/main.go -apiserver=$APISERVER -token=$TOKEN` where `$APISERVER` is the API Server URL and `$TOKEN` is the token.

## Metrics
- [x] The first metric to gather is the number of pods running in the cluster. This is a Gauge Prometheus primitive, and should be simple to collect using the Kubernetes API.

- [x] The next metric to gather will be a Prometheus Counter primitive. A counter is used to represent a monotonically increasing number. The counter for this assessment will be the number of times an endpoint of the application is accessed. This endpoint can be specific for this counter, or it can be the endpoint that the metrics will be exposed on. This endpoint needs to store the counter value in memory or to a file, this can be volatile; it can be reset when the application restarts.

- [x] All metrics gathered must be presented and exposed on the /metrics path. Accessing this path will provide the metrics in the Prometheus Exposition Format.

## Notes
- [prometheus exposition format](https://prometheus.io/docs/instrumenting/exposition_formats)
- [gauge prometheus primitive](https://prometheus.io/docs/concepts/metric_types/#gauge)
- [counter prometheus primitive](https://prometheus.io/docs/concepts/metric_types/#counter)
