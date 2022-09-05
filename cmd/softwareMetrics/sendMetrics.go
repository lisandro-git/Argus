package softwareMetrics

import (
	"argus/cmd/softwareMetrics/nginxCollector"
)

func RegisterMetrics() {
	nginxCollector.Start()
}
