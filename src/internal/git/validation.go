package git

import (
	"github.com/LoyalPotato/sg/src/internal/cli"
)

// Checks if current folder is a git repository
func IsGitRepo() error {
	return cli.RunCmd("git", "rev-parse", "--git-dir")
}
