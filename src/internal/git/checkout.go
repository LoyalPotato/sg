package git

import "github.com/LoyalPotato/sg/src/internal/cli"

func Checkout(branch string) error {
	return cli.RunCmd("git", "checkout", branch)
}
