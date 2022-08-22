package hardware_metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/ricochet2200/go-disk-usage/du"
)

type DiskSize struct {
	mnt   *prometheus.Desc
	slash *prometheus.Desc
}

func (d *DiskSize) Describe(ch chan<- *prometheus.Desc) {
	ch <- d.mnt
	ch <- d.slash
}

func (d *DiskSize) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(d.mnt, prometheus.GaugeValue, GetDiskSize("/mnt"), "/mnt")
	ch <- prometheus.MustNewConstMetric(d.slash, prometheus.GaugeValue, GetDiskSize("/"), "/")
}

func NewDisk() *DiskSize {
	return &DiskSize{
		mnt:   prometheus.NewDesc("mnt_size", "Size of the mounted volume", []string{"Path"}, nil),
		slash: prometheus.NewDesc("root_size", "Size of the whole volume", []string{"Path"}, nil),
	}
}

func GetDiskSize(path string) float64 {
	return float64(du.NewDiskUsage(path).Usage()) * 100
}
