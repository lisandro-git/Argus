package main

import (
	"argus/cmd"
	hm "argus/cmd/hardware_metrics"
	nm "argus/cmd/network_metrics"
	om "argus/cmd/os_metrics"
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
	go func() {
		defer wg.Done()
		for {
			hm.SendMetrics()
			nm.SendMetrics()
			om.SendMetrics()
			time.Sleep(5 * time.Second)
		}
	}()
	wg.Wait()
}
