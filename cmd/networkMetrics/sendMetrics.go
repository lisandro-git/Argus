package networkMetrics

import "argus/cmd"

func RegisterMetrics() {
	cmd.RegisterCollector(NewNetworkClient())
}
