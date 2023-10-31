package gittown


const tokenKey = "git-town.github-token"

func AddGithubToken(token string, logRun bool) error {
	return git.AddConfig(git.GitConfig{
		Key:    tokenKey,
		Value:  token,
		LogRun: logRun,
	})
}

func RemoveGithubToken() (bool, error) {
	exists := git.ConfigExists(tokenKey)

	if !exists {
		return exists, nil
	}

	return exists, git.UnsetConfig(tokenKey)
}
