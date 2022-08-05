package hardware_metrics

import "argus/cmd"

func init() {
	cmd.RegisterCollector(SysMemory)
	cmd.RegisterCollector(cpuUsage)
	cmd.RegisterCollector(hddSize)
}

func SendMetrics() {
	SysMemory.WithLabelValues("RAM").Set(SysMemoryAverage())
	cpuUsage.Set(CpuUsage())
	hddSize.WithLabelValues("/mnt").Set(GetDiskSize("/mnt"))
	hddSize.WithLabelValues("/").Set(GetDiskSize("/"))
}
