package giteaCollector

import (
	"argus/cmd"
	"argus/cmd/softwareMetrics/giteaCollector/collector"
)

func Start() {
	cmd.RegisterCollector(collector.NewgCollector())
}
