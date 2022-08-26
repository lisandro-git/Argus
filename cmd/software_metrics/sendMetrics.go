package software_metrics

import "argus/cmd/software_metrics/nginxCollector"

func RegisterMetrics() {
	nginxCollector.Start()
}
