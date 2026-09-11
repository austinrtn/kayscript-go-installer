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
	tmpPath := path.Join(workDir, "install_files", name)
	return File{
		Name:      name,
		DestPath:  destPath,
		Mode:      mode,
		RootOwned: rootOwned,
		TmpPath: tmpPath,
	}
}

func (file File) ReplaceFileText(old string, new string) error {
	txt, err := os.ReadFile(file.TmpPath)
	if err != nil {
		return fmt.Errorf("read %q: %w", file.Name, err)
	}

	new_txt := strings.ReplaceAll(string(txt), old, new)
	err = os.WriteFile(file.TmpPath, []byte(new_txt), file.Mode)

	if err != nil {
		return fmt.Errorf("write %q: %w", file.DestPath, err)
	}

	return nil
}

func (file File)Install() error {
	var err error = nil
	mode := fmt.Sprintf("%04o", file.Mode.Perm())

	if file.RootOwned {
		err = cmds.InstallFileAsRoot(mode, file.TmpPath, file.DestPath)
	} else {
		err = cmds.InstallFile(mode, file.TmpPath, file.DestPath)
	}

	if err != nil {
		return fmt.Errorf("installing file: %w", err)
	}
	return nil
}

func (file File)Exists() (bool, error) {
	exists, err := cmds.TestSudo(file.DestPath)
	
	if err != nil {
		return false, fmt.Errorf("Error checking file status: %w", err)
	}
	
	return exists, nil 
}

func GetHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("find user home directory: %w", err)
	}
	return home, nil
}
