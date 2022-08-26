package hardware_metrics

import (
	"argus/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"syscall"
)

type SysMemory struct {
	total *prometheus.Desc
	free  *prometheus.Desc
	usage *prometheus.Desc
}

func NewSysMemory() *SysMemory {
	return &SysMemory{
		total: prometheus.NewDesc("sys_memory_total", "Total system memory.", []string{"RAM"}, nil),
		free:  prometheus.NewDesc("sys_memory_free", "Free system memory.", []string{"RAM"}, nil),
		usage: prometheus.NewDesc("sys_memory_usage", "Current system memory usage.", []string{"RAM"}, nil),
	}
}

func (s *SysMemory) Describe(ch chan<- *prometheus.Desc) {
	ch <- s.total
	ch <- s.free
	ch <- s.usage
}

func (s *SysMemory) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(s.total, prometheus.GaugeValue, float64(SysTotalMemory())/cmd.MB, "Total")
	ch <- prometheus.MustNewConstMetric(s.free, prometheus.GaugeValue, float64(SysFreeMemory())/cmd.MB, "Free")
	ch <- prometheus.MustNewConstMetric(s.usage, prometheus.GaugeValue, SysMemoryAverage()/cmd.MB, "Usage")
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
