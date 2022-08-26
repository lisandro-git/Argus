package hardware_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CpuUsage struct {
	CPUValues *prometheus.Desc
}

var (
	// CPUser CPU utilization in Userland
	CPUser = 0
	// CPNice CPU utilization in Userland with CPU High Priority
	CPNice = 1
	// CPSys CPU utilization in Kernel-land
	CPSys = 2
	// CPIntr CPU interruptions
	CPIntr = 3
	// CPIdle When the CPU is idle
	CPIdle = 4
	// CPUStates
	CPUStates = 5
)

func NewCpuUsage() *CpuUsage {
	return &CpuUsage{
		CPUValues: prometheus.NewDesc("cpu_metric_value", "Current CPU Values", []string{"cpu_value"}, nil),
	}
}

func (c *CpuUsage) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.CPUValues
}

func (c *CpuUsage) Collect(ch chan<- prometheus.Metric) {
	percent, _ := cpu.Percent(time.Second, true)
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.GaugeValue, percent[CPUser], "CPUUser")
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.GaugeValue, percent[CPNice], "CPUNice")
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.GaugeValue, percent[CPSys], "CPUSys")
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.GaugeValue, percent[CPIntr], "CPUIntre")
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.GaugeValue, percent[CPIdle], "CPUIdle")
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.GaugeValue, percent[CPUStates], "CPUStates")
}
