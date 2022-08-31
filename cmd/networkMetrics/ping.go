package networkMetrics

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"net"
	"os"
	"time"

	"github.com/digineo/go-ping"
)

type NetworkClient struct {
	networkLatency *prometheus.Desc
}

func NewNetworkClient() *NetworkClient {
	return &NetworkClient{
		networkLatency: prometheus.NewDesc("network_latency", "Network latency", []string{"host"}, nil),
	}
}

func (n *NetworkClient) Describe(ch chan<- *prometheus.Desc) {
	ch <- n.networkLatency
}

func (n *NetworkClient) Collect(ch chan<- prometheus.Metric) {
	for _, host := range Hosts {
		var _, latency = PingClient(true, false, host)
		ch <- prometheus.MustNewConstMetric(n.networkLatency, prometheus.GaugeValue, (float64(latency.Microseconds()) / 1000.0), host)
	}
}

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

func PingClient(proto4, proto6 bool, host string) (*net.IPAddr, time.Duration) {
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
