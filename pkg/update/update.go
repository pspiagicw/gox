package update

import (
	"flag"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/database"
	"github.com/pspiagicw/gox/pkg/help"
	"github.com/pspiagicw/gox/pkg/install"
)

func parseUpdateFlags(args []string) string {
	flag := flag.NewFlagSet("gox update", flag.ExitOnError)

    flag.Usage = help.HelpUpdate
	flag.Parse(args)
	args = flag.Args()

	if len(args) == 0 {
		help.HelpUpdate()
		goreland.LogFatal("No package name provided.")

	}

	return args[0]
}

func UpdatePackage(args []string) {
	name := parseUpdateFlags(args)

	db := database.ParseDatabase()

	entry, exists := db.Packages[name]

	if !exists {
		goreland.LogFatal("Package does not exist!")
	}

	goreland.LogInfo("Using '%s' for [%s]", entry.URL, entry.Name)

	install.InstallPackage([]string{entry.URL})
}
func UpdateAll(args []string) {
	db := database.ParseDatabase()

	for name := range db.Packages {
		goreland.LogInfo("Updating package [%s]", name)
		UpdatePackage([]string{name})
	}
	goreland.LogSuccess("Updated all packages!")
}
