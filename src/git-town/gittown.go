package gittown

import "github.com/LoyalPotato/stacked-guide/src/git"

func AddGithubToken(token string, logRun bool) error {
	return git.AddConfig(git.GitConfig{
		Key:    "git-town.github-token",
		Value:  token,
		LogRun: logRun,
	})
}
