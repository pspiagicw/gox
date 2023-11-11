package install

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-groom/database"
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/compile"
	"github.com/pspiagicw/gox/pkg/help"
)

// TODO: Add directory resolver
const INSTALL_LOCATION = "/home/pspiagicw/.local/share/gox/bin"
const FAKE_INSTALLL = "/home/pspiagicw/.local/bin/"
const DATABASE_LOCATION = "/home/pspiagicw/.local/share/gox/db"

func updateDatabase(filename string, filepath string) {
	err := database.AddPackage(database.Package{
		Name:      filename,
		Path:      filepath,
		Installed: time.Now(),
	})

	if err != nil {
		goreland.LogError("Error updating database: %v", err)
		goreland.LogError("Try installing again!")
		os.Exit(1)
	}
}

func checkReinstall(name string) {
	db, err := database.ParseDatabase()
	if err != nil {
		goreland.LogError("Error reading the database: %v", err)
	}
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

	updateDatabase(entry.Name(), newLocation)
	err = addSymlink(entry.Name(), newLocation)
	if err != nil {
		goreland.LogError("Error symlinking binary: %v", err)
	}

    installSuccessful(entry, newLocation)
}
func installSuccessful(entry fs.DirEntry, location string) {
	goreland.LogSuccess("Installation of '%s' successful!", entry.Name())
	goreland.LogSuccess("'%s' was installed at %s", entry.Name(), location)
	goreland.LogSuccess("You don't have to worry about adding it to your PATH variable, that has been already done")
}

func addSymlink(name string, location string) error {
	fakeLocation := filepath.Join(FAKE_INSTALLL, name)
	removeIfExists(fakeLocation)
	return os.Symlink(location, fakeLocation)
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
	newLocation := filepath.Join(INSTALL_LOCATION, filename)
	oldLocation := filepath.Join(dir, filename)
	removeIfExists(newLocation)

	options := &goreland.InstallFileOptions{
		CreateDir:   true,
		Permissions: 755,
	}

	err := goreland.InstallFile(oldLocation, INSTALL_LOCATION, options)
	if err != nil {
		return "", fmt.Errorf("Error Installing Binary: %v", err)
	}
	return newLocation, nil
}

func parseInstallFlags(args []string) (string) {
	flag := flag.NewFlagSet("gox install", flag.ExitOnError)

	flag.Parse(args)
	args = flag.Args()

	if len(args) == 0 {
		help.HelpInstall()
		goreland.LogFatal("No URL/Path provided.")

	}

    url := args[0]

    return url
}
func InstallPackage(args []string) {
    url := parseInstallFlags(args)

    binDir, err := compile.CompileProject(url)
    if err != nil {
        goreland.LogFatal("Error building the project: %v", err)
    }

    binary := getBinary(binDir)
    installBinary(binDir, binary, url )
}

func getBinary(binDir string) fs.DirEntry {
    entries, err:= os.ReadDir(binDir)

   if err != nil {
        goreland.LogFatal("Error reading the build directory: %v", err)
    }

    if len(entries) != 1 {
        goreland.LogFatal("Binary not found!")
    }
    return entries[0]
}
