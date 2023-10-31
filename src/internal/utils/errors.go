package utils

import (
	"os/exec"
)

func GetExitCode(err error) int {
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}

	return -1
}
