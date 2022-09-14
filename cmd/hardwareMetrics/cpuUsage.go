package hardwareMetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/cpu"
	"io/ioutil"
	"strconv"
	"time"
)

// CpuUsage is a collector for cpu usage
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

// NewCpuUsage returns a new CpuUsage collector
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
	ch <- prometheus.MustNewConstMetric(c.CPUValues, prometheus.CounterValue, cpuTemp(), "CPUTemp")
}

// cpuTemp returns the current CPU temperature from /sys/class/thermal/thermal_zone0/temp
func cpuTemp() float64 {
	temp, err := ioutil.ReadFile("/sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		return 0.0
	}

	y, err := strconv.Atoi(string(temp)[:len(string(temp))-1])
	if err != nil {
		return 0.0
	}
	return float64(y) / 1000
}
