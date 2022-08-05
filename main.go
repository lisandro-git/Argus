package main

import (
	"argus/cmd"
	hm "argus/cmd/hardware_metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var wg sync.WaitGroup

var (
	cpuUsage = cmd.NewGauge("cpu_system_usage", "Current system usage of the CPU.")
	hddSize  = cmd.NewGaugeVec("hdd_size", "Current size of the HDD.", []string{"Path"})
	//sysMemory      = cmd.NewGaugeVec("sys_memory", "Current system memory usage.", []string{"RAM"})
	networkClient  = cmd.NewGaugeVec("network_client", "Current network client usage.", []string{"NetworkClient"})
	uptime         = cmd.NewGaugeVec("uptime", "Current uptime of the system.", []string{"Time"})
	networkLatency = cmd.NewGaugeVec("network_latency", "Current network latency.", []string{"Latency"})
)

func init() {

}

func server() { // https://developpaper.com/implementation-of-prometheus-custom-exporter/
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
	//nm.PingClient(true, false, "192.168.1.240")

	wg.Add(3)
	go func() {
		defer wg.Done()
		server()
	}()
	//go func() {
	//	defer wg.Done()
	//	for {
	//		networkClient.WithLabelValues("NetworkClient").Set(float64(nm.GetNetworkClient()))
	//		time.Sleep(time.Duration(1) * time.Minute)
	//	}
	//}()
	go func() {
		defer wg.Done()
		for {
			//cpuUsage.Set(hm.CpuUsage())
			hddSize.WithLabelValues("/mnt").Set(hm.GetDiskSize("/mnt"))
			hddSize.WithLabelValues("/").Set(hm.GetDiskSize("/"))
			//hm.sysMemory.WithLabelValues("Average").Set(hm.SysMemoryAverage())
			hm.SendMetrics()
			//day, hour := om.Getuptime()
			//uptime.WithLabelValues("Days").Set(day)
			//uptime.WithLabelValues("Hours").Set(hour)
			//_, latency := nm.PingClient(true, false, "192.168.1.240")
			//networkLatency.WithLabelValues("192.168.1.240").Set((float64(latency.Microseconds()) / 1000.0))
			time.Sleep(5 * time.Second)
		}
	}()
	wg.Wait()
}
