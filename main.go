package main

import (
	"argus/cmd"
	hm "argus/cmd/hardwareMetrics"
	om "argus/cmd/osMetrics"
	"argus/cmd/softwareMetrics"
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
		Registry:      cmd.Registry,
	})

	var handler http.Handler = promhttp.InstrumentMetricHandler(cmd.Registry, h)

	http.Handle("/metrics", handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Printf("Error occur when start server %v", err)
		os.Exit(1)
	}
}

func main() {
	hm.RegisterMetrics()
	om.RegisterMetrics()
	go softwareMetrics.RegisterMetrics()
	wg.Add(1)
	go func() {
		defer wg.Done()
		server()
	}()
	wg.Wait()
}
