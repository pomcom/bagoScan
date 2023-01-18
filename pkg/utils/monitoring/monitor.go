package monitoring

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	activeGoroutines = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "active_goroutines",
			Help: "Number of active goroutines",
		}, []string{"Run"},
	)
)

func init() {
	prometheus.MustRegister(activeGoroutines)
}

func IncrementGoroutineCount(labelValues ...string) {
	activeGoroutines.WithLabelValues(labelValues...).Inc()
}

func DecrementGoroutineCount(labelValues ...string) {
	activeGoroutines.WithLabelValues(labelValues...).Dec()
}

func StartPrometheusServer() {
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(":8080", nil)
	}()
}
