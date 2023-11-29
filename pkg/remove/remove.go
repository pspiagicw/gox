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

    flag.Usage = help.HelpRemove
	flag.Parse(args)
	args = flag.Args()

	if len(args) == 0 {
		help.HelpRemove()
		goreland.LogFatal("No package name provided.")

	}

	return args[0]
}

func ensurePackageExists(name string) database.Package {
	db := database.ParseDatabase()

	entry, exists := db.Packages[name]

	if !exists {
		goreland.LogFatal("Package '%s' is not installed!", name)
	}

	return entry

}

func removeSymlink(entry database.Package) {
	binLocation := getBinDir(entry.Name)

	removeFile(binLocation)

	goreland.LogInfo("Deleted '%s'", binLocation)

}
func removeBinary(entry database.Package) {
	originalPath := entry.Path

	removeFile(originalPath)

	goreland.LogInfo("Deleted '%s'", originalPath)
}
func removeFromDB(entry database.Package) {
	database.RemovePackage(entry)

	goreland.LogInfo("Entry deleted from database")

}
func RemovePackage(args []string) {

	name := parseRemoveFlags(args)

	entry := ensurePackageExists(name)

	removeSymlink(entry)
	removeBinary(entry)
	removeFromDB(entry)

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
