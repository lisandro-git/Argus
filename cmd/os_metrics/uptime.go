package os_metrics

import (
	"github.com/shirou/gopsutil/host"
)

func Getuptime() (float64, float64) {
	uptime, _ := host.Uptime()

	var days uint64 = uptime / (60 * 60 * 24)
	var hours uint64 = (uptime - (days * 60 * 60 * 24)) / (60 * 60)
	//var minutes uint64 := ((uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	return float64(days), float64(hours)
}
