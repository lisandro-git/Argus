package softwareMetrics

import "argus/cmd/softwareMetrics/nginxCollector"

// RegisterMetrics registers the new metric to the prometheus registry
func RegisterMetrics() {
	nginxCollector.Start()
}
