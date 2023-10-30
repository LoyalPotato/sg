package git

import (
	"github.com/LoyalPotato/stacked-guide/src/cli"
)

func IsGitRepo() error {
	return cli.RunCmd("git", "rev-parse", "--git-dir")
}
