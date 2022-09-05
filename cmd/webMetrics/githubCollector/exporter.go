package githubCollector

import (
	"argus/cmd"
	"argus/cmd/webMetrics/githubCollector/collector"
)

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
