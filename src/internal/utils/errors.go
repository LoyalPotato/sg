package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Find more about exit codes here: https://tldp.org/LDP/abs/html/exitcodes.html
const ExitCodeInterrupt int = 130

func GetExitCode(err error) int {
	var exitError *exec.ExitError
	if errors.As(err, &exitError) {
		return exitError.ExitCode()
	}

	return -1
}

func Exit(errMsg string, exitCode int8) {
	fmt.Print(errMsg)
	os.Exit(int(exitCode))
}
