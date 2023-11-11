package handle

import (
	"github.com/pspiagicw/goreland"
	"github.com/pspiagicw/gox/pkg/install"
	"github.com/pspiagicw/gox/pkg/remove"
)

func HandleArgs(cmd string, args []string) {

    handlers := map[string]func([]string) {
        "help": printHelp,
        "install": install.InstallPackage,
        "remove": remove.RemovePackage,
        "list": notImplemented,
        "version": notImplemented,
    }

    handlerFunc, exists := handlers[cmd]
    if exists {
        handlerFunc(args)
    } else {
        printHelp([]string{})
    }
}

func printHelp(args []string) {
    // TODO: Helper function to print HELP.
    goreland.LogInfo("Not implemented help printing!")
}

func notImplemented(args []string) {
    goreland.LogFatal("Not implemented this command yet!.")
}
