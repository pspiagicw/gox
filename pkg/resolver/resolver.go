package resolver

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pspiagicw/goreland"
)

func InstallDir() string {
	b := filepath.Join(DataDir(), "bin")
	EnsureExists(b)
	return b
}

func HomeDir() string {
	return xdg.Home
}

func BinDir() string {
	d := filepath.Join(HomeDir(), ".local")
	d = filepath.Join(d, "bin")

	EnsureExists(d)
	return d
}

func DataDir() string {
	location := filepath.Join(xdg.DataHome, "gox")
	EnsureExists(location)
	return location
}

func EnsureExists(location string) {
	if !dirExists(location) {
		err := os.MkdirAll(location, 0755)
		if err != nil {
			goreland.LogFatal("Error creating directory: %s, %v", location, err)
		}
	}
}

func DatabasePath() string {

	d := DataDir()
	d = filepath.Join(d, "db")

	goreland.LogInfo("Using %s for database", d)
	return d

}
func dirExists(dir string) bool {
	_, err := os.Stat(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return false
	} else if err != nil {
		goreland.LogFatal("Error stating directory: %v", err)
	}
	return true
}
