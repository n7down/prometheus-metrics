package prometheus

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/n7down/prometheus-metrics/internal/clients/kubernetes"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	FIVE_MINUTES = 5 * time.Minute
)

type PrometheusMetrics struct {
	NumberPods    prometheus.Gauge
	EndPointCount *prometheus.CounterVec
	client        *kubernetes.KubernetesClient
}

func NewPrometheusMetrics(c *kubernetes.KubernetesClient) *PrometheusMetrics {

	numberPods := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "number_of_pods",
		Help: "The number of pods on the kubernetes cluster.",
	})

	metricsEndPointCount := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "end_point_count",
			Help: "Number of requests made on a end point.",
		},
		[]string{"endpoint"},
	)

	return &PrometheusMetrics{
		NumberPods:    numberPods,
		EndPointCount: metricsEndPointCount,
		client:        c,
	}
}

func (m *PrometheusMetrics) Init() {
	prometheus.MustRegister(m.NumberPods)
	prometheus.MustRegister(m.EndPointCount)
	m.getNumberOfPods()
}

func (m *PrometheusMetrics) getNumberOfPods() {
	numberOfPods := m.client.GetNumberOfPods("default")
	m.NumberPods.Set(float64(numberOfPods))
}

func (m *PrometheusMetrics) IncrementEndPointCount(label string) {
	m.EndPointCount.WithLabelValues(label).Inc()
}

func (m *PrometheusMetrics) MetricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
