package hardware_metrics

import (
	"fmt"
	"github.com/mackerelio/go-osstat/cpu"
	"os"
	"time"
)

func CpuUsage() float64 {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 0.0
	}
	time.Sleep(time.Duration(1) * time.Second)
	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 0.0
	}
	total := float64(after.Total - before.Total)
	return float64(after.System-before.System) / total * 100
}
