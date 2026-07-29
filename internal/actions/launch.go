package actions

import "os/exec"

func LaunchEXE(path string) error {
	// takes FindEXE output and launches the exe file
	cmd := exec.Command(path)
	return cmd.Start()
}
