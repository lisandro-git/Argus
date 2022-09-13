package osMetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/host"
)

// Uptime is a struct that contains the uptime of the system
type Uptime struct {
	Time *prometheus.Desc
}

var (
	Days    = 0
	Hours   = 1
	Minutes = 2
)

// NewUptime returns a new Uptime collector
func NewUptime() *Uptime {
	return &Uptime{
		Time: prometheus.NewDesc("uptime", "Uptime of the system", []string{"uptime"}, nil),
	}
}

func (u *Uptime) Describe(ch chan<- *prometheus.Desc) {
	ch <- u.Time
}

func (u *Uptime) Collect(ch chan<- prometheus.Metric) {
	totalUptime := Getuptime()
	ch <- prometheus.MustNewConstMetric(u.Time, prometheus.CounterValue, totalUptime[Days], "Days")
	ch <- prometheus.MustNewConstMetric(u.Time, prometheus.CounterValue, totalUptime[Hours], "Hours")
	ch <- prometheus.MustNewConstMetric(u.Time, prometheus.CounterValue, totalUptime[Minutes], "Minutes")
}

// Getuptime returns the uptime of the system
func Getuptime() []float64 {
	os_uptime, _ := host.Uptime()

	var days uint64 = os_uptime / (60 * 60 * 24)
	var hours uint64 = (os_uptime - (days * 60 * 60 * 24)) / (60 * 60)
	var minutes uint64 = ((os_uptime - (days * 60 * 60 * 24)) - (hours * 60 * 60)) / 60

	return []float64{float64(days), float64(hours), float64(minutes)}
}
