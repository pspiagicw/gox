package install

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/compile"
	"github.com/pspiagicw/gox/pkg/database"
	"github.com/pspiagicw/gox/pkg/help"
	"github.com/pspiagicw/gox/pkg/resolver"
)

func updateDatabase(filename string, filepath string, url string) {

	err := database.AddPackage(database.Package{
		Name:      filename,
		Path:      filepath,
		Installed: time.Now(),
		URL:       url,
	})

	if err != nil {
		goreland.LogError("Error updating database: %v", err)
		goreland.LogError("Try installing again!")
		os.Exit(1)
	}
}

func checkReinstall(name string) {
	db := database.ParseDatabase()

	for key := range db.Packages {
		if key == name {
			goreland.LogInfo("Package %s already installed, reinstalling...", name)
		}
	}
}
func installBinary(dir string, entry fs.DirEntry, url string) {

	checkReinstall(entry.Name())
	goreland.LogInfo("Installing '%s'", entry.Name())

	newLocation, err := installFile(dir, entry.Name())

	if err != nil {
		goreland.LogError("Error installing '%s': %v", entry.Name(), err)
	}

	updateDatabase(entry.Name(), newLocation, url)

	addSymlink(entry.Name(), newLocation)

	installSuccessful(entry, newLocation)
}
func installSuccessful(entry fs.DirEntry, location string) {
	goreland.LogSuccess("Installation of '%s' successful!", entry.Name())
	goreland.LogSuccess("'%s' was installed at %s", entry.Name(), location)
	goreland.LogSuccess("You don't have to worry about adding it to your PATH variable, that has been already done")
}

func addSymlink(name string, location string) {
	fakeLocation := filepath.Join(resolver.BinDir(), name)
	goreland.LogInfo("Symlinking binary into %s", fakeLocation)
	removeIfExists(fakeLocation)
	err := os.Symlink(location, fakeLocation)
	if err != nil {
		goreland.LogFatal("Error symlinking binary: %v", err)
	}
}

func removeIfExists(file string) {
	_, err := os.Stat(file)
	if err == nil {
		err := os.Remove(file)
		if err != nil {
			goreland.LogError("File %s exists, error removing it: %v", file, err)
		}
	} else if errors.Is(err, fs.ErrNotExist) == false {
		goreland.LogError("Error stating %s: %v", file, err)
	}
}

func installFile(dir, filename string) (string, error) {
	newLocation := filepath.Join(resolver.InstallDir(), filename)
	oldLocation := filepath.Join(dir, filename)
	removeIfExists(newLocation)

	options := &goreland.InstallFileOptions{
		CreateDir:   true,
		Permissions: 0755,
	}

	err := goreland.InstallFile(oldLocation, resolver.InstallDir(), options)
	if err != nil {
		return "", fmt.Errorf("Error Installing Binary: %v", err)
	}
	return newLocation, nil
}

func parseInstallFlags(args []string) string {
	flag := flag.NewFlagSet("gox install", flag.ExitOnError)

    flag.Usage = func() {
    }
	flag.Parse(args)
	args = flag.Args()

	if len(args) == 0 {
		help.HelpInstall()
		goreland.LogFatal("No URL/Path provided.")
	}

	url := args[0]

	return url
}
func getCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		goreland.LogFatal("Error can't resolve current directory: %v", err)
	}
	return dir

}
func InstallPackage(args []string) {
	url := parseInstallFlags(args)

	if url == "." {
		url = getCurrentDir()
	}

	binDir, err := compile.CompileProject(url)
	if err != nil {
		goreland.LogFatal("Error building the project: %v", err)
	}

	binary := getBinary(binDir)
	installBinary(binDir, binary, url)
}

func getBinary(binDir string) fs.DirEntry {
	entries, err := os.ReadDir(binDir)

	if err != nil {
		goreland.LogFatal("Error reading the build directory: %v", err)
	}

	if len(entries) != 1 {
		goreland.LogFatal("Binary not found!")
	}
	return entries[0]
}
