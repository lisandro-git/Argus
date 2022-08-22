package hardware_metrics

import (
	"argus/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"syscall"
)

type sysMemory struct {
	totalMemory *prometheus.Desc
	freeMemory  *prometheus.Desc
}

func (m *sysMemory) Describe(ch chan<- *prometheus.Desc) {
	ch <- m.totalMemory
	ch <- m.freeMemory
}

func (m *sysMemory) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(m.totalMemory, prometheus.GaugeValue, float64(SysTotalMemory()/cmd.MB))
	ch <- prometheus.MustNewConstMetric(m.freeMemory, prometheus.GaugeValue, float64(SysFreeMemory()/cmd.MB))
}

func NewSysMemory() *sysMemory {
	return &sysMemory{
		totalMemory: prometheus.NewDesc("sys_total_memory", "Total RAM memory", nil, nil),
		freeMemory:  prometheus.NewDesc("sys_free_memory", "Free RAM available", nil, nil),
	}
}

func SysTotalMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Totalram) * uint64(in.Unit)
}

func SysFreeMemory() uint64 {
	in := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(in)
	if err != nil {
		return 0
	}
	// If this is a 32-bit system, then these fields are
	// uint32 instead of uint64.
	// So we always convert to uint64 to match signature.
	return uint64(in.Freeram) * uint64(in.Unit)
}

func SysMemoryAverage() float64 {
	return ((float64(SysFreeMemory()) / (1024 * 1024 * 1024)) / (float64(SysTotalMemory()) / (1024 * 1024 * 1024)) * 100)
}
