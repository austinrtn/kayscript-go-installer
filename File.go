package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const RepoUrl string = ""
const ProjectDir string = "/var/lib/kayscript"

type File struct {
	Name string 
	DestPath string 
	Mode os.FileMode
	RootOwned bool 
	TmpPath string 
}

func GetHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Could not find user home directory.")
	}
	return home
}

func DownloadRepo(dir string) error {
	client := &http.Client{Timeout: 30 * time.Second}

	res, err := client.Get(RepoUrl)

	if err != nil {
		return fmt.Errorf("download repositry :%w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("download repo: server returned %s", res.Status)
	}
	
	return nil
}

func NewFile(name string, destPath string, mode os.FileMode, rootOwned bool) File {
	return File {
		Name: name, 
		DestPath: destPath,
		Mode: mode,
		RootOwned: rootOwned,
	}
}

func ReplaceFileText(file File, old string, new string) {
	txt, err := os.ReadFile(file.TmpPath)
	if err != nil {
		log.Fatal(fmt.Errorf("Unable to read file: %s\n%v", file.Name, err))
	}

	new_txt := strings.ReplaceAll(string(txt), old, new)
	err = os.WriteFile(file.DestPath, []byte(new_txt), file.Mode)

	if err != nil {
		log.Fatal(fmt.Errorf("Unable to write to destination: %s\n%v", file.DestPath, err))
	}
}

var SudoersRule = NewFile(
	"sudoers_rule", 
	"/etc/sudoers.d/kayscript-bypass", 
	0o644,
	true,
)

var RootService = NewFile(
	"root_service", 
	"/etc/systemd/system/kayscript-usb.service", 
	0o644,
	true,
)

var UserService = NewFile(
	"user_service", 
	filepath.Join(GetHomeDir(), ".config", "systemd", "user", "kayscript.service"),
	0o644,
	false,
)

var Kayscript = NewFile(
	"kayscript", 
	filepath.Join(GetHomeDir(), "KayScript"),
	0o755,
	true,
)

var Files = []File{ SudoersRule, RootService, UserService, Kayscript }