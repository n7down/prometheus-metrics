package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n7down/prometheus-metrics/internal/metrics/prometheus"
)

type APIServer struct {
	metrics *prometheus.PrometheusMetrics
}

func NewAPIServer(m *prometheus.PrometheusMetrics) *APIServer {
	return &APIServer{
		metrics: m,
	}
}

func (s *APIServer) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func (s *APIServer) metricsCounterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		s.metrics.IncrementEndPointCount("/metrics")
		c.Next()
	}
}

func (s *APIServer) InitV1Routes(router *gin.Engine) error {

	router.Use(s.corsMiddleware())

	metricsGroup := router.Group("/metrics")
	{
		metricsGroup.Use(s.metricsCounterMiddleware())
		metricsGroup.GET("", s.metrics.MetricsHandler())
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return nil
}

func (g *APIServer) Run(router *gin.Engine, port string) error {
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		return err
	}
	return nil
}
