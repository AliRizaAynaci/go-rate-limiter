package prometheus

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	redisTotalRequests = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_total_requests",
			Help: "Total number of requests to Redis",
		},
	)

	redisRequestsPerEndpoint = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "redis_requests_per_endpoint",
			Help: "Requests to Redis grouped by endpoint",
		},
		[]string{"endpoint"},
	)

	redisRateLimitViolations = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "redis_rate_limit_violations",
			Help: "Number of rate limit violations",
		},
	)
)

func InitMetrics() {
	prometheus.MustRegister(redisTotalRequests)
	prometheus.MustRegister(redisRequestsPerEndpoint)
	prometheus.MustRegister(redisRateLimitViolations)
}

func RecordRedisRequest(endpoint string) {
	fmt.Println("Metric recorded for:", endpoint) // Debug için
	redisTotalRequests.Inc()
	redisRequestsPerEndpoint.WithLabelValues(endpoint).Inc()
}

func RecordRateLimitViolation() {
	fmt.Println("Rate limit violation recorded!") // Debug için
	redisRateLimitViolations.Inc()
}
