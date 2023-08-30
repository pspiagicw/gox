package installer

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/golang-groom/database"
	"github.com/pspiagicw/goreland"
)

const INSTALL_LOCATION = "/home/pspiagicw/.local/share/groom/bin"
const DATABASE_LOCATION = "/home/pspiagicw/.local/share/groom/db"

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
func checkIfAlreadyExists(name string) {
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
func InstallBinaries(dir string, entries []fs.DirEntry) {
	for _, entry := range entries {
		checkIfAlreadyExists(entry.Name())
		goreland.LogInfo("Installing '%s'", entry.Name())
		newLocation, err := installFile(dir, entry.Name())
		if err != nil {
			goreland.LogError("Error installing '%s': %v", entry.Name(), err)
		}
		updateDatabase(entry.Name(), newLocation)
        goreland.LogSuccess("Installation of '%s' successful!", entry.Name())
        goreland.LogSuccess("'%s' was installed at %s", entry.Name(), newLocation)
        goreland.LogSuccess("You don't have to worry about adding it to your PATH variable, that has been already done")
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
