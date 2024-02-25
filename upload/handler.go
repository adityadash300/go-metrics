package upload

import (
	"go-metrics/metrics"
	"math/rand"
	"net/http"
	"time"
)

func (u PacketUploadHandler) Handler() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		success := false
		startTime := time.Now()

		defer func(success *bool, startTime time.Time) {
			if *success {
				metrics.UpdatePacketUploadSuccess()
			} else {
				metrics.UpdatePacketUploadFailure()
			}
		}(&success, startTime)

		sleepTime := time.Duration(rand.Intn(800)+600) * time.Millisecond
		time.Sleep(sleepTime)

		success = true

		res.Write([]byte("Request completed"))
	}
}
