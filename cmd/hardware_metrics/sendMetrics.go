package hardware_metrics

import "argus/cmd"

func RegisterMetrics() {
	cmd.RegisterCollector(NewCpuUsage())
	cmd.RegisterCollector(NewSysMemory())
	cmd.RegisterCollector(NewDiskSize())
}
