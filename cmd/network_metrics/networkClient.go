package network_metrics

import (
	"argus/cmd"
	"fmt"
	ps "github.com/kotakanbe/go-pingscanner"
)

var networkClient = cmd.NewGaugeVec("network_client", "Current network client usage.", []string{"NetworkClient"})

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
