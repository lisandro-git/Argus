package hardware_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type CpuUsage struct {
	CPUser    *prometheus.Desc
	CPNice    *prometheus.Desc
	CPSys     *prometheus.Desc
	CPIntr    *prometheus.Desc
	CPIdle    *prometheus.Desc
	CPUStates *prometheus.Desc
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
		CPUser:    prometheus.NewDesc("userland_cpu_usage", "Current system usage of the CPU in userland", []string{"CPU"}, nil),
		CPNice:    prometheus.NewDesc("userland_cpu_usage_nice", "Current system usage of the CPU in userland with High Priority", []string{"CPU"}, nil),
		CPSys:     prometheus.NewDesc("kernel_land_cpu_usage", "Current system usage of the CPU in kernel-land", []string{"CPU"}, nil),
		CPIntr:    prometheus.NewDesc("cpu_interruptions", "Current number of CPU Interruptions", []string{"CPU"}, nil),
		CPIdle:    prometheus.NewDesc("cpu_idle", "CPU idle time since latest swipe", []string{"CPU"}, nil),
		CPUStates: prometheus.NewDesc("cpu_states", "CPU states", []string{"CPU"}, nil),
	}
}

func (c *CpuUsage) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.CPUser
	ch <- c.CPNice
	ch <- c.CPSys
	ch <- c.CPIntr
	ch <- c.CPIdle
	ch <- c.CPUStates
}

func (c *CpuUsage) Collect(ch chan<- prometheus.Metric) {
	percent, _ := cpu.Percent(time.Second, true)
	ch <- prometheus.MustNewConstMetric(c.CPUser, prometheus.GaugeValue, percent[CPUser], "CPU")
	ch <- prometheus.MustNewConstMetric(c.CPNice, prometheus.GaugeValue, percent[CPNice], "CPU")
	ch <- prometheus.MustNewConstMetric(c.CPSys, prometheus.GaugeValue, percent[CPSys], "CPU")
	ch <- prometheus.MustNewConstMetric(c.CPIntr, prometheus.GaugeValue, percent[CPIntr], "CPU")
	ch <- prometheus.MustNewConstMetric(c.CPIdle, prometheus.GaugeValue, percent[CPIdle], "CPU")
	ch <- prometheus.MustNewConstMetric(c.CPUStates, prometheus.GaugeValue, percent[CPUStates], "CPU")
}
