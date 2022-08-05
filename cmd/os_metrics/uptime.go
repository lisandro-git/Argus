package os_metrics

import (
	"argus/cmd"
	"github.com/shirou/gopsutil/host"
)

var uptime = cmd.NewGaugeVec("uptime", "Current uptime of the system.", []string{"Time"})

func Getuptime() (float64, float64) {
	os_uptime, _ := host.Uptime()

	var days uint64 = os_uptime / (60 * 60 * 24)
	var hours uint64 = (os_uptime - (days * 60 * 60 * 24)) / (60 * 60)
	//var minutes uint64 := ((os_uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	return float64(days), float64(hours)
}
