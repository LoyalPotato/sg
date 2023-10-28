package git

import (
	"os/exec"
)

func IsGitRepo() error {
	subProcess := exec.Command("git", "rev-parse", "--git-dir")
	return subProcess.Run()
}
