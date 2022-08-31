package hardwareMetrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ricochet2200/go-disk-usage/du"
)

var path = []string{
	"/",
	"/mnt",
}

type DiskSize struct {
	root *prometheus.Desc
}

func NewDiskSize() *DiskSize {
	return &DiskSize{
		root: prometheus.NewDesc("hdd_size", "Current size of the disk.", []string{"Path"}, nil),
	}
}

func (d *DiskSize) Describe(ch chan<- *prometheus.Desc) {
	ch <- d.root
}

func (d *DiskSize) Collect(ch chan<- prometheus.Metric) {
	for _, p := range path {
		ch <- prometheus.MustNewConstMetric(d.root, prometheus.GaugeValue, GetDiskSize(p), p)
	}
}

func GetDiskSize(path string) float64 {
	return float64(du.NewDiskUsage(path).Usage()) * 100
}
