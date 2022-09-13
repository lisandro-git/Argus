package main

import (
	"argus/cmd"
	hm "argus/cmd/hardwareMetrics"
	nm "argus/cmd/networkMetrics"
	om "argus/cmd/osMetrics"
	sm "argus/cmd/softwareMetrics"
	wm "argus/cmd/webMetrics"
	"fmt"
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
	listener, err := net.Listen(TYPE, ADDR+":"+PORT)
	fmt.Println("Listening on " + ADDR + ":" + PORT)
	if err != nil {
		log.Fatal(err)
	}
	http.Handle("/", handler)
	http.Handle("/metrics", handler)

	err = server.Serve(listener)
}

func main() {
	hm.RegisterMetrics()
	om.RegisterMetrics()
	sm.RegisterMetrics()
	nm.RegisterMetrics()
	wm.RegisterMetrics()

	wg.Add(1)
	go func() {
		defer wg.Done()
		startListener()
	}()
	wg.Wait()
}
