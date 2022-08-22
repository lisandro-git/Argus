package hardware_metrics

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"github.com/prometheus/client_golang/prometheus"
	"os"
	"time"
)

type cpuUsage struct {
	Desc *prometheus.Desc
}

func (c *cpuUsage) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Desc
}

func (c *cpuUsage) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(c.Desc, prometheus.GaugeValue, CpuUsage())
}

func NewCpuUsage() *cpuUsage {
	return &cpuUsage{
		Desc: prometheus.NewDesc("cpu_usage", "CPU usage", nil, nil),
	}
}

func CpuUsage() float64 {
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
