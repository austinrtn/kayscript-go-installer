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
	option := os.Args[1]

	if option == "--i" || option == "" {
		install()
	} else if option == "--r" {
		uninstall()
	} else {
		panic("Invalid argument.\n--i: install kayscript\n--r: uninstall kayscript")
	}
}

func install() {
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
	fmt.Print("Kayscript Installed!\n")
}

func uninstall() {
	err := cmds.SudoV()
	targets, err := repo.NewTargets("")
	if err != nil {
		panic("Unable to initialize target files")
	}

	fmt.Print("Removing Files...\n")
	for _, file := range targets.All() {
		err := cmds.ForceRmSudo(file.DestPath)
		if err != nil {
			panic(fmt.Sprintf("Unable to remove file: %s", file.Name))
		}
	}
	fmt.Print("Files Removed!\n")
}

func setup() {
	// Need to use sudo 
	err := cmds.MkDirSudo(ProjectDir)
	if err != nil {
		panic(err)
	}
	
	cmds.Cls()
	if err = cmds.SudoV(); err != nil {
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
