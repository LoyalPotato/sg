package cli

import "os/exec"

func RunCmd(cmd string, args ...string) error {
	subProcess := exec.Command(cmd, args...)
	return subProcess.Run()
}

func RunCmdWithOutput(cmd string, args ...string) ([]byte, error) {
	subProcess := exec.Command(cmd, args...)
	return subProcess.Output()
}
