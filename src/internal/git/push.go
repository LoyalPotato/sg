package git

import "github.com/LoyalPotato/sg/src/internal/cli"

func Push() error {
	return cli.RunCmd("git", "push")
}
