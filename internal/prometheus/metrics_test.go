package prometheus_test

import (
	"rate-limiter/internal/prometheus"
	"testing"
)

func TestMetricsRecording(t *testing.T) {
	prometheus.InitMetrics()

	prometheus.RecordRedisRequest("/test")
	prometheus.RecordRateLimitViolation()
}
