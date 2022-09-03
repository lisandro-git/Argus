package main

import (
	"argus/cmd"
	hm "argus/cmd/hardwareMetrics"
	om "argus/cmd/osMetrics"
	sm "argus/cmd/softwareMetrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

var (
	TYPE = "tcp4"
	ADDR = ""
	PORT = "8080"
)

func startListener() {
	h := promhttp.HandlerFor(cmd.Gatherer, promhttp.HandlerOpts{
		ErrorHandling: promhttp.ContinueOnError,
		Registry:      cmd.Registry,
	})

	var handler http.Handler = promhttp.InstrumentMetricHandler(cmd.Registry, h)

	server := &http.Server{Handler: handler}
	l, err := net.Listen(TYPE, ADDR+":"+PORT)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", handler)
	http.Handle("/metrics", handler)

	err = server.Serve(l)
}

func main() {
	hm.RegisterMetrics()
	om.RegisterMetrics()
	sm.RegisterMetrics()
	//networkMetrics.RegisterMetrics()
	wg.Add(1)
	go func() {
		defer wg.Done()
		startListener()
	}()
	wg.Wait()
}
