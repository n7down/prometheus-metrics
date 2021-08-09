package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/n7down/prometheus-metrics/internal/api"
	"github.com/n7down/prometheus-metrics/internal/clients/kubernetes"
	"github.com/n7down/prometheus-metrics/internal/metrics/prometheus"
)

const (
	port = "8080"
)

var (
	prometheusMetrics *prometheus.PrometheusMetrics
)

func init() {
	var (
		token        *string
		apiServerURL *string
	)

	token = flag.String("token", "", "The token for the API Server")
	apiServerURL = flag.String("apiserver", "", "The URL to the API Server")
	flag.Parse()

	kubernetesClient := kubernetes.NewKubernetesClient(*apiServerURL, *token)

	prometheusMetrics = prometheus.NewPrometheusMetrics(kubernetesClient)
	prometheusMetrics.Init()
}

func main() {
	r := gin.Default()

	apiServer := api.NewAPIServer(prometheusMetrics)

	apiServer.InitV1Routes(r)

	err := apiServer.Run(r, port)
	if err != nil {
		log.Fatal(err)
	}
}
