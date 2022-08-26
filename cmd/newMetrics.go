package cmd

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	Registry *prometheus.Registry = NewRegisterer()
	Gatherer prometheus.Gatherer  = NewGatherer(Registry)
)

func NewGaugeVec(name, help string, labels []string) *prometheus.GaugeVec {
	return prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	}, labels)
}

func NewGauge(name, help string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
}

func NewRegisterer() *prometheus.Registry {
	return prometheus.NewPedanticRegistry()
}

func NewGatherer(reg *prometheus.Registry) prometheus.Gatherer {
	return prometheus.Gatherers{reg}
}

func RegisterCollector(C interface{ prometheus.Collector }) {
	Registry.MustRegister(C)
}
