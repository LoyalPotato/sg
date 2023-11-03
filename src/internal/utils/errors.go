package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func GetExitCode(err error) int {
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}

	return -1
}

func Exit(errMsg string, exitCode int8) {
	fmt.Print(errMsg)
	os.Exit(int(exitCode))
}
