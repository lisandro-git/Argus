package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/digineo/go-ping"
)

var (
	attempts       uint = 3
	timeout             = time.Second
	proto4, proto6 bool
	size           uint = 56
	bind           string

	destination string
	remoteAddr  *net.IPAddr
	pinger      *ping.Pinger
)

func PingClient(proto4, proto6 bool, host string) {
	if proto4 == proto6 {
		log.Fatalf("need exactly one of -4 and -6 flags")
	}

	if bind == "" {
		if proto4 {
			bind = "0.0.0.0"
		} else if proto6 {
			bind = "::"
		}
	}

	if proto4 {
		if r, err := net.ResolveIPAddr("ip4", host); err != nil {
			panic(err)
		} else {
			remoteAddr = r
		}

		if p, err := ping.New(bind, ""); err != nil {
			panic(err)
		} else {
			pinger = p
		}
	} else if proto6 {
		if r, err := net.ResolveIPAddr("ip6", host); err != nil {
			panic(err)
		} else {
			remoteAddr = r
		}

		if p, err := ping.New("", bind); err != nil {
			panic(err)
		} else {
			pinger = p
		}
	}
	defer pinger.Close()

	if pinger.PayloadSize() != uint16(size) {
		pinger.SetPayloadSize(uint16(size))
	}

	unicastPing()
}

func unicastPing() {
	rtt, err := pinger.PingAttempts(remoteAddr, timeout, int(attempts))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("ping %s (%s) rtt=%v\n", destination, remoteAddr, rtt)
}
