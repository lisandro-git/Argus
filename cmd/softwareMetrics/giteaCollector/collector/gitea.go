package collector

import (
	"argus/cmd"
	"code.gitea.io/sdk/gitea"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"time"
)

type gCollector struct {
	client  *[]ParsedRepos
	metrics map[string]*prometheus.Desc
	// lisandro : add a metric to check if the process is up or down ?
}

func NewgCollector() *gCollector {
	return &gCollector{
		metrics: map[string]*prometheus.Desc{
			"Repository": prometheus.NewDesc("gitea_repository", "Repository", []string{"ID", "Name", "Owner", "RepoSize", "Email"}, nil),
			"HTMLURL":    prometheus.NewDesc("url", "url", []string{"HTMLURL"}, nil),
			"Created":    prometheus.NewDesc("createDate", "cd", []string{"Created"}, nil),
			"Updated":    prometheus.NewDesc("lastUpdate", "lu", []string{"Updated"}, nil),
		},
	}
}

func (g *gCollector) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range g.metrics {
		ch <- m
	}
}

type ParsedRepos struct {
	ID       int64
	Name     string
	RepoSize int
	Owner    string
	Email    string
	HTMLURL  string
	Created  time.Time
	Updated  time.Time
}

func getRepos(client *gitea.Client) *[]ParsedRepos {
	var parsedRepo []ParsedRepos
	result, _, err := client.ListMyRepos(gitea.ListReposOptions{})
	if err != nil {
		return &parsedRepo
	}
	for _, x := range result {
		parsedRepo = append(parsedRepo, ParsedRepos{
			ID:       x.ID,
			Name:     x.Name,
			RepoSize: x.Size,
			Owner:    x.Owner.UserName,
			Email:    x.Owner.Email,
			HTMLURL:  x.HTMLURL,
			Created:  x.Created,
		})
	}
	return &parsedRepo
}

func (g *gCollector) Collect(ch chan<- prometheus.Metric) {
	client, err := gitea.NewClient(cmd.GiteaLink, gitea.SetToken(cmd.GiteaApiToken))
	if err != nil {
		panic(err)
	}
	g.client = getRepos(client)
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
