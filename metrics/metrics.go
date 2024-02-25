package metrics

import (
	"time"
)

type Metrics struct {
	packetSync   packetSyncMetrics
	packetUpload packetUploadMetrics
}

var m Metrics

func Init() {
	m = Metrics{
		packetSync:   NewPacketSyncMetrics(),
		packetUpload: NewPacketUploadMetrics(),
	}
}

func UpdatePacketSyncSuccess(response time.Duration, lettertype string) {
	m.packetSync.successCount.WithLabelValues(lettertype).Inc()
	m.packetSync.successResponseTime.Observe(float64(response / time.Millisecond))
}

func UpdatePacketSyncFailure(response time.Duration, lettertype string) {
	m.packetSync.failureCount.WithLabelValues(lettertype).Inc()
	m.packetSync.failureResponseTime.Observe(float64(response / time.Millisecond))
}

func UpdatePacketUploadSuccess() {
	m.packetUpload.successCount.Inc()
}

func UpdatePacketUploadFailure() {
	m.packetUpload.failureCount.Inc()
}
