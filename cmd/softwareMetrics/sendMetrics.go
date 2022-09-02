package softwareMetrics

import (
	"argus/cmd/softwareMetrics/giteaCollector"
	"argus/cmd/softwareMetrics/nginxCollector"
)

func RegisterMetrics() {
	nginxCollector.Start()
	giteaCollector.Start()
}
