package git

import (
	"fmt"

	"github.com/LoyalPotato/sg/src/internal/cli"
	"github.com/LoyalPotato/sg/src/internal/styles"
	"github.com/LoyalPotato/sg/src/internal/utils"
)

type Config struct {
	Key, Value string
	LogRun     bool
}

// Add item to local git config.
func AddConfig(gitConfig Config) error {
	args := []string{"config", "--add", gitConfig.Key, gitConfig.Value}
	if gitConfig.LogRun {
		fmt.Printf("Inserting %s into config\n", styles.Faint(args[2]))
	}

	return cli.RunCmd("git", args...)
}

// Get local single git config.
func GetConfig(key string) (string, error) {
	out, err := cli.RunCmdWithOutput("git", "config", "--get", key)

	return string(out), err
}

// Check if local git config exists.
func ConfigExists(key string) bool {
	_, err := cli.RunCmdWithOutput("git", "config", "--get", key)

	return utils.GetExitCode(err) != 1
}

// Delete local git config.
func UnsetConfig(key string) error {
	return cli.RunCmd("git", "config", "--unset-all", key)
}
