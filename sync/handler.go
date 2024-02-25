package sync

import (
	"go-metrics/metrics"
	"math/rand"
	"net/http"
	"time"
)

func (s PacketSyncHandler) Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		success := false
		startTime := time.Now()

		lettertype := req.URL.Query().Get("lettertype")

		defer func(success *bool, startTime time.Time) {
			endTime := time.Now()
			if *success {
				metrics.UpdatePacketSyncSuccess(endTime.Sub(startTime), lettertype)
			} else {
				metrics.UpdatePacketSyncFailure(endTime.Sub(startTime), lettertype)
			}
		}(&success, startTime)

		sleepTime := time.Duration(rand.Intn(800)+200) * time.Millisecond
		time.Sleep(sleepTime)

		success = true

		res.Write([]byte("Request completed"))
	}
}
