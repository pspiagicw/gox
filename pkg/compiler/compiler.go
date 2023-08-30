package compiler

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/pspiagicw/goreland"
)

func CompileProject(url string) (string, []fs.DirEntry, error) {

	dir := getTemp()
	environments := getEnvironments(dir)
	err := goreland.Execute("go", []string{"install", url}, environments)
	if err != nil {
		return "", nil, fmt.Errorf("Error executing 'go install': %v", err)
	}

	binDir := getBinDir(dir)
	entries, err := os.ReadDir(binDir)
	if err != nil {
		return "", nil, fmt.Errorf("Error while reading temp 'bin' directory: %v", err)
	}

	return binDir, entries, nil
}

func getEnvironments(dir string) []string {
	binDir := fmt.Sprintf("GOBIN=%s", getBinDir(dir))
	cacheDir := fmt.Sprintf("GOCACHE=%s", getCacheDir(dir))
	gopath := fmt.Sprintf("GOPATH=%s", dir)
	gomodcache := fmt.Sprintf("GOMODCACHE=%s", getModDir(dir))

	return []string{binDir, cacheDir, gopath, gomodcache, "GIT_SSL_NO_VERIFY=1"}
}
func getTemp() string {
	dir, err := os.MkdirTemp("", "")

	if err != nil {
		goreland.LogError("Error creating temp directory: %v", err)
		os.Exit(1)
	}

	return dir
}
func getBinDir(dir string) string {
	binDir := filepath.Join(dir, "bin")
	createIfNeeded(binDir)
	return binDir
}

func getCacheDir(dir string) string {
	cachedir := filepath.Join(dir, "cache")
	createIfNeeded(cachedir)
	return cachedir
}
func getModDir(dir string) string {
	modpath := filepath.Join(dir, "pkg")
	modpath = filepath.Join(modpath, "mod")
	createIfNeeded(modpath)
	return modpath
}

func dirExists(dir string) bool {
	_, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	} else if err != nil {
		goreland.LogError("Error stating directory: %v", err)
		return false
	}
	return true
}
func createIfNeeded(dir string) {
	if dirExists(dir) == false {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			goreland.LogError("There was a error creating the bin directory: %v", err)
			os.Exit(1)
		}
	}

}
