package main

import (
	"fmt"
	"kayscript-installer/cmds"
	"kayscript-installer/repo"
	"os"
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

	fmt.Print("Installing...")

	targets, err := repo.NewTargets(workDir)
	if err != nil {
		panic(err)
	}

	for _, file := range targets.All() {
		if err := repo.Install(file); err != nil {
			panic(err)
		}
	}

	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		fmt.Printf("%s\n", entry.Name())
	}
}

func setup() {
	err := os.MkdirAll(ProjectDir, 0o755)
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
