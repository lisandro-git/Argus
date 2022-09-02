package giteaCollector

import (
	"argus/cmd"
	"argus/cmd/softwareMetrics/giteaCollector/collector"
	"code.gitea.io/sdk/gitea"
	"fmt"
	"time"
)

var apiToken = "a81b5d25ba82d21309e85cd18f62e9bba74b6f64"

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

func displayStarredRepos(repos []*gitea.Repository) {
	for r, re := range repos {
		fmt.Println(r, re)
	}
}

func getRepoCount(client *gitea.Client) int {
	result, _, err := client.ListMyRepos(gitea.ListReposOptions{})
	if err != nil {
		return 0
	}
	return len(result)
}

func getRepos(client *gitea.Client) []ParsedRepos {
	var parsedRepo []ParsedRepos
	result, _, err := client.ListMyRepos(gitea.ListReposOptions{})
	if err != nil {
		return parsedRepo
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
	return parsedRepo
}

func Start() {
	cmd.RegisterCollector(collector.NewgCollector())

	//displayStarredRepos(repos)
	//fmt.Println(repos, r)

}
