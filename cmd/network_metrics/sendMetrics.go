package network_metrics

import "argus/cmd"

func RegisterMetrics() {
	cmd.RegisterCollector(NewNetworkLatency())
	cmd.RegisterCollector(NewNetworkClient())
}
