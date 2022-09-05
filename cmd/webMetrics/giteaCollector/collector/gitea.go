package collector

import (
	"argus/cmd"
	"argus/cmd/webMetrics/giteaCollector/client"
	"code.gitea.io/sdk/gitea"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

type gCollector struct {
	client   *[]client.ParsedRepos
	metrics  map[string]*prometheus.Desc
	upMetric prometheus.Gauge
}

func NewgCollector() *gCollector {
	return &gCollector{
		metrics: map[string]*prometheus.Desc{
			"Repository": prometheus.NewDesc("gitea_repository", "Repository", []string{"ID", "Name", "Owner", "RepoSize", "Email"}, nil),
			"RepoCount":  prometheus.NewDesc("gitea_repo_count", "RepoCount", nil, nil),
			"HTMLURL":    prometheus.NewDesc("url", "url", []string{"HTMLURL"}, nil),
			"Created":    prometheus.NewDesc("createDate", "cd", []string{"Created"}, nil),
			"Updated":    prometheus.NewDesc("lastUpdate", "lu", []string{"Updated"}, nil),
		},
		upMetric: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "gitea_up",
			Help: "Gitea service is up",
		}),
	}
}

func (g *gCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- g.upMetric.Desc()

	for _, m := range g.metrics {
		ch <- m
	}
}

func (g *gCollector) Collect(ch chan<- prometheus.Metric) {
	c, err := gitea.NewClient(cmd.GiteaLink, gitea.SetToken(cmd.GiteaApiToken))
	if err != nil {
		g.upMetric.Set(1)
		ch <- g.upMetric
		panic(err)
	}
	g.upMetric.Set(1)
	ch <- g.upMetric

	var repoCount float64
	g.client, repoCount = client.GetRepos(c)
	ch <- prometheus.MustNewConstMetric(g.metrics["RepoCount"], prometheus.GaugeValue, repoCount)

	for _, repo := range *g.client {
		ch <- prometheus.MustNewConstMetric(g.metrics["Repository"],
			prometheus.CounterValue,
			0,
			strconv.Itoa(int(repo.ID)),
			repo.Name,
			repo.Owner,
			strconv.Itoa(repo.RepoSize),
			repo.Email,
		)
	}

}
