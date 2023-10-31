package git

import (
	"github.com/LoyalPotato/sg/src/cli"
)

func IsGitRepo() error {
	return cli.RunCmd("git", "rev-parse", "--git-dir")
}
