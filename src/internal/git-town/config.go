package gittown

import (
	"github.com/LoyalPotato/sg/src/internal/git"
)

const MainBranchKey = "git-town.main-branch-name"

func AddMainBranchConfig(brachName string, logRun bool) error {
	return git.AddConfig(git.GitConfig{
		Key:    MainBranchKey,
		Value:  brachName,
		LogRun: logRun,
	})
}

func GetMainBranchConfig() (string, error) {
	return git.GetConfig(MainBranchKey)
}

func RemoveMainBranchConfig() (bool, error) {
	exists := git.ConfigExists(MainBranchKey)

	if !exists {
		return exists, nil
	}

	return exists, git.UnsetConfig(MainBranchKey)
}
