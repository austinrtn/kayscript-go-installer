package main

import (
	"log"
	"os"
)

func main() {
	dir, err := os.MkdirTemp("", "kayscript-*")

	if err != nil {
		log.Fatal("Failed to create temporary directory: %w", err)
	}

	os.Chdir(dir)

	DownloadRepo(dir)
}