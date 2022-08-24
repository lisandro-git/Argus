package main

import (
	"argus/cmd"
	hm "argus/cmd/hardware_metrics"
	om "argus/cmd/os_metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"sync"
)

var wg sync.WaitGroup

func server() {
	h := promhttp.HandlerFor(cmd.Gatherer, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
		Registry:      cmd.Register,
	})

	handler := promhttp.InstrumentMetricHandler(cmd.Register, h)

	http.Handle("/metrics", handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Printf("Error occur when start server %v", err)
		os.Exit(1)
	}
}

func main() {
	hm.RegisterMetrics()
	om.RegisterMetrics()
	wg.Add(1)
	go func() {
		defer wg.Done()
		server()
	}()
	wg.Wait()
}
