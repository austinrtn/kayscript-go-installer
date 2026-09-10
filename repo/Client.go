package repo

import (
	"fmt"
	"io"
	"os"
	"net/http"
	"time"
)
const RepoUrl string = "https://github.com/austinrtn/KayScript/archive/refs/tags/1.0.tar.gz"

func DownloadRepo() error {
	client := &http.Client{Timeout: 30 * time.Second}

	res, err := client.Get(RepoUrl)

	if err != nil {
		return fmt.Errorf("download repositry :%w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("download repo: server returned %s", res.Status)
	}

	file, err := os.Create("kayscript.tar.gz")
	if err != nil {
		return fmt.Errorf("download repo: unable to create file: %w", err)
	}
	defer file.Close() 

	if _, err := io.Copy(file, res.Body); nil != err {
		return fmt.Errorf("download repo: unable to save download contents: %w", err)
	}
	
	return nil
}