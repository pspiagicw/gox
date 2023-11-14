package remove

import (
	"flag"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/database"
	"github.com/pspiagicw/gox/pkg/help"
)

const FAKE_INSTALLL = "/home/pspiagicw/.local/bin/"

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

	goreland.LogInfo("Removing file %s", entry.Path)
}
