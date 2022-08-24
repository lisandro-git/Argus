package os_metrics

import "argus/cmd"

func RegisterMetrics() {
	cmd.RegisterCollector(NewUptime())
}
