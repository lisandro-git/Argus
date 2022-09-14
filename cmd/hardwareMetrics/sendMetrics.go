package hardwareMetrics

import "argus/cmd"

// RegisterMetrics registers the new metric to the prometheus registry
func RegisterMetrics() {
	cmd.RegisterCollector(NewCpuUsage())
	cmd.RegisterCollector(NewSysMemory())
	cmd.RegisterCollector(NewDiskSize())
}
