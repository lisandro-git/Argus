package collector

import (
	"argus/cmd/webMetrics/githubCollector/client"
	"github.com/google/go-github/v47/github"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
)

// gCollector is a custom collector for GitHub metrics
type gCollector struct {
	metrics map[string]*prometheus.Desc
}

// NewgCollector returns a new gCollector
func NewgCollector() *gCollector {
	return &gCollector{
		metrics: map[string]*prometheus.Desc{
			"Repository": prometheus.NewDesc(
				"github_repository",
				"Repository",
				[]string{
					"ID",
					"RepoName",
					"Owner",
					"RepoSize",
					"DefaultBranch",
					"CloneURL",
					"MainLanguage",
					"Description",
					"Visibility",
					"Created",
					"Updated",
				},
				nil,
			),
			"RepoCount": prometheus.NewDesc("github_repo_count", "RepoCount", nil, nil),
		},
	}
}

func (g *gCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range g.metrics {
		ch <- m
	}
}

func (g *gCollector) Collect(ch chan<- prometheus.Metric) {
	// Getting the repositories
	var ghRepo []*github.Repository = client.GetGithubRepos()

	// Parsing and sending the metrics to prometheus
	for _, repo := range ghRepo {
		ch <- prometheus.MustNewConstMetric(
			g.metrics["Repository"],
			prometheus.GaugeValue,
			1,
			strconv.Itoa(int(repo.GetID())),
			repo.GetName(),
			repo.Owner.GetLogin(),
			strconv.Itoa(repo.GetSize()),
			repo.GetDefaultBranch(),
			repo.GetCloneURL(),
			repo.GetLanguage(),
			repo.GetDescription(),
			repo.GetVisibility(),
			repo.GetCreatedAt().String(),
			repo.GetUpdatedAt().String(),
		)
	}
	// Sending the repo count
	ch <- prometheus.MustNewConstMetric(
		g.metrics["RepoCount"],
		prometheus.GaugeValue,
		float64(len(ghRepo)),
	)
}
