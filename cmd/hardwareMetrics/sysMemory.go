package hardwareMetrics

import (
	"argus/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"syscall"
)

type SysMemory struct {
	ram *prometheus.Desc
}

func NewSysMemory() *SysMemory {
	return &SysMemory{
		ram: prometheus.NewDesc("sys_memory", "Total system memory.", []string{"ram_usage"}, nil),
	}
}

func (s *SysMemory) Describe(ch chan<- *prometheus.Desc) {
	ch <- s.ram
}

func (s *SysMemory) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(s.ram, prometheus.GaugeValue, float64(SysTotalMemory())/cmd.MB, "Total")
	ch <- prometheus.MustNewConstMetric(s.ram, prometheus.GaugeValue, float64(SysFreeMemory())/cmd.MB, "Free")
	ch <- prometheus.MustNewConstMetric(s.ram, prometheus.GaugeValue, float64(SysMemoryUsage())/cmd.MB, "Usage")
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

func SysMemoryUsage() float64 {
	return float64(SysTotalMemory()) - float64(SysFreeMemory())
}
