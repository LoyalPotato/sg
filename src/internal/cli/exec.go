package cli

import (
	"io"
	"os"
	"os/exec"
)

func RunCmd(cmd string, args ...string) error {
	subProcess := exec.Command(cmd, args...)
	subProcess.Stdin = os.Stdin
	return subProcess.Run()
}

func RunCmdCustomStdOut(cmd string, out io.Writer, err io.Writer, args ...string) error {
	subProcess := exec.Command(cmd, args...)
	subProcess.Stdin = os.Stdin
	subProcess.Stdout = out
	subProcess.Stderr = err
	return subProcess.Run()
}

func RunCmdWithOutput(cmd string, args ...string) ([]byte, error) {
	subProcess := exec.Command(cmd, args...)
	return subProcess.Output()
}
