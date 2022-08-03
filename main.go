package main

import (
	"argus/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
	TB = 1024 * GB
)

var wg sync.WaitGroup

var (
	cpuTemp = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_temperature_celsius",
			Help: "Current temperature of the CPU.",
		})
	hddSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "size_of_hdd",
			Help: "Size of the HDD.",
		},
		[]string{"Path"},
	)
	sysMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "RAM_memory",
			Help: "Free space available in the RAM",
		},
		[]string{"RAM"},
	)
)

func init() {
	prometheus.MustRegister(sysMemory)
	prometheus.MustRegister(hddSize)
	prometheus.MustRegister(cpuTemp)
}

func server() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
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
			cpuTemp.Set(float64(50))
			hddSize.WithLabelValues("/mnt").Set(cmd.GetDiskSize("/mnt"))
			hddSize.WithLabelValues("/").Set(cmd.GetDiskSize("/"))
			sysMemory.WithLabelValues("Average").Set(cmd.SysMemoryAverage())
			time.Sleep(5 * time.Second)
		}
	}()
	wg.Wait()
}
