package cmds

import (
	"fmt"
	"os/exec"
)

func RunCmd(input ...string) (*exec.Cmd, error) {
	if len(input) == 0 {
		return nil, fmt.Errorf("command subprocess: no command provided")
	}

	command := exec.Command(input[0], input[1:]...)

	res, err := command.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("command subprocess: %w: %s", err, res)
	}
	
	return command, nil
}

func Unarchive(path string) error {
	_, err := RunCmd(
		"tar",
		"-xzf",
		path,
		"--strip-components=1",
	)
	
	return err
}

func SudoV() error {
	_, err := RunCmd(
		"sudo",
		"-v",
	)
	return err
}

func InstallFile(mode string, target string, dest string) error {
	_, err := RunCmd(
		"install",
		"-m", mode,
		target,
		dest,
	)
	return err
}

func InstallFileAsRoot(mode string, target string, dest string) error {
	_, err := RunCmd(
		"sudo",
		"install",
		"-o", "root",
		"-g", "root",
		"-m", mode,
		target,
		dest,
	)
	return err 
}

func MkDirSudo(path string) error {
	_, err := RunCmd(
		"sudo",
		"mkdir",
		"-p",
		path,
	)
	return err
}

func ForceRmSudo(path string) error {
	_, err := RunCmd(
		"sudo",
		"rm",
		"-f", 
		
		path,
	)
	return err 
}

func TestSudo(path string) (bool, error) {
	cmd, err := RunCmd(
		"sudo",
		"test",
		"-f",
		path,
	)

	if err != nil {
		return false, err
	}

	return cmd.ProcessState.ExitCode() == 0, nil
}

func Cls() {
	fmt.Print("\033[H\033[2J")
}