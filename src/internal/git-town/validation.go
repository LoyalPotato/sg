package gittown

import (
	"os/exec"
)

func IsInstalled() bool {
	path, err := exec.LookPath("git-town")
	return path != "" && err == nil
}
