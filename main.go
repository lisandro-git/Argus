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
	cpuUsage = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "cpu_system_usage",
			Help: "Current system usage of the CPU.",
		})
	hddSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "HDDSize",
			Help: "Size of the HDD.",
		},
		[]string{"Path"},
	)
	sysMemory = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "RAMMemory",
			Help: "Free space available in the RAM",
		},
		[]string{"RAM"},
	)
	networkClient = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "NetworkClient",
			Help: "Number of clients connected to the local network",
		},
		[]string{"NetworkClient"},
	)
)

func init() {
	prometheus.MustRegister(sysMemory)
	prometheus.MustRegister(hddSize)
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(networkClient)
}

func server() {
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	wg.Add(3)
	go func() {
		defer wg.Done()
		server()
	}()
	go func() {
		defer wg.Done()
		for {
			networkClient.WithLabelValues("NetworkClient").Set(float64(cmd.GetNetworkClient()))
			time.Sleep(time.Duration(1) * time.Minute)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			cpuUsage.Set(cmd.CpuUsage())
			hddSize.WithLabelValues("/mnt").Set(cmd.GetDiskSize("/mnt"))
			hddSize.WithLabelValues("/").Set(cmd.GetDiskSize("/"))
			sysMemory.WithLabelValues("Average").Set(cmd.SysMemoryAverage())
			time.Sleep(5 * time.Second)
		}
	}()
	wg.Wait()
}
