package remove

import (
	"flag"
	"os"

	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/help"
)
const FAKE_INSTALLL = "/home/pspiagicw/.local/bin/"

func parseRemoveFlags(args []string) string {
    flag := flag.NewFlagSet("gox remove", flag.ExitOnError)

    flag.Parse()
    args := flag.Args()

    if len(args) == 0 {
        help.HelpRemove()
        goreland.LogFatal("No package name provided.")

    }

    return args[0]
}

func RemovePackage(args []string) {
    name := parseRemoveFlags(args)
    // TODO: Implement Removal of pacakge.
    goreland.LogFatal("NOT IMPLELMENTED YET!")

}
