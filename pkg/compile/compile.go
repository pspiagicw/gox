package compile

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/resolver"
)

func CompileProject(url string) (string, error) {

	dir := getTemp()

	environments := getEnvironments(dir)

    err := goreland.ExecuteWithoutOutput("go", []string{"install", url}, environments)

	if err != nil {
		return "", fmt.Errorf("Error executing 'go install': %v", err)
	}

	binDir := getBinDir(dir)

    goreland.LogSuccess("Package compiled succesfully!")

	return binDir, nil
}
func startSpinner() func() {
    s := spinner.New(spinner.CharSets[14], 100 * time.Millisecond)
    s.Suffix = " Compiling something awesome! ✨"
    s.Start()
    return func() {
        s.Stop()
    }

}

func getEnvironments(dir string) []string {
	binDir := fmt.Sprintf("GOBIN=%s", getBinDir(dir))
	cacheDir := fmt.Sprintf("GOCACHE=%s", getCacheDir(dir))
	gopath := fmt.Sprintf("GOPATH=%s", dir)
	gomodcache := fmt.Sprintf("GOMODCACHE=%s", getModDir(dir))

	return []string{binDir, cacheDir, gopath, gomodcache, "GIT_SSL_NO_VERIFY=1"}
}

func getTemp() string {
	dir, err := os.MkdirTemp("", "gox-build")

	if err != nil {
		goreland.LogFatal("Error creating temp directory: %v", err)
	}

	return dir
}

func getBinDir(dir string) string {
	binDir := filepath.Join(dir, "bin")
	resolver.EnsureExists(binDir)
	return binDir
}

func getCacheDir(dir string) string {
	cacheDir := filepath.Join(dir, "cache")
	resolver.EnsureExists(cacheDir)
	return cacheDir
}
func getModDir(dir string) string {
	modpath := filepath.Join(dir, "pkg")
	modpath = filepath.Join(modpath, "mod")
	resolver.EnsureExists(modpath)
	return modpath
}
