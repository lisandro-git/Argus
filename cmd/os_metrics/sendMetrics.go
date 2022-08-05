package os_metrics

import "argus/cmd"

func init() {
	cmd.RegisterCollector(uptime)
}

func SendMetrics() {
	day, hour := Getuptime()
	uptime.WithLabelValues("Days").Set(day)
	uptime.WithLabelValues("Hours").Set(hour)
}
