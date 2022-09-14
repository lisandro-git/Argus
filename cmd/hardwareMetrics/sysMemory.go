package hardwareMetrics

import (
	"argus/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"syscall"
)

// SysMemory is a collector for system memory metrics.
type SysMemory struct {
	ram *prometheus.Desc
}

// NewSysMemory returns a new SysMemory collector.
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

// SysTotalMemory returns the total amount of usable memory (RAM) in bytes.
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

// SysFreeMemory returns the amount of free memory (RAM) in bytes.
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

// SysMemoryUsage returns the amount of used memory (RAM) in bytes.
func SysMemoryUsage() float64 {
	return float64(SysTotalMemory()) - float64(SysFreeMemory())
}
