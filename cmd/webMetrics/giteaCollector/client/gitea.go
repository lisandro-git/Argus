package client

import (
	"code.gitea.io/sdk/gitea"
	"time"
)

// ParsedRepos is a struct that contains the data that we want to send to prometheus
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

// GetRepos returns a list of all the repos in the gitea instance
func GetRepos(client *gitea.Client) (*[]ParsedRepos, float64) {
	var parsedRepo []ParsedRepos
	result, _, err := client.ListMyRepos(gitea.ListReposOptions{})
	if err != nil {
		return &parsedRepo, 0.0
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
	return &parsedRepo, float64(len(parsedRepo))
}

func GetRepoCount(client *gitea.Client) int {
	result, _, err := client.ListMyRepos(gitea.ListReposOptions{})
	if err != nil {
		return 0
	}
	return len(result)
}
