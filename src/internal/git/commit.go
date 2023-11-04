package git

import "github.com/LoyalPotato/sg/src/internal/cli"

// Perform a git commit
func Commit(msg string) error {
	return cli.RunCmd("git", "commit", "-m", msg)
}
