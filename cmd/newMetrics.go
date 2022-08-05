package cmd

import "github.com/prometheus/client_golang/prometheus"

var (
	Register *prometheus.Registry = NewRegisterer()
	Gatherer prometheus.Gatherer  = NewGatherer(Register)
)

type Deck[C any] struct {
	cards []C
}

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

func RegisterGauge[C Deck](gauge C) {
	x := prometheus.Desc{}
	y := prometheus.MetricVec{}
	Register.MustRegister(a)
}
