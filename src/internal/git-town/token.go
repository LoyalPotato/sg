package gittown

import (
	"github.com/LoyalPotato/sg/src/internal/git"
)

const TokenKey = "git-town.github-token"

func AddGithubToken(token string, logRun bool) error {
	return git.AddConfig(git.GitConfig{
		Key:    TokenKey,
		Value:  token,
		LogRun: logRun,
	})
}

func RemoveGithubToken() (bool, error) {
	exists := git.ConfigExists(TokenKey)

	if !exists {
		return exists, nil
	}

	return exists, git.UnsetConfig(TokenKey)
}
