package cmd

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Registry *prometheus.Registry = NewRegisterer()
	Gatherer prometheus.Gatherer  = NewGatherer(Registry)
)

// NewGaugeVec creates a new prometheus.GaugeVec
func NewGaugeVec(name, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labels)
}

// NewGauge creates a new prometheus.Gauge
func NewGauge(name, help string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
}

// NewRegisterer creates a new prometheus.Registerer
func NewRegisterer() *prometheus.Registry {
	return prometheus.NewPedanticRegistry()
}

// NewGatherer creates a new prometheus.Gatherer
func NewGatherer(reg *prometheus.Registry) prometheus.Gatherer {
	return prometheus.Gatherers{reg}
}

// RegisterCollector registers a prometheus.Collector
func RegisterCollector(C interface{ prometheus.Collector }) {
	Registry.MustRegister(C)
}
