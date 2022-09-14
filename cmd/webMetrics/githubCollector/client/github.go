package client

import (
	"argus/cmd"
	"context"
	"github.com/google/go-github/v47/github"
	"golang.org/x/oauth2"
)

// GetGithubRepos returns a list of all repositories for the authenticated user.
func GetGithubRepos() []*github.Repository {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: cmd.GithubToken,
		},
	)

	client := github.NewClient(oauth2.NewClient(ctx, ts))

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, cmd.GithubUsername, nil)
	if err != nil {
		var repositories []*github.Repository
		return repositories
	}
	return repos
}
