package metrics

import (
	"github.com/gin-gonic/gin"
)

type Metrics interface {
	Init()
	SetNumberPods(n float64)
	IncrementEndPointCount(label string)
	MetricsHandler() gin.HandlerFunc
}
