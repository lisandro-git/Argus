package softwareMetrics

import (
	"argus/cmd/webMetrics/giteaCollector"
	"argus/cmd/webMetrics/githubCollector"
	"argus/cmd/webMetrics/piholeCollector"
)

func RegisterMetrics() {
	giteaCollector.Start()
	githubCollector.Start()
	piholeCollector.S()
}
