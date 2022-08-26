package osMetrics

import "argus/cmd"

func RegisterMetrics() {
	cmd.RegisterCollector(NewUptime())
}
