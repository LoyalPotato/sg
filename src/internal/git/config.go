package git

import (
	"fmt"

	"github.com/LoyalPotato/sg/src/cli"
	"github.com/LoyalPotato/sg/src/internal/styles"
	"github.com/LoyalPotato/sg/src/internal/utils"
)

type GitConfig struct {
	Key, Value string
	LogRun     bool
}

func AddConfig(gitConfig GitConfig) error {
	args := []string{"config", "--add", gitConfig.Key, gitConfig.Value}
	if gitConfig.LogRun {
		fmt.Printf("Inserting %s into config\n", styles.Faint(args[2]))
	}

	return cli.RunCmd("git", args...)
}

func GetConfig(key string) (string, error) {
	out, err := cli.RunCmdWithOutput("git", "config", "--get", key)

	return string(out), err
}

func ConfigExists(key string) bool {
	_, err := cli.RunCmdWithOutput("git", "config", "--get", key)

	return utils.GetExitCode(err) != 1
}

func UnsetConfig(key string) error {
	return cli.RunCmd("git", "config", "--unset-all", key)
}
