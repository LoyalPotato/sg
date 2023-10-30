package git

import (
	"fmt"

	"github.com/LoyalPotato/stacked-guide/src/cli"
	"github.com/LoyalPotato/stacked-guide/src/styles"
)

type GitConfig struct {
	Key, Value string
	LogRun     bool
}

func AddConfig(gitConfig GitConfig) error {
	args := []string{"config", "--add", gitConfig.Key, gitConfig.Value}
	if gitConfig.LogRun {
		fmt.Printf("Inserting %s into config\n", styles.FaintItalic(args[2]))
	}

	return cli.RunCmd("git", args...)
}
