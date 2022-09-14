package hardwareMetrics

import (
	"argus/cmd"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/disk"
)

// DiskSize is a collector for disk size metrics.
type DiskSize struct {
	total *prometheus.Desc
}

// NewDiskSize returns a new DiskSize collector.
func NewDiskSize() *DiskSize {
	return &DiskSize{
		total: prometheus.NewDesc("hdd_size", "Size of a partition given a path and a usage (total, Usage...)", []string{"Path", "Usage"}, nil),
	}
}

func (d *DiskSize) Describe(ch chan<- *prometheus.Desc) {
	ch <- d.total
}

func (d *DiskSize) Collect(ch chan<- prometheus.Metric) {
	// Getting all partitions
	parts, _ := disk.Partitions(true)

	// Iterating through all partitions and sending the needed values to prometheus
	for _, p := range parts {
		device := p.Mountpoint
		s, err := disk.Usage(device)
		if err != nil || s.Total == 0 {
			_ = fmt.Errorf("Error getting disk usage for %s: %s", device, err)
			continue
		}

		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, float64(s.Total)/cmd.GB, s.Path, "total")
		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, float64(s.Free)/cmd.GB, s.Path, "Free")
		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, float64(s.Used)/cmd.GB, s.Path, "Used")
		ch <- prometheus.MustNewConstMetric(d.total, prometheus.GaugeValue, s.UsedPercent, s.Path, "Percentage")

	}
}
