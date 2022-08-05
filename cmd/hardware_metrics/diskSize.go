package hardware_metrics

import (
	"argus/cmd"
	"github.com/ricochet2200/go-disk-usage/du"
)

var hddSize = cmd.NewGaugeVec("hdd_size", "Current size of the HDD.", []string{"Path"})

func GetDiskSize(path string) float64 {
	return float64(du.NewDiskUsage(path).Usage()) * 100
}
