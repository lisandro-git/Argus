package piholeCollector

import (
	"argus/cmd/webMetrics/piholeCollector/collector"
	"github.com/eko/pihole-exporter/cmd/config"
	"github.com/eko/pihole-exporter/cmd/pihole"
	"log"
)

func S() {
	envConf, clientConfigs, err := config.Load()
	if err != nil {
		log.Fatal(err.Error())
	}

	collector.Start()

	//serverDead := make(chan struct{})

	clients := buildClients(clientConfigs, envConf)

	for _, client := range clients {
		go client.CollectMetricsAsync()
	}

	for _, client := range clients {
		status := <-client.Status
		if status.Status == pihole.MetricsCollectionError {
			log.Printf("An error occured while contacting %s: %s", client.GetHostname(), status.Err.Error())
		}
	}

	//s := server.NewServer(envConf.Port, clients)
	//go func() {
	//	s.ListenAndServe()
	//	close(serverDead)
	//}()

}

func buildClients(clientConfigs []config.Config, envConfig *config.EnvConfig) []*pihole.Client {
	clients := make([]*pihole.Client, 0, len(clientConfigs))
	for i := range clientConfigs {
		clientConfig := &clientConfigs[i]

		client := pihole.NewClient(clientConfig, envConfig)
		clients = append(clients, client)
	}
	return clients
}
