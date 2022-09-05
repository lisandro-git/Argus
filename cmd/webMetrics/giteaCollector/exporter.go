package giteaCollector

import (
	"argus/cmd"
	"argus/cmd/webMetrics/giteaCollector/collector"
)

func Start() {
	cmd.RegisterCollector(collector.NewgCollector())
}
