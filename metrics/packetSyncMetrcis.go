package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type packetSyncMetrics struct {
	successCount        *prometheus.CounterVec
	failureCount        *prometheus.CounterVec
	successResponseTime prometheus.Histogram
	failureResponseTime prometheus.Histogram
}

func NewPacketSyncMetrics() packetSyncMetrics {
	return packetSyncMetrics{
		successCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ens",
			Name:      "packetsync_success_count",
			Help:      "Sucess count of Packet Sync",
		}, []string{"lettertype"}),
		failureCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "ens",
			Name:      "packetsync_failure_count",
			Help:      "Failure count of Packet Sync",
		}, []string{"lettertype"}),
		successResponseTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "ens",
			Name:      "packetsync_success_response_time",
			Help:      "Success response time for packet sync",
			Buckets:   []float64{200, 400, 600, 800, 1000, 1200, 1400, 1600, 1800},
		}),
		failureResponseTime: promauto.NewHistogram(prometheus.HistogramOpts{
			Namespace: "ens",
			Name:      "packetsync_failure_response_time",
			Help:      "Failure response time for packet sync",
			Buckets:   []float64{200, 400, 600, 800, 1000, 1200, 1400, 1600, 1800},
		}),
	}
}
