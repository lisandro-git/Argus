package networkMetrics

import (
	"fmt"
	ps "github.com/kotakanbe/go-pingscanner"
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"os"
	"time"

	"github.com/digineo/go-ping"
)

var (
	attempts       uint = 3
	timeout             = time.Second
	size           uint = 56
	proto4, proto6 bool
	bind           string
	destination    string
	remoteAddr     *net.IPAddr
	pinger         *ping.Pinger
)

type NetworkClient struct {
	networkLatency *prometheus.Desc
	//connectedClient *prometheus.Desc
}

func NewNetworkClient() *NetworkClient {
	return &NetworkClient{
		networkLatency: prometheus.NewDesc("network_latency", "Network latency", []string{"host"}, nil),
		//connectedClient: prometheus.NewDesc("connected_client", "Number of connected clients", []string{"host"}, nil),
	}
}

func (n *NetworkClient) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.networkLatency
	//ch <- n.connectedClient
}

func (n *NetworkClient) Collect(ch chan<- prometheus.Metric) {
	for _, host := range Hosts {
		var _, latency = pingClient(true, false, host)
		ch <- prometheus.MustNewConstMetric(n.networkLatency, prometheus.GaugeValue, (float64(latency.Microseconds()) / 1000.0), host)
	}
	//ch <- prometheus.MustNewConstMetric(n.connectedClient, prometheus.CounterValue, float64(getConnectedClient()), "connectedClient")
}

func pingClient(proto4, proto6 bool, host string) (*net.IPAddr, time.Duration) {
	var network string
	if bind == "" {
		if proto4 {
			bind = "0.0.0.0"
			network = "ip4"
		} else if proto6 {
			bind = "::"
			network = "ip6"
		}
	}

	if proto4 {
		if r, err := net.ResolveIPAddr(network, host); err != nil {
			panic(err)
		} else {
			remoteAddr = r
		}

		if p, err := ping.New(bind, ""); err != nil {
			panic(err)
		} else {
			pinger = p
		}
	}
	defer pinger.Close()

	if pinger.PayloadSize() != uint16(size) {
		pinger.SetPayloadSize(uint16(size))
	}

	return unicastPing()
}

func unicastPing() (*net.IPAddr, time.Duration) {
	rtt, err := pinger.PingAttempts(remoteAddr, timeout, int(attempts))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return remoteAddr, rtt
}

func getConnectedClient() int {
	scanner := ps.PingScanner{
		CIDR: SUBNET,
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
