package cmds

import (
	"fmt"
	"os/exec"
)

func RunCmd(input ...string) error {
	if len(input) == 0 {
		return fmt.Errorf("command subprocess: no command provided")
	}

	command := exec.Command(input[0], input[1:]...)

	res, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("command subprocess: %w: %s", err, res)
	}

	return nil
}

func Unarchive(path string) error {
	return RunCmd(
		"tar",
		"-xzf",
		path,
		"--strip-components=1",
	)
}

func Sudo() error {
	return RunCmd(
		"sudo",
		"-v",
	)
}

func InstallFile(mode string, target string, dest string) error {
	return RunCmd(
		"install",
		"-m", mode,
		target,
		dest,
	)
}

func InstallFileAsRoot(mode string, target string, dest string) error {
	return RunCmd(
		"sudo",
		"install",
		"-o", "root",
		"-g", "root",
		"-m", mode,
		target,
		dest,
	)
}

func Cls() {
	fmt.Print("\033[H\033[2J")
}
