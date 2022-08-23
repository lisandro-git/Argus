package network_metrics

import (
	"fmt"
	ps "github.com/kotakanbe/go-pingscanner"
	"github.com/prometheus/client_golang/prometheus"
)

type NetworkClient struct {
	networkClient *prometheus.Desc
}

func NewNetworkClient() *NetworkClient {
	return &NetworkClient{
		networkClient: prometheus.NewDesc("network_client", "Number of network clients", nil, nil),
	}
}

func (n *NetworkClient) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.networkClient
}

func (n *NetworkClient) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(n.networkClient, prometheus.GaugeValue, float64(GetNetworkClient()))
}

func GetNetworkClient() int {
	scanner := ps.PingScanner{
		CIDR: "192.168.1.1/24",
		PingOptions: []string{
			"-c1",
			"-t1",
		},
		NumOfConcurrency: 254,
	}
	aliveIPs, err := scanner.Scan()
	if err != nil {
		fmt.Println(err)
	}
	return len(aliveIPs)
}
