package softwareMetrics

import "argus/cmd/webMetrics/githubCollector"

func RegisterMetrics() {
	//giteaCollector.Start()
	githubCollector.Start()
	//piholeCollector.S()
}
