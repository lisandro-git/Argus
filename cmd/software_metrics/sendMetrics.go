package software_metrics

import (
	"argus/cmd"
	"argus/cmd/software_metrics/client"
	"argus/cmd/software_metrics/collector"
	"log"
	"net/http"
)

func RegisterMetrics(x *http.Client, apiEndpoint string) {
	ossClient, err := createClientWithRetries(func() (interface{}, error) {
		return client.NewNginxClient(x, apiEndpoint)
	}, 0, 0)
	if err != nil {
		log.Fatalf("Could not create Nginx Client: %v", err)
	}
	var lab = map[string]string{"nginx_labels: ": "nginx_labels: "}
	cmd.RegisterCollector(collector.NewNginxCollector(ossClient.(*client.NginxClient), "nginx", lab))
}
