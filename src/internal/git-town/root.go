package gittown

import (
	"os"

	"github.com/LoyalPotato/sg/src/internal/cli"
)

const gittownRoot = "git-town"

func runGitTown(args ...string) error {
	return cli.RunCmdCustomStdOut(gittownRoot, os.Stdout, os.Stderr, args...)
}
