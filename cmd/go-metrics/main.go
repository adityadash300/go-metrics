package main

import (
	"fmt"
	"go-metrics/metrics"
	"go-metrics/sync"
	"go-metrics/upload"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	metrics.Init()

	pMux := http.NewServeMux()
	pMux.Handle("/metrics", promhttp.Handler())
	fmt.Println("Application started")

	sMux := http.NewServeMux()
	sMux.Handle("/health", http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("I am running"))
	}))

	packetSyncHandler := sync.Init()
	packetUploadHandler := upload.Init()

	sMux.Handle("/packetsync", packetSyncHandler.Handler())
	sMux.Handle("/packetupload", packetUploadHandler.Handler())

	// Start metrics server
	go func() {
		if err := http.ListenAndServe(":9091", pMux); err != nil {
			fmt.Println("error starting server", err)
		}
	}()

	startTime := time.Now()

	sleepTime := time.Duration(rand.Intn(400)+200) * time.Millisecond
	time.Sleep(sleepTime)

	endTime := time.Now()
	fmt.Println((endTime.Sub(startTime)) / time.Millisecond)

	if err := http.ListenAndServe(":8080", sMux); err != nil {
		fmt.Println("error starting server", err)
	}
}
