package githubCollector

import (
	"argus/cmd"
	"argus/cmd/webMetrics/githubCollector/collector"
)

// RepositoryInfo is a struct that contains all the information about a repository
var RepositoryInfo = []string{
	"ID",
	"RepoName",
	"Owner",
	"RepoSize",
	"DefaultBranch",
	"CloneURL",
	"MainLanguage",
	"Description",
	"Visibility",
}

func Start() {
	cmd.RegisterCollector(collector.NewgCollector())
}
