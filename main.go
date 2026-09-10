package main

import (
	"fmt"
	"kayscript-installer/cmds"
	"kayscript-installer/repo"
	"os"
	"path"
	"path/filepath"
)

const ProjectDir string = "/var/lib/kayscript"
var workDir string

func main() {
	setup()
	defer deinit()

	err := os.Chdir(workDir)
	if err != nil {
		panic(err)
	}

	fmt.Print("Downloading...\n")
	if err = repo.DownloadRepo(); err != nil {
		panic(err)
	}

	fmt.Print("Extracting...\n")
	kayscript_dest_path, err := filepath.Abs("kayscript.tar.gz")
	if err != nil {
		panic(err)
	}
	if err := cmds.Unarchive((kayscript_dest_path)); err != nil {
		panic(err)
	}

	fmt.Print("Installing...\n")
	os.Chdir(path.Join(workDir, "install_files"))
	targets, err := repo.NewTargets(workDir)
	if err != nil {
		panic(err)
	}

	for _, file := range targets.All() {
		if err := repo.Install(file); err != nil {
			panic(err)
		}
	}
}

func setup() {
	// Need to use sudo 
	err := cmds.MkDirSudo(ProjectDir)
	if err != nil {
		panic(err)
	}
	
	cmds.Cls()
	if err = cmds.Sudo(); err != nil {
		panic(err)
	}

	workDir, err = os.MkdirTemp("", "kayscript-*")
	if err != nil {
		panic(err)
	}
}

func deinit() {
	os.RemoveAll(workDir)
}

// func replaceFilePlaceholders() {
// 	fmt.Append()
// }
