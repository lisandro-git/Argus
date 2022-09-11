package collector

import (
	"argus/cmd"
	"github.com/eko/pihole-exporter/cmd/metrics"
)

func Start() {
	cmd.RegisterCollector(metrics.DomainsBlocked)
	cmd.RegisterCollector(metrics.DNSQueriesToday)
	cmd.RegisterCollector(metrics.AdsBlockedToday)
	cmd.RegisterCollector(metrics.AdsPercentageToday)
	cmd.RegisterCollector(metrics.UniqueDomains)
	cmd.RegisterCollector(metrics.QueriesForwarded)
	cmd.RegisterCollector(metrics.QueriesCached)
	cmd.RegisterCollector(metrics.ClientsEverSeen)
	cmd.RegisterCollector(metrics.UniqueClients)
	cmd.RegisterCollector(metrics.DNSQueriesAllTypes)
	cmd.RegisterCollector(metrics.Reply)
	cmd.RegisterCollector(metrics.TopQueries)
	cmd.RegisterCollector(metrics.TopAds)
	cmd.RegisterCollector(metrics.TopSources)
	cmd.RegisterCollector(metrics.ForwardDestinations)
	cmd.RegisterCollector(metrics.QueryTypes)
	cmd.RegisterCollector(metrics.Status)
}
