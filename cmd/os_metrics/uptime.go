package os_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/host"
)

type Uptime struct {
	Days    *prometheus.Desc
	Hours   *prometheus.Desc
	Minutes *prometheus.Desc
}

var (
	Days    = 0
	Hours   = 1
	Minutes = 2
)

func NewUptime() *Uptime {
	return &Uptime{
		Days:    prometheus.NewDesc("uptime_days", "Current uptime of the system.", []string{"Time"}, nil),
		Hours:   prometheus.NewDesc("uptime_hours", "Current uptime of the system.", []string{"Time"}, nil),
		Minutes: prometheus.NewDesc("uptime_minutes", "Current uptime of the system.", []string{"Time"}, nil),
	}
}

func (u *Uptime) Describe(ch chan<- *prometheus.Desc) {
	ch <- u.Days
	ch <- u.Hours
	ch <- u.Minutes
}

func (u *Uptime) Collect(ch chan<- prometheus.Metric) {
	totalUptime := Getuptime()
	ch <- prometheus.MustNewConstMetric(u.Days, prometheus.GaugeValue, totalUptime[Days], "Days")
	ch <- prometheus.MustNewConstMetric(u.Hours, prometheus.GaugeValue, totalUptime[Hours], "Hours")
	ch <- prometheus.MustNewConstMetric(u.Minutes, prometheus.GaugeValue, totalUptime[Minutes], "Minutes")
}

func Getuptime() []float64 {
	os_uptime, _ := host.Uptime()

	var days uint64 = os_uptime / (60 * 60 * 24)
	var hours uint64 = (os_uptime - (days * 60 * 60 * 24)) / (60 * 60)
	var minutes uint64 = ((os_uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	return []float64{float64(days), float64(hours), float64(minutes)}
}
