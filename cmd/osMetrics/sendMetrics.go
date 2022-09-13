package osMetrics

import "argus/cmd"

// RegisterMetrics registers the new metric to the prometheus registry
func RegisterMetrics() {
	cmd.RegisterCollector(NewUptime())
}
