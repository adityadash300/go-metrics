package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type packetUploadMetrics struct {
	successCount prometheus.Counter
	failureCount prometheus.Counter
}

func NewPacketUploadMetrics() packetUploadMetrics {
	return packetUploadMetrics{
		successCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "ens",
			Name:      "packetupload_success_count",
			Help:      "Sucess count of Packet Upload",
		}),
		failureCount: promauto.NewCounter(prometheus.CounterOpts{
			Namespace: "ens",
			Name:      "packetupload_failure_count",
			Help:      "Failure count of Packet Upload",
		}),
	}
}
