package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"net/http"
)

type emr struct {
	cpuUsage uint64
}

var cpuStatus = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "cpu_usage",
		Help: "Percentage of CPU used by the system.",
	},
	[]string{"cpuUsage", "usedCpuMemory"},
)

func init() {
	// we need to register the counter so prometheus can collect this metric
	prometheus.MustRegister(cpuStatus)
	prometheus.Unregister(collectors.NewGoCollector())
	prometheus.Unregister(collectors.NewBuildInfoCollector())
}

func server(w http.ResponseWriter, r *http.Request) {
	var mr emr
	json.NewDecoder(r.Body).Decode(&mr)
	var status string = "OK"
	var user string = string(mr.cpuUsage)

	cpuStatus.WithLabelValues(user, status).Inc()
}

func P(data uint64) {
	postBody, _ := json.Marshal(
		emr{
			cpuUsage: data,
		},
	)
	requestBody := bytes.NewBuffer(postBody)

	fmt.Println(requestBody)
	_, err := http.Post("http://localhost:8080", "application/json", requestBody)
	if err != nil {
		return
	}

}
