package hardwareMetrics

import (
	"argus/cmd"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/disk"
)

type DiskSize struct {
	total *prometheus.Desc
}

func NewDiskSize() *DiskSize {
	return &DiskSize{
		total: prometheus.NewDesc("hdd_size", "Size of a partition given a path and a usage (total, Usage...)", []string{"Path", "Usage"}, nil),
	}
}

func (d *DiskSize) Describe(ch chan<- *prometheus.Desc) {
	ch <- d.total
}

func (d *DiskSize) Collect(ch chan<- prometheus.Metric) {
	parts, _ := disk.Partitions(true)
	for _, p := range parts {
		device := p.Mountpoint
		s, _ := disk.Usage(device)
		if s.Total == 0 {
			continue
		}

		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, float64(s.Total)/cmd.GB, s.Path, "total")
		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, float64(s.Free)/cmd.GB, s.Path, "Free")
		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, float64(s.Used)/cmd.GB, s.Path, "Used")
		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, s.UsedPercent, s.Path, "Percentage")

	}
}
