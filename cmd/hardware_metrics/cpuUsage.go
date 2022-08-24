package hardware_metrics

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"time"
)

//var cpuUsage = cmd.NewGauge("cpu_system_usage", "Current system usage of the CPU.")

type CpuUsage struct {
	load *prometheus.Desc
}

func NewCpuUsage() *CpuUsage {
	return &CpuUsage{
		load: prometheus.NewDesc("cpu_system_usage", "Current system usage of the CPU.", nil, nil),
	}
}

func (c *CpuUsage) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.load
}

func (c *CpuUsage) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.load, prometheus.GaugeValue, cpuUsage())
}

func cpuUsage() float64 {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 0.0
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 0.0
	}
	total := float64(after.Total - before.Total)
	return float64(after.System-before.System) / total * 100
}
