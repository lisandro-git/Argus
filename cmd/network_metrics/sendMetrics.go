package network_metrics

import "argus/cmd"

func init() {
	cmd.RegisterCollector(networkClient)
	cmd.RegisterCollector(networkLatency)
}

func SendMetrics() {
	networkClient.WithLabelValues("NetworkClient").Set(float64(GetNetworkClient()))

	var _, latency = PingClient(true, false, "192.168.1.240")
	networkLatency.WithLabelValues("192.168.1.240").Set((float64(latency.Microseconds()) / 1000.0))
}
