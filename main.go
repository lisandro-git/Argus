package main

import (
	"argus/cmd"
	"argus/cmd/hardware_metrics"
	"argus/cmd/network_metrics"
	"argus/cmd/os_metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

func server() {
	h := promhttp.HandlerFor(cmd.Gatherer, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("Error occur when start server %v", err)
		os.Exit(1)
	}
}

func main() {
	wg.Add(2)
	go func() {
		defer wg.Done()
		server()
	}()
	hardware_metrics.RegisterMetrics()
	os_metrics.RegisterMetrics()
	network_metrics.RegisterMetrics()
	go func() {
		defer wg.Done()
		for {

			//hm.SendMetrics()
			//nm.SendMetrics()
			//om.SendMetrics()
			time.Sleep(5 * time.Second)
		}
	}()
	wg.Wait()
}
