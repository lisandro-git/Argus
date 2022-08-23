package hardware_metrics

import "argus/cmd"

func RegisterMetrics() {
	cmd.RegisterCollector(NewCpuUsage())
	cmd.RegisterCollector(NewDisk())
	cmd.RegisterCollector(NewSysMemory())
}
