package utils

import (
	"fmt"
	"os"
	"os/exec"
)

// Find more about exit codes here: https://tldp.org/LDP/abs/html/exitcodes.html
const EXIT_CODE_INTERRUPT int = 130

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
