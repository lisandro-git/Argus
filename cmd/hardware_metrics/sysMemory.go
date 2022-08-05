package hardware_metrics

import (
	"argus/cmd"
	"syscall"
)

var SysMemory = cmd.NewGaugeVec("sys_memory", "Current system memory usage.", []string{"RAM"})

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
