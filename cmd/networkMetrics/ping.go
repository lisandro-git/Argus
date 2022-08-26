package networkMetrics

import (
	"argus/cmd"
	"fmt"
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
var networkLatency = cmd.NewGaugeVec("network_latency", "Current network latency.", []string{"Latency"})

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
