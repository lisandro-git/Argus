package os_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/host"
)

type Uptime struct {
	DayUptime  *prometheus.Desc
	HourUptime *prometheus.Desc
}

func (u *Uptime) Describe(ch chan<- *prometheus.Desc) {
	ch <- u.DayUptime
	ch <- u.HourUptime
}

func (u *Uptime) Collect(ch chan<- prometheus.Metric) {
	days, hours := Getuptime()
	ch <- prometheus.MustNewConstMetric(u.DayUptime, prometheus.GaugeValue, days)
	ch <- prometheus.MustNewConstMetric(u.HourUptime, prometheus.GaugeValue, hours)
}

func NewUptime() *Uptime {
	return &Uptime{
		DayUptime:  prometheus.NewDesc("DayUptime", "Uptime", nil, nil),
		HourUptime: prometheus.NewDesc("HourUptime", "Uptime", nil, nil),
	}
}

func Getuptime() (float64, float64) {
	os_uptime, _ := host.Uptime()

	var days uint64 = os_uptime / (60 * 60 * 24)
	var hours uint64 = (os_uptime - (days * 60 * 60 * 24)) / (60 * 60)
	//var minutes uint64 := ((os_uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	return float64(days), float64(hours)
}
