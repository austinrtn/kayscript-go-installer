package repo

import (
	"fmt"
	"kayscript-installer/cmds"
	"os"
	"path"
	"strings"
)

type File struct {
	Name      string
	DestPath  string
	Mode      os.FileMode
	RootOwned bool
	TmpPath   string
}

func NewFile(name string, destPath string, mode os.FileMode, rootOwned bool, workDir string) File {
	tmpPath := path.Join(workDir, name)
	return File{
		Name:      name,
		DestPath:  destPath,
		Mode:      mode,
		RootOwned: rootOwned,
		TmpPath: tmpPath,
	}
}

func ReplaceFileText(file File, old string, new string) error {
	txt, err := os.ReadFile(file.TmpPath)
	if err != nil {
		return fmt.Errorf("read %q: %w", file.Name, err)
	}

	new_txt := strings.ReplaceAll(string(txt), old, new)
	err = os.WriteFile(file.DestPath, []byte(new_txt), file.Mode)

	if err != nil {
		return fmt.Errorf("write %q: %w", file.DestPath, err)
	}

	return nil
}

func Install(file File) error {
	var err error = nil

	if file.RootOwned {
		err = cmds.InstallFileAsRoot(file.Mode.String(), file.TmpPath, file.DestPath)
	} else {
		err = cmds.InstallFile(file.Mode.String(), file.TmpPath, file.DestPath)
	}

	if err != nil {
		return fmt.Errorf("installing file: %w", err)
	}
	return nil
}

func GetHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("find user home directory: %w", err)
	}
	return home, nil
}
