package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	nginxUp   = 1
	nginxDown = 0
)

// newGlobalMetric creates a new prometheus.Desc for a global metric.
func newGlobalMetric(namespace string, metricName string, docString string, constLabels map[string]string) *prometheus.Desc {
	return prometheus.NewDesc(namespace+"_"+metricName, docString, nil, constLabels)
}

// newUpMetric creates a new prometheus.Gauge for the up metric.
func newUpMetric(namespace string, constLabels map[string]string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace:   namespace,
		Name:        "up",
		Help:        "Status of the last metric scrape",
		ConstLabels: constLabels,
	})
}
