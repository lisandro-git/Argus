package softwareMetrics

import (
	"argus/cmd/webMetrics/giteaCollector"
)

func RegisterMetrics() {
	giteaCollector.Start()
}
