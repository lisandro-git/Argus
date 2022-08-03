package cmd

import (
	"github.com/ricochet2200/go-disk-usage/du"
)

func GetDiskSize(path string) float64 {
	return float64(du.NewDiskUsage(path).Usage()) * 100
}
