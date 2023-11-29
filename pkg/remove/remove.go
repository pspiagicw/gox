package remove

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/database"
	"github.com/pspiagicw/gox/pkg/help"
	"github.com/pspiagicw/gox/pkg/resolver"
)

func parseRemoveFlags(args []string) string {
	flag := flag.NewFlagSet("gox remove", flag.ExitOnError)

	flag.Parse(args)
	args = flag.Args()

	if len(args) == 0 {
		help.HelpRemove()
		goreland.LogFatal("No package name provided.")

	}

	return args[0]
}

func RemovePackage(args []string) {
	name := parseRemoveFlags(args)
	// TODO: Implement Removal of pacakge.
	// goreland.LogFatal("NOT IMPLELMENTED YET!")
	db := database.ParseDatabase()

	entry, exists := db.Packages[name]

	if !exists {
		goreland.LogFatal("Package is not installed!")
	}

	binLocation := getBinDir(entry.Name)

	removeFile(binLocation)

	goreland.LogInfo("Deleted '%s'", binLocation)

	originalPath := entry.Path

	removeFile(originalPath)

	goreland.LogInfo("Deleted '%s'", originalPath)

	database.RemovePackage(entry)

	goreland.LogInfo("Entry deleted from database")

	goreland.LogSuccess("Package deleted!")

}
func getBinDir(name string) string {
	binDir := resolver.BinDir()

	return filepath.Join(binDir, name)
}

func removeFile(filepath string) {
	err := os.Remove(filepath)
	if err != nil {
		goreland.LogFatal("Error deleting '%s': %v", filepath, err)
	}
}
