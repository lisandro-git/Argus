package client

import (
	"code.gitea.io/sdk/gitea"
	"time"
)

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

func GetRepos(client *gitea.Client) *[]ParsedRepos {
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

func GetRepoCount(client *gitea.Client) int {
	result, _, err := client.ListMyRepos(gitea.ListReposOptions{})
	if err != nil {
		return 0
	}
	return len(result)
}
