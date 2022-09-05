package softwareMetrics

import (
	"argus/cmd/webMetrics/giteaCollector"
	"argus/cmd/webMetrics/githubCollector"
)

func RegisterMetrics() {
	giteaCollector.Start()
	githubCollector.Start()
}
